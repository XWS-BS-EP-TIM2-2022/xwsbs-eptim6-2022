// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: job_offers_service/job_offers_service.proto

package job_offers_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{0}
}

type CreateJobOfferMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offer *CreateJobOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
}

func (x *CreateJobOfferMessage) Reset() {
	*x = CreateJobOfferMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobOfferMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOfferMessage) ProtoMessage() {}

func (x *CreateJobOfferMessage) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOfferMessage.ProtoReflect.Descriptor instead.
func (*CreateJobOfferMessage) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobOfferMessage) GetOffer() *CreateJobOffer {
	if x != nil {
		return x.Offer
	}
	return nil
}

type JobOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offer *JobOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
}

func (x *JobOfferResponse) Reset() {
	*x = JobOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOfferResponse) ProtoMessage() {}

func (x *JobOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOfferResponse.ProtoReflect.Descriptor instead.
func (*JobOfferResponse) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{2}
}

func (x *JobOfferResponse) GetOffer() *JobOffer {
	if x != nil {
		return x.Offer
	}
	return nil
}

type JobOffersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offers []*JobOffer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *JobOffersResponse) Reset() {
	*x = JobOffersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffersResponse) ProtoMessage() {}

func (x *JobOffersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffersResponse.ProtoReflect.Descriptor instead.
func (*JobOffersResponse) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{3}
}

func (x *JobOffersResponse) GetOffers() []*JobOffer {
	if x != nil {
		return x.Offers
	}
	return nil
}

type UsernameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UsernameMessage) Reset() {
	*x = UsernameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameMessage) ProtoMessage() {}

func (x *UsernameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameMessage.ProtoReflect.Descriptor instead.
func (*UsernameMessage) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{4}
}

func (x *UsernameMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateJobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position          string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Experience        string `protobuf:"bytes,3,opt,name=experience,proto3" json:"experience,omitempty"`
	CreatedOn         string `protobuf:"bytes,4,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	ValidUntil        string `protobuf:"bytes,5,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
	WorkScheduleTitle string `protobuf:"bytes,9,opt,name=workScheduleTitle,proto3" json:"workScheduleTitle,omitempty"`
	WorkScheduleHours string `protobuf:"bytes,10,opt,name=workScheduleHours,proto3" json:"workScheduleHours,omitempty"`
	JobOfferUrl       string `protobuf:"bytes,11,opt,name=jobOfferUrl,proto3" json:"jobOfferUrl,omitempty"`
}

func (x *CreateJobOffer) Reset() {
	*x = CreateJobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobOffer) ProtoMessage() {}

func (x *CreateJobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobOffer.ProtoReflect.Descriptor instead.
func (*CreateJobOffer) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateJobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *CreateJobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateJobOffer) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

func (x *CreateJobOffer) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *CreateJobOffer) GetValidUntil() string {
	if x != nil {
		return x.ValidUntil
	}
	return ""
}

func (x *CreateJobOffer) GetWorkScheduleTitle() string {
	if x != nil {
		return x.WorkScheduleTitle
	}
	return ""
}

func (x *CreateJobOffer) GetWorkScheduleHours() string {
	if x != nil {
		return x.WorkScheduleHours
	}
	return ""
}

func (x *CreateJobOffer) GetJobOfferUrl() string {
	if x != nil {
		return x.JobOfferUrl
	}
	return ""
}

type JobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username          string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Position          string `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Description       string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Experience        string `protobuf:"bytes,5,opt,name=experience,proto3" json:"experience,omitempty"`
	CreatedOn         string `protobuf:"bytes,6,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	ValidUntil        string `protobuf:"bytes,8,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
	WorkScheduleTitle string `protobuf:"bytes,9,opt,name=WorkScheduleTitle,proto3" json:"WorkScheduleTitle,omitempty"`
	WorkScheduleHours string `protobuf:"bytes,10,opt,name=WorkScheduleHours,proto3" json:"WorkScheduleHours,omitempty"`
	JobOfferUrl       string `protobuf:"bytes,11,opt,name=jobOfferUrl,proto3" json:"jobOfferUrl,omitempty"`
}

func (x *JobOffer) Reset() {
	*x = JobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offers_service_job_offers_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffer) ProtoMessage() {}

func (x *JobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_job_offers_service_job_offers_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffer.ProtoReflect.Descriptor instead.
func (*JobOffer) Descriptor() ([]byte, []int) {
	return file_job_offers_service_job_offers_service_proto_rawDescGZIP(), []int{6}
}

func (x *JobOffer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobOffer) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *JobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *JobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobOffer) GetExperience() string {
	if x != nil {
		return x.Experience
	}
	return ""
}

func (x *JobOffer) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *JobOffer) GetValidUntil() string {
	if x != nil {
		return x.ValidUntil
	}
	return ""
}

