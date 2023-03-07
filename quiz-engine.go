package main

import "fmt"

func New(t SurveyType) (survey *Survey) {
	survey = &Survey{
		pos:        0,
		surveyType: t,
		scores:     make(map[Aggregator]int),
	}
	switch t {
	case ADULT:
		survey.questions = adultQs
	case CHILD:
		survey.questions = childQs
	default:
		survey = nil
	}
	return
}

func (s *Survey) Next() string {
	if s.pos >= len(s.questions) {
		return ""
	}
	q := fmt.Sprintf("%d. %s\n\tA. %s\n\tB. %s\nYour answer: ",
		s.pos+1,
		s.questions[s.pos].q,
		s.questions[s.pos].options[0],
		s.questions[s.pos].options[1])
	s.pos++
	return q
}

func (s *Survey) Add(agg Aggregator) {
	s.scores[agg]++
}

func (s *Survey) PrintResult() {
	fmt.Printf("Scores: %v", s.scores)
	fmt.Printf("PmB = %d\tPmG = %d\n", s.scores[PmB], s.scores[PmG])
	fmt.Printf("PvB = %d\tPvG = %d\n", s.scores[PvB], s.scores[PvG])
	fmt.Printf("HoB = %d\n", s.scores[PmB]+s.scores[PvB])
	fmt.Printf("PsB = %d\tPsG = %d\n", s.scores[PsB], s.scores[PsG])
	totalB := s.scores[PmB] + s.scores[PvB] + s.scores[PsB]
	totalG := s.scores[PmG] + s.scores[PvG] + s.scores[PsG]
	fmt.Printf("Total B = %d\tTotal G = %d\n", totalB, totalG)
	fmt.Printf("G - B = %d\n", totalG-totalB)
}

type Survey struct {
	pos        int
	questions  []Question
	surveyType SurveyType
	scores     map[Aggregator]int
}

type Question struct {
	q       string
	options []Option
}

type Option struct {
	txt string
	agg Aggregator
}

type Aggregator string

type SurveyType int

const (
	ADULT = 1
	CHILD = 2
)

const (
	PsB = "PsB"
	PsG = "PsG"
	PmB = "PmB"
	PmG = "PmG"
	PvB = "PvB"
	PvG = "PvG"
	NIL = ""
)
