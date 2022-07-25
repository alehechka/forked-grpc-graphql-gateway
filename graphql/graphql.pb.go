// graphql.proto
//
// Copyright (c) 2020 ysugimoto
//
// Released under the MIT license.
// see https://opensource.org/licenses/MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: graphql.proto

package graphql

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// explicit schema declaration enum
type GraphqlType int32

const (
	// schema will generate as Query
	GraphqlType_QUERY GraphqlType = 0
	// schema will generate as Mutation
	GraphqlType_MUTATION GraphqlType = 1
	// schema will generate as Resolver. Resolver behaves not listed in query, but can resolve nested field.
	GraphqlType_RESOLVER GraphqlType = 2
)

// Enum value maps for GraphqlType.
var (
	GraphqlType_name = map[int32]string{
		0: "QUERY",
		1: "MUTATION",
		2: "RESOLVER",
	}
	GraphqlType_value = map[string]int32{
		"QUERY":    0,
		"MUTATION": 1,
		"RESOLVER": 2,
	}
)

func (x GraphqlType) Enum() *GraphqlType {
	p := new(GraphqlType)
	*p = x
	return p
}

func (x GraphqlType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphqlType) Descriptor() protoreflect.EnumDescriptor {
	return file_graphql_proto_enumTypes[0].Descriptor()
}

func (GraphqlType) Type() protoreflect.EnumType {
	return &file_graphql_proto_enumTypes[0]
}

func (x GraphqlType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphqlType.Descriptor instead.
func (GraphqlType) EnumDescriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{0}
}

// Base GraphqlService that gateway is built upon.
type GraphqlService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GraphqlService) Reset() {
	*x = GraphqlService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlService) ProtoMessage() {}

func (x *GraphqlService) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlService.ProtoReflect.Descriptor instead.
func (*GraphqlService) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{0}
}

// Extend MethodOptions in order to define GraphQL Query or Mutation.
// User can use this option as following:
//
// service Greeter {
//    rpc SayHello(HelloRequest) returns (HelloReply) {
//      option (graphql.schema) = {
//        type: QUERY    // declare as Query
//        name: "hello"  // query name
//      }
//    }
// }
//
// Since gRPC reason, it has limitation that the response could not be repeated.
// it's dificcurl to respond array response, so that we accept "response.pluck"
// in order to expose repeated fields in response message.
//
// For instance:
//
// message Member {
//   string name = 1;
// }
//
// message ListMembersResponse {
//   repeated Member members = 1; -- could be array response
// }
//
// message ListMembersRequest {
// }
//
// service MemberService {
//    rpc ListMembers(ListMembersRequest) returns (ListMembersResponse) {
//      option (graphql.schema) = {
//        type: QUERY
//        name: "members"
//        response {
//          repeated : true
//          pluck: "members" // Query will respond [Member] instead of ListMembersResponse
//        }
//      }
//    }
// }
//
// In mutation declaration:
//
// service MemberService {
//    rpc CreateMember(CreateMemberRequest) returns (Member) {
//      option (graphql.schema) = {
//        type: MUTATION        // declare as Mutation
//        name: "cretemember"   // mutation name
//      }
//    }
// }
//
// The Mutation's input always becomes an input object, so you need to declare argument name.
//
// message Member {
//   string name = 1;
// }
//
// message CreateMemberRequest {
//   string name = 1;
// }
//
// service MemberService {
//    rpc CreateMember(CreateMemberRequest) returns (Member) {
//      option (graphql.schema) = {
//        type: MUTATION
//        name: "createmember"
//        request {
//          name: "member" // this is equivalent to createbook(member: Member): Member in GraphQL
//        }
//      }
//    }
// }
//
// Finally, user can access this query via /graphql?query={members{name}}
type GraphqlSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// graphql type. Enum of QUERY or MUTATION is valid value
	Type GraphqlType `protobuf:"varint,1,opt,name=type,proto3,enum=graphql.GraphqlType" json:"type,omitempty"`
	// query name. this field is required
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Query request object configuration
	Request *GraphqlRequest `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// Query response object configuration
	Response *GraphqlResponse `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GraphqlSchema) Reset() {
	*x = GraphqlSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlSchema) ProtoMessage() {}

