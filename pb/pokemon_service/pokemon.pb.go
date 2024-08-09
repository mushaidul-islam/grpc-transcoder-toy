// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: pokemon.proto

package pokemon_service

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
		mi := &file_pokemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[0]
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
	return file_pokemon_proto_rawDescGZIP(), []int{0}
}

type PokemonListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PokemonData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PokemonListResponse) Reset() {
	*x = PokemonListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonListResponse) ProtoMessage() {}

func (x *PokemonListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonListResponse.ProtoReflect.Descriptor instead.
func (*PokemonListResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{1}
}

func (x *PokemonListResponse) GetData() []*PokemonData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PokemonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PokemonData) Reset() {
	*x = PokemonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonData) ProtoMessage() {}

func (x *PokemonData) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonData.ProtoReflect.Descriptor instead.
func (*PokemonData) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{2}
}

func (x *PokemonData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PokemonData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_pokemon_proto protoreflect.FileDescriptor

var file_pokemon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0b, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x6a,
	0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e,
	0x73, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a,
	0x2f, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pokemon_proto_rawDescOnce sync.Once
	file_pokemon_proto_rawDescData = file_pokemon_proto_rawDesc
)

func file_pokemon_proto_rawDescGZIP() []byte {
	file_pokemon_proto_rawDescOnce.Do(func() {
		file_pokemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pokemon_proto_rawDescData)
	})
	return file_pokemon_proto_rawDescData
}

var file_pokemon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pokemon_proto_goTypes = []any{
	(*EmptyRequest)(nil),        // 0: pokemon.EmptyRequest
	(*PokemonListResponse)(nil), // 1: pokemon.PokemonListResponse
	(*PokemonData)(nil),         // 2: pokemon.PokemonData
}
var file_pokemon_proto_depIdxs = []int32{
	2, // 0: pokemon.PokemonListResponse.data:type_name -> pokemon.PokemonData
	0, // 1: pokemon.Pokemon.GetLegendaryPokemons:input_type -> pokemon.EmptyRequest
	1, // 2: pokemon.Pokemon.GetLegendaryPokemons:output_type -> pokemon.PokemonListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pokemon_proto_init() }
func file_pokemon_proto_init() {
	if File_pokemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pokemon_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_pokemon_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PokemonListResponse); i {
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
		file_pokemon_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PokemonData); i {
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
			RawDescriptor: file_pokemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pokemon_proto_goTypes,
		DependencyIndexes: file_pokemon_proto_depIdxs,
		MessageInfos:      file_pokemon_proto_msgTypes,
	}.Build()
	File_pokemon_proto = out.File
	file_pokemon_proto_rawDesc = nil
	file_pokemon_proto_goTypes = nil
	file_pokemon_proto_depIdxs = nil
}