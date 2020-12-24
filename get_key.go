package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
	"time"
	"encoding/base64"
	"io/ioutil"
	//"math"
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
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Lease ID: ", secret.LeaseID)

	var duration = time.Duration(secret.LeaseDuration * 1000) * time.Millisecond

	fmt.Println("Lease Duration: ", duration.String())
	fmt.Println("Lease Renewable: ", secret.Renewable)
	fmt.Println("Key Algorithm: ", secret.Data["key_algorithm"])
	fmt.Println("Key Type: ", secret.Data["key_type"])
	key, _ := base64.StdEncoding.DecodeString(fmt.Sprintf("%v", secret.Data["private_key_data"]))
	ioutil.WriteFile("new_key.json", key, 0644)
	fmt.Println("Successfully outputted the key to new_key.json")
}
