// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: cbws/iam/policy/v1alpha1/policy.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Defines an Identity and Access Management (IAM) policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` is a collection of `bindings`. A `binding` binds one or more
// `members` to a single `role`. Members can be user accounts, service accounts,
// Google groups, and domains (such as G Suite). A `role` is a named list of
// permissions (defined by IAM or configured by users). A `binding` can
// optionally specify a `condition`, which is a logic expression that further
// constrains the role binding based on attributes about the request and/or
// target resource.
//
// **JSON Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/resourcemanager.organizationAdmin",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//             "serviceAccount:my-project-id@appspot.gserviceaccount.com"
//           ]
//         },
//         {
//           "role": "roles/resourcemanager.organizationViewer",
//           "members": ["user:eve@example.com"],
//           "condition": {
//             "title": "expirable access",
//             "description": "Does not grant access after Sep 2020",
//             "expression": "request.time <
//             timestamp('2020-10-01T00:00:00.000Z')",
//           }
//         }
//       ]
//     }
//
// **YAML Example**
//
//     bindings:
//     - members:
//       - user:mike@example.com
//       - group:admins@example.com
//       - domain:google.com
//       - serviceAccount:my-project-id@appspot.gserviceaccount.com
//       role: roles/resourcemanager.organizationAdmin
//     - members:
//       - user:eve@example.com
//       role: roles/resourcemanager.organizationViewer
//       condition:
//         title: expirable access
//         description: Does not grant access after Sep 2020
//         expression: request.time < timestamp('2020-10-01T00:00:00.000Z')
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam/docs).
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Associates a list of `members` to a `role`. Optionally may specify a
	// `condition` that determines when binding is in effect.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `protobuf:"bytes,4,rep,name=bindings,proto3" json:"bindings,omitempty"`
	// `etag` is used for optimistic concurrency control as a way to help
	// prevent simultaneous updates of a policy from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in the
	// read-modify-write cycle to perform policy updates in order to avoid race
	// conditions: An `etag` is returned in the response to `getIamPolicy`, and
	// systems are expected to put that etag in the request to `setIamPolicy` to
	// ensure that their change will be applied to the same version of the policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the existing
	// policy is overwritten. Due to blind-set semantics of an etag-less policy,
	// 'setIamPolicy' will not fail even if the incoming policy version does not
	// meet the requirements for modifying the stored policy.
	Etag []byte `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_cbws_iam_policy_v1alpha1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetBindings() []*Binding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

func (x *Policy) GetEtag() []byte {
	if x != nil {
		return x.Etag
	}
	return nil
}

// Associates `members` with a `role`.
type Binding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role        string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Specifies the identities requesting access for a Cloud Platform resource.
	// `members` can have the following values:
	//
	// * `anonymous`: A special identifier that represents anyone who is
	//    on the internet
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    authenticated with a Cloudbear account or a service account.
	//
	// * `user:{emailid}`: An email address that represents a specific Google
	//    account. For example, `alice@example.com` .
	//
	// * `serviceAccount:{emailid}`: An email address that represents a service
	//    account. For example, `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google group.
	//    For example, `admins@example.com`.
	//
	// * `principal:{id}`: Identifier of the specific principal, can be used if
	//    access to either the user, serviceAccount or group is not available and
	//    only the principal id is known. This member type is only available when
	//    setting a IAM policy and is not returned when getting a IAM policy.
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Binding) Reset() {
	*x = Binding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binding) ProtoMessage() {}

func (x *Binding) ProtoReflect() protoreflect.Message {
	mi := &file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binding.ProtoReflect.Descriptor instead.
func (*Binding) Descriptor() ([]byte, []int) {
	return file_cbws_iam_policy_v1alpha1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Binding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Binding) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Binding) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_cbws_iam_policy_v1alpha1_policy_proto protoreflect.FileDescriptor

var file_cbws_iam_policy_v1alpha1_policy_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x62, 0x77, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x62, 0x77, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x22, 0x5b, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x62, 0x77, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x59,
	0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0xa0, 0x01, 0x0a, 0x1c, 0x78, 0x79,
	0x7a, 0x2e, 0x63, 0x62, 0x77, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x62, 0x77, 0x73, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x62, 0x77, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x62, 0x77, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x1d, 0x43, 0x62, 0x77, 0x73, 0x5c, 0x49, 0x41, 0x4d, 0x5c, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x47, 0x52, 0x50, 0x43, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x43, 0x62, 0x77, 0x73, 0x5c, 0x49, 0x41, 0x4d, 0x5c, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x47, 0x52, 0x50, 0x43, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cbws_iam_policy_v1alpha1_policy_proto_rawDescOnce sync.Once
	file_cbws_iam_policy_v1alpha1_policy_proto_rawDescData = file_cbws_iam_policy_v1alpha1_policy_proto_rawDesc
)

func file_cbws_iam_policy_v1alpha1_policy_proto_rawDescGZIP() []byte {
	file_cbws_iam_policy_v1alpha1_policy_proto_rawDescOnce.Do(func() {
		file_cbws_iam_policy_v1alpha1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_cbws_iam_policy_v1alpha1_policy_proto_rawDescData)
	})
	return file_cbws_iam_policy_v1alpha1_policy_proto_rawDescData
}

var file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cbws_iam_policy_v1alpha1_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),  // 0: cbws.iam.policy.v1alpha1.Policy
	(*Binding)(nil), // 1: cbws.iam.policy.v1alpha1.Binding
}
var file_cbws_iam_policy_v1alpha1_policy_proto_depIdxs = []int32{
	1, // 0: cbws.iam.policy.v1alpha1.Policy.bindings:type_name -> cbws.iam.policy.v1alpha1.Binding
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cbws_iam_policy_v1alpha1_policy_proto_init() }
func file_cbws_iam_policy_v1alpha1_policy_proto_init() {
	if File_cbws_iam_policy_v1alpha1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binding); i {
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
			RawDescriptor: file_cbws_iam_policy_v1alpha1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cbws_iam_policy_v1alpha1_policy_proto_goTypes,
		DependencyIndexes: file_cbws_iam_policy_v1alpha1_policy_proto_depIdxs,
		MessageInfos:      file_cbws_iam_policy_v1alpha1_policy_proto_msgTypes,
	}.Build()
	File_cbws_iam_policy_v1alpha1_policy_proto = out.File
	file_cbws_iam_policy_v1alpha1_policy_proto_rawDesc = nil
	file_cbws_iam_policy_v1alpha1_policy_proto_goTypes = nil
	file_cbws_iam_policy_v1alpha1_policy_proto_depIdxs = nil
}
