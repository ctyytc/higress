// Copyright Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
// Modified by Higress Authors

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: api/extensions/v1alpha1/wasm.proto

// $schema: higress.extensions.v1alpha1.WasmPlugin
// $title: WasmPlugin
// $description: Extend the functionality provided by the envoy through WebAssembly filters.

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The phase in the filter chain where the plugin will be injected.
type PluginPhase int32

const (
	// Control plane decides where to insert the plugin. This will generally
	// be at the end of the filter chain, right before the Router.
	// Do not specify `PluginPhase` if the plugin is independent of others.
	PluginPhase_UNSPECIFIED_PHASE PluginPhase = 0
	// Insert plugin before Istio authentication filters.
	PluginPhase_AUTHN PluginPhase = 1
	// Insert plugin before Istio authorization filters and after Istio authentication filters.
	PluginPhase_AUTHZ PluginPhase = 2
	// Insert plugin before Istio stats filters and after Istio authorization filters.
	PluginPhase_STATS PluginPhase = 3
)

// Enum value maps for PluginPhase.
var (
	PluginPhase_name = map[int32]string{
		0: "UNSPECIFIED_PHASE",
		1: "AUTHN",
		2: "AUTHZ",
		3: "STATS",
	}
	PluginPhase_value = map[string]int32{
		"UNSPECIFIED_PHASE": 0,
		"AUTHN":             1,
		"AUTHZ":             2,
		"STATS":             3,
	}
)

func (x PluginPhase) Enum() *PluginPhase {
	p := new(PluginPhase)
	*p = x
	return p
}

func (x PluginPhase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginPhase) Descriptor() protoreflect.EnumDescriptor {
	return file_api_extensions_v1alpha1_wasm_proto_enumTypes[0].Descriptor()
}

func (PluginPhase) Type() protoreflect.EnumType {
	return &file_api_extensions_v1alpha1_wasm_proto_enumTypes[0]
}

