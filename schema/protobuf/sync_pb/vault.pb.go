// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Trusted vault protos to communicate with backend written in proto3 to avoid
// subtle differences between enum fields.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: vault.proto

package sync_pb

import (
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

type SecurityDomainMember_MemberType int32

const (
	SecurityDomainMember_MEMBER_TYPE_UNSPECIFIED     SecurityDomainMember_MemberType = 0
	SecurityDomainMember_MEMBER_TYPE_PHYSICAL_DEVICE SecurityDomainMember_MemberType = 1
)

// Enum value maps for SecurityDomainMember_MemberType.
var (
	SecurityDomainMember_MemberType_name = map[int32]string{
		0: "MEMBER_TYPE_UNSPECIFIED",
		1: "MEMBER_TYPE_PHYSICAL_DEVICE",
	}
	SecurityDomainMember_MemberType_value = map[string]int32{
		"MEMBER_TYPE_UNSPECIFIED":     0,
		"MEMBER_TYPE_PHYSICAL_DEVICE": 1,
	}
)

func (x SecurityDomainMember_MemberType) Enum() *SecurityDomainMember_MemberType {
	p := new(SecurityDomainMember_MemberType)
	*p = x
	return p
}

func (x SecurityDomainMember_MemberType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityDomainMember_MemberType) Descriptor() protoreflect.EnumDescriptor {
	return file_vault_proto_enumTypes[0].Descriptor()
}

func (SecurityDomainMember_MemberType) Type() protoreflect.EnumType {
	return &file_vault_proto_enumTypes[0]
}

func (x SecurityDomainMember_MemberType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityDomainMember_MemberType.Descriptor instead.
func (SecurityDomainMember_MemberType) EnumDescriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{4, 0}
}

type SharedMemberKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch       int32  `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	WrappedKey  []byte `protobuf:"bytes,2,opt,name=wrapped_key,json=wrappedKey,proto3" json:"wrapped_key,omitempty"`
	MemberProof []byte `protobuf:"bytes,3,opt,name=member_proof,json=memberProof,proto3" json:"member_proof,omitempty"`
}

func (x *SharedMemberKey) Reset() {
	*x = SharedMemberKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedMemberKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedMemberKey) ProtoMessage() {}

func (x *SharedMemberKey) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedMemberKey.ProtoReflect.Descriptor instead.
func (*SharedMemberKey) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{0}
}

func (x *SharedMemberKey) GetEpoch() int32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *SharedMemberKey) GetWrappedKey() []byte {
	if x != nil {
		return x.WrappedKey
	}
	return nil
}

func (x *SharedMemberKey) GetMemberProof() []byte {
	if x != nil {
		return x.MemberProof
	}
	return nil
}

type RotationProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewEpoch      int32  `protobuf:"varint,1,opt,name=new_epoch,json=newEpoch,proto3" json:"new_epoch,omitempty"`
	RotationProof []byte `protobuf:"bytes,2,opt,name=rotation_proof,json=rotationProof,proto3" json:"rotation_proof,omitempty"`
}

func (x *RotationProof) Reset() {
	*x = RotationProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotationProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotationProof) ProtoMessage() {}

func (x *RotationProof) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotationProof.ProtoReflect.Descriptor instead.
func (*RotationProof) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{1}
}

func (x *RotationProof) GetNewEpoch() int32 {
	if x != nil {
		return x.NewEpoch
	}
	return 0
}

func (x *RotationProof) GetRotationProof() []byte {
	if x != nil {
		return x.RotationProof
	}
	return nil
}

type SecurityDomainDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncDetails *SecurityDomainDetails_SyncDetails `protobuf:"bytes,1,opt,name=sync_details,json=syncDetails,proto3" json:"sync_details,omitempty"`
}

func (x *SecurityDomainDetails) Reset() {
	*x = SecurityDomainDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomainDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomainDetails) ProtoMessage() {}

func (x *SecurityDomainDetails) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomainDetails.ProtoReflect.Descriptor instead.
func (*SecurityDomainDetails) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{2}
}

func (x *SecurityDomainDetails) GetSyncDetails() *SecurityDomainDetails_SyncDetails {
	if x != nil {
		return x.SyncDetails
	}
	return nil
}

type SecurityDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CurrentEpoch          int32                  `protobuf:"varint,2,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	SecurityDomainDetails *SecurityDomainDetails `protobuf:"bytes,3,opt,name=security_domain_details,json=securityDomainDetails,proto3" json:"security_domain_details,omitempty"`
}

func (x *SecurityDomain) Reset() {
	*x = SecurityDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomain) ProtoMessage() {}

