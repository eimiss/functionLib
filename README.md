# functionLib

Image-to-ASCII art generator with optional colored output and edge detection. Includes HTTP API support.

## Usage (Go code)

```go
import "github.com/eimiss/functionLib/function"

func main() {
    fn := function.ImageToASCIIFunction{}
    result, err := fn.Execute("path/to/image.png", 200000, 50, true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
}