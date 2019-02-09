package main

type record struct {
	key  string
	name string
	seq  string
}

/*func (c ByArea) Len() int           { return len(c) }
func (c ByArea) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByArea) Less(i, j int) bool { return c[i].Area > c[j].Area }*/

type times struct {
	strLen  int
	runTime int
}

type runTimesArr []times

func (c runTimesArr) Len() int           { return len(c) }
func (c runTimesArr) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c runTimesArr) Less(i, j int) bool { return c[i].strLen > c[j].strLen }