func (x *SecurityDomain) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomain.ProtoReflect.Descriptor instead.
func (*SecurityDomain) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{3}
}

func (x *SecurityDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityDomain) GetCurrentEpoch() int32 {
	if x != nil {
		return x.CurrentEpoch
	}
	return 0
}

func (x *SecurityDomain) GetSecurityDomainDetails() *SecurityDomainDetails {
	if x != nil {
		return x.SecurityDomainDetails
	}
	return nil
}

type SecurityDomainMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey   []byte                                           `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Memberships []*SecurityDomainMember_SecurityDomainMembership `protobuf:"bytes,3,rep,name=memberships,proto3" json:"memberships,omitempty"`
	MemberType  SecurityDomainMember_MemberType                  `protobuf:"varint,4,opt,name=member_type,json=memberType,proto3,enum=sync_pb.SecurityDomainMember_MemberType" json:"member_type,omitempty"`
}

func (x *SecurityDomainMember) Reset() {
	*x = SecurityDomainMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomainMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomainMember) ProtoMessage() {}

func (x *SecurityDomainMember) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomainMember.ProtoReflect.Descriptor instead.
func (*SecurityDomainMember) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{4}
}

func (x *SecurityDomainMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityDomainMember) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SecurityDomainMember) GetMemberships() []*SecurityDomainMember_SecurityDomainMembership {
	if x != nil {
		return x.Memberships
	}
	return nil
}

func (x *SecurityDomainMember) GetMemberType() SecurityDomainMember_MemberType {
	if x != nil {
		return x.MemberType
	}
	return SecurityDomainMember_MEMBER_TYPE_UNSPECIFIED
}

type JoinSecurityDomainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain       *SecurityDomain       `protobuf:"bytes,1,opt,name=security_domain,json=securityDomain,proto3" json:"security_domain,omitempty"`
	SecurityDomainMember *SecurityDomainMember `protobuf:"bytes,2,opt,name=security_domain_member,json=securityDomainMember,proto3" json:"security_domain_member,omitempty"`
	SharedMemberKey      []*SharedMemberKey    `protobuf:"bytes,3,rep,name=shared_member_key,json=sharedMemberKey,proto3" json:"shared_member_key,omitempty"`
	MemberTypeHint       int32                 `protobuf:"varint,4,opt,name=member_type_hint,json=memberTypeHint,proto3" json:"member_type_hint,omitempty"`
}

func (x *JoinSecurityDomainsRequest) Reset() {
	*x = JoinSecurityDomainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSecurityDomainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSecurityDomainsRequest) ProtoMessage() {}

func (x *JoinSecurityDomainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSecurityDomainsRequest.ProtoReflect.Descriptor instead.
func (*JoinSecurityDomainsRequest) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{5}
}

func (x *JoinSecurityDomainsRequest) GetSecurityDomain() *SecurityDomain {
	if x != nil {
		return x.SecurityDomain
	}
	return nil
}

func (x *JoinSecurityDomainsRequest) GetSecurityDomainMember() *SecurityDomainMember {
	if x != nil {
		return x.SecurityDomainMember
	}
	return nil
}

func (x *JoinSecurityDomainsRequest) GetSharedMemberKey() []*SharedMemberKey {
	if x != nil {
		return x.SharedMemberKey
	}
	return nil
}

func (x *JoinSecurityDomainsRequest) GetMemberTypeHint() int32 {
	if x != nil {
		return x.MemberTypeHint
	}
	return 0
}

type JoinSecurityDomainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain *SecurityDomain `protobuf:"bytes,1,opt,name=security_domain,json=securityDomain,proto3" json:"security_domain,omitempty"`
}

func (x *JoinSecurityDomainsResponse) Reset() {
	*x = JoinSecurityDomainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSecurityDomainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSecurityDomainsResponse) ProtoMessage() {}

func (x *JoinSecurityDomainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSecurityDomainsResponse.ProtoReflect.Descriptor instead.
func (*JoinSecurityDomainsResponse) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{6}
}

func (x *JoinSecurityDomainsResponse) GetSecurityDomain() *SecurityDomain {
	if x != nil {
		return x.SecurityDomain
	}
	return nil
}

type JoinSecurityDomainsErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlreadyExistsResponse *JoinSecurityDomainsResponse `protobuf:"bytes,1,opt,name=already_exists_response,json=alreadyExistsResponse,proto3" json:"already_exists_response,omitempty"`
}

func (x *JoinSecurityDomainsErrorDetail) Reset() {
	*x = JoinSecurityDomainsErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSecurityDomainsErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSecurityDomainsErrorDetail) ProtoMessage() {}

func (x *JoinSecurityDomainsErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSecurityDomainsErrorDetail.ProtoReflect.Descriptor instead.
func (*JoinSecurityDomainsErrorDetail) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{7}
}

func (x *JoinSecurityDomainsErrorDetail) GetAlreadyExistsResponse() *JoinSecurityDomainsResponse {
	if x != nil {
		return x.AlreadyExistsResponse
	}
	return nil
}

// TODO(crbug.com/1234719): figure out how to link google.protobuf.Any and use
// it instead.
type Proto3Any struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Value   []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Proto3Any) Reset() {
	*x = Proto3Any{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto3Any) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto3Any) ProtoMessage() {}

func (x *Proto3Any) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto3Any.ProtoReflect.Descriptor instead.
func (*Proto3Any) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{8}
}

func (x *Proto3Any) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Proto3Any) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Forked version of google.rpc.Status.
type RPCStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*Proto3Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *RPCStatus) Reset() {
	*x = RPCStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCStatus) ProtoMessage() {}

func (x *RPCStatus) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCStatus.ProtoReflect.Descriptor instead.
func (*RPCStatus) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{9}
}

func (x *RPCStatus) GetDetails() []*Proto3Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type SecurityDomainDetails_SyncDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DegradedRecoverability bool `protobuf:"varint,1,opt,name=degraded_recoverability,json=degradedRecoverability,proto3" json:"degraded_recoverability,omitempty"`
}

func (x *SecurityDomainDetails_SyncDetails) Reset() {
	*x = SecurityDomainDetails_SyncDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomainDetails_SyncDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomainDetails_SyncDetails) ProtoMessage() {}

func (x *SecurityDomainDetails_SyncDetails) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomainDetails_SyncDetails.ProtoReflect.Descriptor instead.
func (*SecurityDomainDetails_SyncDetails) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SecurityDomainDetails_SyncDetails) GetDegradedRecoverability() bool {
	if x != nil {
		return x.DegradedRecoverability
	}
	return false
}

type SecurityDomainMember_SecurityDomainMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain string             `protobuf:"bytes,1,opt,name=security_domain,json=securityDomain,proto3" json:"security_domain,omitempty"`
	Keys           []*SharedMemberKey `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	RotationProofs []*RotationProof   `protobuf:"bytes,4,rep,name=rotation_proofs,json=rotationProofs,proto3" json:"rotation_proofs,omitempty"`
}

