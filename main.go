package main

import (
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "Hello",
		Config: &docker.Config{
			Image:     "debian:8",
			StdinOnce: true,
			OpenStdin: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(container)
	fmt.Println("Container created!")
}
