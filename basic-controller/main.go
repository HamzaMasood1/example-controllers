package main

import (
	"fmt"
	mygroupv1alpha1 "github.com/HamzaMasood1/example-controllers/basic-controller/github.com/myid/myresource-crd/pkg/apis/mygroup.example.com/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	fmt.Println("hello")
	scheme := runtime.NewScheme()

	clientgoscheme.AddToScheme(scheme)
	mygroupv1alpha1.AddToScheme(scheme)

	mgr, err := manager.New(
		config.GetConfigOrDie(),
		manager.Options{
			Scheme: scheme,
		})
	panicIf(err)

}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
