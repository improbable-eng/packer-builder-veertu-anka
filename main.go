package main

import (
	"github.com/veertuinc/packer-builder-veertu-anka/builder/anka"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(anka.Builder))
	server.Serve()
}
