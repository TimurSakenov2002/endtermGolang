package calculatorPB

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type PrimeNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeNumberRequest) Reset() {
	*x = PrimeNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorPB_protocol_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberRequest) ProtoMessage() {}

func (x *PrimeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorPB_protocol_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorPB_protocol_rawDescGZIP(), []int{0}
}

func (x *PrimeNumberRequest) sumNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumberResponse) Reset() {
	*x = PrimeNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorPB_protocol_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberResponse) ProtoMessage() {}

func (x *PrimeNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorPB_protocol_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorPB_protocol_rawDescGZIP(), []int{1}
}

func (x *PrimeNumberResponse) getResultat() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AvgNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *AvgNumberRequest) Reset() {
	*x = AvgNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorPB_protocol_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvgNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvgNumberRequest) ProtoMessage() {}

func (x *AvgNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorPB_protocol_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AvgNumberRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorPB_protocol_rawDescGZIP(), []int{2}
}

func (x *AvgNumberRequest) sumNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AvgNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res float32 `protobuf:"fixed32,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *AvgNumberResponse) Reset() {
	*x = AvgNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorPB_protocol_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvgNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvgNumberResponse) ProtoMessage() {}

func (x *AvgNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorPB_protocol_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AvgNumberResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorPB_protocol_rawDescGZIP(), []int{3}
}

func (x *AvgNumberResponse) resultat() float32 {
	if x != nil {
		return x.Res
	}
	return 0
}

var File_calculator_calculatorPB_protocol_ protoreflect.FileDescriptor

var (
	file_calculator_calculatorPB_protocol_rawDescOnce sync.Once
	file_calculator_calculatorPB_protocol_rawDescData = file_calculator_calculatorPB_protocol_rawDesc
)
func init() { file_calculator_calculatorPB_protocol_init() }

func file_calculator_calculatorPB_protocol_init() {
	if File_calculator_calculatorPB_protocol != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculatorPB_protocol_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberRequest); i {
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
		file_calculator_calculatorPB_protocol_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberResponse); i {
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
		file_calculator_calculatorPB_protocol_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvgNumberRequest); i {
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
		file_calculator_calculatorPB_protocol_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvgNumberResponse); i {
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
	/*
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_calculatorPB_protocol_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculatorPB_protocol_goTypes,
		DependencyIndexes: file_calculator_calculatorPB_protocol_depIdxs,
		MessageInfos:      file_calculator_calculatorPB_protocol_msgTypes,
	}.Build()
	File_calculator_calculatorPB_protocol = out.File
	file_calculator_calculatorPB_protocol_rawDesc = nil
	file_calculator_calculatorPB_protocol_goTypes = nil
	file_calculator_calculatorPB_protocol_depIdxs = nil
	 */
}