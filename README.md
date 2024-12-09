# GoBreach
A Golang library for BreachDirectory's REST API.

## Installation
```
$ go get github.com/ibnaleem/gobreach
```
## Usage
Get a valid [API Key](https://rapidapi.com/rohan-patra/api/breachdirectory/)
```go
package main

import (
	"fmt"
	"log"
	"github.com/ibnaleem/gobreach"
)

func main() {
	
	client, err := gobreach.NewBreachDirectoryClient("your-api-key")
	if err != nil {
		log.Fatal(err)
	}

	email := "example@example.com"


	response, err := client.SearchEmail(email)
	if err != nil {
		log.Fatal(err)
	}


	if response.Found > 0 {
		fmt.Printf("Found %d breaches for %s:\n", response.Found, email)
		for _, entry := range response.Result {
			fmt.Printf("Email: %s, Sources: %s\n", entry.Email, entry.Sources)
		}
	} else {
		fmt.Printf("No breaches found for %s.\n", email)
	}
}
```
