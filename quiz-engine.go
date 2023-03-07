package main

import "fmt"

func New(t SurveyType) (survey *Survey) {
	survey = &Survey{
		pos:        -1,
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
	s.pos++
	if s.pos == len(s.questions) {
		return ""
	}
	q := fmt.Sprintf("\n%d. %s\n\tA. %s\n\tB. %s\nYour answer (enter a or b): ",
		s.pos+1,
		s.questions[s.pos].q,
		s.questions[s.pos].options[0].txt,
		s.questions[s.pos].options[1].txt)
	return q
}

func (s *Survey) Add(selection rune) bool {
	if s.pos < 0 || s.pos >= len(s.questions) {
		return false
	}
	switch selection {
	case 'a', 'A':
		s.scores[s.questions[s.pos].options[0].agg]++
		return true
	case 'b', 'B':
		s.scores[s.questions[s.pos].options[1].agg]++
		return true
	default:
		return false
	}
}

func (s *Survey) PrintResult() {
	// fmt.Printf("Scores: %v", s.scores)
	fmt.Println("\n\nYour score:")
	fmt.Printf("PmB = %d\tPmG = %d\n", s.scores[PmB], s.scores[PmG])
	fmt.Printf("PvB = %d\tPvG = %d\n", s.scores[PvB], s.scores[PvG])
	fmt.Printf("HoB = %d\n", s.scores[PmB]+s.scores[PvB])
	fmt.Printf("PsB = %d\tPsG = %d\n", s.scores[PsB], s.scores[PsG])
	totalB := s.scores[PmB] + s.scores[PvB] + s.scores[PsB]
	totalG := s.scores[PmG] + s.scores[PvG] + s.scores[PsG]
	fmt.Printf("Total B = %d\tTotal G = %d\n", totalB, totalG)
	fmt.Printf("G - B = %d\n\n", totalG-totalB)

	fmt.Print("PvB score interpretation: ")
	switch s.scores[PvB] {
	case 0, 1:
		fmt.Println("very optimistic")
	case 2, 3:
		fmt.Println("moderately optimistic")
	case 4:
		fmt.Println("average")
	case 5, 6:
		fmt.Println("moderately pessimistic")
	case 7, 8:
		fmt.Println("very pessimistic")
	default:
		fmt.Println("error")
	}

	fmt.Print("PvG score interpretation: ")
	switch s.scores[PvG] {
	case 8, 7:
		fmt.Println("very optimistic")
	case 6:
		fmt.Println("moderately optimistic")
	case 5, 4:
		fmt.Println("average")
	case 3:
		fmt.Println("moderately pessimistic")
	case 2, 1, 0:
		fmt.Println("very pessimistic")
	default:
		fmt.Println("error")
	}

	fmt.Print("HoB score interpretation: ")
	switch s.scores[PvB] + s.scores[PmB] {
	case 0, 1, 2:
		fmt.Println("extraordinarily hopeful")
	case 3, 4, 5, 6:
		fmt.Println("moderately hopeful")
	case 7, 8:
		fmt.Println("average")
	case 9, 10, 11:
		fmt.Println("moderately hopeless")
	case 12, 13, 14, 15, 16:
		fmt.Println("severely hopeless")
	default:
		fmt.Println("error")
	}

	fmt.Print("PsB score interpretation: ")
	switch s.scores[PsB] {
	case 0, 1:
		fmt.Println("very high self-esteem")
	case 2, 3:
		fmt.Println("moderate self-esteem")
	case 4:
		fmt.Println("average")
	case 5, 6:
		fmt.Println("moderately low self-esteem")
	case 7, 8:
		fmt.Println("very low self-esteem")
	default:
		fmt.Println("error")
	}

	fmt.Print("PsG score interpretation: ")
	switch s.scores[PsG] {
	case 8, 7:
		fmt.Println("very optimistic")
	case 6:
		fmt.Println("moderately optimistic")
	case 5, 4:
		fmt.Println("average")
	case 3:
		fmt.Println("moderately pessimistic")
	case 2, 1, 0:
		fmt.Println("very pessimistic")
	default:
		fmt.Println("error")
	}

	fmt.Print("Total B score interpretation: ")
	switch {
	case totalB >= 3 && totalB <= 6:
		fmt.Println("marvelously optimistic. You do not need \"Changing\" chapters")
	case totalB >= 7 && totalB <= 9:
		fmt.Println("moderately optimistic")
	case totalB >= 10 && totalB <= 11:
		fmt.Println("average")
	case totalB >= 12 && totalB <= 14:
		fmt.Println("moderately pessimistic")
	case totalB > 14:
		fmt.Println("crying out for change!")
	default:
		fmt.Println("error")
	}

	fmt.Print("Total G score interpretation: ")
	switch {
	case totalG >= 19:
		fmt.Println("you think about good events very optimistically")
	case totalG >= 17 && totalG <= 18:
		fmt.Println("your thinking is moderately optimistic")
	case totalG >= 14 && totalG <= 16:
		fmt.Println("average")
	case totalG >= 11 && totalG <= 13:
		fmt.Println("you think quite pessimistic")
	case totalG < 11:
		fmt.Println("indicates great pessimism")
	default:
		fmt.Println("error")
	}

	fmt.Print("Total G - B score interpretation: ")
	GmB := totalG - totalB
	switch {
	case GmB > 8:
		fmt.Println("you are very optimistic across the board")
	case GmB >= 6 && GmB <= 8:
		fmt.Println("your are moderately optimistic")
	case GmB >= 3 && GmB <= 5:
		fmt.Println("average")
	case GmB >= 1 && GmB <= 2:
		fmt.Println("moderately pessimistic")
	case GmB < 1:
		fmt.Println("very pessimistic")
	default:
		fmt.Println("error")
	}
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
