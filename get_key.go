package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
//	"encoding/json"
	"encoding/base64"
	"io/ioutil"
)

var token = os.Getenv("VAULT_TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := client.Logical()
	secret, err := c.Read("gcp/key/my-key-roleset")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Lease ID: ", secret.LeaseID)
	key, _ := base64.StdEncoding.DecodeString(fmt.Sprintf("%v", secret.Data["private_key_data"]))
	ioutil.WriteFile("new_key.json", key, 0644)
	fmt.Println("Successfully outputted the key to new_key.json")
}
