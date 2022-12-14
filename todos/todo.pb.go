// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: todos/todo.proto

package todos

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

type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todos_todo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTodoRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetAllTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllTodosRequest) Reset() {
	*x = GetAllTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTodosRequest) ProtoMessage() {}

func (x *GetAllTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTodosRequest.ProtoReflect.Descriptor instead.
func (*GetAllTodosRequest) Descriptor() ([]byte, []int) {
	return file_todos_todo_proto_rawDescGZIP(), []int{1}
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todos_todo_proto_rawDescGZIP(), []int{2}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetAllTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *GetAllTodosResponse) Reset() {
	*x = GetAllTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todos_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTodosResponse) ProtoMessage() {}

func (x *GetAllTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todos_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTodosResponse.ProtoReflect.Descriptor instead.
func (*GetAllTodosResponse) Descriptor() ([]byte, []int) {
	return file_todos_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTodosResponse) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_todos_todo_proto protoreflect.FileDescriptor

var file_todos_todo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40,
	0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x32, 0xc7, 0x01, 0x0a, 0x0b, 0x54,
	0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73,
	0x12, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todos_todo_proto_rawDescOnce sync.Once
	file_todos_todo_proto_rawDescData = file_todos_todo_proto_rawDesc
)

func file_todos_todo_proto_rawDescGZIP() []byte {
	file_todos_todo_proto_rawDescOnce.Do(func() {
		file_todos_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todos_todo_proto_rawDescData)
	})
	return file_todos_todo_proto_rawDescData
}

var file_todos_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_todos_todo_proto_goTypes = []interface{}{
	(*CreateTodoRequest)(nil),   // 0: todos.CreateTodoRequest
	(*GetAllTodosRequest)(nil),  // 1: todos.GetAllTodosRequest
	(*Todo)(nil),                // 2: todos.Todo
	(*GetAllTodosResponse)(nil), // 3: todos.GetAllTodosResponse
}
var file_todos_todo_proto_depIdxs = []int32{
	2, // 0: todos.GetAllTodosResponse.todos:type_name -> todos.Todo
	0, // 1: todos.TodoService.CreateTodo:input_type -> todos.CreateTodoRequest
	1, // 2: todos.TodoService.GetAllTodos:input_type -> todos.GetAllTodosRequest
	1, // 3: todos.TodoService.StreamTodos:input_type -> todos.GetAllTodosRequest
	2, // 4: todos.TodoService.CreateTodo:output_type -> todos.Todo
	3, // 5: todos.TodoService.GetAllTodos:output_type -> todos.GetAllTodosResponse
	2, // 6: todos.TodoService.StreamTodos:output_type -> todos.Todo
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todos_todo_proto_init() }
func file_todos_todo_proto_init() {
	if File_todos_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todos_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
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
		file_todos_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTodosRequest); i {
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
		file_todos_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_todos_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTodosResponse); i {
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
			RawDescriptor: file_todos_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todos_todo_proto_goTypes,
		DependencyIndexes: file_todos_todo_proto_depIdxs,
		MessageInfos:      file_todos_todo_proto_msgTypes,
	}.Build()
	File_todos_todo_proto = out.File
	file_todos_todo_proto_rawDesc = nil
	file_todos_todo_proto_goTypes = nil
	file_todos_todo_proto_depIdxs = nil
}
