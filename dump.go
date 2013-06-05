package main

import (
	proto "code.google.com/p/goprotobuf/proto"
	"flag"
	"fmt"
	"github.com/rdarder/golumio/google_protobuf"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	schema := flag.String("schema", "", "The Schema within the proto file")
	flag.Parse()
	if err != nil {
		fmt.Println("Error reading standard input")
	}
	proto_file_descriptors := &google_protobuf.FileDescriptorSet{}
	if err := proto.Unmarshal(data, proto_file_descriptors); err != nil {
		fmt.Println("Error unmarshaling")
	}

	for _, file_descriptor := range proto_file_descriptors.GetFile() {
		for _, descriptor := range file_descriptor.GetMessageType() {
			if *schema == descriptor.GetName() {
				fmt.Println(*schema)
			}
		}
	}
}
