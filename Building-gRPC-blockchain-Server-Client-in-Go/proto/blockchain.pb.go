package proto

import (
	fmt "fmt"

	proto1 "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AddBlockRequest struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *AddBlockRequest) Reset()                    { *m = AddBlockRequest{} }
func (m *AddBlockRequest) String() string            { return proto1.CompactTextString(m) }
func (*AddBlockRequest) ProtoMessage()               {}
func (*AddBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddBlockRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AddBlockResponse struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *AddBlockResponse) Reset()                    { *m = AddBlockResponse{} }
func (m *AddBlockResponse) String() string            { return proto1.CompactTextString(m) }
func (*AddBlockResponse) ProtoMessage()               {}
func (*AddBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddBlockResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Block struct {
	PrevBlockHash string `protobuf:"bytes,1,opt,name=prevBlockHash" json:"prevBlockHash,omitempty"`
	Data          string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Hash          string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto1.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Block) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *Block) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockchainRequest struct {
}

func (m *GetBlockchainRequest) Reset()                    { *m = GetBlockchainRequest{} }
func (m *GetBlockchainRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetBlockchainRequest) ProtoMessage()               {}
func (*GetBlockchainRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetBlockchainResponse struct {
	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *GetBlockchainResponse) Reset()                    { *m = GetBlockchainResponse{} }
func (m *GetBlockchainResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetBlockchainResponse) ProtoMessage()               {}
func (*GetBlockchainResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetBlockchainResponse) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto1.RegisterType((*AddBlockRequest)(nil), "proto.AddBlockRequest")
	proto1.RegisterType((*AddBlockResponse)(nil), "proto.AddBlockResponse")
	proto1.RegisterType((*Block)(nil), "proto.Block")
	proto1.RegisterType((*GetBlockchainRequest)(nil), "proto.GetBlockchainRequest")
	proto1.RegisterType((*GetBlockchainResponse)(nil), "proto.GetBlockchainResponse")
}

var _ context.Context
var _ grpc.ClientConn

const _ = grpc.SupportPackageIsVersion4

type BlockchainClient interface {
	AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error)
	GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error)
}

type blockchainClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainClient(cc *grpc.ClientConn) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error) {
	out := new(AddBlockResponse)
	err := grpc.Invoke(ctx, "/proto.Blockchain/AddBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error) {
	out := new(GetBlockchainResponse)
	err := grpc.Invoke(ctx, "/proto.Blockchain/GetBlockchain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type BlockchainServer interface {
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
	GetBlockchain(context.Context, *GetBlockchainRequest) (*GetBlockchainResponse, error)
}

func RegisterBlockchainServer(s *grpc.Server, srv BlockchainServer) {
	s.RegisterService(&_Blockchain_serviceDesc, srv)
}

func _Blockchain_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blockchain/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).AddBlock(ctx, req.(*AddBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blockchain/GetBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockchain(ctx, req.(*GetBlockchainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blockchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlock",
			Handler:    _Blockchain_AddBlock_Handler,
		},
		{
			MethodName: "GetBlockchain",
			Handler:    _Blockchain_GetBlockchain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blockchain.proto",
}

func init() { proto1.RegisterFile("proto/blockchain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x03, 0x0b, 0x08,
	0xb1, 0x82, 0x29, 0x25, 0x55, 0x2e, 0x7e, 0xc7, 0x94, 0x14, 0x27, 0x90, 0x6c, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x98, 0xad, 0xa4, 0xc6, 0x25, 0x80, 0x50, 0x56, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x0a, 0x52, 0x97, 0x91, 0x58, 0x9c, 0x01, 0x53, 0x07, 0x62, 0x2b, 0x85, 0x72, 0xb1, 0x82, 0x15,
	0x09, 0xa9, 0x70, 0xf1, 0x16, 0x14, 0xa5, 0x96, 0x81, 0x39, 0x1e, 0x08, 0x55, 0xa8, 0x82, 0x70,
	0xab, 0x98, 0x10, 0x56, 0xc1, 0x8d, 0x65, 0x46, 0x32, 0x56, 0x8c, 0x4b, 0xc4, 0x3d, 0xb5, 0xc4,
	0x09, 0xee, 0x07, 0xa8, 0x53, 0x95, 0x6c, 0xb9, 0x44, 0xd1, 0xc4, 0xa1, 0x6e, 0x53, 0xe1, 0x62,
	0x03, 0xfb, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x07, 0xe2, 0x6b, 0x3d, 0x88,
	0x0f, 0xa0, 0x72, 0x46, 0x33, 0x19, 0xb9, 0xb8, 0x10, 0x9a, 0x85, 0x6c, 0xb9, 0x38, 0x60, 0x9e,
	0x14, 0x12, 0x83, 0x6a, 0x40, 0x0b, 0x1c, 0x29, 0x71, 0x0c, 0x71, 0x88, 0x8d, 0x4a, 0x0c, 0x42,
	0x3e, 0x5c, 0xbc, 0x28, 0x8e, 0x11, 0x92, 0x86, 0xaa, 0xc5, 0xe6, 0x74, 0x29, 0x19, 0xec, 0x92,
	0x30, 0xd3, 0x92, 0xd8, 0xc0, 0xd2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xaa, 0xb7,
	0xd5, 0xc0, 0x01, 0x00, 0x00,
}
