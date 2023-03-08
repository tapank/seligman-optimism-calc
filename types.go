package main

// ASQ is Attributional Style Questionnaire
type ASQ struct {
	pos       int
	questions []ASQuestion
	subject   Target
	scores    map[Aggregator]int
}

type ASQuestion struct {
	q       string
	options []ASQOption
}

type ASQOption struct {
	txt string
	agg Aggregator
}

type Aggregator string

type Target int

const (
	ADULT = iota
	BOY
	GIRL
)

type CESDTest struct {
	pos       int
	questions []CESDQuestion
	score     int
}

type CESDQuestion struct {
	q   string
	opt []CESDOption
}

type CESDOption struct {
	txt   string
	score int
}

const (
	PsB = "PsB"
	PsG = "PsG"
	PmB = "PmB"
	PmG = "PmG"
	PvB = "PvB"
	PvG = "PvG"
	NIL = ""
)
