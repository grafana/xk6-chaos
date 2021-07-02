package podkillers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/simskij/xk6-chaos/internal/experiments/summary"
	"github.com/simskij/xk6-chaos/internal/k8s/pods"
	"github.com/simskij/xk6-chaos/pkg/k8s/client"
)

// This exposes podkiller metadata for use in displaying results.
type Podkillers struct {
	Id              int    `json:"id"`
	ExperimentType  string `json:"expType"`
	NumOfPodsBefore int    `json:"numPodsBefore"`
	NumOfPodsAfter  int    `json:"numPodsAfter"`
	Victims         string `json:"victims"`
	pod             *pods.Pods
}

var ExperimentNumber = 1
var PodkillerSummary = ""

// var victims []string = []string{}

// New creates a new podkiller
func New() *Podkillers {
	experimentNum := ExperimentNumber
	ExperimentNumber++
	c, err := client.New()
	p := pods.New(c)
	if err != nil {
		return nil
	}
	// var expType = "Podkiller"
	// var desc = ""
	// podkiller := Podkillers{
	// 	Id:              experimentNum,
	// 	ExperimentType:  "Pod termination",
	// 	NumOfPodsBefore: 0,
	// 	NumOfPodsAfter:  0,
	// 	Victims:         "",
	// 	pod:             p,
	// }
	return &Podkillers{experimentNum, "Pod termination", 0, 0, "", p}
}

// AddResults saves the names of pods selected to be terminated.
func (p *Podkillers) AddResults(namespace string, victim string) {
	p.GetNumOfPods(namespace)

	sum := summary.GetSummary()
	sum.AddResult(summary.PodkillerResult{
		Victim:    victim,
		Timestamp: p.Timestamp(),
		PodCount: summary.PodCount{
			Before: p.NumOfPodsBefore,
			After:  p.NumOfPodsAfter,
		},
	})
}

// GetVictims returns the string array of all pods selected to be terminated.
// func (p *Podkillers) GetVictims() []string {
// 	victims := p.Victims
// 	return victims
// }

// SetStartingPods saves the number of pods at the beginning of a test as the "before" state.
func (p *Podkillers) SetStartingPods(number int) {
	p.NumOfPodsBefore = number
}

// GetStartingPods returns and saves the number of pods at the beginning of the test for reporting purposes.
func (p *Podkillers) GetStartingPods(namespace string) int {
	var podsAlive, _ = p.pod.List(context.Background(), namespace)
	p.NumOfPodsBefore = len(podsAlive)
	fmt.Println(p.Timestamp() + "Num of pods before: " + strconv.Itoa(p.NumOfPodsBefore))
	return p.NumOfPodsBefore
}

// GetNumOfPods returns and saves the current number of pods.
func (p *Podkillers) GetNumOfPods(namespace string) int {
	time.Sleep(5 * time.Second)
	var podsAlive, _ = p.pod.List(context.Background(), namespace)
	p.NumOfPodsAfter = len(podsAlive)
	fmt.Println(p.Timestamp() + "Num of pods after: " + strconv.Itoa(p.NumOfPodsAfter))
	return p.NumOfPodsAfter
}

// KillPod terminates a k8s pod identified by name.
func (p *Podkillers) KillPod(namespace string, podName string) error {
	p.GetStartingPods(namespace)
	err := p.pod.KillByName(context.Background(), namespace, podName)
	p.AddResults(namespace, podName)
	fmt.Println(p.Timestamp() + "Pod " + podName + " terminated.")
	return err
}

// KillPodLike terminates a k8s pod whose name contains string.
func (p *Podkillers) KillPodLike(namespace string, keyword string) error {
	p.GetStartingPods(namespace)
	podName, err := p.pod.KillByKeyword(context.Background(), namespace, keyword)
	p.AddResults(namespace, podName)
	fmt.Println(p.Timestamp() + "Pod " + podName + " containing keyword '" + keyword + "' terminated.")
	return err
}

// KillRandomPod terminates a pod at random.
func (p *Podkillers) KillRandomPod(namespace string) error {
	p.GetStartingPods(namespace)
	podName, err := p.pod.KillRandom(context.Background(), namespace)
	p.AddResults(namespace, podName)
	fmt.Println(p.Timestamp() + "Random pod " + podName + " terminated.")
	return err
}

// Timestamp constructs the format of a timestamp for logging purposes
func (p *Podkillers) Timestamp() string {
	dt := time.Now()
	tsMsg := "[xk6-chaos-Podkiller" + strconv.Itoa(p.Id) + "-" + dt.Format("2006-Jan-02 15:04:05") + "] "
	return tsMsg
}

func (p *Podkillers) GenerateSummary() string {
	sum := summary.GetSummary()
	output := "\n"
	for i, result := range sum.Results {
		output += fmt.Sprintf(" Victim #%d: %s at %s\n", i, result.Victim, result.Timestamp)
	}
	return output + "\n"
	// PodkillerSummary = "Podkiller id: " + strconv.Itoa(p.Id) + "\nPodkiller NumofPodsBefore: " + strconv.Itoa(p.NumOfPodsBefore) + "\nPodkiller NumOfPodsAfter: " + strconv.Itoa(p.NumOfPodsAfter) + "\nVictims: " + p.Victims
	// summary.ChaosSummary = PodkillerSummary
	// fmt.Println("PodkillerSummary: " + PodkillerSummary)
	// return PodkillerSummary
}
