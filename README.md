# ProPublica Congress API Go SDK

A Golang API client for interacting with the [ProPublica Congress API](https://www.propublica.org/datastore/api/propublica-congress-api), which provides access to some of the most extensive data available about the United States Congress, both its current sessions and going decades into the past. **This library is in no was affiliated with ProPublica or Congress.**

## Getting started

Add `import "github.com/rabdill/go-congress/congress"` to the top of your script or application file, with the rest of your imports. From there, all you need to start making calls is a client:

```go
c := congress.Client{
    Endpoint: "https://api.propublica.org/congress/v1",
    Key:      "KEYgoesHERE123",
}
```

(You can request a free API key from ProPublica [on their website](https://www.propublica.org/datastore/api/propublica-congress-api), which also has the [API documentation](https://projects.propublica.org/api-docs/congress-api/).)

Authenticating to the Congress API is simple and happens without intervention; you can jump right into making calls now:

```go
members, _ := c.GetDepartingMembers(115, "senate")
```

If you just want to test things out (and make sure your key is working), you can create a file called `main.go` in the root of this repository and paste this code in:

```go
package main

import (
	"fmt"

	"github.com/rabdill/go-congress/congress"
)

func main() {
	c := congress.Client{
		Endpoint: "https://api.propublica.org/congress/v1",
		Key:      "KEYgoesHERE123",
	}
	// answer, err := c.GetMembers(115, "senate")
	// answer, err := c.GetMember("K000388")

	// answer, err := c.GetChamberMembersByState("nj", "house")
	// answer, err := c.GetChamberMembersByDistrict("nj", 1, "senate")
	// answer, err := c.GetMembersByState("nj")

	// answer, err := c.GetNewMembers()
	answer, err := c.GetDepartingMembers(115, "senate")
	if err != nil {
		fmt.Printf("\nOh no!\n|%s|\n", err)
    }
    
    fmt.Printf("\nRESULT COUNT: |%v|", len(answer))
	fmt.Printf("\n\n\n!!!RESULTS!!!\n\n%+v\n", answer)
}
```

From there, all you'll need to do is run `go run main.go` and it should print out a list of members departing the current Senate. Uncommenting any of the other calls will give you different sets of results.