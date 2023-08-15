package main

import (
	"fmt"
	"net/http"
	"time"
)

// The porpuse of wich is just verify wheter or not the domain is responding to traffic.
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Channels are used to communcate in between different running go routines.
	c := make(chan string)

	for _, link := range links {
		// Go routine - create a new thread go routine and run this function with it to running concurrently.
		go checkLink(link, c)
	}

	/*
		   // Statement to receive the data from the channel. Listen for each message and print it out.
		   	for i := 0; i < len(links); i++ {
		   		fmt.Println(<-c)
			}
	*/

	/* Infinite loop, if you want make sure that we just continually fetch link from now to through eternity...
	for {
		go checkLink(<-c, c)
	}
	*/
	// Return some value and after that, assign it to this variable
	for l := range c {
		/*
			go checkLink(l, c)
				Instead, we're going to place a function literal, equivalent to a Anonymous Functon in Javascript and PHP, or Lambda in Ruby, Python and C#.
				In Go, a Function Literal is an unnamed function that we use to wrap some little chunk of code, so we can execute it at some point in the
				future. We have to add a () after the function literal to invoke.
		*/
		go func(link string) {
			// Integrate a little pause between each check link call wiht 'func Sleep(d Duration)'
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) /* We need to provide 'l' in checkLink(l, c), the link we're trying to fetch, as an argument to this function literal and we do have to
		receive it as an argument inside the actual function definition. We call 'link' to not confusing.
		That string is copied in memory and then goroutine has access to that copy as opposed to the original value to 'l', so now 'l' can changes
		as much as it pleases and we don't have to worry about still having our go routine referecing that some copy or same address in memory.
		*/

	}
}

// We need do make this function communicate with the channel passing as an argument and type
func checkLink(link string, c chan string) {
	// We get two values back from calling this function. The first is a struct representing the actual response that came back, and then the second is an error if one occurred.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "migh be down!")
		/* Send a message through our channel
		c <- "Might be down, I think"
		*/
		/// Or for kinda status checker, we take that link and push it into our channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	/* Send a message through our channel
	c <- "Yep, it's up"
	*/
	c <- link
}
