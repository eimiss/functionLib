package function

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

// ImageToASCIIFunction is a stateless function that converts images (PNG or JPEG) to ASCII art.
// It supports optional colored output using ANSI escape sequences.
type ImageToASCIIFunction struct{}

var asciiChars = "@%#*+=-:. "

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Function to initialize the ascii slice with proper dimensions
func initializeAscii(height, width int) [][]rune {
	ascii := make([][]rune, height)
	for y := 0; y < height; y++ {
		ascii[y] = make([]rune, width)
	}
	return ascii
}

// Simple Euclidean distance in RGB space
func colorDifference(c1, c2 color.Color) int {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	return abs(int(r1)-int(r2)) + abs(int(g1)-int(g2)) + abs(int(b1)-int(b2))
}

// Function to check if a pixel is an 'O'
func isO(y, x, height, width int, ascii [][]rune) bool {
	if y < 0 || y >= height || x < 0 || x >= width {
		return false
	}
	return ascii[y][x] == 'O'
}

// Function to determine which contour symbol to use based on adjacent 'O's
func getContourSymbol(ascii [][]rune, y, x, height, width int) rune {
	left := isO(y, x-1, height, width, ascii)
	right := isO(y, x+1, height, width, ascii)
	up := isO(y-1, x, height, width, ascii)
	down := isO(y+1, x, height, width, ascii)
	upLeft := isO(y-1, x-1, height, width, ascii)
	upRight := isO(y-1, x+1, height, width, ascii)
	downLeft := isO(y+1, x-1, height, width, ascii)
	downRight := isO(y+1, x+1, height, width, ascii)

	// Decision logic for contour symbols
	switch {
	case (upRight || downLeft):
		return '╱'
	case (upLeft || downRight) && !(upRight || downLeft):
		return '╲'
	case left || right:
		return '─'
	case up || down:
		return '│'
	default:
		return 'O'
	}
}

func refineContours(ascii [][]rune, height, width int) {
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			curr := ascii[y][x]

			// ┘: connects from left (─) and up (╱)
			if curr == '╱' && ascii[y][x-1] == '─' {
				ascii[y][x] = '┘'
			}

			// └: connects from right (─) and up-left (╲)
			if curr == '╲' && ascii[y][x+1] == '─' {
				ascii[y][x] = '└'
			}

			// ┐: connects from left (─) and down-right (╲)
			if curr == '╲' && ascii[y][x-1] == '─' {
				ascii[y][x] = '┐'
			}

			// ┌: connects from right (─) and down-left (╱)
			if curr == '╱' && ascii[y][x+1] == '─' {
				ascii[y][x] = '┌'
			}
		}
	}
}

// Function to replace 'O' with contour symbols
func replaceWithContours(originalASCII, ascii [][]rune, height, width int) {
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if originalASCII[y][x] == 'O' {
				ascii[y][x] = getContourSymbol(originalASCII, y, x, height, width)
			}
		}
	}
}

// Function to mark the edges of O
func markEdgeO(ascii [][]rune, height, width int) [][]bool {
	edgeMask := make([][]bool, height)
	for y := 0; y < height; y++ {
		edgeMask[y] = make([]bool, width)
	}

	// Mark edge O
	for y := 1; y < height-1; y++ {
		// Left to right
		for x := 1; x < width-1; x++ {
			if ascii[y][x] == 'O' {
				edgeMask[y][x] = true
				break
			}
		}
		// Right to left
		for x := width - 2; x >= 1; x-- {
			if ascii[y][x] == 'O' {
				edgeMask[y][x] = true
				break
			}
		}
	}

	for x := 1; x < width-1; x++ {
		// Top to bottom
		for y := 1; y < height-1; y++ {
			if ascii[y][x] == 'O' {
				edgeMask[y][x] = true
				break
			}
		}
		// Bottom to top
		for y := height - 2; y >= 1; y-- {
			if ascii[y][x] == 'O' {
				edgeMask[y][x] = true
				break
			}
		}
	}

	return edgeMask
}

// Execute converts an image to ASCII art.
//
// Parameters:
//   - inputPath: path to the image (relative or absolute).
//   - euclideanDistance: threshold for edge detection.
//   - widthImage: desired output width in characters.
//   - isColored: whether the ASCII art should be colorized.
//
// Returns:
//   - ASCII string representation of the image.
//   - An error if decoding or processing fails.
func (f ImageToASCIIFunction) Execute(inputPath string, euclideanDistance int, widthImage int, isColored bool) (string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return "", fmt.Errorf("could not open image: %w", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		file.Seek(0, 0)
		img, err = jpeg.Decode(file)
		if err != nil {
			return "", fmt.Errorf("could not decode as PNG or JPEG: %w", err)
		}
	}

	// Resize image based on width
	aspectRatio := float64(img.Bounds().Dy()) / float64(img.Bounds().Dx())
	newHeight := uint(float64(widthImage) * aspectRatio * 0.5)
	resized := resize.Resize(uint(widthImage), newHeight, img, resize.Lanczos3)

	bounds := resized.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Initialize the ascii array
	ascii := initializeAscii(height, width)

	// Generate ASCII + O for edge detection
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			center := resized.At(x, y)
			diff := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					neighbor := resized.At(x+dx, y+dy)
					diff += colorDifference(center, neighbor)
				}
			}
			if diff > euclideanDistance {
				ascii[y][x] = 'O'
			} else {
				gray := color.GrayModel.Convert(center).(color.Gray)
				index := int(gray.Y) * (len(asciiChars) - 1) / 255
				// Convert to grayscale char
				ascii[y][x] = rune(asciiChars[index])
			}
		}
	}

	// Mark edges 'O'
	edgeMask := markEdgeO(ascii, height, width)
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if ascii[y][x] == 'O' && !edgeMask[y][x] {
				gray := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
				index := int(gray.Y) * (len(asciiChars) - 1) / 255
				ascii[y][x] = rune(asciiChars[index]) // Replace 'O' with grayscale chars
			}
		}
	}

	originalASCII := initializeAscii(height, width)
	for y := range ascii {
		copy(originalASCII[y], ascii[y])
	}

	// Replace 'O' with contour symbols
	replaceWithContours(originalASCII, ascii, height, width)
	refineContours(ascii, height, width)

	// Fill edges if needed (replace any 0-value with ' ')
	for y := 0; y < height; y++ {
		if ascii[y] == nil {
			ascii[y] = make([]rune, width)
		}
		for x := 0; x < width; x++ {
			if ascii[y][x] == 'O' {
				ascii[y][x] = '.'
			}
		}
	}

	// Convert the 2D rune array into a string
	result := ""
	if isColored {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				char := ascii[y][x]
				r, g, b, _ := resized.At(x, y).RGBA()
				r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

				// Optional: highlight contour characters in color (e.g. red)
				if char == '╱' || char == '╲' || char == '─' || char == '│' ||
					char == '┐' || char == '┘' || char == '└' || char == '┌' {
					result += fmt.Sprintf("\x1b[38;2;255;0;0m%c\x1b[0m", char) // Red
				} else {
					result += fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", r8, g8, b8, char)
				}
			}
			result += "\n"
		}
	} else {
		for _, row := range ascii {
			result += string(row) + "\n"
		}

	}
	return result, nil
}
