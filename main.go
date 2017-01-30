package main

import (
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
	uuid "github.com/satori/go.uuid"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: fmt.Sprintf("gofn-%s", uuid.NewV4().String()),
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
