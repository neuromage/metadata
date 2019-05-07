// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Workspace struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workspace.Unmarshal(m, b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
}
func (m *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(m, src)
}
func (m *Workspace) XXX_Size() int {
	return xxx_messageInfo_Workspace.Size(m)
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Namespace struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ArtifactType struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *Namespace       `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TypeProperties       map[string]*Type `protobuf:"bytes,4,rep,name=type_properties,json=typeProperties,proto3" json:"type_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description          string           `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactType) Reset()         { *m = ArtifactType{} }
func (m *ArtifactType) String() string { return proto.CompactTextString(m) }
func (*ArtifactType) ProtoMessage()    {}
func (*ArtifactType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *ArtifactType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactType.Unmarshal(m, b)
}
func (m *ArtifactType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactType.Marshal(b, m, deterministic)
}
func (m *ArtifactType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactType.Merge(m, src)
}
func (m *ArtifactType) XXX_Size() int {
	return xxx_messageInfo_ArtifactType.Size(m)
}
func (m *ArtifactType) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactType.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactType proto.InternalMessageInfo

func (m *ArtifactType) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArtifactType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactType) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ArtifactType) GetTypeProperties() map[string]*Type {
	if m != nil {
		return m.TypeProperties
	}
	return nil
}

func (m *ArtifactType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ExecutionType struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *Namespace       `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	TypeProperties       map[string]*Type `protobuf:"bytes,4,rep,name=type_properties,json=typeProperties,proto3" json:"type_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description          string           `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExecutionType) Reset()         { *m = ExecutionType{} }
func (m *ExecutionType) String() string { return proto.CompactTextString(m) }
func (*ExecutionType) ProtoMessage()    {}
func (*ExecutionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{3}
}

func (m *ExecutionType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionType.Unmarshal(m, b)
}
func (m *ExecutionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionType.Marshal(b, m, deterministic)
}
func (m *ExecutionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionType.Merge(m, src)
}
func (m *ExecutionType) XXX_Size() int {
	return xxx_messageInfo_ExecutionType.Size(m)
}
func (m *ExecutionType) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionType.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionType proto.InternalMessageInfo

func (m *ExecutionType) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExecutionType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExecutionType) GetNamespace() *Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ExecutionType) GetTypeProperties() map[string]*Type {
	if m != nil {
		return m.TypeProperties
	}
	return nil
}

func (m *ExecutionType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Artifact struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Workspace            *Workspace           `protobuf:"bytes,3,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Properties           map[string]*Value    `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CustomProperties     map[string]*Value    `protobuf:"bytes,6,rep,name=custom_properties,json=customProperties,proto3" json:"custom_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string    `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{4}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetWorkspace() *Workspace {
	if m != nil {
		return m.Workspace
	}
	return nil
}

func (m *Artifact) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Artifact) GetCustomProperties() map[string]*Value {
	if m != nil {
		return m.CustomProperties
	}
	return nil
}

func (m *Artifact) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Artifact) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Artifact) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type Execution struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Workspace            *Workspace           `protobuf:"bytes,3,opt,name=workspace,proto3" json:"workspace,omitempty"`
	Properties           map[string]*Value    `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CustomProperties     map[string]*Value    `protobuf:"bytes,5,rep,name=custom_properties,json=customProperties,proto3" json:"custom_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string    `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{5}
}

func (m *Execution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execution.Unmarshal(m, b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
}
func (m *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(m, src)
}
func (m *Execution) XXX_Size() int {
	return xxx_messageInfo_Execution.Size(m)
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func (m *Execution) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Execution) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Execution) GetWorkspace() *Workspace {
	if m != nil {
		return m.Workspace
	}
	return nil
}

func (m *Execution) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Execution) GetCustomProperties() map[string]*Value {
	if m != nil {
		return m.CustomProperties
	}
	return nil
}

