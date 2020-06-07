// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cbws/projects/v1alpha1/project.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Project struct {
	// The name of the project in the format: projects/test-project
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// @OutputOnly The id of the project that owns the service account.
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	// @OutputOnly The UUID of the project
	UniqueId             string   `protobuf:"bytes,4,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_11308aab321ad7d3, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Project) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *Project) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "cbws.projects.v1alpha1.Project")
}

func init() {
	proto.RegisterFile("cbws/projects/v1alpha1/project.proto", fileDescriptor_11308aab321ad7d3)
}

var fileDescriptor_11308aab321ad7d3 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x4e, 0x2a, 0x2f,
	0xd6, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8,
	0x48, 0x34, 0x84, 0x89, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x81, 0x54, 0xe9, 0xc1,
	0x54, 0xe9, 0xc1, 0x54, 0x49, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64,
	0xea, 0x17, 0xa5, 0x16, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0x42, 0xb4, 0x28, 0x6d, 0x65, 0xe4, 0x62,
	0x0f, 0x80, 0x68, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x14, 0xb9, 0x78, 0x52, 0x32, 0x8b, 0x0b, 0x72, 0x12, 0x2b, 0xe3,
	0xc1, 0x72, 0x4c, 0x60, 0x39, 0x6e, 0xa8, 0x98, 0x1f, 0x48, 0x89, 0x12, 0x17, 0x4f, 0x7e, 0x51,
	0x7a, 0x62, 0x5e, 0x66, 0x55, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x33, 0x58, 0x09, 0x8a, 0x98,
	0x90, 0x34, 0x17, 0x67, 0x69, 0x5e, 0x66, 0x61, 0x69, 0x6a, 0x7c, 0x66, 0x8a, 0x04, 0x0b, 0x58,
	0x01, 0x07, 0x44, 0xc0, 0x33, 0xc5, 0xca, 0xe8, 0x95, 0xa3, 0x3e, 0x97, 0x24, 0xdc, 0xd9, 0x60,
	0x4f, 0x54, 0x54, 0x56, 0xe9, 0xc3, 0xdd, 0x05, 0xf7, 0x77, 0x35, 0x94, 0x55, 0xeb, 0x64, 0x1a,
	0x65, 0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x0e, 0x9d, 0xf4,
	0x7c, 0x5d, 0x10, 0xad, 0x9b, 0x5e, 0x54, 0x90, 0xac, 0x8f, 0x3d, 0xbc, 0x92, 0xd8, 0xc0, 0xbe,
	0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x58, 0xde, 0x70, 0x50, 0x01, 0x00, 0x00,
}
