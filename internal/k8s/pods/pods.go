package pods

import (
	"context"
	"math/rand"
	"strings"

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
func (pods *Pods) List(ctx context.Context, namespace string) ([]string, error) {
	podList, err := pods.client.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
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
func (pods *Pods) KillByName(ctx context.Context, namespace string, podName string) error {
	podsInNamespace := pods.client.CoreV1().Pods(namespace)
	err := podsInNamespace.Delete(ctx, podName, v1.DeleteOptions{})
	return err
}

// KillByKeyword kills the first pod with a name that contains the specified keyword
func (pods *Pods) KillByKeyword(ctx context.Context, namespace string, podKeyword string) (string, error) {
	var podsList, err = pods.List(ctx, namespace)
	var podToCheck = ""
	for i := 0; i < len(podsList); i++ {
		podToCheck = podsList[i]
		if strings.Contains(podToCheck, podKeyword) {
			break
		}
	}
	var podName = podToCheck
	err = pods.KillByName(ctx, namespace, podName)
	return podName, err
}

// KillRandom kills a random pod within a namespace
func (pods *Pods) KillRandom(ctx context.Context, namespace string) (string, error) {
	var podsList, err = pods.List(ctx, namespace)
	var randomNum = rand.Intn(len(podsList) - 1)
	var randomPod = podsList[randomNum]

	err = pods.KillByName(ctx, namespace, randomPod)
	return randomPod, err
}

// Status of a pod with a specific name in a specific namespace
func (pods *Pods) Status(ctx context.Context, namespace string, podName string) (coreV1.PodStatus, error) {
	pod, err := pods.client.CoreV1().Pods(namespace).Get(ctx, podName, v1.GetOptions{})

	return pod.Status, err
}
