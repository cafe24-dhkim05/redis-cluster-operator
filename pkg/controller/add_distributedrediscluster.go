package controller

import (
	"github.com/cafe24-dhkim05/redis-cluster-operator/pkg/controller/distributedrediscluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, distributedrediscluster.Add)
}
