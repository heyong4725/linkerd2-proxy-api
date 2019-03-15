// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

package identity // import "github.com/linkerd/linkerd2-proxy-api/go/identity"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CertifyRequest struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// Proof of the requester's identity.
	//
	// In Kubernetes, for instance, this is the contents of a service account token.
	Token []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// A PEM-encoded x509 Certificate Signing Request.
	CertificateSigningRequest []byte   `protobuf:"bytes,3,opt,name=certificate_signing_request,json=certificateSigningRequest,proto3" json:"certificate_signing_request,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *CertifyRequest) Reset()         { *m = CertifyRequest{} }
func (m *CertifyRequest) String() string { return proto.CompactTextString(m) }
func (*CertifyRequest) ProtoMessage()    {}
func (*CertifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_6e4256bddd2bbdea, []int{0}
}
func (m *CertifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertifyRequest.Unmarshal(m, b)
}
func (m *CertifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertifyRequest.Marshal(b, m, deterministic)
}
func (dst *CertifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertifyRequest.Merge(dst, src)
}
func (m *CertifyRequest) XXX_Size() int {
	return xxx_messageInfo_CertifyRequest.Size(m)
}
func (m *CertifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertifyRequest proto.InternalMessageInfo

func (m *CertifyRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CertifyRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *CertifyRequest) GetCertificateSigningRequest() []byte {
	if m != nil {
		return m.CertificateSigningRequest
	}
	return nil
}

type CertifyResponse struct {
	// A PEM-encoded x509 Certificate.
	LeafCertificate []byte `protobuf:"bytes,1,opt,name=leaf_certificate,json=leafCertificate,proto3" json:"leaf_certificate,omitempty"`
	// A list of PEM-encoded x509 Certificates that establish the trust chain
	// between the leaf_certificate and the well-known trust anchors.
	IntermediateCertificates [][]byte             `protobuf:"bytes,2,rep,name=intermediate_certificates,json=intermediateCertificates,proto3" json:"intermediate_certificates,omitempty"`
	ValidUntil               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}             `json:"-"`
	XXX_unrecognized         []byte               `json:"-"`
	XXX_sizecache            int32                `json:"-"`
}

func (m *CertifyResponse) Reset()         { *m = CertifyResponse{} }
func (m *CertifyResponse) String() string { return proto.CompactTextString(m) }
func (*CertifyResponse) ProtoMessage()    {}
func (*CertifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_6e4256bddd2bbdea, []int{1}
}
func (m *CertifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertifyResponse.Unmarshal(m, b)
}
func (m *CertifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertifyResponse.Marshal(b, m, deterministic)
}
func (dst *CertifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertifyResponse.Merge(dst, src)
}
func (m *CertifyResponse) XXX_Size() int {
	return xxx_messageInfo_CertifyResponse.Size(m)
}
func (m *CertifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertifyResponse proto.InternalMessageInfo

func (m *CertifyResponse) GetLeafCertificate() []byte {
	if m != nil {
		return m.LeafCertificate
	}
	return nil
}

func (m *CertifyResponse) GetIntermediateCertificates() [][]byte {
	if m != nil {
		return m.IntermediateCertificates
	}
	return nil
}

func (m *CertifyResponse) GetValidUntil() *timestamp.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func init() {
	proto.RegisterType((*CertifyRequest)(nil), "io.linkerd.proxy.identity.CertifyRequest")
	proto.RegisterType((*CertifyResponse)(nil), "io.linkerd.proxy.identity.CertifyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	// Requests that a time-bounded certificate be signed.
	//
	// The requester must provide a token that verifies the client's identity and
	// a Certificate Signing Request that adheres to the service naming rules.
	//
	// Errors are returned when the provided request is invalid or when
	// authentication cannot be performed.
	Certify(ctx context.Context, in *CertifyRequest, opts ...grpc.CallOption) (*CertifyResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) Certify(ctx context.Context, in *CertifyRequest, opts ...grpc.CallOption) (*CertifyResponse, error) {
	out := new(CertifyResponse)
	err := c.cc.Invoke(ctx, "/io.linkerd.proxy.identity.Identity/Certify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	// Requests that a time-bounded certificate be signed.
	//
	// The requester must provide a token that verifies the client's identity and
	// a Certificate Signing Request that adheres to the service naming rules.
	//
	// Errors are returned when the provided request is invalid or when
	// authentication cannot be performed.
	Certify(context.Context, *CertifyRequest) (*CertifyResponse, error)
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_Certify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Certify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.linkerd.proxy.identity.Identity/Certify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Certify(ctx, req.(*CertifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.linkerd.proxy.identity.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certify",
			Handler:    _Identity_Certify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor_identity_6e4256bddd2bbdea) }

var fileDescriptor_identity_6e4256bddd2bbdea = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0xff, 0x85, 0xfc, 0x15, 0x17, 0x02, 0x66, 0xe3, 0xa1, 0xd4, 0x0b, 0xe1, 0x04, 0x26,
	0x6c, 0x23, 0x1c, 0x49, 0x3c, 0xc8, 0xc9, 0x6b, 0xd5, 0x8b, 0x97, 0xa6, 0xa5, 0x43, 0x9d, 0xd0,
	0xee, 0xd6, 0xee, 0xd6, 0xc8, 0xd5, 0x0f, 0xe5, 0xe7, 0x33, 0xdd, 0xed, 0x42, 0x3d, 0x98, 0x78,
	0x6a, 0x67, 0xf2, 0x7b, 0x6f, 0x66, 0xde, 0x92, 0x21, 0x26, 0xc0, 0x15, 0xaa, 0x03, 0x2b, 0x4a,
	0xa1, 0x04, 0x1d, 0xa3, 0x60, 0x19, 0xf2, 0x3d, 0x94, 0x49, 0xdd, 0xf9, 0x38, 0x30, 0x0b, 0x78,
	0xf3, 0x54, 0x88, 0x34, 0x03, 0x5f, 0x83, 0x71, 0xb5, 0xf3, 0x15, 0xe6, 0x20, 0x55, 0x94, 0x17,
	0xa7, 0x3f, 0xe3, 0x32, 0xfd, 0x74, 0xc8, 0x70, 0x03, 0xa5, 0xc2, 0xdd, 0x21, 0x80, 0xb7, 0x0a,
	0xa4, 0xa2, 0x1e, 0xe9, 0x59, 0x27, 0xd7, 0x99, 0x38, 0xb3, 0x8b, 0xe0, 0x58, 0xd3, 0x2b, 0xf2,
	0x5f, 0x89, 0x3d, 0x70, 0xb7, 0x33, 0x71, 0x66, 0x83, 0xc0, 0x14, 0xf4, 0x8e, 0x5c, 0x6f, 0xb5,
	0x07, 0x6e, 0x23, 0x05, 0xa1, 0xc4, 0x94, 0x23, 0x4f, 0xc3, 0xd2, 0x18, 0xba, 0x5d, 0xcd, 0x8e,
	0x5b, 0xc8, 0xa3, 0x21, 0x9a, 0x89, 0xd3, 0x2f, 0x87, 0x8c, 0x8e, 0x4b, 0xc8, 0x42, 0x70, 0x09,
	0x74, 0x4e, 0x2e, 0x33, 0x88, 0x76, 0x61, 0x4b, 0xa5, 0xb7, 0x19, 0x04, 0xa3, 0xba, 0xbf, 0x39,
	0xb5, 0xe9, 0x9a, 0x8c, 0x91, 0x2b, 0x28, 0x73, 0x48, 0xb0, 0x9e, 0xdf, 0x92, 0x48, 0xb7, 0x33,
	0xe9, 0xce, 0x06, 0x81, 0xdb, 0x06, 0x5a, 0x5a, 0x49, 0xd7, 0xa4, 0xff, 0x1e, 0x65, 0x98, 0x84,
	0x15, 0x57, 0x98, 0xe9, 0x5d, 0xfb, 0x4b, 0x8f, 0x99, 0x04, 0x99, 0x4d, 0x90, 0x3d, 0xd9, 0xdc,
	0x02, 0xa2, 0xf1, 0xe7, 0x9a, 0x5e, 0x72, 0xd2, 0x7b, 0xb0, 0xd1, 0xc4, 0xe4, 0xbc, 0xb9, 0x81,
	0xce, 0xd9, 0xaf, 0x6f, 0xc3, 0x7e, 0x86, 0xed, 0xdd, 0xfc, 0x05, 0x35, 0x91, 0x4c, 0xff, 0xdd,
	0xaf, 0x5e, 0x6e, 0x53, 0x54, 0xaf, 0x55, 0xcc, 0xb6, 0x22, 0xf7, 0x1b, 0x99, 0xfd, 0x2e, 0x17,
	0x5a, 0xbf, 0x88, 0x0a, 0xf4, 0x53, 0xe1, 0x5b, 0x9b, 0xf8, 0x4c, 0x1f, 0xb1, 0xfa, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xce, 0xd7, 0xaa, 0x04, 0x41, 0x02, 0x00, 0x00,
}
