// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ConfigType represents the available field types from:
// https://github.com/spf13/viper#getting-values-from-viper
type ConfigType int32

const (
	ConfigType_BOOL         ConfigType = 0
	ConfigType_FLOAT64      ConfigType = 1
	ConfigType_INT          ConfigType = 2
	ConfigType_INT_SLICE    ConfigType = 3
	ConfigType_STRING       ConfigType = 4
	ConfigType_STRING_MAP   ConfigType = 5
	ConfigType_STRING_SLICE ConfigType = 6
	ConfigType_TIME         ConfigType = 7
	ConfigType_DURATION     ConfigType = 8
)

var ConfigType_name = map[int32]string{
	0: "BOOL",
	1: "FLOAT64",
	2: "INT",
	3: "INT_SLICE",
	4: "STRING",
	5: "STRING_MAP",
	6: "STRING_SLICE",
	7: "TIME",
	8: "DURATION",
}

var ConfigType_value = map[string]int32{
	"BOOL":         0,
	"FLOAT64":      1,
	"INT":          2,
	"INT_SLICE":    3,
	"STRING":       4,
	"STRING_MAP":   5,
	"STRING_SLICE": 6,
	"TIME":         7,
	"DURATION":     8,
}

func (x ConfigType) String() string {
	return proto.EnumName(ConfigType_name, int32(x))
}