func (m *Execution) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Execution) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Execution) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Execution) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Execution) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type StringType struct {
	Validator            string   `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringType) Reset()         { *m = StringType{} }
func (m *StringType) String() string { return proto.CompactTextString(m) }
func (*StringType) ProtoMessage()    {}
func (*StringType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{6}
}

func (m *StringType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringType.Unmarshal(m, b)
}
func (m *StringType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringType.Marshal(b, m, deterministic)
}
func (m *StringType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringType.Merge(m, src)
}
func (m *StringType) XXX_Size() int {
	return xxx_messageInfo_StringType.Size(m)
}
func (m *StringType) XXX_DiscardUnknown() {
	xxx_messageInfo_StringType.DiscardUnknown(m)
}

var xxx_messageInfo_StringType proto.InternalMessageInfo

func (m *StringType) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type IntType struct {
	Validator            string   `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntType) Reset()         { *m = IntType{} }
func (m *IntType) String() string { return proto.CompactTextString(m) }
func (*IntType) ProtoMessage()    {}
func (*IntType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{7}
}

func (m *IntType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntType.Unmarshal(m, b)
}
func (m *IntType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntType.Marshal(b, m, deterministic)
}
func (m *IntType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntType.Merge(m, src)
}
func (m *IntType) XXX_Size() int {
	return xxx_messageInfo_IntType.Size(m)
}
func (m *IntType) XXX_DiscardUnknown() {
	xxx_messageInfo_IntType.DiscardUnknown(m)
}

var xxx_messageInfo_IntType proto.InternalMessageInfo

func (m *IntType) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type DoubleType struct {
	Validator            string   `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleType) Reset()         { *m = DoubleType{} }
func (m *DoubleType) String() string { return proto.CompactTextString(m) }
func (*DoubleType) ProtoMessage()    {}
func (*DoubleType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{8}
}

func (m *DoubleType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleType.Unmarshal(m, b)
}
func (m *DoubleType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleType.Marshal(b, m, deterministic)
}
func (m *DoubleType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleType.Merge(m, src)
}
func (m *DoubleType) XXX_Size() int {
	return xxx_messageInfo_DoubleType.Size(m)
}
func (m *DoubleType) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleType.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleType proto.InternalMessageInfo

func (m *DoubleType) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

type StringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}
func (*StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{9}
}

func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (m *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(m, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IntValue struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntValue) Reset()         { *m = IntValue{} }
func (m *IntValue) String() string { return proto.CompactTextString(m) }
func (*IntValue) ProtoMessage()    {}
func (*IntValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{10}
}

func (m *IntValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntValue.Unmarshal(m, b)
}
func (m *IntValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntValue.Marshal(b, m, deterministic)
}
func (m *IntValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntValue.Merge(m, src)
}
func (m *IntValue) XXX_Size() int {
	return xxx_messageInfo_IntValue.Size(m)
}
func (m *IntValue) XXX_DiscardUnknown() {
	xxx_messageInfo_IntValue.DiscardUnknown(m)
}

var xxx_messageInfo_IntValue proto.InternalMessageInfo

func (m *IntValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type DoubleValue struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleValue) Reset()         { *m = DoubleValue{} }
func (m *DoubleValue) String() string { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()    {}
func (*DoubleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{11}
}

func (m *DoubleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleValue.Unmarshal(m, b)
}
func (m *DoubleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleValue.Marshal(b, m, deterministic)
}
func (m *DoubleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleValue.Merge(m, src)
}
func (m *DoubleValue) XXX_Size() int {
	return xxx_messageInfo_DoubleValue.Size(m)
}
func (m *DoubleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleValue.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleValue proto.InternalMessageInfo

func (m *DoubleValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Type struct {
	// Types that are valid to be assigned to Type:
	//	*Type_StringType
	//	*Type_IntegerType
	//	*Type_DoubleType
	Type                 isType_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{12}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

type isType_Type interface {
	isType_Type()
}

type Type_StringType struct {
	StringType *StringType `protobuf:"bytes,1,opt,name=string_type,json=stringType,proto3,oneof"`
}

type Type_IntegerType struct {
	IntegerType *IntType `protobuf:"bytes,2,opt,name=integer_type,json=integerType,proto3,oneof"`
}

type Type_DoubleType struct {
	DoubleType *DoubleType `protobuf:"bytes,3,opt,name=double_type,json=doubleType,proto3,oneof"`
}

func (*Type_StringType) isType_Type() {}

func (*Type_IntegerType) isType_Type() {}

func (*Type_DoubleType) isType_Type() {}

func (m *Type) GetType() isType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type) GetStringType() *StringType {
	if x, ok := m.GetType().(*Type_StringType); ok {
		return x.StringType
	}
	return nil
}

func (m *Type) GetIntegerType() *IntType {
	if x, ok := m.GetType().(*Type_IntegerType); ok {
		return x.IntegerType
	}
	return nil
}

func (m *Type) GetDoubleType() *DoubleType {
	if x, ok := m.GetType().(*Type_DoubleType); ok {
		return x.DoubleType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Type) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Type_StringType)(nil),
		(*Type_IntegerType)(nil),
		(*Type_DoubleType)(nil),
	}
}

type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_StringValue
	//	*Value_IntValue
	//	*Value_DoubleValue
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{13}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_StringValue struct {
	StringValue *StringValue `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_IntValue struct {
	IntValue *IntValue `protobuf:"bytes,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue *DoubleValue `protobuf:"bytes,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

func (*Value_StringValue) isValue_Value() {}

func (*Value_IntValue) isValue_Value() {}

func (*Value_DoubleValue) isValue_Value() {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetStringValue() *StringValue {
	if x, ok := m.GetValue().(*Value_StringValue); ok {
		return x.StringValue
	}
	return nil
}

func (m *Value) GetIntValue() *IntValue {
	if x, ok := m.GetValue().(*Value_IntValue); ok {
		return x.IntValue
	}
	return nil
}

func (m *Value) GetDoubleValue() *DoubleValue {
	if x, ok := m.GetValue().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_StringValue)(nil),
		(*Value_IntValue)(nil),
		(*Value_DoubleValue)(nil),
	}
}

type Resource struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{14}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResourceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourceRequest) Reset()         { *m = GetResourceRequest{} }
func (m *GetResourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourceRequest) ProtoMessage()    {}
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{15}
}

