package main

import (
	"fmt"
	mygroupv1alpha1 "github.com/HamzaMasood1/example-controllers/basic-controller/github.com/myid/myresource-crd/pkg/apis/mygroup.example.com/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

func main() {
	fmt.Println("hello")
	scheme := runtime.NewScheme()
	err := clientgoscheme.AddToScheme(scheme)
	if err != nil {
		fmt.Println(err)
	}

	mygroupv1alpha1.AddToScheme()
}