func (x PluginPhase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginPhase.Descriptor instead.
func (PluginPhase) EnumDescriptor() ([]byte, []int) {
	return file_api_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{0}
}

// The pull behaviour to be applied when fetching an OCI image,
// mirroring K8s behaviour.
//
// <!--
// buf:lint:ignore ENUM_VALUE_UPPER_SNAKE_CASE
// -->
type PullPolicy int32

const (
	// Defaults to IfNotPresent, except for OCI images with tag `latest`, for which
	// the default will be Always.
	PullPolicy_UNSPECIFIED_POLICY PullPolicy = 0
	// If an existing version of the image has been pulled before, that
	// will be used. If no version of the image is present locally, we
	// will pull the latest version.
	PullPolicy_IfNotPresent PullPolicy = 1
	// We will always pull the latest version of an image when applying
	// this plugin.
	PullPolicy_Always PullPolicy = 2
)

// Enum value maps for PullPolicy.
var (
	PullPolicy_name = map[int32]string{
		0: "UNSPECIFIED_POLICY",
		1: "IfNotPresent",
		2: "Always",
	}
	PullPolicy_value = map[string]int32{
		"UNSPECIFIED_POLICY": 0,
		"IfNotPresent":       1,
		"Always":             2,
	}
)

func (x PullPolicy) Enum() *PullPolicy {
	p := new(PullPolicy)
	*p = x
	return p
}

func (x PullPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PullPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_api_extensions_v1alpha1_wasm_proto_enumTypes[1].Descriptor()
}

func (PullPolicy) Type() protoreflect.EnumType {
	return &file_api_extensions_v1alpha1_wasm_proto_enumTypes[1]
}

func (x PullPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PullPolicy.Descriptor instead.
func (PullPolicy) EnumDescriptor() ([]byte, []int) {
	return file_api_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{1}
}

// <!-- crd generation tags
// +cue-gen:WasmPlugin:groupName:extensions.higress.io
// +cue-gen:WasmPlugin:version:v1alpha1
// +cue-gen:WasmPlugin:storageVersion
// +cue-gen:WasmPlugin:annotations:helm.sh/resource-policy=keep
// +cue-gen:WasmPlugin:subresource:status
// +cue-gen:WasmPlugin:scope:Namespaced
// +cue-gen:WasmPlugin:resource:categories=higress-io,extensions-higress-io
// +cue-gen:WasmPlugin:preserveUnknownFields:pluginConfig,defaultConfig,matchRules.[].config
// +cue-gen:WasmPlugin:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=extensions.higress.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WasmPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL of a Wasm module or OCI container. If no scheme is present,
	// defaults to `oci://`, referencing an OCI image. Other valid schemes
	// are `file://` for referencing .wasm module files present locally
	// within the proxy container, and `http[s]://` for .wasm module files
	// hosted remotely.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// SHA256 checksum that will be used to verify Wasm module or OCI container.
	// If the `url` field already references a SHA256 (using the `@sha256:`
	// notation), it must match the value of this field. If an OCI image is
	// referenced by tag and this field is set, its checksum will be verified
	// against the contents of this field after pulling.
	Sha256 string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// The pull behaviour to be applied when fetching an OCI image. Only
	// relevant when images are referenced by tag instead of SHA. Defaults
	// to IfNotPresent, except when an OCI image is referenced in the `url`
	// and the `latest` tag is used, in which case `Always` is the default,
	// mirroring K8s behaviour.
	// Setting is ignored if `url` field is referencing a Wasm module directly
	// using `file://` or `http[s]://`
	ImagePullPolicy PullPolicy `protobuf:"varint,4,opt,name=image_pull_policy,json=imagePullPolicy,proto3,enum=higress.extensions.v1alpha1.PullPolicy" json:"image_pull_policy,omitempty"`
	// Credentials to use for OCI image pulling.
	// Name of a K8s Secret in the same namespace as the `WasmPlugin` that
	// contains a docker pull secret which is to be used to authenticate
	// against the registry when pulling the image.
	ImagePullSecret string `protobuf:"bytes,5,opt,name=image_pull_secret,json=imagePullSecret,proto3" json:"image_pull_secret,omitempty"`
	// Public key that will be used to verify signatures of signed OCI images
	// or Wasm modules. Must be supplied in PEM format.
	VerificationKey string `protobuf:"bytes,6,opt,name=verification_key,json=verificationKey,proto3" json:"verification_key,omitempty"`
	// The configuration that will be passed on to the plugin.
	PluginConfig *structpb.Struct `protobuf:"bytes,7,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
	// The plugin name to be used in the Envoy configuration (used to be called
	// `rootID`). Some .wasm modules might require this value to select the Wasm
	// plugin to execute.
	PluginName string `protobuf:"bytes,8,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// Determines where in the filter chain this `WasmPlugin` is to be injected.
	Phase PluginPhase `protobuf:"varint,9,opt,name=phase,proto3,enum=higress.extensions.v1alpha1.PluginPhase" json:"phase,omitempty"`
	// Determines ordering of `WasmPlugins` in the same `phase`.
	// When multiple `WasmPlugins` are applied to the same workload in the
	// same `phase`, they will be applied by priority, in descending order.
	// If `priority` is not set, or two `WasmPlugins` exist with the same
	// value, the ordering will be deterministically derived from name and
	// namespace of the `WasmPlugins`. Defaults to `0`.
	Priority *wrapperspb.Int32Value `protobuf:"bytes,10,opt,name=priority,proto3" json:"priority,omitempty"`
	// Extended by Higress, the default configuration takes effect globally
	DefaultConfig *structpb.Struct `protobuf:"bytes,101,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	// Extended by Higress, matching rules take effect
	MatchRules []*MatchRule `protobuf:"bytes,102,rep,name=match_rules,json=matchRules,proto3" json:"match_rules,omitempty"`
	// disable the default config
	DefaultConfigDisable bool `protobuf:"varint,103,opt,name=default_config_disable,json=defaultConfigDisable,proto3" json:"default_config_disable,omitempty"`
}

func (x *WasmPlugin) Reset() {
	*x = WasmPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_extensions_v1alpha1_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmPlugin) ProtoMessage() {}

func (x *WasmPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_api_extensions_v1alpha1_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmPlugin.ProtoReflect.Descriptor instead.
func (*WasmPlugin) Descriptor() ([]byte, []int) {
	return file_api_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *WasmPlugin) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WasmPlugin) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *WasmPlugin) GetImagePullPolicy() PullPolicy {
	if x != nil {
		return x.ImagePullPolicy
	}
	return PullPolicy_UNSPECIFIED_POLICY
}

func (x *WasmPlugin) GetImagePullSecret() string {
	if x != nil {
		return x.ImagePullSecret
	}
	return ""
}

func (x *WasmPlugin) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

func (x *WasmPlugin) GetPluginConfig() *structpb.Struct {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

func (x *WasmPlugin) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *WasmPlugin) GetPhase() PluginPhase {
	if x != nil {
		return x.Phase
	}
	return PluginPhase_UNSPECIFIED_PHASE
}

func (x *WasmPlugin) GetPriority() *wrapperspb.Int32Value {
	if x != nil {
		return x.Priority
	}
	return nil
}

func (x *WasmPlugin) GetDefaultConfig() *structpb.Struct {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

func (x *WasmPlugin) GetMatchRules() []*MatchRule {
	if x != nil {
		return x.MatchRules
	}
	return nil
}

func (x *WasmPlugin) GetDefaultConfigDisable() bool {
	if x != nil {
		return x.DefaultConfigDisable
	}
	return false
}

// Extended by Higress
type MatchRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ingress       []string         `protobuf:"bytes,1,rep,name=ingress,proto3" json:"ingress,omitempty"`
	Domain        []string         `protobuf:"bytes,2,rep,name=domain,proto3" json:"domain,omitempty"`
	Config        *structpb.Struct `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	ConfigDisable bool             `protobuf:"varint,4,opt,name=config_disable,json=configDisable,proto3" json:"config_disable,omitempty"`
}

func (x *MatchRule) Reset() {
	*x = MatchRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_extensions_v1alpha1_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRule) ProtoMessage() {}

func (x *MatchRule) ProtoReflect() protoreflect.Message {
	mi := &file_api_extensions_v1alpha1_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRule.ProtoReflect.Descriptor instead.
func (*MatchRule) Descriptor() ([]byte, []int) {
	return file_api_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{1}
}

func (x *MatchRule) GetIngress() []string {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *MatchRule) GetDomain() []string {
	if x != nil {
		return x.Domain
	}
	return nil
}

func (x *MatchRule) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *MatchRule) GetConfigDisable() bool {
	if x != nil {
		return x.ConfigDisable
	}
	return false
}

var File_api_extensions_v1alpha1_wasm_proto protoreflect.FileDescriptor

var file_api_extensions_v1alpha1_wasm_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x68, 0x69, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf9, 0x04, 0x0a, 0x0a, 0x57, 0x61, 0x73, 0x6d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x53, 0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x68, 0x69, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2a, 0x0a,
	0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x68, 0x69, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x05, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x0e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0d, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x0b,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x66, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x68, 0x69, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x09,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x2a, 0x45, 0x0a, 0x0b, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x68, 0x61,
	0x73, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x54,
	0x48, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x54, 0x48, 0x5a, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x03, 0x2a, 0x42, 0x0a, 0x0a, 0x50, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x49, 0x66, 0x4e, 0x6f, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x10, 0x02, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x69,
	0x62, 0x61, 0x62, 0x61, 0x2f, 0x68, 0x69, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_extensions_v1alpha1_wasm_proto_rawDescOnce sync.Once
	file_api_extensions_v1alpha1_wasm_proto_rawDescData = file_api_extensions_v1alpha1_wasm_proto_rawDesc
)

