package main

import (
	"github.com/pulsar-sigs/bk-client-go/pkg/stream"
)

func main() {

	client := &stream.StorageClient{}
	client.Hello()
}
