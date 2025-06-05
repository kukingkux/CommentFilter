package main

type comment struct {
	id int
	sender string
	text string
	status int
	isHidden bool
}

type keywordScore struct {
	word string
	score int
}

type modifierWord struct {
	word string
	multiplier float64
}

const NMAX = 100000
const MAX_P_KEYWORDS = 50
const MAX_N_KEYWORDS = 50
const MAX_NEGATION_WORDS = 10
const MAX_INTENSIFIER_WORDS = 10
const MAX_DIMINISHER_WORDS = 10