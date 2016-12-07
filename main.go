package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/swag"

	apiclient "github.com/shirou/swagger_sample/client"
	"github.com/shirou/swagger_sample/client/user"
)

func main() {

	// make the request to get all items
	p := &user.GetSearchParams{
		UserID: swag.Int64(10),
	}

	// なんかruntime.DefaultTimeoutをちゃんと使ってくれてないので明示的に指定
	resp, err := apiclient.Default.User.GetSearch(p.WithTimeout(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Payload.Name)
}
