// Code generated by protoc-gen-go.
// source: google/spanner/v1/type.proto
// DO NOT EDIT!

package spanner

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `TypeCode` is used as part of [Type][google.spanner.v1.Type] to
// indicate the type of a Cloud Spanner value.
//
// Each legal value of a type can be encoded to or decoded from a JSON
// value, using the encodings described below. All Cloud Spanner values can
// be `null`, regardless of type; `null`s are always encoded as a JSON
// `null`.
type TypeCode int32

const (
	// Not specified.
	TypeCode_TYPE_CODE_UNSPECIFIED TypeCode = 0
	// Encoded as JSON `true` or `false`.
	TypeCode_BOOL TypeCode = 1
	// Encoded as `string`, in decimal format.
	TypeCode_INT64 TypeCode = 2
	// Encoded as `number`, or the strings `"NaN"`, `"Infinity"`, or
	// `"-Infinity"`.
	TypeCode_FLOAT64 TypeCode = 3
	// Encoded as `string` in RFC 3339 timestamp format. The time zone
	// must be present, and must be `"Z"`.
	TypeCode_TIMESTAMP TypeCode = 4
	// Encoded as `string` in RFC 3339 date format.
	TypeCode_DATE TypeCode = 5
	// Encoded as `string`.
	TypeCode_STRING TypeCode = 6
	// Encoded as a base64-encoded `string`, as described in RFC 4648,
	// section 4.
	TypeCode_BYTES TypeCode = 7
	// Encoded as `list`, where the list elements are represented
	// according to [array_element_type][google.spanner.v1.Type.array_element_type].
	TypeCode_ARRAY TypeCode = 8
	// Encoded as `list`, where list element `i` is represented according
	// to [struct_type.fields[i]][google.spanner.v1.StructType.fields].
	TypeCode_STRUCT TypeCode = 9
)

var TypeCode_name = map[int32]string{
	0: "TYPE_CODE_UNSPECIFIED",
	1: "BOOL",
	2: "INT64",
	3: "FLOAT64",
	4: "TIMESTAMP",
	5: "DATE",
	6: "STRING",
	7: "BYTES",
	8: "ARRAY",
	9: "STRUCT",
}
var TypeCode_value = map[string]int32{
	"TYPE_CODE_UNSPECIFIED": 0,
	"BOOL":                  1,
	"INT64":                 2,
	"FLOAT64":               3,
	"TIMESTAMP":             4,
	"DATE":                  5,
	"STRING":                6,
	"BYTES":                 7,
	"ARRAY":                 8,
	"STRUCT":                9,
}

