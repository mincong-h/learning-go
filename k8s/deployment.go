package k8s

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteDeployments(ctx context.Context, client kubernetes.Interface) error {
	return client.AppsV1().Deployments("my-namespace").DeleteCollection(ctx, v1.DeleteOptions{}, v1.ListOptions{
		LabelSelector: "tag=foo",
	})
}