func (x *GraphqlSchema) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlSchema.ProtoReflect.Descriptor instead.
func (*GraphqlSchema) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{1}
}

func (x *GraphqlSchema) GetType() GraphqlType {
	if x != nil {
		return x.Type
	}
	return GraphqlType_QUERY
}

func (x *GraphqlSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphqlSchema) GetRequest() *GraphqlRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *GraphqlSchema) GetResponse() *GraphqlResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

// configuration option for request
type GraphqlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define input name.
	// This field enables only for mutation and note that if this field is specified,
	// the gRPC request message will be dealt with an input.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Define pluck message fields
	Plucks []string `protobuf:"bytes,2,rep,name=plucks,proto3" json:"plucks,omitempty"`
}

func (x *GraphqlRequest) Reset() {
	*x = GraphqlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlRequest) ProtoMessage() {}

func (x *GraphqlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlRequest.ProtoReflect.Descriptor instead.
func (*GraphqlRequest) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{2}
}

func (x *GraphqlRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphqlRequest) GetPlucks() []string {
	if x != nil {
		return x.Plucks
	}
	return nil
}

// configuration option for response
type GraphqlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, this response object is required
	// But when you declare "pluck", we respect expose field definition.
	Required bool `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	// Define pluck message field.
	// Note that this field IS NOT repeated, just single string field.
	// It means the response could only be single.
	Pluck string `protobuf:"bytes,2,opt,name=pluck,proto3" json:"pluck,omitempty"`
}

func (x *GraphqlResponse) Reset() {
	*x = GraphqlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlResponse) ProtoMessage() {}

func (x *GraphqlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlResponse.ProtoReflect.Descriptor instead.
func (*GraphqlResponse) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{3}
}

func (x *GraphqlResponse) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *GraphqlResponse) GetPluck() string {
	if x != nil {
		return x.Pluck
	}
	return ""
}

// GraphqlField is FieldOptions in protobuf in order to define type field attribute.
// User can use this option as following:
//
// message Member {
//   string name = 1 [(graphql.field) = {required: true}]; // this field is required in GraphQL, it equivalent to String! on GraphQL
// }
//
// message CreateMemberRequest {
//   string name = 1; [(grahpql.field) = {default: "anonymous"}]; // use default value on input or query
// }
//
// Note that in protobuf, all fields are dealt with optional
// so the same as it, all GraphQL fields are optional as default.
// If you need to be required, use 'required: true' option
type GraphqlField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, this field is required.
	Required bool `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	// Use as other field name (not recommend)
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Define default value on input.
	Default string `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	// Omit this field from graphql definition
	Omit bool `protobuf:"varint,4,opt,name=omit,proto3" json:"omit,omitempty"`
	// Resolve this field by nested query with additional RPC
	Resolver string `protobuf:"bytes,5,opt,name=resolver,proto3" json:"resolver,omitempty"`
}

func (x *GraphqlField) Reset() {
	*x = GraphqlField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlField) ProtoMessage() {}

func (x *GraphqlField) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlField.ProtoReflect.Descriptor instead.
func (*GraphqlField) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{4}
}

func (x *GraphqlField) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *GraphqlField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphqlField) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *GraphqlField) GetOmit() bool {
	if x != nil {
		return x.Omit
	}
	return false
}

func (x *GraphqlField) GetResolver() string {
	if x != nil {
		return x.Resolver
	}
	return ""
}

var file_graphql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*GraphqlService)(nil),
		Field:         1079,
		Name:          "graphql.service",
		Tag:           "bytes,1079,opt,name=service",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*GraphqlField)(nil),
		Field:         1079,
		Name:          "graphql.field",
		Tag:           "bytes,1079,opt,name=field",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*GraphqlSchema)(nil),
		Field:         1079,
		Name:          "graphql.schema",
		Tag:           "bytes,1079,opt,name=schema",
		Filename:      "graphql.proto",
	},
}

// Extension fields to descriptor.ServiceOptions.
var (
	// optional graphql.GraphqlService service = 1079;
	E_Service = &file_graphql_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional graphql.GraphqlField field = 1079;
	E_Field = &file_graphql_proto_extTypes[1]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional graphql.GraphqlSchema schema = 1079;
	E_Schema = &file_graphql_proto_extTypes[2]
)

var File_graphql_proto protoreflect.FileDescriptor

var file_graphql_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb6, 0x01, 0x0a,
	0x0d, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x28,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6c, 0x75, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75,
	0x63, 0x6b, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x75, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x6c, 0x75, 0x63, 0x6b, 0x22, 0x88, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x6f, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2a, 0x34, 0x0a, 0x0b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x10, 0x02, 0x3a, 0x53, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x4b,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4f, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x73, 0x75, 0x67, 0x69,
	0x6d, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graphql_proto_rawDescOnce sync.Once
	file_graphql_proto_rawDescData = file_graphql_proto_rawDesc
)

func file_graphql_proto_rawDescGZIP() []byte {
	file_graphql_proto_rawDescOnce.Do(func() {
		file_graphql_proto_rawDescData = protoimpl.X.CompressGZIP(file_graphql_proto_rawDescData)
	})
	return file_graphql_proto_rawDescData
}

var file_graphql_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_graphql_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_graphql_proto_goTypes = []interface{}{
	(GraphqlType)(0),                  // 0: graphql.GraphqlType
	(*GraphqlService)(nil),            // 1: graphql.GraphqlService
	(*GraphqlSchema)(nil),             // 2: graphql.GraphqlSchema
	(*GraphqlRequest)(nil),            // 3: graphql.GraphqlRequest
	(*GraphqlResponse)(nil),           // 4: graphql.GraphqlResponse
	(*GraphqlField)(nil),              // 5: graphql.GraphqlField
	(*descriptor.ServiceOptions)(nil), // 6: google.protobuf.ServiceOptions
	(*descriptor.FieldOptions)(nil),   // 7: google.protobuf.FieldOptions
	(*descriptor.MethodOptions)(nil),  // 8: google.protobuf.MethodOptions
}
var file_graphql_proto_depIdxs = []int32{
	0, // 0: graphql.GraphqlSchema.type:type_name -> graphql.GraphqlType
	3, // 1: graphql.GraphqlSchema.request:type_name -> graphql.GraphqlRequest
	4, // 2: graphql.GraphqlSchema.response:type_name -> graphql.GraphqlResponse
	6, // 3: graphql.service:extendee -> google.protobuf.ServiceOptions
	7, // 4: graphql.field:extendee -> google.protobuf.FieldOptions
	8, // 5: graphql.schema:extendee -> google.protobuf.MethodOptions
	1, // 6: graphql.service:type_name -> graphql.GraphqlService
	5, // 7: graphql.field:type_name -> graphql.GraphqlField
	2, // 8: graphql.schema:type_name -> graphql.GraphqlSchema
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	6, // [6:9] is the sub-list for extension type_name
	3, // [3:6] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_graphql_proto_init() }
func file_graphql_proto_init() {
	if File_graphql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graphql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_graphql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlSchema); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_graphql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_graphql_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_graphql_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_graphql_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_graphql_proto_goTypes,
		DependencyIndexes: file_graphql_proto_depIdxs,
		EnumInfos:         file_graphql_proto_enumTypes,
		MessageInfos:      file_graphql_proto_msgTypes,
		ExtensionInfos:    file_graphql_proto_extTypes,
	}.Build()
	File_graphql_proto = out.File
	file_graphql_proto_rawDesc = nil
	file_graphql_proto_goTypes = nil
	file_graphql_proto_depIdxs = nil
}
