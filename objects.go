package main

type record struct {
	key  string
	name string
	seq  string
}

type times struct {
	strLen  int
	runTime int
}

type runTimesArr []times

func (c runTimesArr) Len() int           { return len(c) }
func (c runTimesArr) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c runTimesArr) Less(i, j int) bool { return c[i].strLen > c[j].strLen }

type scores struct {
	key   string
	score int
}

type scoresArr []scores
