// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/restaurant/restauran.proto

package restaurant

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

type RestaurantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantIds []int32 `protobuf:"varint,1,rep,packed,name=restaurantIds,proto3" json:"restaurantIds,omitempty"`
}

func (x *RestaurantRequest) Reset() {
	*x = RestaurantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurant_restauran_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantRequest) ProtoMessage() {}

func (x *RestaurantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_restauran_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantRequest.ProtoReflect.Descriptor instead.
func (*RestaurantRequest) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_restauran_proto_rawDescGZIP(), []int{0}
}

func (x *RestaurantRequest) GetRestaurantIds() []int32 {
	if x != nil {
		return x.RestaurantIds
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Role      string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurant_restauran_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_restauran_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_restauran_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url       string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Width     int32  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height    int32  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	CloudName string `protobuf:"bytes,5,opt,name=cloudName,proto3" json:"cloudName,omitempty"`
	Extension string `protobuf:"bytes,6,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurant_restauran_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_restauran_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_restauran_proto_rawDescGZIP(), []int{2}
}

func (x *Image) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Image) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Image) GetCloudName() string {
	if x != nil {
		return x.CloudName
	}
	return ""
}

func (x *Image) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Addr      string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Logo      string   `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	Cover     []string `protobuf:"bytes,5,rep,name=cover,proto3" json:"cover,omitempty"`
	LikeCount int32    `protobuf:"varint,6,opt,name=likeCount,proto3" json:"likeCount,omitempty"`
	Owner     *User    `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurant_restauran_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_restauran_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_restauran_proto_rawDescGZIP(), []int{3}
}

func (x *Restaurant) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Restaurant) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Restaurant) GetCover() []string {
	if x != nil {
		return x.Cover
	}
	return nil
}

func (x *Restaurant) GetLikeCount() int32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *Restaurant) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

type RestaurantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurants []*Restaurant `protobuf:"bytes,1,rep,name=restaurants,proto3" json:"restaurants,omitempty"`
}

func (x *RestaurantResponse) Reset() {
	*x = RestaurantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurant_restauran_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantResponse) ProtoMessage() {}

func (x *RestaurantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_restauran_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantResponse.ProtoReflect.Descriptor instead.
func (*RestaurantResponse) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_restauran_proto_rawDescGZIP(), []int{4}
}

func (x *RestaurantResponse) GetRestaurants() []*Restaurant {
	if x != nil {
		return x.Restaurants
	}
	return nil
}

var File_proto_restaurant_restauran_proto protoreflect.FileDescriptor

var file_proto_restaurant_restauran_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0x64,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb9, 0x01, 0x0a, 0x0a, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x32, 0x74, 0x0a, 0x11, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_restaurant_restauran_proto_rawDescOnce sync.Once
	file_proto_restaurant_restauran_proto_rawDescData = file_proto_restaurant_restauran_proto_rawDesc
)

func file_proto_restaurant_restauran_proto_rawDescGZIP() []byte {
	file_proto_restaurant_restauran_proto_rawDescOnce.Do(func() {
		file_proto_restaurant_restauran_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_restaurant_restauran_proto_rawDescData)
	})
	return file_proto_restaurant_restauran_proto_rawDescData
}

var file_proto_restaurant_restauran_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_restaurant_restauran_proto_goTypes = []interface{}{
	(*RestaurantRequest)(nil),  // 0: restaurantproto.RestaurantRequest
	(*User)(nil),               // 1: restaurantproto.User
	(*Image)(nil),              // 2: restaurantproto.Image
	(*Restaurant)(nil),         // 3: restaurantproto.Restaurant
	(*RestaurantResponse)(nil), // 4: restaurantproto.RestaurantResponse
}
var file_proto_restaurant_restauran_proto_depIdxs = []int32{
	1, // 0: restaurantproto.Restaurant.owner:type_name -> restaurantproto.User
	3, // 1: restaurantproto.RestaurantResponse.restaurants:type_name -> restaurantproto.Restaurant
	0, // 2: restaurantproto.RestaurantService.GetRestaurantByIds:input_type -> restaurantproto.RestaurantRequest
	4, // 3: restaurantproto.RestaurantService.GetRestaurantByIds:output_type -> restaurantproto.RestaurantResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_restaurant_restauran_proto_init() }
func file_proto_restaurant_restauran_proto_init() {
	if File_proto_restaurant_restauran_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_restaurant_restauran_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantRequest); i {
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
		file_proto_restaurant_restauran_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_restaurant_restauran_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_proto_restaurant_restauran_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_proto_restaurant_restauran_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantResponse); i {
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
			RawDescriptor: file_proto_restaurant_restauran_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_restaurant_restauran_proto_goTypes,
		DependencyIndexes: file_proto_restaurant_restauran_proto_depIdxs,
		MessageInfos:      file_proto_restaurant_restauran_proto_msgTypes,
	}.Build()
	File_proto_restaurant_restauran_proto = out.File
	file_proto_restaurant_restauran_proto_rawDesc = nil
	file_proto_restaurant_restauran_proto_goTypes = nil
	file_proto_restaurant_restauran_proto_depIdxs = nil
}