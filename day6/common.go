package day6

const (
	_ = iota
	directionUp
	directionRight
	directionDown
	directionLeft
)

type guardState struct {
	CurrCoord     coordinate
	CurrDirection int
}

type coordinate struct {
	x int
	y int
}
