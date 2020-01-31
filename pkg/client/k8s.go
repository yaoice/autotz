package client

import (
	tzclientset "github.com/yaoice/autotz/pkg/generated/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetInClusterClientSet() (*kubernetes.Clientset, *tzclientset.Clientset, error) {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}

	tzClient, err := tzclientset.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	return kubeClient, tzClient, err
}

func GetClusterClientSetWithKC(masterURL, kubeconfig string) (*kubernetes.Clientset, *tzclientset.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		return nil, nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	tzClient, err := tzclientset.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	return kubeClient, tzClient, err
}
