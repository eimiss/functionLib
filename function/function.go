package function

// The Function interface defines a single method Execute that:
// Takes four parameters: input (string), euclideanDistance (int), widthImage (int), and isColored (bool)
// Returns two values: a string and an error
// The purpose of the Execute method is not explicitly defined in this interface, but based on the context of the codebase, it likely executes a function that converts an image to ASCII art.
type Function interface {
	Execute(input string, euclideanDistance int, widthImage int, isColored bool) (string, error)
}
