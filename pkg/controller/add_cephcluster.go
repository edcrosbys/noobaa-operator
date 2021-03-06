package controller

import "github.com/noobaa/noobaa-operator/v2/pkg/controller/cephcluster"

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cephcluster.Add)
}