func file_api_extensions_v1alpha1_wasm_proto_rawDescGZIP() []byte {
	file_api_extensions_v1alpha1_wasm_proto_rawDescOnce.Do(func() {
		file_api_extensions_v1alpha1_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_extensions_v1alpha1_wasm_proto_rawDescData)
	})
	return file_api_extensions_v1alpha1_wasm_proto_rawDescData
}

var file_api_extensions_v1alpha1_wasm_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_extensions_v1alpha1_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_extensions_v1alpha1_wasm_proto_goTypes = []interface{}{
	(PluginPhase)(0),              // 0: higress.extensions.v1alpha1.PluginPhase
	(PullPolicy)(0),               // 1: higress.extensions.v1alpha1.PullPolicy
	(*WasmPlugin)(nil),            // 2: higress.extensions.v1alpha1.WasmPlugin
	(*MatchRule)(nil),             // 3: higress.extensions.v1alpha1.MatchRule
	(*structpb.Struct)(nil),       // 4: google.protobuf.Struct
	(*wrapperspb.Int32Value)(nil), // 5: google.protobuf.Int32Value
}
var file_api_extensions_v1alpha1_wasm_proto_depIdxs = []int32{
	1, // 0: higress.extensions.v1alpha1.WasmPlugin.image_pull_policy:type_name -> higress.extensions.v1alpha1.PullPolicy
	4, // 1: higress.extensions.v1alpha1.WasmPlugin.plugin_config:type_name -> google.protobuf.Struct
	0, // 2: higress.extensions.v1alpha1.WasmPlugin.phase:type_name -> higress.extensions.v1alpha1.PluginPhase
	5, // 3: higress.extensions.v1alpha1.WasmPlugin.priority:type_name -> google.protobuf.Int32Value
	4, // 4: higress.extensions.v1alpha1.WasmPlugin.default_config:type_name -> google.protobuf.Struct
	3, // 5: higress.extensions.v1alpha1.WasmPlugin.match_rules:type_name -> higress.extensions.v1alpha1.MatchRule
	4, // 6: higress.extensions.v1alpha1.MatchRule.config:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_extensions_v1alpha1_wasm_proto_init() }
func file_api_extensions_v1alpha1_wasm_proto_init() {
	if File_api_extensions_v1alpha1_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_extensions_v1alpha1_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasmPlugin); i {
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
		file_api_extensions_v1alpha1_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRule); i {
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
			RawDescriptor: file_api_extensions_v1alpha1_wasm_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_extensions_v1alpha1_wasm_proto_goTypes,
		DependencyIndexes: file_api_extensions_v1alpha1_wasm_proto_depIdxs,
		EnumInfos:         file_api_extensions_v1alpha1_wasm_proto_enumTypes,
		MessageInfos:      file_api_extensions_v1alpha1_wasm_proto_msgTypes,
	}.Build()
	File_api_extensions_v1alpha1_wasm_proto = out.File
	file_api_extensions_v1alpha1_wasm_proto_rawDesc = nil
	file_api_extensions_v1alpha1_wasm_proto_goTypes = nil
	file_api_extensions_v1alpha1_wasm_proto_depIdxs = nil
}
