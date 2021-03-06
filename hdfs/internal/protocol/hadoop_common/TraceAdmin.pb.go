// Code generated by protoc-gen-go.
// source: TraceAdmin.proto
// DO NOT EDIT!

package hadoop_common

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ListSpanReceiversRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListSpanReceiversRequestProto) Reset()         { *m = ListSpanReceiversRequestProto{} }
func (m *ListSpanReceiversRequestProto) String() string { return proto.CompactTextString(m) }
func (*ListSpanReceiversRequestProto) ProtoMessage()    {}

type SpanReceiverListInfo struct {
	Id               *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ClassName        *string `protobuf:"bytes,2,req,name=className" json:"className,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SpanReceiverListInfo) Reset()         { *m = SpanReceiverListInfo{} }
func (m *SpanReceiverListInfo) String() string { return proto.CompactTextString(m) }
func (*SpanReceiverListInfo) ProtoMessage()    {}

func (m *SpanReceiverListInfo) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SpanReceiverListInfo) GetClassName() string {
	if m != nil && m.ClassName != nil {
		return *m.ClassName
	}
	return ""
}

type ListSpanReceiversResponseProto struct {
	Descriptions     []*SpanReceiverListInfo `protobuf:"bytes,1,rep,name=descriptions" json:"descriptions,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListSpanReceiversResponseProto) Reset()         { *m = ListSpanReceiversResponseProto{} }
func (m *ListSpanReceiversResponseProto) String() string { return proto.CompactTextString(m) }
func (*ListSpanReceiversResponseProto) ProtoMessage()    {}

func (m *ListSpanReceiversResponseProto) GetDescriptions() []*SpanReceiverListInfo {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

type ConfigPair struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConfigPair) Reset()         { *m = ConfigPair{} }
func (m *ConfigPair) String() string { return proto.CompactTextString(m) }
func (*ConfigPair) ProtoMessage()    {}

func (m *ConfigPair) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ConfigPair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type AddSpanReceiverRequestProto struct {
	ClassName        *string       `protobuf:"bytes,1,req,name=className" json:"className,omitempty"`
	Config           []*ConfigPair `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *AddSpanReceiverRequestProto) Reset()         { *m = AddSpanReceiverRequestProto{} }
func (m *AddSpanReceiverRequestProto) String() string { return proto.CompactTextString(m) }
func (*AddSpanReceiverRequestProto) ProtoMessage()    {}

func (m *AddSpanReceiverRequestProto) GetClassName() string {
	if m != nil && m.ClassName != nil {
		return *m.ClassName
	}
	return ""
}

func (m *AddSpanReceiverRequestProto) GetConfig() []*ConfigPair {
	if m != nil {
		return m.Config
	}
	return nil
}

type AddSpanReceiverResponseProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AddSpanReceiverResponseProto) Reset()         { *m = AddSpanReceiverResponseProto{} }
func (m *AddSpanReceiverResponseProto) String() string { return proto.CompactTextString(m) }
func (*AddSpanReceiverResponseProto) ProtoMessage()    {}

func (m *AddSpanReceiverResponseProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type RemoveSpanReceiverRequestProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveSpanReceiverRequestProto) Reset()         { *m = RemoveSpanReceiverRequestProto{} }
func (m *RemoveSpanReceiverRequestProto) String() string { return proto.CompactTextString(m) }
func (*RemoveSpanReceiverRequestProto) ProtoMessage()    {}

func (m *RemoveSpanReceiverRequestProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type RemoveSpanReceiverResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveSpanReceiverResponseProto) Reset()         { *m = RemoveSpanReceiverResponseProto{} }
func (m *RemoveSpanReceiverResponseProto) String() string { return proto.CompactTextString(m) }
func (*RemoveSpanReceiverResponseProto) ProtoMessage()    {}

func init() {
}
