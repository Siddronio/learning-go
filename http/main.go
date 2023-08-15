package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// A custom writer
type logWriter struct{}

// HTTP resquest
func main() {
	// Check documentation on Package Net/http, Index and func Get(url string) (resp *Response, err error)
	resp, err := http.Get("https://google.com")
	// Handler the error
	if err != nil {
		fmt.Println("Error:", err)
		// Exit the program
		os.Exit(1)
	}
	/*
		We need to print out the response body implementing the type Reader:

		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
	/*
		bs := make([]byte, 99999)
		// Now, the response body should take the byte slice, take it's HTML that is contained inside the body and read that data into the byte slice.
		// Essentially, it's saving us from having to write a bungh of different custom function, specifically for working with something that's coming out of the body with custom type.
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
	// Replacing the code above with one line of code using the type Writer, the difference between type Reader, is the byte size is being used truly as a source of input.
	// Using the func Copy(dst Writer, src Reader) (written int64, err error), the first argument (os.Stdout) is a value of type 'File' that implements the Writer interface, while the second
	// argument (resp.Body) is our body field of the response struct that implements the Reader interface. We're passing two values of types that implement the respective interfaces.
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// Defining a writer function that takes the logWriter as a receiver
// Interfaces are a contract to help us manage types
// Just defining this function and associating it with the logWriter, now the logWriter is now implementing the writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
