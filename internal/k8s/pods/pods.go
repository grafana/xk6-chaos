package pods

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

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
	dt := time.Now()
	fmt.Println("Pod " + podName + " terminated at " + dt.Format("2006-Jan-02 15:04:05"))
	return err
}

// KillByKeyword kills the first pod with a name that contains the specified keyword
func (pods *Pods) KillByKeyword(ctx context.Context, namespace string, podKeyword string) (string, error) {
	// Iterate through podnames in list and find one that matches keyword
	var podsList, err = pods.List(ctx, namespace)
	var podToCheck = ""
	for i := 0; i < len(podsList); i++ {
		podToCheck = podsList[i]
		if strings.Contains(podToCheck, podKeyword) {
			fmt.Println(podToCheck + " contains keyword '" + podKeyword + "' and will be terminated.")
			break
		}
	}
	var podName = podToCheck
	err = pods.KillByName(ctx, namespace, podName)
	return podName, err
}

// KillRandom kills a random pod within a namespace
func (pods *Pods) KillRandom(ctx context.Context, namespace string) (string, error) {
	// Randomly choose a pod from the given namespace
	var podsList, err = pods.List(ctx, namespace)
	var randomNum = rand.Intn(len(podsList) - 1)
	var randomPod = podsList[randomNum]
	fmt.Println(randomPod + " has been selected for termination.")
	// Kill the random pod by name
	err = pods.KillByName(ctx, namespace, randomPod)
	return randomPod, err
}

// Status of a pod with a specific name in a specific namespace
func (pods *Pods) Status(ctx context.Context, namespace string, podName string) (coreV1.PodStatus, error) {
	pod, err := pods.client.CoreV1().Pods(namespace).Get(ctx, podName, v1.GetOptions{})

	return pod.Status, err
}