func (x *SecurityDomainMember_SecurityDomainMembership) Reset() {
	*x = SecurityDomainMember_SecurityDomainMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDomainMember_SecurityDomainMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDomainMember_SecurityDomainMembership) ProtoMessage() {}

func (x *SecurityDomainMember_SecurityDomainMembership) ProtoReflect() protoreflect.Message {
	mi := &file_vault_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityDomainMember_SecurityDomainMembership.ProtoReflect.Descriptor instead.
func (*SecurityDomainMember_SecurityDomainMembership) Descriptor() ([]byte, []int) {
	return file_vault_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SecurityDomainMember_SecurityDomainMembership) GetSecurityDomain() string {
	if x != nil {
		return x.SecurityDomain
	}
	return ""
}

func (x *SecurityDomainMember_SecurityDomainMembership) GetKeys() []*SharedMemberKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *SecurityDomainMember_SecurityDomainMembership) GetRotationProofs() []*RotationProof {
	if x != nil {
		return x.RotationProofs
	}
	return nil
}

var File_vault_proto protoreflect.FileDescriptor

var file_vault_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x6b, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x22, 0x53, 0x0a, 0x0d, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x70, 0x6f, 0x63,
	0x68, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x72, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0xae, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x4d, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x1a, 0x46, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x37, 0x0a, 0x17, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x16, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x56, 0x0a, 0x17, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x15, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xef, 0x03,
	0x0a, 0x14, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x58, 0x0a, 0x0b, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x12, 0x49, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x1a, 0xb2,
	0x01, 0x0a, 0x18, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x52, 0x0e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x73, 0x22, 0x4a, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f,
	0x0a, 0x1b, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48,
	0x59, 0x53, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x01, 0x22,
	0xa3, 0x02, 0x0a, 0x1a, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x53, 0x0a, 0x16, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x14, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x69, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x1b, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x7e, 0x0a, 0x1e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x5c, 0x0a, 0x17, 0x61, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x15, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x41, 0x6e, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x52, 0x50, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2c, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x33, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42,
	0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vault_proto_rawDescOnce sync.Once
	file_vault_proto_rawDescData = file_vault_proto_rawDesc
)

