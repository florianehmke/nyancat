Operator SDK

```
export GO111MODULE=on                                                                                                                                                        ✖ ✹

oc new-project nyancat
operator-sdk new nyancat-miner-operator --type=go

operator-sdk add api --api-version=nyancat.florianehmke.github.com/v1alpha1 --kind=NyanCatMiner
# edit types and regenerate afterwards

operator-sdk generate k8s
operator-sdk generate openapi


oc create -f deploy/crds/nyancat_v1alpha1_nyancatminer_crd.yaml
# if it already exists: 
# oc replace -f deploy/crds/nyancat_v1alpha1_nyancatminer_crd.yaml


operator-sdk add controller --api-version=nyancat/v1alpha1 --kind=NyanCatMiner
# edit controller

operator-sdk up local --namespace nyancat

# edit custom resource file and deploy:
oc create -f deploy/crds/nyancat_v1alpha1_nyancatminer_cr.yaml

```
