package summary

var summary *Summary

func GetSummary() *Summary {
	if summary == nil {
		summary = &Summary{}
	}
	return summary
}

func (s *Summary) AddResult(result PodkillerResult) {
	s.Results = append(s.Results, result)
}

func (s *Summary) GetResults() []PodkillerResult {
	return s.Results
}

type Summary struct {
	Results []PodkillerResult
}

type PodkillerResult struct {
	Victim    string
	PodCount  PodCount
	Timestamp string
}

type PodCount struct {
	Before int
	After  int
}
