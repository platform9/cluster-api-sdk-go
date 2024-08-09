package capi

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"
	clusterctlClient "sigs.k8s.io/cluster-api/cmd/clusterctl/client"
	cluster "sigs.k8s.io/cluster-api/cmd/clusterctl/client/cluster"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ClusterctlClientInterface interface {
	GetKubeconfig(ctx context.Context, clusterName, namespace string)
	GetWorkloadClusterClient(ctx context.Context, clusterName, namespace string)
}

type ClusterctlClient struct {
	Client     clusterctlClient.Client
	Kubeconfig cluster.Kubeconfig
}

func (c *ClusterctlClient) GetKubeconfig(ctx context.Context, workloadClusterName string, namespace string) (*string, error) {
	clientKubeconfig := clusterctlClient.Kubeconfig{Path: c.Kubeconfig.Path}
	options := clusterctlClient.GetKubeconfigOptions{
		Kubeconfig:          clientKubeconfig,
		Namespace:           namespace,
		WorkloadClusterName: workloadClusterName,
	}
	kubeconfig, err := c.Client.GetKubeconfig(ctx, options)
	if err != nil {
		return nil, err
	}
	return &kubeconfig, nil
}

func (c *ClusterctlClient) GetWorkloadClusterClient(ctx context.Context, workloadClusterName string, namespace string) (client.Client, error) {
	var kubeconfig *string

	kubeconfig, err := c.GetKubeconfig(ctx, workloadClusterName, namespace)
	if err != nil {
		return nil, err
	}

	// use the current context in kubeconfig
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(*kubeconfig))
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := client.New(config, client.Options{})
	if err != nil {
		return nil, err
	}
	return clientset, err
}