func file_vault_proto_rawDescGZIP() []byte {
	file_vault_proto_rawDescOnce.Do(func() {
		file_vault_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_proto_rawDescData)
	})
	return file_vault_proto_rawDescData
}

var file_vault_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_vault_proto_goTypes = []interface{}{
	(SecurityDomainMember_MemberType)(0),                  // 0: sync_pb.SecurityDomainMember.MemberType
	(*SharedMemberKey)(nil),                               // 1: sync_pb.SharedMemberKey
	(*RotationProof)(nil),                                 // 2: sync_pb.RotationProof
	(*SecurityDomainDetails)(nil),                         // 3: sync_pb.SecurityDomainDetails
	(*SecurityDomain)(nil),                                // 4: sync_pb.SecurityDomain
	(*SecurityDomainMember)(nil),                          // 5: sync_pb.SecurityDomainMember
	(*JoinSecurityDomainsRequest)(nil),                    // 6: sync_pb.JoinSecurityDomainsRequest
	(*JoinSecurityDomainsResponse)(nil),                   // 7: sync_pb.JoinSecurityDomainsResponse
	(*JoinSecurityDomainsErrorDetail)(nil),                // 8: sync_pb.JoinSecurityDomainsErrorDetail
	(*Proto3Any)(nil),                                     // 9: sync_pb.Proto3Any
	(*RPCStatus)(nil),                                     // 10: sync_pb.RPCStatus
	(*SecurityDomainDetails_SyncDetails)(nil),             // 11: sync_pb.SecurityDomainDetails.SyncDetails
	(*SecurityDomainMember_SecurityDomainMembership)(nil), // 12: sync_pb.SecurityDomainMember.SecurityDomainMembership
}
var file_vault_proto_depIdxs = []int32{
	11, // 0: sync_pb.SecurityDomainDetails.sync_details:type_name -> sync_pb.SecurityDomainDetails.SyncDetails
	3,  // 1: sync_pb.SecurityDomain.security_domain_details:type_name -> sync_pb.SecurityDomainDetails
	12, // 2: sync_pb.SecurityDomainMember.memberships:type_name -> sync_pb.SecurityDomainMember.SecurityDomainMembership
	0,  // 3: sync_pb.SecurityDomainMember.member_type:type_name -> sync_pb.SecurityDomainMember.MemberType
	4,  // 4: sync_pb.JoinSecurityDomainsRequest.security_domain:type_name -> sync_pb.SecurityDomain
	5,  // 5: sync_pb.JoinSecurityDomainsRequest.security_domain_member:type_name -> sync_pb.SecurityDomainMember
	1,  // 6: sync_pb.JoinSecurityDomainsRequest.shared_member_key:type_name -> sync_pb.SharedMemberKey
	4,  // 7: sync_pb.JoinSecurityDomainsResponse.security_domain:type_name -> sync_pb.SecurityDomain
	7,  // 8: sync_pb.JoinSecurityDomainsErrorDetail.already_exists_response:type_name -> sync_pb.JoinSecurityDomainsResponse
	9,  // 9: sync_pb.RPCStatus.details:type_name -> sync_pb.Proto3Any
	1,  // 10: sync_pb.SecurityDomainMember.SecurityDomainMembership.keys:type_name -> sync_pb.SharedMemberKey
	2,  // 11: sync_pb.SecurityDomainMember.SecurityDomainMembership.rotation_proofs:type_name -> sync_pb.RotationProof
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_vault_proto_init() }
func file_vault_proto_init() {
	if File_vault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedMemberKey); i {
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
		file_vault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotationProof); i {
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
		file_vault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomainDetails); i {
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
		file_vault_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomain); i {
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
		file_vault_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomainMember); i {
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
		file_vault_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSecurityDomainsRequest); i {
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
		file_vault_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSecurityDomainsResponse); i {
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
		file_vault_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSecurityDomainsErrorDetail); i {
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
		file_vault_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto3Any); i {
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
		file_vault_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCStatus); i {
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
		file_vault_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomainDetails_SyncDetails); i {
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
		file_vault_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityDomainMember_SecurityDomainMembership); i {
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
			RawDescriptor: file_vault_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vault_proto_goTypes,
		DependencyIndexes: file_vault_proto_depIdxs,
		EnumInfos:         file_vault_proto_enumTypes,
		MessageInfos:      file_vault_proto_msgTypes,
	}.Build()
	File_vault_proto = out.File
	file_vault_proto_rawDesc = nil
	file_vault_proto_goTypes = nil
	file_vault_proto_depIdxs = nil
}
