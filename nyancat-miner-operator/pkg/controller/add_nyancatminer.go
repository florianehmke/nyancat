package controller

import (
	"github.com/florianehmke/nyancat/nyancat-miner-operator/pkg/controller/nyancatminer"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nyancatminer.Add)
}