func (ConfigType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

// VerifyRequest represents a message for an AuthenticationVerifier to verify.
type VerifyRequest struct {
	// username represents the username portion of the basic auth received from
	// the docker client.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// password represents the password portion of the basic auth received from
	// the docker client.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// host represents the host of the image that the docker client is attempting
	// to pull. e.g. docker-proxy.example.org
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// repository represents the repository of the image that the docker client is
	// attempting to pull. e.g. my-repository
	Repository string `protobuf:"bytes,4,opt,name=repository,proto3" json:"repository,omitempty"`
	// image represents the image that the docker client is attempting to pull.
	// e.g. my-image:latest or my-image@sha256:digest
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *VerifyRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *VerifyRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *VerifyRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *VerifyRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

// ProvideRequest represents a message for an AuthenticationProvider to return
// credentials for.
type ProvideRequest struct {
	// host represents the host of the image that the docker client is attempting
	// to pull. e.g. docker-proxy.example.org
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// repository represents the repository of the image that the docker client is
	// attempting to pull. e.g. my-repository
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// image represents the image that the docker client is attempting to pull.
	// e.g. my-image:latest or my-image@sha256:digest
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvideRequest) Reset()         { *m = ProvideRequest{} }
func (m *ProvideRequest) String() string { return proto.CompactTextString(m) }
func (*ProvideRequest) ProtoMessage()    {}
func (*ProvideRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1}
}

func (m *ProvideRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvideRequest.Unmarshal(m, b)
}
func (m *ProvideRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvideRequest.Marshal(b, m, deterministic)
}
func (m *ProvideRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvideRequest.Merge(m, src)
}
func (m *ProvideRequest) XXX_Size() int {
	return xxx_messageInfo_ProvideRequest.Size(m)
}
func (m *ProvideRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvideRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProvideRequest proto.InternalMessageInfo

func (m *ProvideRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ProvideRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ProvideRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

// ProvideResponse represents the returned credentials.
type ProvideResponse struct {
	// Inlined what we use from github.com/cli/cli/config/types
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Auth     string `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	// IdentityToken is used to authenticate the user and get
	// an access token for the registry.
	IdentityToken string `protobuf:"bytes,4,opt,name=identity_token,json=identityToken,proto3" json:"identity_token,omitempty"`
	// RegistryToken is a bearer token to be sent to a registry
	RegistryToken        string   `protobuf:"bytes,5,opt,name=registry_token,json=registryToken,proto3" json:"registry_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProvideResponse) Reset()         { *m = ProvideResponse{} }
func (m *ProvideResponse) String() string { return proto.CompactTextString(m) }
func (*ProvideResponse) ProtoMessage()    {}
func (*ProvideResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{2}
}

func (m *ProvideResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProvideResponse.Unmarshal(m, b)
}
func (m *ProvideResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProvideResponse.Marshal(b, m, deterministic)
}
func (m *ProvideResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvideResponse.Merge(m, src)
}
func (m *ProvideResponse) XXX_Size() int {
	return xxx_messageInfo_ProvideResponse.Size(m)
}
func (m *ProvideResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvideResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProvideResponse proto.InternalMessageInfo

func (m *ProvideResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ProvideResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ProvideResponse) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *ProvideResponse) GetIdentityToken() string {
	if m != nil {
		return m.IdentityToken
	}
	return ""
}

func (m *ProvideResponse) GetRegistryToken() string {
	if m != nil {
		return m.RegistryToken
	}
	return ""
}

// ConfigurationSchema represents the configuration for a plugin.
type ConfigurationSchema struct {
	// key is attribute name
	Attributes           map[string]*ConfigurationAttribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ConfigurationSchema) Reset()         { *m = ConfigurationSchema{} }
func (m *ConfigurationSchema) String() string { return proto.CompactTextString(m) }
func (*ConfigurationSchema) ProtoMessage()    {}
func (*ConfigurationSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{3}
}

func (m *ConfigurationSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigurationSchema.Unmarshal(m, b)
}
func (m *ConfigurationSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigurationSchema.Marshal(b, m, deterministic)
}
func (m *ConfigurationSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationSchema.Merge(m, src)
}
func (m *ConfigurationSchema) XXX_Size() int {
	return xxx_messageInfo_ConfigurationSchema.Size(m)
}
func (m *ConfigurationSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationSchema proto.InternalMessageInfo

func (m *ConfigurationSchema) GetAttributes() map[string]*ConfigurationAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ConfigurationAttribute struct {
	AttributeType        ConfigType `protobuf:"varint,1,opt,name=attribute_type,json=attributeType,proto3,enum=vjftw.dockerregistryproxy.v1.ConfigType" json:"attribute_type,omitempty"`
	Description          string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConfigurationAttribute) Reset()         { *m = ConfigurationAttribute{} }
func (m *ConfigurationAttribute) String() string { return proto.CompactTextString(m) }
func (*ConfigurationAttribute) ProtoMessage()    {}
func (*ConfigurationAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{4}
}

func (m *ConfigurationAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigurationAttribute.Unmarshal(m, b)
}
func (m *ConfigurationAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigurationAttribute.Marshal(b, m, deterministic)
}
func (m *ConfigurationAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationAttribute.Merge(m, src)
}
func (m *ConfigurationAttribute) XXX_Size() int {
	return xxx_messageInfo_ConfigurationAttribute.Size(m)
}
func (m *ConfigurationAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationAttribute proto.InternalMessageInfo

func (m *ConfigurationAttribute) GetAttributeType() ConfigType {
	if m != nil {
		return m.AttributeType
	}
	return ConfigType_BOOL
}

func (m *ConfigurationAttribute) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ConfigurationAttributeValue struct {
	AttributeType        ConfigType `protobuf:"varint,1,opt,name=attribute_type,json=attributeType,proto3,enum=vjftw.dockerregistryproxy.v1.ConfigType" json:"attribute_type,omitempty"`
	Value                []byte     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConfigurationAttributeValue) Reset()         { *m = ConfigurationAttributeValue{} }
func (m *ConfigurationAttributeValue) String() string { return proto.CompactTextString(m) }
func (*ConfigurationAttributeValue) ProtoMessage()    {}
func (*ConfigurationAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{5}
}

func (m *ConfigurationAttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigurationAttributeValue.Unmarshal(m, b)
}
func (m *ConfigurationAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigurationAttributeValue.Marshal(b, m, deterministic)
}
func (m *ConfigurationAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationAttributeValue.Merge(m, src)
}
func (m *ConfigurationAttributeValue) XXX_Size() int {
	return xxx_messageInfo_ConfigurationAttributeValue.Size(m)
}
func (m *ConfigurationAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationAttributeValue proto.InternalMessageInfo

func (m *ConfigurationAttributeValue) GetAttributeType() ConfigType {
	if m != nil {
		return m.AttributeType
	}
	return ConfigType_BOOL
}

func (m *ConfigurationAttributeValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ConfigureRequest struct {
	Attributes           map[string]*ConfigurationAttributeValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ConfigureRequest) Reset()         { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()    {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{6}
}

func (m *ConfigureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureRequest.Unmarshal(m, b)
}
func (m *ConfigureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureRequest.Marshal(b, m, deterministic)
}
func (m *ConfigureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureRequest.Merge(m, src)
}
func (m *ConfigureRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigureRequest.Size(m)
}
func (m *ConfigureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureRequest proto.InternalMessageInfo

func (m *ConfigureRequest) GetAttributes() map[string]*ConfigurationAttributeValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterEnum("vjftw.dockerregistryproxy.v1.ConfigType", ConfigType_name, ConfigType_value)
	proto.RegisterType((*VerifyRequest)(nil), "vjftw.dockerregistryproxy.v1.VerifyRequest")
	proto.RegisterType((*ProvideRequest)(nil), "vjftw.dockerregistryproxy.v1.ProvideRequest")
	proto.RegisterType((*ProvideResponse)(nil), "vjftw.dockerregistryproxy.v1.ProvideResponse")
	proto.RegisterType((*ConfigurationSchema)(nil), "vjftw.dockerregistryproxy.v1.ConfigurationSchema")
	proto.RegisterMapType((map[string]*ConfigurationAttribute)(nil), "vjftw.dockerregistryproxy.v1.ConfigurationSchema.AttributesEntry")
	proto.RegisterType((*ConfigurationAttribute)(nil), "vjftw.dockerregistryproxy.v1.ConfigurationAttribute")
	proto.RegisterType((*ConfigurationAttributeValue)(nil), "vjftw.dockerregistryproxy.v1.ConfigurationAttributeValue")
	proto.RegisterType((*ConfigureRequest)(nil), "vjftw.dockerregistryproxy.v1.ConfigureRequest")
	proto.RegisterMapType((map[string]*ConfigurationAttributeValue)(nil), "vjftw.dockerregistryproxy.v1.ConfigureRequest.AttributesEntry")
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor_d0dbc99083440df2) }

var fileDescriptor_d0dbc99083440df2 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x66, 0xf3, 0xdf, 0x49, 0x93, 0x5a, 0x43, 0x15, 0x45, 0x29, 0x42, 0x55, 0x24, 0xa4, 0x8a,
	0x1f, 0x57, 0x0d, 0x15, 0x02, 0x0e, 0x48, 0x69, 0x09, 0x55, 0x50, 0x9b, 0x54, 0x8e, 0xe9, 0xa1,
	0x07, 0x22, 0x37, 0xd9, 0xa4, 0xa6, 0x8d, 0xd7, 0xac, 0xd7, 0x69, 0x7d, 0x03, 0x89, 0x1b, 0x27,
	0x9e, 0x83, 0x27, 0xe1, 0x5d, 0xe0, 0x1d, 0xd0, 0xfa, 0x8f, 0xa4, 0x4a, 0xa2, 0x14, 0xb8, 0xed,
	0xce, 0x8c, 0x3f, 0xcf, 0x7c, 0xdf, 0xec, 0x07, 0xeb, 0x86, 0x2b, 0xce, 0xa9, 0x25, 0xcc, 0x9e,
	0x21, 0x4c, 0x66, 0xa9, 0x36, 0x67, 0x82, 0xe1, 0xbd, 0xf1, 0x87, 0x81, 0xb8, 0x52, 0xfb, 0xac,
	0x77, 0x41, 0x39, 0xa7, 0x43, 0xd3, 0x11, 0xdc, 0xb3, 0x39, 0xbb, 0xf6, 0xd4, 0xf1, 0x4e, 0x65,
	0x63, 0xc8, 0xd8, 0xf0, 0x92, 0x6e, 0xfb, 0xb5, 0x67, 0xee, 0x60, 0x9b, 0x8e, 0x6c, 0xe1, 0x05,
	0x9f, 0x56, 0xbf, 0x11, 0x28, 0x9c, 0x50, 0x6e, 0x0e, 0x3c, 0x8d, 0x7e, 0x74, 0xa9, 0x23, 0xb0,
	0x02, 0x39, 0xd7, 0xa1, 0xdc, 0x32, 0x46, 0xb4, 0x4c, 0x36, 0xc9, 0xd6, 0x8a, 0x16, 0xdf, 0x65,
	0xce, 0x36, 0x1c, 0xe7, 0x8a, 0xf1, 0x7e, 0x39, 0x11, 0xe4, 0xa2, 0x3b, 0x22, 0xa4, 0xce, 0x99,
	0x23, 0xca, 0x49, 0x3f, 0xee, 0x9f, 0xf1, 0x3e, 0x00, 0xa7, 0x36, 0x73, 0x4c, 0xc1, 0xb8, 0x57,
	0x4e, 0xf9, 0x99, 0x89, 0x08, 0xae, 0x43, 0xda, 0x1c, 0x19, 0x43, 0x5a, 0x4e, 0xfb, 0xa9, 0xe0,
	0x52, 0x3d, 0x85, 0xe2, 0x31, 0x67, 0x63, 0xb3, 0x4f, 0xa3, 0x9e, 0x22, 0x6c, 0x32, 0x17, 0x3b,
	0x31, 0x1f, 0x3b, 0x39, 0x89, 0xfd, 0x9d, 0xc0, 0x5a, 0x0c, 0xee, 0xd8, 0xcc, 0x72, 0xe8, 0xbf,
	0x4c, 0x2c, 0xe5, 0x88, 0x26, 0x96, 0x67, 0x7c, 0x00, 0x45, 0xb3, 0x2f, 0x05, 0x12, 0x5e, 0x57,
	0xb0, 0x0b, 0x6a, 0x85, 0x53, 0x17, 0xa2, 0xa8, 0x2e, 0x83, 0xb2, 0x2c, 0xd2, 0x29, 0x2c, 0x0b,
	0x18, 0x28, 0x44, 0x51, 0xbf, 0xac, 0xfa, 0x8b, 0xc0, 0xdd, 0x7d, 0x66, 0x0d, 0xcc, 0xa1, 0xcb,
	0x7d, 0xc1, 0x3b, 0xbd, 0x73, 0x3a, 0x32, 0xd0, 0x00, 0x30, 0x84, 0xe0, 0xe6, 0x99, 0x2b, 0xa8,
	0x53, 0x26, 0x9b, 0xc9, 0xad, 0x7c, 0xad, 0xae, 0x2e, 0xda, 0x02, 0x75, 0x06, 0x8c, 0x5a, 0x8f,
	0x31, 0x1a, 0x96, 0xe0, 0x9e, 0x36, 0x01, 0x5a, 0x71, 0x60, 0xed, 0x46, 0x1a, 0x15, 0x48, 0x5e,
	0x50, 0x2f, 0xa4, 0x48, 0x1e, 0xf1, 0x2d, 0xa4, 0xc7, 0xc6, 0xa5, 0x4b, 0x7d, 0x6a, 0xf2, 0xb5,
	0xdd, 0x5b, 0xb4, 0x10, 0x83, 0x6b, 0x01, 0xc4, 0xcb, 0xc4, 0x73, 0x52, 0xfd, 0x4a, 0xa0, 0x34,
	0xbb, 0x0a, 0xdb, 0x50, 0x8c, 0xbb, 0xeb, 0x0a, 0xcf, 0x0e, 0xa4, 0x2a, 0xd6, 0xb6, 0x96, 0xf9,
	0xa7, 0xee, 0xd9, 0x54, 0x2b, 0xc4, 0xdf, 0xcb, 0x2b, 0x6e, 0x42, 0xbe, 0x4f, 0x9d, 0x1e, 0x37,
	0x6d, 0xf9, 0xa3, 0x50, 0xdc, 0xc9, 0x50, 0xf5, 0x0b, 0x81, 0x8d, 0xd9, 0xdd, 0x9c, 0xc8, 0x8e,
	0xff, 0x7f, 0x4b, 0xeb, 0x93, 0x74, 0xae, 0x86, 0xc4, 0x54, 0x7f, 0x12, 0x50, 0xa2, 0x36, 0xe2,
	0x17, 0xf1, 0x7e, 0xc6, 0x06, 0xbc, 0x5a, 0x8e, 0xfe, 0x08, 0x63, 0xa1, 0xfc, 0xd7, 0xcb, 0xc8,
	0xdf, 0x9e, 0x96, 0xff, 0xc5, 0xdf, 0xc8, 0xef, 0x53, 0x39, 0xb1, 0x03, 0x0f, 0x3f, 0x13, 0x80,
	0x3f, 0x14, 0x61, 0x0e, 0x52, 0x7b, 0xed, 0xf6, 0xa1, 0x72, 0x07, 0xf3, 0x90, 0x7d, 0x73, 0xd8,
	0xae, 0xeb, 0xcf, 0x76, 0x15, 0x82, 0x59, 0x48, 0x36, 0x5b, 0xba, 0x92, 0xc0, 0x02, 0xac, 0x34,
	0x5b, 0x7a, 0xb7, 0x73, 0xd8, 0xdc, 0x6f, 0x28, 0x49, 0x04, 0xc8, 0x74, 0x74, 0xad, 0xd9, 0x3a,
	0x50, 0x52, 0x58, 0x04, 0x08, 0xce, 0xdd, 0xa3, 0xfa, 0xb1, 0x92, 0x46, 0x05, 0x56, 0xc3, 0x7b,
	0x50, 0x9d, 0x91, 0xe0, 0x7a, 0xf3, 0xa8, 0xa1, 0x64, 0x71, 0x15, 0x72, 0xaf, 0xdf, 0x69, 0x75,
	0xbd, 0xd9, 0x6e, 0x29, 0xb9, 0xda, 0x10, 0x4a, 0xf5, 0x29, 0xa3, 0xf5, 0x2d, 0xd2, 0xa4, 0x1c,
	0x8f, 0x20, 0x13, 0xd8, 0x25, 0x3e, 0x5a, 0x3c, 0xed, 0x94, 0xa9, 0x56, 0x4a, 0x6a, 0x60, 0xc2,
	0x6a, 0x64, 0xc2, 0x6a, 0x43, 0x9a, 0x70, 0xed, 0x07, 0x81, 0xc2, 0x14, 0x2f, 0xd8, 0x83, 0xd2,
	0x01, 0x15, 0xb3, 0x1e, 0xfd, 0x1c, 0x8c, 0xca, 0xce, 0xad, 0x1f, 0x3e, 0x76, 0x60, 0x25, 0xde,
	0x06, 0x54, 0x6f, 0xb7, 0x36, 0x73, 0x67, 0xf9, 0x44, 0x6e, 0xb2, 0x16, 0x1a, 0x2d, 0xc7, 0x01,
	0x64, 0xc3, 0x33, 0x3e, 0x5e, 0xfc, 0xb7, 0x69, 0xe3, 0xaf, 0x3c, 0x59, 0xb2, 0x3a, 0x70, 0xf2,
	0xbd, 0xd4, 0x69, 0x62, 0xbc, 0x73, 0x96, 0xf1, 0x1b, 0x7b, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x24, 0x75, 0x7c, 0x2d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationVerifierClient is the client API for AuthenticationVerifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationVerifierClient interface {
	// Verify takes the raw Base64 encoded BasicAuth credential supplied to
	// the Docker Registry Proxy from the Docker client and authenticates
	// with the implementation.
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authenticationVerifierClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationVerifierClient(cc grpc.ClientConnInterface) AuthenticationVerifierClient {
	return &authenticationVerifierClient{cc}
}

func (c *authenticationVerifierClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/vjftw.dockerregistryproxy.v1.AuthenticationVerifier/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationVerifierServer is the server API for AuthenticationVerifier service.
type AuthenticationVerifierServer interface {
	// Verify takes the raw Base64 encoded BasicAuth credential supplied to
	// the Docker Registry Proxy from the Docker client and authenticates
	// with the implementation.
	Verify(context.Context, *VerifyRequest) (*empty.Empty, error)
}

// UnimplementedAuthenticationVerifierServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationVerifierServer struct {
}

func (*UnimplementedAuthenticationVerifierServer) Verify(ctx context.Context, req *VerifyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterAuthenticationVerifierServer(s *grpc.Server, srv AuthenticationVerifierServer) {
	s.RegisterService(&_AuthenticationVerifier_serviceDesc, srv)
}

func _AuthenticationVerifier_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationVerifierServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vjftw.dockerregistryproxy.v1.AuthenticationVerifier/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationVerifierServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationVerifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vjftw.dockerregistryproxy.v1.AuthenticationVerifier",
	HandlerType: (*AuthenticationVerifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _AuthenticationVerifier_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	// GetConfigurationSchema returns the schema for the plugin
	GetConfigurationSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConfigurationSchema, error)
	// Configure configures the plugin
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type configurationClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationClient(cc grpc.ClientConnInterface) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) GetConfigurationSchema(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConfigurationSchema, error) {
	out := new(ConfigurationSchema)
	err := c.cc.Invoke(ctx, "/vjftw.dockerregistryproxy.v1.Configuration/GetConfigurationSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/vjftw.dockerregistryproxy.v1.Configuration/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	// GetConfigurationSchema returns the schema for the plugin
	GetConfigurationSchema(context.Context, *empty.Empty) (*ConfigurationSchema, error)
	// Configure configures the plugin
	Configure(context.Context, *ConfigureRequest) (*empty.Empty, error)
}

// UnimplementedConfigurationServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (*UnimplementedConfigurationServer) GetConfigurationSchema(ctx context.Context, req *empty.Empty) (*ConfigurationSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigurationSchema not implemented")
}
func (*UnimplementedConfigurationServer) Configure(ctx context.Context, req *ConfigureRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_GetConfigurationSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetConfigurationSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vjftw.dockerregistryproxy.v1.Configuration/GetConfigurationSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetConfigurationSchema(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vjftw.dockerregistryproxy.v1.Configuration/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vjftw.dockerregistryproxy.v1.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigurationSchema",
			Handler:    _Configuration_GetConfigurationSchema_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Configuration_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

// AuthenticationProviderClient is the client API for AuthenticationProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationProviderClient interface {
	Provide(ctx context.Context, in *ProvideRequest, opts ...grpc.CallOption) (*ProvideResponse, error)
}

type authenticationProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationProviderClient(cc grpc.ClientConnInterface) AuthenticationProviderClient {
	return &authenticationProviderClient{cc}
}

func (c *authenticationProviderClient) Provide(ctx context.Context, in *ProvideRequest, opts ...grpc.CallOption) (*ProvideResponse, error) {
	out := new(ProvideResponse)
	err := c.cc.Invoke(ctx, "/vjftw.dockerregistryproxy.v1.AuthenticationProvider/Provide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationProviderServer is the server API for AuthenticationProvider service.
type AuthenticationProviderServer interface {
	Provide(context.Context, *ProvideRequest) (*ProvideResponse, error)
}

// UnimplementedAuthenticationProviderServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationProviderServer struct {
}

func (*UnimplementedAuthenticationProviderServer) Provide(ctx context.Context, req *ProvideRequest) (*ProvideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provide not implemented")
}

func RegisterAuthenticationProviderServer(s *grpc.Server, srv AuthenticationProviderServer) {
	s.RegisterService(&_AuthenticationProvider_serviceDesc, srv)
}

func _AuthenticationProvider_Provide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationProviderServer).Provide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vjftw.dockerregistryproxy.v1.AuthenticationProvider/Provide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationProviderServer).Provide(ctx, req.(*ProvideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vjftw.dockerregistryproxy.v1.AuthenticationProvider",
	HandlerType: (*AuthenticationProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provide",
			Handler:    _AuthenticationProvider_Provide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
