package pods

import (
	"context"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Pods exposes methods to interact with k8s pods at runtime
type Pods struct {
	client *kubernetes.Clientset
}

// New creates a new pod struct
func New(client *kubernetes.Clientset) *Pods {
	return &Pods{client}
}

// List pods in a specific namespace
func (pods *Pods) List(_ context.Context, namespace string) ([]string, error) {
	podList, err := pods.client.CoreV1().Pods(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	alivePods := make([]string, 0)
	for _, pod := range podList.Items {
		if pod.DeletionTimestamp != nil {
			continue
		}
		alivePods = append(alivePods, pod.Name)
	}

	return alivePods, nil
}

// KillByName kills a specific pod in the specified namespace
func (pods *Pods) KillByName(_ context.Context, namespace string, podName string) error {
	podsInNamespace := pods.client.CoreV1().Pods(namespace)
	err := podsInNamespace.Delete(podName, &v1.DeleteOptions{})

	return err
}

// Status of a pod with a specific name in a specific namespace
func (pods *Pods) Status(_ context.Context, namespace string, podName string) (coreV1.PodStatus, error) {
	pod, err := pods.client.CoreV1().Pods(namespace).Get(podName, v1.GetOptions{})

	return pod.Status, err
}
