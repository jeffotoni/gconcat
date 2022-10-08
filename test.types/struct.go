package testtypes

type Y struct {
	i string
	X string
}

func (Y) WorkingInternalVar(i string) (a Y) {
	a.i = i
	a.X = "exported"
	return
}

func (Y) Fries(string) string {
	return "good"
}
