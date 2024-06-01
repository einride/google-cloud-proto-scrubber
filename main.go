package main

import (
	"flag"
	"log"
	"os"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "file descriptor set to scrub")
	flag.Parse()
	if filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var fileDescriptorSet descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(input, &fileDescriptorSet); err != nil {
		log.Fatal(err)
	}
	for _, file := range fileDescriptorSet.GetFile() {
		file.Dependency = scrubDependencies(file.GetDependency())
		scrubResourceDefinitions(file)
		for _, message := range file.GetMessageType() {
			scrubResource(message)
			scrubResourceReferences(message)
			scrubFieldBehaviors(message)
			scrubMessageFields(message)
		}
		for _, service := range file.GetService() {
			scrubServiceOptions(service)
			for _, method := range service.GetMethod() {
				scrubLongRunningOperationInfo(method)
				scrubMethodSignature(method)
			}
		}
	}
	output, err := proto.Marshal(&fileDescriptorSet)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(filename, output, 0o600); err != nil {
		log.Fatal(err)
	}
}

func scrubDependencies(dependencies []string) []string {
	var n int
	for _, dependency := range dependencies {
		switch dependency {
		case "google/api/client.proto",
			"google/api/resource.proto",
			"google/api/field_behavior.proto":
			continue
		}
		dependencies[n] = dependency
		n++
	}
	return dependencies[:n]
}

func scrubResourceDefinitions(file *descriptorpb.FileDescriptorProto) {
	if file.GetOptions() != nil {
		proto.ClearExtension(file.GetOptions(), annotations.E_ResourceDefinition)
	}
}

func scrubResource(message *descriptorpb.DescriptorProto) {
	if message.GetOptions() != nil {
		proto.ClearExtension(message.GetOptions(), annotations.E_Resource)
	}
}

func scrubServiceOptions(service *descriptorpb.ServiceDescriptorProto) {
	if service.GetOptions() != nil {
		proto.ClearExtension(service.GetOptions(), annotations.E_DefaultHost)
	}
}

func scrubMessageFields(message *descriptorpb.DescriptorProto) {
	n := 0
	for _, field := range message.GetField() {
		if field.GetTypeName() != ".google.api.ResourceDescriptor" {
			message.Field[n] = field
			n++
		}
	}
	message.Field = message.GetField()[:n]
	for _, nestedMessage := range message.GetNestedType() {
		scrubMessageFields(nestedMessage)
	}
}

func scrubResourceReferences(message *descriptorpb.DescriptorProto) {
	for _, field := range message.GetField() {
		if field.GetOptions() != nil {
			proto.ClearExtension(field.GetOptions(), annotations.E_ResourceReference)
		}
	}
	for _, nestedMessage := range message.GetNestedType() {
		scrubResourceReferences(nestedMessage)
	}
}

func scrubFieldBehaviors(message *descriptorpb.DescriptorProto) {
	for _, field := range message.GetField() {
		if field.GetOptions() != nil {
			proto.ClearExtension(field.GetOptions(), annotations.E_FieldBehavior)
		}
	}
	for _, nestedMessage := range message.GetNestedType() {
		scrubFieldBehaviors(nestedMessage)
	}
}

func scrubLongRunningOperationInfo(method *descriptorpb.MethodDescriptorProto) {
	if method.GetOptions() != nil {
		proto.ClearExtension(method.GetOptions(), longrunningpb.E_OperationInfo)
	}
}

func scrubMethodSignature(method *descriptorpb.MethodDescriptorProto) {
	if method.GetOptions() != nil {
		proto.ClearExtension(method.GetOptions(), annotations.E_MethodSignature)
	}
}
