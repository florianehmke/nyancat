package nyancatminer

import (
	"context"
	"reflect"

	nyancatv1alpha1 "github.com/florianehmke/nyancat/operator/pkg/apis/nyancat/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_nyancatminer")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new NyanCatMiner Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNyanCatMiner{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("nyancatminer-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource NyanCatMiner
	err = c.Watch(&source.Kind{Type: &nyancatv1alpha1.NyanCatMiner{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner NyanCatMiner
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &nyancatv1alpha1.NyanCatMiner{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileNyanCatMiner implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNyanCatMiner{}

// ReconcileNyanCatMiner reconciles a NyanCatMiner object
type ReconcileNyanCatMiner struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a NyanCatMiner object and makes changes based on the state read
// and what is in the NyanCatMiner.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNyanCatMiner) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NyanCatMiner")

	// Fetch the NyanCatMiner instance
	instance := &nyancatv1alpha1.NyanCatMiner{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// List all pods owned by this instance
	podList := &corev1.PodList{}
	lbs := map[string]string{
		"app":     instance.Name,
		"version": "v0.1",
	}
	labelSelector := labels.SelectorFromSet(lbs)
	listOps := &client.ListOptions{Namespace: instance.Namespace, LabelSelector: labelSelector}
	if err = r.client.List(context.TODO(), listOps, podList); err != nil {
		return reconcile.Result{}, err
	}

	// Count the pods that are pending or running as available
	var available []corev1.Pod
	for _, pod := range podList.Items {
		if pod.ObjectMeta.DeletionTimestamp != nil {
			continue
		}
		if pod.Status.Phase == corev1.PodRunning || pod.Status.Phase == corev1.PodPending {
			available = append(available, pod)
		}
	}
	numAvailable := int32(len(available))
	availableNames := []string{}
	for _, pod := range available {
		availableNames = append(availableNames, pod.ObjectMeta.Name)
	}

	// Update the status if necessary
	status := nyancatv1alpha1.NyanCatMinerStatus{
		PodNames: availableNames,
	}
	if !reflect.DeepEqual(instance.Status, status) {
		instance.Status = status
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "Failed to update status")
			return reconcile.Result{}, err
		}
	}

	if numAvailable > instance.Spec.Replicas {
		reqLogger.Info("Scaling down pods", "Currently available", numAvailable, "Required replicas", instance.Spec.Replicas)
		diff := numAvailable - instance.Spec.Replicas
		dpods := available[:diff]
		for _, dpod := range dpods {
			err = r.client.Delete(context.TODO(), &dpod)
			if err != nil {
				reqLogger.Error(err, "Failed to delete pod", "pod.name", dpod.Name)
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{Requeue: true}, nil
	}

	if numAvailable < instance.Spec.Replicas {
		reqLogger.Info("Scaling up pods", "Currently available", numAvailable, "Required replicas", instance.Spec.Replicas)
		// Define a new Pod object
		pod := newPodForCR(instance)
		// Set instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			reqLogger.Error(err, "Failed to create pod", "pod.name", pod.Name)
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *nyancatv1alpha1.NyanCatMiner) *corev1.Pod {
	labels := map[string]string{
		"app":     cr.Name,
		"version": "v0.1",
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cr.Name + "-pod",
			Namespace:    cr.Namespace,
			Labels:       labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
