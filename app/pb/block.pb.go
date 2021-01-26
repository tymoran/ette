// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: block.proto

package proto

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash                string         `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Number              uint64         `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Time                uint64         `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	ParentHash          string         `protobuf:"bytes,4,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Difficulty          string         `protobuf:"bytes,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	GasUsed             uint64         `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	GasLimit            uint64         `protobuf:"varint,7,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Nonce               uint64         `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Miner               string         `protobuf:"bytes,9,opt,name=miner,proto3" json:"miner,omitempty"`
	Size                float64        `protobuf:"fixed64,10,opt,name=size,proto3" json:"size,omitempty"`
	TransactionRootHash string         `protobuf:"bytes,11,opt,name=transaction_root_hash,json=transactionRootHash,proto3" json:"transaction_root_hash,omitempty"`
	ReceiptRootHash     string         `protobuf:"bytes,12,opt,name=receipt_root_hash,json=receiptRootHash,proto3" json:"receipt_root_hash,omitempty"`
	Transactions        []*Transaction `protobuf:"bytes,13,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Block) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Block) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *Block) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Block) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *Block) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Block) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Block) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *Block) GetSize() float64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Block) GetTransactionRootHash() string {
	if x != nil {
		return x.TransactionRootHash
	}
	return ""
}

func (x *Block) GetReceiptRootHash() string {
	if x != nil {
		return x.ReceiptRootHash
	}
	return ""
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x92, 0x03, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67,
	0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x7a, 0x6d, 0x65, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x2f, 0x65,
	0x74, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_block_proto_goTypes = []interface{}{
	(*Block)(nil),       // 0: Block
	(*Transaction)(nil), // 1: Transaction
}
var file_block_proto_depIdxs = []int32{
	1, // 0: Block.transactions:type_name -> Transaction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
