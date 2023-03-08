package main

import "fmt"

func NewASQ(t Target) (asq *ASQ) {
	asq = &ASQ{
		pos:     -1,
		subject: t,
		scores:  make(map[Aggregator]int),
	}
	switch t {
	case ADULT:
		asq.questions = adultQs
	case BOY, GIRL:
		asq.questions = childQs
	default:
		asq = nil
	}
	return
}

func (s *ASQ) Next() string {
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

func (s *ASQ) Add(selection rune) bool {
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

func (s *ASQ) PrintResult() {
	HoB := s.scores[PmB] + s.scores[PvB]
	totalB := s.scores[PmB] + s.scores[PvB] + s.scores[PsB]
	totalG := s.scores[PmG] + s.scores[PvG] + s.scores[PsG]
	GmB := totalG - totalB

	fmt.Println("\n\nYour score:")
	fmt.Printf("PmB = %d\tPmG = %d\n", s.scores[PmB], s.scores[PmG])
	fmt.Printf("PvB = %d\tPvG = %d\n", s.scores[PvB], s.scores[PvG])
	fmt.Printf("HoB = %d\n", HoB)
	fmt.Printf("PsB = %d\tPsG = %d\n", s.scores[PsB], s.scores[PsG])
	fmt.Printf("Total B = %d\tTotal G = %d\n", totalB, totalG)
	fmt.Printf("G - B = %d\n\n", GmB)

	switch s.subject {
	case ADULT:
		s.PrintAdultSummary(HoB, totalB, totalG, GmB)
	case BOY:
		s.PrintBoySummary(totalB, totalG, GmB)
	case GIRL:
		s.PrintGirlSummary(totalB, totalG, GmB)
	}
}

func (s *ASQ) PrintBoySummary(totalB int, totalG int, GmB int) {
	fmt.Println("Score interpretation for boys")
	fmt.Print("Total B score interpretation: ")
	switch {
	case totalB >= 12:
		fmt.Println("very pessimistic")
	case totalB >= 8:
		fmt.Println("average")
	default:
		fmt.Println("optimistic")
	}

	fmt.Print("Total G score interpretation: ")
	switch {
	case totalG > 14:
		fmt.Println("optimistic")
	case totalG >= 10 && totalG <= 14:
		fmt.Println("average")
	case totalG < 10:
		fmt.Println("great pessimism")
	}

	fmt.Print("Total G - B score interpretation: ")
	switch {
	case GmB <= 1:
		fmt.Println("very pessimistic; risk of depression")
	case GmB <= 3:
		fmt.Println("somewhat pessimistic")
	case GmB <= 5:
		fmt.Println("average")
	default:
		fmt.Println("optimistic")
	}
}

func (s *ASQ) PrintGirlSummary(totalB int, totalG int, GmB int) {
	fmt.Println("Score interpretation for girls")
	fmt.Print("Total B score interpretation: ")
	switch {
	case totalB >= 11:
		fmt.Println("very pessimistic")
	case totalB >= 7:
		fmt.Println("average")
	default:
		fmt.Println("optimistic")
	}

	fmt.Print("Total G score interpretation: ")
	switch {
	case totalG > 14:
		fmt.Println("optimistic")
	case totalG >= 10 && totalG <= 14:
		fmt.Println("average")
	case totalG < 10:
		fmt.Println("great pessimism")
	}

	fmt.Print("Total G - B score interpretation: ")
	switch {
	case GmB <= 2:
		fmt.Println("very pessimistic; risk of depression")
	case GmB <= 4:
		fmt.Println("somewhat pessimistic")
	case GmB <= 7:
		fmt.Println("average")
	default:
		fmt.Println("optimistic")
	}
}

func (s *ASQ) PrintAdultSummary(HoB int, totalB int, totalG int, GmB int) {
	fmt.Println("Score interpretation for adults")
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
	switch HoB {
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
