// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: product.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Product message definition
type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl      string                 `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Stock         int32                  `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Request messages
type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl      string                 `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Category      string                 `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Stock         int32                  `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      string                 `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	MinPrice      float64                `protobuf:"fixed64,2,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice      float64                `protobuf:"fixed64,3,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Page          int32                  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	mi := &file_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *ListProductsRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ListProductsRequest) GetMinPrice() float64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *ListProductsRequest) GetMaxPrice() float64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *ListProductsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl      string                 `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Stock         int32                  `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *UpdateProductRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response messages
type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

func (x *ListProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ListProductsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_product_proto protoreflect.FileDescriptor

const file_product_proto_rawDesc = "" +
	"\n" +
	"\rproduct.proto\x12\tinventory\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\xaa\x02\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x1b\n" +
	"\timage_url\x18\x05 \x01(\tR\bimageUrl\x12\x1a\n" +
	"\bcategory\x18\x06 \x01(\tR\bcategory\x12\x14\n" +
	"\x05stock\x18\a \x01(\x05R\x05stock\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\xb1\x01\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\x12\x1b\n" +
	"\timage_url\x18\x04 \x01(\tR\bimageUrl\x12\x1a\n" +
	"\bcategory\x18\x05 \x01(\tR\bcategory\x12\x14\n" +
	"\x05stock\x18\x06 \x01(\x05R\x05stock\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x95\x01\n" +
	"\x13ListProductsRequest\x12\x1a\n" +
	"\bcategory\x18\x01 \x01(\tR\bcategory\x12\x1b\n" +
	"\tmin_price\x18\x02 \x01(\x01R\bminPrice\x12\x1b\n" +
	"\tmax_price\x18\x03 \x01(\x01R\bmaxPrice\x12\x12\n" +
	"\x04page\x18\x04 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x05 \x01(\x05R\x05limit\"\xc1\x01\n" +
	"\x14UpdateProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x1b\n" +
	"\timage_url\x18\x05 \x01(\tR\bimageUrl\x12\x1a\n" +
	"\bcategory\x18\x06 \x01(\tR\bcategory\x12\x14\n" +
	"\x05stock\x18\a \x01(\x05R\x05stock\"&\n" +
	"\x14DeleteProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\\\n" +
	"\x14ListProductsResponse\x12.\n" +
	"\bproducts\x18\x01 \x03(\v2\x12.inventory.ProductR\bproducts\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total2\x81\x03\n" +
	"\x0eProductService\x12F\n" +
	"\rCreateProduct\x12\x1f.inventory.CreateProductRequest\x1a\x12.inventory.Product\"\x00\x12@\n" +
	"\n" +
	"GetProduct\x12\x1c.inventory.GetProductRequest\x1a\x12.inventory.Product\"\x00\x12Q\n" +
	"\fListProducts\x12\x1e.inventory.ListProductsRequest\x1a\x1f.inventory.ListProductsResponse\"\x00\x12F\n" +
	"\rUpdateProduct\x12\x1f.inventory.UpdateProductRequest\x1a\x12.inventory.Product\"\x00\x12J\n" +
	"\rDeleteProduct\x12\x1f.inventory.DeleteProductRequest\x1a\x16.google.protobuf.Empty\"\x00B\x10Z\x0eADV3/proto/genb\x06proto3"

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData []byte
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_proto_rawDesc), len(file_product_proto_rawDesc)))
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_proto_goTypes = []any{
	(*Product)(nil),               // 0: inventory.Product
	(*CreateProductRequest)(nil),  // 1: inventory.CreateProductRequest
	(*GetProductRequest)(nil),     // 2: inventory.GetProductRequest
	(*ListProductsRequest)(nil),   // 3: inventory.ListProductsRequest
	(*UpdateProductRequest)(nil),  // 4: inventory.UpdateProductRequest
	(*DeleteProductRequest)(nil),  // 5: inventory.DeleteProductRequest
	(*ListProductsResponse)(nil),  // 6: inventory.ListProductsResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_product_proto_depIdxs = []int32{
	7, // 0: inventory.Product.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: inventory.Product.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: inventory.ListProductsResponse.products:type_name -> inventory.Product
	1, // 3: inventory.ProductService.CreateProduct:input_type -> inventory.CreateProductRequest
	2, // 4: inventory.ProductService.GetProduct:input_type -> inventory.GetProductRequest
	3, // 5: inventory.ProductService.ListProducts:input_type -> inventory.ListProductsRequest
	4, // 6: inventory.ProductService.UpdateProduct:input_type -> inventory.UpdateProductRequest
	5, // 7: inventory.ProductService.DeleteProduct:input_type -> inventory.DeleteProductRequest
	0, // 8: inventory.ProductService.CreateProduct:output_type -> inventory.Product
	0, // 9: inventory.ProductService.GetProduct:output_type -> inventory.Product
	6, // 10: inventory.ProductService.ListProducts:output_type -> inventory.ListProductsResponse
	0, // 11: inventory.ProductService.UpdateProduct:output_type -> inventory.Product
	8, // 12: inventory.ProductService.DeleteProduct:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_proto_rawDesc), len(file_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
