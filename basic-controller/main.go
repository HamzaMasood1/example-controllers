package main

import (
	"context"
	"fmt"
	mygroupv1alpha1 "github.com/HamzaMasood1/example-controllers/basic-controller/github.com/myid/myresource-crd/pkg/apis/mygroup.example.com/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	controller "sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
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

	controller, err := controller.New(
		"my-operator",
		mgr,
		controller.Options{
			Reconciler: &MyReconciler{},
		})
	panicIf(err)

	err = controller.Watch(
		&source.Kind{
			Type: &mygroupv1alpha1.MyResource{},
		},
		&handler.EnqueueRequestForObject{},
	)

	panicIf(err)

	err = controller.Watch(
		&source.Kind{
			Type: &corev1.Pod{},
		},
		&handler.EnqueueRequestForOwner{
			OwnerType:    &corev1.Pod{},
			IsController: true,
		},
	)
	panicIf(err)

	err = mgr.Start(context.Background())
	panicIf(err)

}

type MyReconciler struct{}

func (m *MyReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	//TODO implement me
	fmt.Printf("reconcile %v\n", request)
	return reconcile.Result{}, nil
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
