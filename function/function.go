package function

type Function interface {
	Execute(input string, euclideanDistance int, widthImage int, isColored bool) (string, error)
}
