package k8s

import (
	"context"
	"testing"

	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8sTesting "k8s.io/client-go/testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func TestDeleteDeployments(t *testing.T) {
	// Given
	deployment1 := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "deployment-1",
			Namespace: "my-namespace",
			Labels:    map[string]string{"tag": "foo"},
		},
	}
	deployment2 := v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "deployment-2",
			Namespace: "my-namespace",
			Labels:    map[string]string{"tag": "bar"},
		},
	}

	client := fake.NewSimpleClientset(&deployment1, &deployment2)

	assertionsCalled := false
	client.AddReactor("delete-collection", "*", func(action k8sTesting.Action) (handled bool, ret runtime.Object, err error) {
		deletion := action.(k8sTesting.DeleteCollectionAction)
		assert.Equal(t, "my-namespace", deletion.GetNamespace())
		assert.Equal(t, "tag=foo", deletion.GetListRestrictions().Labels.String())
		assertionsCalled = true
		return false, nil, nil
	})

	// When
	err := DeleteDeployments(context.TODO(), client)

	// Then
	assert.NoError(t, err)
	assert.True(t, assertionsCalled)

	remainingDeployments, err := client.AppsV1().Deployments("my-namespace").List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	assert.Equal(t, []v1.Deployment{deployment1, deployment2}, remainingDeployments.Items)
}