func (x TypeCode) String() string {
	return proto.EnumName(TypeCode_name, int32(x))
}
func (TypeCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// `Type` indicates the type of a Cloud Spanner value, as might be stored in a
// table cell or returned from an SQL query.
type Type struct {
	// Required. The [TypeCode][google.spanner.v1.TypeCode] for this type.
	Code TypeCode `protobuf:"varint,1,opt,name=code,enum=google.spanner.v1.TypeCode" json:"code,omitempty"`
	// If [code][google.spanner.v1.Type.code] == [ARRAY][google.spanner.v1.TypeCode.ARRAY], then `array_element_type`
	// is the type of the array elements.
	ArrayElementType *Type `protobuf:"bytes,2,opt,name=array_element_type,json=arrayElementType" json:"array_element_type,omitempty"`
	// If [code][google.spanner.v1.Type.code] == [STRUCT][google.spanner.v1.TypeCode.STRUCT], then `struct_type`
	// provides type information for the struct's fields.
	StructType *StructType `protobuf:"bytes,3,opt,name=struct_type,json=structType" json:"struct_type,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Type) GetCode() TypeCode {
	if m != nil {
		return m.Code
	}
	return TypeCode_TYPE_CODE_UNSPECIFIED
}

func (m *Type) GetArrayElementType() *Type {
	if m != nil {
		return m.ArrayElementType
	}
	return nil
}

func (m *Type) GetStructType() *StructType {
	if m != nil {
		return m.StructType
	}
	return nil
}

// `StructType` defines the fields of a [STRUCT][google.spanner.v1.TypeCode.STRUCT] type.
type StructType struct {
	// The list of fields that make up this struct. Order is
	// significant, because values of this struct type are represented as
	// lists, where the order of field values matches the order of
	// fields in the [StructType][google.spanner.v1.StructType]. In turn, the order of fields
	// matches the order of columns in a read request, or the order of
	// fields in the `SELECT` clause of a query.
	Fields []*StructType_Field `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
}

func (m *StructType) Reset()                    { *m = StructType{} }
func (m *StructType) String() string            { return proto.CompactTextString(m) }
func (*StructType) ProtoMessage()               {}
func (*StructType) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *StructType) GetFields() []*StructType_Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Message representing a single field of a struct.
type StructType_Field struct {
	// The name of the field. For reads, this is the column name. For
	// SQL queries, it is the column alias (e.g., `"Word"` in the
	// query `"SELECT 'hello' AS Word"`), or the column name (e.g.,
	// `"ColName"` in the query `"SELECT ColName FROM Table"`). Some
	// columns might have an empty name (e.g., !"SELECT
	// UPPER(ColName)"`). Note that a query result can contain
	// multiple fields with the same name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The type of the field.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *StructType_Field) Reset()                    { *m = StructType_Field{} }
func (m *StructType_Field) String() string            { return proto.CompactTextString(m) }
func (*StructType_Field) ProtoMessage()               {}
func (*StructType_Field) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1, 0} }

func (m *StructType_Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StructType_Field) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "google.spanner.v1.Type")
	proto.RegisterType((*StructType)(nil), "google.spanner.v1.StructType")
	proto.RegisterType((*StructType_Field)(nil), "google.spanner.v1.StructType.Field")
	proto.RegisterEnum("google.spanner.v1.TypeCode", TypeCode_name, TypeCode_value)
}

