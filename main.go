package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/kovetskiy/go-vkcom"
)

const (
	usage = `vk-group-members

Usage:
	vk-group-members -i <id>
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, "1.0", true, true)
	if err != nil {
		panic(err)
	}

	group := args["<id>"].(string)

	api := vk.Api{}

	offset := 0
	for {
		payload := map[string]string{
			"group_id": group,
			"sort":     "id_asc",
			"offset":   fmt.Sprintf("%d", offset),
			"count":    "1000",
		}

		response, err := api.Request("groups.getMembers", payload)
		if err != nil {
			log.Fatal(err)
		}

		users, ok := response["response"].(map[string]interface{})["users"].([]interface{})
		if !ok {
			log.Printf("response: %#v\n", response)
			break
		}

		if len(users) == 0 {
			break
		}

		for _, user := range users {
			fmt.Printf("%.0f\n", user)
		}

		offset = offset + 1000

		log.Printf("%d", offset)
	}
}