func (m *GetResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourceRequest.Unmarshal(m, b)
}
func (m *GetResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourceRequest.Marshal(b, m, deterministic)
}
func (m *GetResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceRequest.Merge(m, src)
}
func (m *GetResourceRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourceRequest.Size(m)
}
func (m *GetResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceRequest proto.InternalMessageInfo

func (m *GetResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Workspace)(nil), "api.Workspace")
	proto.RegisterType((*Namespace)(nil), "api.Namespace")
	proto.RegisterType((*ArtifactType)(nil), "api.ArtifactType")
	proto.RegisterMapType((map[string]*Type)(nil), "api.ArtifactType.TypePropertiesEntry")
	proto.RegisterType((*ExecutionType)(nil), "api.ExecutionType")
	proto.RegisterMapType((map[string]*Type)(nil), "api.ExecutionType.TypePropertiesEntry")
	proto.RegisterType((*Artifact)(nil), "api.Artifact")
	proto.RegisterMapType((map[string]*Value)(nil), "api.Artifact.CustomPropertiesEntry")
	proto.RegisterMapType((map[string]string)(nil), "api.Artifact.LabelsEntry")
	proto.RegisterMapType((map[string]*Value)(nil), "api.Artifact.PropertiesEntry")
	proto.RegisterType((*Execution)(nil), "api.Execution")
	proto.RegisterMapType((map[string]*Value)(nil), "api.Execution.CustomPropertiesEntry")
	proto.RegisterMapType((map[string]string)(nil), "api.Execution.LabelsEntry")
	proto.RegisterMapType((map[string]*Value)(nil), "api.Execution.PropertiesEntry")
	proto.RegisterType((*StringType)(nil), "api.StringType")
	proto.RegisterType((*IntType)(nil), "api.IntType")
	proto.RegisterType((*DoubleType)(nil), "api.DoubleType")
	proto.RegisterType((*StringValue)(nil), "api.StringValue")
	proto.RegisterType((*IntValue)(nil), "api.IntValue")
	proto.RegisterType((*DoubleValue)(nil), "api.DoubleValue")
	proto.RegisterType((*Type)(nil), "api.Type")
	proto.RegisterType((*Value)(nil), "api.Value")
	proto.RegisterType((*Resource)(nil), "api.Resource")
	proto.RegisterType((*GetResourceRequest)(nil), "api.GetResourceRequest")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xdb, 0x6e, 0xe3, 0x44,
	0x18, 0xde, 0xc4, 0x39, 0xf9, 0x77, 0xda, 0x94, 0x01, 0xb4, 0x59, 0x6b, 0xd9, 0x44, 0x2e, 0x87,
	0x68, 0xb5, 0x72, 0x54, 0xa3, 0x4a, 0x2c, 0x08, 0x24, 0x0a, 0x2b, 0xb6, 0xd2, 0xb2, 0xbb, 0x78,
	0xab, 0x22, 0x71, 0x13, 0x4d, 0xec, 0x69, 0x18, 0x35, 0xb1, 0x8d, 0x3d, 0x4e, 0x89, 0x10, 0x37,
	0xbc, 0x02, 0x2f, 0x80, 0xb8, 0x43, 0xe2, 0x6d, 0x78, 0x05, 0x24, 0xae, 0x79, 0x03, 0x34, 0x33,
	0x3e, 0x8c, 0x43, 0xda, 0x06, 0x50, 0x25, 0xb4, 0x37, 0x55, 0xfc, 0xcf, 0xf7, 0xfd, 0x87, 0x6f,
	0xbe, 0xbf, 0x36, 0xec, 0xe0, 0x88, 0x8e, 0x71, 0x44, 0xed, 0x28, 0x0e, 0x59, 0x88, 0x34, 0x1c,
	0x51, 0xf3, 0xee, 0x2c, 0x0c, 0x67, 0x73, 0x32, 0x16, 0x47, 0x41, 0x10, 0x32, 0xcc, 0x68, 0x18,
	0x24, 0x12, 0x62, 0x0e, 0xb2, 0x53, 0xf1, 0x34, 0x4d, 0xcf, 0xc6, 0x8c, 0x2e, 0x48, 0xc2, 0xf0,
	0x22, 0x92, 0x00, 0x6b, 0x00, 0xfa, 0x97, 0x61, 0x7c, 0x9e, 0x44, 0xd8, 0x23, 0x08, 0x41, 0x23,
	0xc0, 0x0b, 0xd2, 0xaf, 0x0d, 0x6b, 0x23, 0xdd, 0x15, 0xbf, 0x39, 0xe0, 0x29, 0x5e, 0x90, 0xcb,
	0x01, 0x3f, 0xd5, 0xa1, 0xfb, 0x71, 0xcc, 0xe8, 0x19, 0xf6, 0xd8, 0xc9, 0x2a, 0x22, 0x68, 0x17,
	0xea, 0xd4, 0x17, 0x10, 0xcd, 0xad, 0x53, 0xbf, 0x20, 0xd5, 0x4b, 0x12, 0x7a, 0x00, 0x7a, 0x90,
	0x67, 0xed, 0x6b, 0xc3, 0xda, 0xc8, 0x70, 0x76, 0x6d, 0x3e, 0x59, 0x51, 0xcb, 0x2d, 0x01, 0xe8,
	0x29, 0xf4, 0xd8, 0x2a, 0x22, 0x93, 0x28, 0x0e, 0x23, 0x12, 0x33, 0x4a, 0x92, 0x7e, 0x63, 0xa8,
	0x8d, 0x0c, 0xe7, 0x2d, 0xc1, 0x51, 0xab, 0xdb, 0xfc, 0xcf, 0xf3, 0x02, 0xf7, 0x28, 0x60, 0xf1,
	0xca, 0xdd, 0x65, 0x95, 0x20, 0x1a, 0x82, 0xe1, 0x93, 0xc4, 0x8b, 0x69, 0xc4, 0xb5, 0xea, 0x37,
	0x45, 0x63, 0x6a, 0xc8, 0x7c, 0x02, 0xaf, 0x6e, 0x48, 0x84, 0xf6, 0x40, 0x3b, 0x27, 0xab, 0x6c,
	0x7c, 0xfe, 0x13, 0x0d, 0xa0, 0xb9, 0xc4, 0xf3, 0x54, 0x4e, 0x67, 0x38, 0xba, 0x68, 0x88, 0x53,
	0x5d, 0x19, 0x7f, 0xbf, 0xfe, 0x5e, 0xcd, 0xfa, 0xb9, 0x0e, 0x3b, 0x8f, 0xbe, 0x25, 0x5e, 0xca,
	0x73, 0xdf, 0x90, 0x46, 0xcf, 0x2e, 0xd3, 0xe8, 0x6d, 0xc1, 0xa9, 0x94, 0xff, 0x5f, 0x8a, 0xf4,
	0x67, 0x03, 0x3a, 0xf9, 0x4d, 0x6e, 0xab, 0xcf, 0x45, 0x6e, 0xdd, 0x8a, 0x3e, 0x85, 0xa1, 0xdd,
	0x12, 0x80, 0x3e, 0x04, 0x50, 0xa4, 0x69, 0x0a, 0x69, 0xde, 0xa8, 0xd8, 0xc7, 0x5e, 0x57, 0x44,
	0x21, 0xa0, 0xe7, 0xf0, 0x8a, 0x97, 0x26, 0x2c, 0x5c, 0xa8, 0x02, 0xb7, 0x44, 0x96, 0xfd, 0x6a,
	0x96, 0x4f, 0x04, 0x6c, 0x3d, 0xd7, 0x9e, 0xb7, 0x16, 0x46, 0x07, 0xd0, 0x9a, 0xe3, 0x29, 0x99,
	0x27, 0xfd, 0xb6, 0x48, 0x73, 0xa7, 0x9a, 0xe6, 0x89, 0x38, 0x93, 0xe4, 0x0c, 0x88, 0x3e, 0x00,
	0xc3, 0x8b, 0x09, 0x66, 0x64, 0xc2, 0xd7, 0xb8, 0xdf, 0x11, 0x33, 0x9b, 0xb6, 0xdc, 0x71, 0x3b,
	0xdf, 0x71, 0xfb, 0x24, 0xdf, 0x71, 0x17, 0x24, 0x9c, 0x07, 0x38, 0x39, 0x8d, 0xfc, 0x82, 0xac,
	0x5f, 0x4f, 0x96, 0x70, 0x1e, 0x30, 0x8f, 0xa1, 0x77, 0xfd, 0x35, 0x0f, 0xab, 0xd7, 0x0c, 0x62,
	0xa0, 0x53, 0x1e, 0x51, 0xee, 0xd9, 0x7c, 0x06, 0xaf, 0x6f, 0x94, 0xe8, 0x5f, 0x27, 0x7c, 0x08,
	0x86, 0x22, 0xd6, 0x86, 0x34, 0xaf, 0xa9, 0x69, 0x74, 0xd5, 0x73, 0x7f, 0x34, 0x41, 0x2f, 0x36,
	0xe3, 0x06, 0x4c, 0xf7, 0x51, 0xc5, 0x74, 0x72, 0x1f, 0xef, 0x55, 0xf7, 0xf1, 0x4a, 0xd7, 0x7d,
	0xb1, 0xc9, 0x75, 0xd2, 0xbb, 0x6f, 0xae, 0xa5, 0xd9, 0xd6, 0x76, 0x4e, 0x61, 0x3b, 0xe9, 0x5e,
	0x73, 0x2d, 0xcf, 0x16, 0xbe, 0x6b, 0xff, 0x17, 0xdf, 0x75, 0xfe, 0x89, 0xef, 0xd0, 0x43, 0x80,
	0x84, 0xe1, 0x98, 0x6d, 0xeb, 0x59, 0x5d, 0xa0, 0x05, 0xf5, 0x10, 0x3a, 0x24, 0xf0, 0x25, 0x11,
	0xae, 0x25, 0xb6, 0x49, 0xe0, 0xbf, 0xcc, 0x4e, 0xbf, 0x0f, 0xf0, 0x82, 0xc5, 0x34, 0x98, 0x89,
	0xd7, 0xcf, 0x5d, 0xd0, 0x97, 0x78, 0x4e, 0x7d, 0xcc, 0xc2, 0x38, 0xe3, 0x97, 0x01, 0xeb, 0x1d,
	0x68, 0x1f, 0x07, 0x6c, 0x0b, 0xe0, 0x7d, 0x80, 0x4f, 0xc3, 0x74, 0x3a, 0x27, 0x5b, 0x60, 0xf7,
	0xc1, 0x90, 0x0d, 0x88, 0xa9, 0xca, 0x4e, 0x6b, 0x4a, 0xa7, 0xd6, 0x10, 0x3a, 0xc7, 0x01, 0xdb,
	0x80, 0xd0, 0x72, 0xc4, 0x3e, 0x18, 0xb2, 0xe4, 0x55, 0xa0, 0x5f, 0x6a, 0xd0, 0x10, 0x2d, 0x39,
	0x60, 0x24, 0xa2, 0xe8, 0x84, 0xbf, 0xdc, 0x04, 0xc8, 0x70, 0x7a, 0x42, 0xdc, 0x52, 0x8d, 0xc7,
	0xb7, 0x5c, 0x48, 0x4a, 0x6d, 0x0e, 0xa0, 0x4b, 0x03, 0x46, 0x66, 0x24, 0x96, 0x24, 0x79, 0x23,
	0x5d, 0x41, 0xca, 0x64, 0x79, 0x7c, 0xcb, 0x35, 0x32, 0x4c, 0x5e, 0xc6, 0x17, 0x4d, 0x49, 0x86,
	0xa6, 0x94, 0x29, 0xf5, 0xe1, 0x65, 0xfc, 0xe2, 0xe9, 0xa8, 0x05, 0x0d, 0x0e, 0xb6, 0x7e, 0xad,
	0x41, 0x53, 0xce, 0x72, 0x08, 0xdd, 0xac, 0xd9, 0x72, 0x24, 0xc3, 0xd9, 0x53, 0xba, 0x15, 0x38,
	0x5e, 0x3c, 0x51, 0x94, 0x7c, 0x00, 0x3a, 0x0d, 0xd8, 0x44, 0xb5, 0xcf, 0x4e, 0xde, 0x6c, 0x4e,
	0xe8, 0xd0, 0x5c, 0xd5, 0x43, 0xe8, 0x66, 0xad, 0x4a, 0x82, 0xa6, 0x14, 0x51, 0x84, 0xe5, 0x45,
	0xfc, 0xf2, 0xf1, 0xa8, 0x9d, 0xe9, 0x6c, 0xdd, 0x83, 0x8e, 0x4b, 0x92, 0x30, 0x8d, 0x2f, 0xf9,
	0x1a, 0x1c, 0x01, 0xfa, 0x8c, 0xb0, 0x1c, 0xe2, 0x92, 0x6f, 0x52, 0x92, 0xb0, 0x4d, 0x48, 0x87,
	0x42, 0xef, 0x73, 0xc2, 0xb0, 0x8f, 0x19, 0x7e, 0x41, 0xe2, 0x25, 0xf5, 0x08, 0x3a, 0x05, 0x43,
	0x21, 0xa3, 0xdb, 0xa2, 0xab, 0xbf, 0xa7, 0x33, 0xe5, 0x7c, 0x79, 0xd4, 0x1a, 0xfc, 0xf0, 0xdb,
	0xef, 0x3f, 0xd6, 0xef, 0xa0, 0xdb, 0xe2, 0x23, 0x78, 0x79, 0x30, 0x8e, 0xb3, 0x93, 0xf1, 0x77,
	0xbc, 0xd2, 0xf7, 0x47, 0xd6, 0x57, 0xc3, 0x19, 0x65, 0x5f, 0xa7, 0x53, 0xdb, 0x0b, 0x17, 0xe3,
	0xf3, 0x74, 0x4a, 0xce, 0xe6, 0xe1, 0xc5, 0x78, 0x91, 0x95, 0xe7, 0xb4, 0x69, 0x4b, 0xfc, 0x53,
	0x78, 0xf7, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0xd0, 0xe2, 0x9c, 0x64, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataServiceClient interface {
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error)
}

type metadataServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetadataServiceClient(cc *grpc.ClientConn) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/api.MetadataService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
type MetadataServiceServer interface {
	GetResource(context.Context, *GetResourceRequest) (*Resource, error)
}

func RegisterMetadataServiceServer(s *grpc.Server, srv MetadataServiceServer) {
	s.RegisterService(&_MetadataService_serviceDesc, srv)
}

func _MetadataService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MetadataService/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _MetadataService_GetResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