func init() { proto.RegisterFile("google/spanner/v1/type.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x9d, 0x36, 0xed, 0x36, 0xa7, 0x28, 0xe3, 0xc0, 0xb2, 0x75, 0x55, 0x28, 0xeb, 0x4d,
	0x51, 0x48, 0x68, 0x15, 0x11, 0x16, 0x84, 0x34, 0x9d, 0xae, 0x81, 0xdd, 0x36, 0x24, 0xb3, 0x42,
	0xbd, 0x29, 0x63, 0x3b, 0x86, 0x40, 0x3a, 0x13, 0x92, 0xec, 0x62, 0x5f, 0xc2, 0x1b, 0xdf, 0xc2,
	0x87, 0xf0, 0xd9, 0x64, 0x26, 0x59, 0x15, 0xaa, 0xe2, 0xdd, 0x9f, 0xfc, 0xff, 0x77, 0xce, 0x99,
	0xc3, 0x81, 0x27, 0x89, 0x52, 0x49, 0x26, 0xdc, 0x32, 0xe7, 0x52, 0x8a, 0xc2, 0xbd, 0x1d, 0xbb,
	0xd5, 0x3e, 0x17, 0x4e, 0x5e, 0xa8, 0x4a, 0x91, 0x87, 0xb5, 0xeb, 0x34, 0xae, 0x73, 0x3b, 0x3e,
	0xbd, 0x03, 0x78, 0x9e, 0xba, 0x5c, 0x4a, 0x55, 0xf1, 0x2a, 0x55, 0xb2, 0xac, 0x81, 0xb3, 0xef,
	0x08, 0x2c, 0xb6, 0xcf, 0x05, 0x71, 0xc1, 0xda, 0xa8, 0xad, 0x18, 0xa0, 0x21, 0x1a, 0x3d, 0x98,
	0x3c, 0x76, 0x0e, 0x0a, 0x39, 0x3a, 0xe6, 0xab, 0xad, 0x88, 0x4c, 0x90, 0x50, 0x20, 0xbc, 0x28,
	0xf8, 0x7e, 0x2d, 0x32, 0xb1, 0x13, 0xb2, 0x5a, 0xeb, 0x31, 0x06, 0xad, 0x21, 0x1a, 0xf5, 0x27,
	0x27, 0x7f, 0xc1, 0x23, 0x6c, 0x10, 0x5a, 0x13, 0xa6, 0xef, 0x5b, 0xe8, 0x97, 0x55, 0x71, 0xb3,
	0x69, 0xf8, 0xb6, 0xe1, 0x9f, 0xfe, 0x81, 0x8f, 0x4d, 0xca, 0x54, 0x81, 0xf2, 0xa7, 0x3e, 0xfb,
	0x8a, 0x00, 0x7e, 0x59, 0xe4, 0x1c, 0xba, 0x9f, 0x52, 0x91, 0x6d, 0xcb, 0x01, 0x1a, 0xb6, 0x47,
	0xfd, 0xc9, 0xb3, 0x7f, 0x56, 0x72, 0xe6, 0x3a, 0x1b, 0x35, 0xc8, 0xe9, 0x3b, 0xe8, 0x98, 0x1f,
	0x84, 0x80, 0x25, 0xf9, 0xae, 0x5e, 0x86, 0x1d, 0x19, 0x4d, 0x5e, 0x80, 0xf5, 0x3f, 0x2f, 0x34,
	0xa1, 0xe7, 0x5f, 0x10, 0xf4, 0xee, 0xf6, 0x45, 0x1e, 0xc1, 0x31, 0x5b, 0x85, 0x74, 0xed, 0x2f,
	0x67, 0x74, 0x7d, 0xbd, 0x88, 0x43, 0xea, 0x07, 0xf3, 0x80, 0xce, 0xf0, 0x3d, 0xd2, 0x03, 0x6b,
	0xba, 0x5c, 0x5e, 0x62, 0x44, 0x6c, 0xe8, 0x04, 0x0b, 0xf6, 0xfa, 0x15, 0x6e, 0x91, 0x3e, 0x1c,
	0xcd, 0x2f, 0x97, 0x9e, 0xfe, 0x68, 0x93, 0xfb, 0x60, 0xb3, 0xe0, 0x8a, 0xc6, 0xcc, 0xbb, 0x0a,
	0xb1, 0xa5, 0x81, 0x99, 0xc7, 0x28, 0xee, 0x10, 0x80, 0x6e, 0xcc, 0xa2, 0x60, 0x71, 0x81, 0xbb,
	0x1a, 0x9e, 0xae, 0x18, 0x8d, 0xf1, 0x91, 0x96, 0x5e, 0x14, 0x79, 0x2b, 0xdc, 0x6b, 0x12, 0xd7,
	0x3e, 0xc3, 0xf6, 0xf4, 0x33, 0x1c, 0x6f, 0xd4, 0xee, 0x70, 0xe8, 0xa9, 0xad, 0xc7, 0x0c, 0xf5,
	0x2d, 0x84, 0xe8, 0xc3, 0x9b, 0xc6, 0x4f, 0x54, 0xc6, 0x65, 0xe2, 0xa8, 0x22, 0x71, 0x13, 0x21,
	0xcd, 0xa5, 0xb8, 0xb5, 0xc5, 0xf3, 0xb4, 0xfc, 0xed, 0xf6, 0xce, 0x1b, 0xf9, 0xad, 0x75, 0x72,
	0x51, 0xa3, 0x7e, 0xa6, 0x6e, 0xb6, 0x4e, 0xdc, 0x34, 0x78, 0x3f, 0xfe, 0xd8, 0x35, 0xf8, 0xcb,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xaa, 0xcb, 0xef, 0xb9, 0x02, 0x00, 0x00,
}