func (x *JobOffer) GetWorkScheduleTitle() string {
	if x != nil {
		return x.WorkScheduleTitle
	}
	return ""
}

func (x *JobOffer) GetWorkScheduleHours() string {
	if x != nil {
		return x.WorkScheduleHours
	}
	return ""
}

func (x *JobOffer) GetJobOfferUrl() string {
	if x != nil {
		return x.JobOfferUrl
	}
	return ""
}

var File_job_offers_service_job_offers_service_proto protoreflect.FileDescriptor

var file_job_offers_service_job_offers_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x51, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x22, 0x46, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x11, 0x4a, 0x6f,
	0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xaa, 0x02, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x4f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e,
	0x74, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x22, 0xd0, 0x02, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55,
	0x6e, 0x74, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x6f, 0x75,
	0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x32, 0xa2, 0x03, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x91,
	0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x25, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x12, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x24, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x3a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x57, 0x53, 0x2d, 0x42, 0x53, 0x2d, 0x45,
	0x50, 0x2d, 0x54, 0x49, 0x4d, 0x32, 0x2d, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x78, 0x77, 0x73, 0x62,
	0x73, 0x2d, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x36, 0x2d, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_offers_service_job_offers_service_proto_rawDescOnce sync.Once
	file_job_offers_service_job_offers_service_proto_rawDescData = file_job_offers_service_job_offers_service_proto_rawDesc
)

func file_job_offers_service_job_offers_service_proto_rawDescGZIP() []byte {
	file_job_offers_service_job_offers_service_proto_rawDescOnce.Do(func() {
		file_job_offers_service_job_offers_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_offers_service_job_offers_service_proto_rawDescData)
	})
	return file_job_offers_service_job_offers_service_proto_rawDescData
}

var file_job_offers_service_job_offers_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_job_offers_service_job_offers_service_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),          // 0: job_offers_service.EmptyRequest
	(*CreateJobOfferMessage)(nil), // 1: job_offers_service.CreateJobOfferMessage
	(*JobOfferResponse)(nil),      // 2: job_offers_service.JobOfferResponse
	(*JobOffersResponse)(nil),     // 3: job_offers_service.JobOffersResponse
	(*UsernameMessage)(nil),       // 4: job_offers_service.UsernameMessage
	(*CreateJobOffer)(nil),        // 5: job_offers_service.CreateJobOffer
	(*JobOffer)(nil),              // 6: job_offers_service.JobOffer
}
var file_job_offers_service_job_offers_service_proto_depIdxs = []int32{
	5, // 0: job_offers_service.CreateJobOfferMessage.offer:type_name -> job_offers_service.CreateJobOffer
	6, // 1: job_offers_service.JobOfferResponse.offer:type_name -> job_offers_service.JobOffer
	6, // 2: job_offers_service.JobOffersResponse.offers:type_name -> job_offers_service.JobOffer
	0, // 3: job_offers_service.JobOffersService.GetAllJobOffers:input_type -> job_offers_service.EmptyRequest
	4, // 4: job_offers_service.JobOffersService.GetAllJobOffersByUsername:input_type -> job_offers_service.UsernameMessage
	1, // 5: job_offers_service.JobOffersService.CreateNewJobOffer:input_type -> job_offers_service.CreateJobOfferMessage
	3, // 6: job_offers_service.JobOffersService.GetAllJobOffers:output_type -> job_offers_service.JobOffersResponse
	3, // 7: job_offers_service.JobOffersService.GetAllJobOffersByUsername:output_type -> job_offers_service.JobOffersResponse
	2, // 8: job_offers_service.JobOffersService.CreateNewJobOffer:output_type -> job_offers_service.JobOfferResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_job_offers_service_job_offers_service_proto_init() }
func file_job_offers_service_job_offers_service_proto_init() {
	if File_job_offers_service_job_offers_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_offers_service_job_offers_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_job_offers_service_job_offers_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobOfferMessage); i {
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
		file_job_offers_service_job_offers_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOfferResponse); i {
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
		file_job_offers_service_job_offers_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffersResponse); i {
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
		file_job_offers_service_job_offers_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameMessage); i {
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
		file_job_offers_service_job_offers_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobOffer); i {
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
		file_job_offers_service_job_offers_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffer); i {
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
			RawDescriptor: file_job_offers_service_job_offers_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_offers_service_job_offers_service_proto_goTypes,
		DependencyIndexes: file_job_offers_service_job_offers_service_proto_depIdxs,
		MessageInfos:      file_job_offers_service_job_offers_service_proto_msgTypes,
	}.Build()
	File_job_offers_service_job_offers_service_proto = out.File
	file_job_offers_service_job_offers_service_proto_rawDesc = nil
	file_job_offers_service_job_offers_service_proto_goTypes = nil
	file_job_offers_service_job_offers_service_proto_depIdxs = nil
}
