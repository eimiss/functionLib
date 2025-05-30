# functionLib
Image-to-ASCII art generator with optional colored output and edge detection via difference in colour. Includes HTTP API support.

# Using as API (for curl/postman)
Clone this repository

Configure your golang (1.20 version needs to be installed)

Check if all libraries are installed

Write command in terminal: `go run .\main.go`

API usage example:

`curl --output - "http://localhost:8080/ascii?input=path_to_image&distance=200000&width=50&colored=true"`

Example of path_to_image: `C:/images/cat.png`

# Using as a library
Create a folder where you will store your golang project

Create new Go module `go mod init <name>`

Install the library `go get github.com/eimiss/functionLib`

Use the example code shown below.

Run the client `go run main.go`

## Testing examples
IF you want to test the functionality without any additional work, there is an example folder.

Just run example file and you will get some of the results.

## Variables
These are the variables of the component. Structure: name:name when using in html/curl:type - description
•	inputPath:input:string – image input path.
•	euclideanDistance:distance:int – euclidean distance.
•	widthImage:width:int – photo width.
•	isColored:colored:bool – is colored required boolean.

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
```
## Example
![examplephoto](https://github.com/user-attachments/assets/59fe44b3-c94d-4501-a266-8dff36f29217)
