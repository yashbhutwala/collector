// Code generated by protoc-gen-go.
// source: reports/bloat.proto
// DO NOT EDIT!

package pganalyze_collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BloatLookupMethod int32

const (
	BloatLookupMethod_ESTIMATE_FAST BloatLookupMethod = 0
	BloatLookupMethod_ESTIMATE_SLOW BloatLookupMethod = 1
	BloatLookupMethod_FULL_SCAN     BloatLookupMethod = 2
)

var BloatLookupMethod_name = map[int32]string{
	0: "ESTIMATE_FAST",
	1: "ESTIMATE_SLOW",
	2: "FULL_SCAN",
}
var BloatLookupMethod_value = map[string]int32{
	"ESTIMATE_FAST": 0,
	"ESTIMATE_SLOW": 1,
	"FULL_SCAN":     2,
}

func (x BloatLookupMethod) String() string {
	return proto.EnumName(BloatLookupMethod_name, int32(x))
}
func (BloatLookupMethod) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type BloatReport struct {
	ReportRunUuid           string                    `protobuf:"bytes,1,opt,name=report_run_uuid,json=reportRunUuid" json:"report_run_uuid,omitempty"`
	DatabaseReferences      []*DatabaseReference      `protobuf:"bytes,10,rep,name=database_references,json=databaseReferences" json:"database_references,omitempty"`
	RelationReferences      []*RelationReference      `protobuf:"bytes,11,rep,name=relation_references,json=relationReferences" json:"relation_references,omitempty"`
	IndexReferences         []*IndexReference         `protobuf:"bytes,12,rep,name=index_references,json=indexReferences" json:"index_references,omitempty"`
	RelationBloatStatistics []*RelationBloatStatistic `protobuf:"bytes,20,rep,name=relation_bloat_statistics,json=relationBloatStatistics" json:"relation_bloat_statistics,omitempty"`
	IndexBloatStatistics    []*IndexBloatStatistic    `protobuf:"bytes,21,rep,name=index_bloat_statistics,json=indexBloatStatistics" json:"index_bloat_statistics,omitempty"`
}

func (m *BloatReport) Reset()                    { *m = BloatReport{} }
func (m *BloatReport) String() string            { return proto.CompactTextString(m) }
func (*BloatReport) ProtoMessage()               {}
func (*BloatReport) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BloatReport) GetDatabaseReferences() []*DatabaseReference {
	if m != nil {
		return m.DatabaseReferences
	}
	return nil
}

func (m *BloatReport) GetRelationReferences() []*RelationReference {
	if m != nil {
		return m.RelationReferences
	}
	return nil
}

func (m *BloatReport) GetIndexReferences() []*IndexReference {
	if m != nil {
		return m.IndexReferences
	}
	return nil
}

func (m *BloatReport) GetRelationBloatStatistics() []*RelationBloatStatistic {
	if m != nil {
		return m.RelationBloatStatistics
	}
	return nil
}

func (m *BloatReport) GetIndexBloatStatistics() []*IndexBloatStatistic {
	if m != nil {
		return m.IndexBloatStatistics
	}
	return nil
}

type RelationBloatStatistic struct {
	RelationIdx       int32             `protobuf:"varint,1,opt,name=relation_idx,json=relationIdx" json:"relation_idx,omitempty"`
	BloatLookupMethod BloatLookupMethod `protobuf:"varint,2,opt,name=bloat_lookup_method,json=bloatLookupMethod,enum=pganalyze.collector.BloatLookupMethod" json:"bloat_lookup_method,omitempty"`
	TotalBytes        int64             `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	BloatBytes        int64             `protobuf:"varint,4,opt,name=bloat_bytes,json=bloatBytes" json:"bloat_bytes,omitempty"`
	LiveTupleBytes    int64             `protobuf:"varint,5,opt,name=live_tuple_bytes,json=liveTupleBytes" json:"live_tuple_bytes,omitempty"`
	LiveTupleCount    int64             `protobuf:"varint,6,opt,name=live_tuple_count,json=liveTupleCount" json:"live_tuple_count,omitempty"`
	DeadTupleBytes    int64             `protobuf:"varint,7,opt,name=dead_tuple_bytes,json=deadTupleBytes" json:"dead_tuple_bytes,omitempty"`
	DeadTupleCount    int64             `protobuf:"varint,8,opt,name=dead_tuple_count,json=deadTupleCount" json:"dead_tuple_count,omitempty"`
}

func (m *RelationBloatStatistic) Reset()                    { *m = RelationBloatStatistic{} }
func (m *RelationBloatStatistic) String() string            { return proto.CompactTextString(m) }
func (*RelationBloatStatistic) ProtoMessage()               {}
func (*RelationBloatStatistic) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type IndexBloatStatistic struct {
	IndexIdx          int32             `protobuf:"varint,1,opt,name=index_idx,json=indexIdx" json:"index_idx,omitempty"`
	BloatLookupMethod BloatLookupMethod `protobuf:"varint,2,opt,name=bloat_lookup_method,json=bloatLookupMethod,enum=pganalyze.collector.BloatLookupMethod" json:"bloat_lookup_method,omitempty"`
	TotalBytes        int64             `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes" json:"total_bytes,omitempty"`
	BloatBytes        int64             `protobuf:"varint,4,opt,name=bloat_bytes,json=bloatBytes" json:"bloat_bytes,omitempty"`
}

func (m *IndexBloatStatistic) Reset()                    { *m = IndexBloatStatistic{} }
func (m *IndexBloatStatistic) String() string            { return proto.CompactTextString(m) }
func (*IndexBloatStatistic) ProtoMessage()               {}
func (*IndexBloatStatistic) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func init() {
	proto.RegisterType((*BloatReport)(nil), "pganalyze.collector.BloatReport")
	proto.RegisterType((*RelationBloatStatistic)(nil), "pganalyze.collector.RelationBloatStatistic")
	proto.RegisterType((*IndexBloatStatistic)(nil), "pganalyze.collector.IndexBloatStatistic")
	proto.RegisterEnum("pganalyze.collector.BloatLookupMethod", BloatLookupMethod_name, BloatLookupMethod_value)
}

var fileDescriptor2 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x69, 0xbb, 0x8d, 0xf5, 0xa4, 0xdd, 0x52, 0x67, 0x8c, 0x00, 0x17, 0x8c, 0x22, 0x4d,
	0x15, 0x48, 0x45, 0x82, 0x27, 0xe8, 0xc6, 0x26, 0x2a, 0x75, 0x43, 0x72, 0x3a, 0x76, 0x47, 0xe4,
	0xc4, 0x66, 0x8b, 0x08, 0x71, 0x64, 0x3b, 0xa8, 0xe3, 0x8a, 0x07, 0xe4, 0x95, 0x90, 0x88, 0xed,
	0x76, 0x6b, 0xd3, 0x6c, 0xd7, 0x5c, 0xfa, 0xd3, 0xe7, 0xff, 0x3f, 0xf5, 0xa9, 0x02, 0x9e, 0x60,
	0x39, 0x17, 0x4a, 0xbe, 0x8b, 0x52, 0x4e, 0xd4, 0x30, 0x17, 0x5c, 0x71, 0xe4, 0xe5, 0x57, 0x24,
	0x23, 0xe9, 0xcd, 0x2f, 0x36, 0x8c, 0x79, 0x9a, 0xb2, 0x58, 0x71, 0xf1, 0xbc, 0x23, 0xaf, 0x89,
	0x60, 0xd4, 0x2a, 0xfd, 0xdf, 0x1b, 0xe0, 0x1c, 0xe9, 0x2b, 0xd8, 0xdc, 0x47, 0x87, 0xb0, 0x6b,
	0x93, 0x42, 0x51, 0x64, 0x61, 0x51, 0x24, 0xd4, 0x6f, 0x1c, 0x34, 0x06, 0x6d, 0xdc, 0xb5, 0x18,
	0x17, 0xd9, 0x45, 0x09, 0xd1, 0x25, 0x78, 0x94, 0x28, 0x12, 0x11, 0xc9, 0x42, 0xc1, 0xbe, 0x31,
	0xc1, 0xb2, 0x98, 0x49, 0x1f, 0x0e, 0x5a, 0x03, 0xe7, 0xfd, 0xe1, 0xb0, 0xa6, 0x78, 0xf8, 0x71,
	0xee, 0xe3, 0x85, 0x8e, 0x11, 0xad, 0x22, 0xa9, 0x83, 0x05, 0x4b, 0x89, 0x4a, 0x78, 0xb6, 0x1c,
	0xec, 0x3c, 0x10, 0x8c, 0xe7, 0xfe, 0x52, 0xb0, 0xa8, 0x22, 0x89, 0xce, 0xc1, 0x4d, 0x32, 0xca,
	0x66, 0xcb, 0xa9, 0x1d, 0x93, 0xfa, 0xba, 0x36, 0x75, 0xac, 0xe5, 0xbb, 0xc8, 0xdd, 0x64, 0xe5,
	0x2c, 0xd1, 0x15, 0x3c, 0xbb, 0x1d, 0xd4, 0x3c, 0x7a, 0x28, 0x55, 0x79, 0x92, 0x2a, 0x89, 0xa5,
	0xbf, 0x67, 0x82, 0xdf, 0x3e, 0x38, 0xae, 0x79, 0xf6, 0x60, 0x71, 0x07, 0x3f, 0x15, 0xb5, 0x5c,
	0xa2, 0xaf, 0xb0, 0x6f, 0x07, 0x5f, 0x6b, 0x79, 0x62, 0x5a, 0x06, 0xf7, 0x8f, 0x5f, 0xa9, 0xd8,
	0x4b, 0xd6, 0xa1, 0xec, 0xff, 0x6d, 0xc2, 0x7e, 0xfd, 0x4c, 0xe8, 0x15, 0x74, 0x6e, 0x7f, 0x63,
	0x42, 0x67, 0xe6, 0xaf, 0xb0, 0x89, 0x9d, 0x05, 0x1b, 0xd3, 0x19, 0xfa, 0x02, 0x9e, 0x9d, 0x2b,
	0xe5, 0xfc, 0x7b, 0x91, 0x87, 0x3f, 0x98, 0xba, 0xe6, 0xd4, 0x6f, 0x96, 0xe6, 0xce, 0x3d, 0xfb,
	0x32, 0x25, 0x13, 0xa3, 0x9f, 0x19, 0x1b, 0xf7, 0xa2, 0x2a, 0x42, 0x2f, 0xc1, 0x51, 0x5c, 0x91,
	0x34, 0x8c, 0x6e, 0x54, 0xb9, 0xa9, 0x56, 0x99, 0xd7, 0xc2, 0x60, 0xd0, 0x91, 0x26, 0x5a, 0xb0,
	0xc5, 0x56, 0xd8, 0xb0, 0x82, 0x41, 0x56, 0x18, 0x80, 0x9b, 0x26, 0x3f, 0x59, 0xa8, 0x8a, 0x3c,
	0x65, 0x73, 0x6b, 0xd3, 0x58, 0x3b, 0x9a, 0x4f, 0x35, 0xae, 0x33, 0x63, 0x5e, 0x64, 0xca, 0xdf,
	0xaa, 0x98, 0xc7, 0x9a, 0x6a, 0x93, 0x32, 0x42, 0x57, 0x32, 0x1f, 0x5b, 0x53, 0xf3, 0xd5, 0xcc,
	0x25, 0xd3, 0x66, 0x6e, 0x57, 0x4c, 0x93, 0xd9, 0xff, 0xd3, 0x00, 0xaf, 0x66, 0x5b, 0xe8, 0x05,
	0xb4, 0xed, 0xde, 0xef, 0x5e, 0x7e, 0xdb, 0x80, 0xff, 0xfa, 0xd9, 0xdf, 0x7c, 0x82, 0xde, 0x5a,
	0x13, 0xea, 0x41, 0xf7, 0x24, 0x98, 0x8e, 0xcf, 0x46, 0xd3, 0x93, 0xf0, 0x74, 0x14, 0x4c, 0xdd,
	0x47, 0x2b, 0x28, 0x98, 0x7c, 0xbe, 0x74, 0x1b, 0xa8, 0x0b, 0xed, 0xd3, 0x8b, 0xc9, 0x24, 0x0c,
	0x8e, 0x47, 0xe7, 0x6e, 0x33, 0xda, 0x32, 0x9f, 0xa8, 0x0f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0xd8, 0xa2, 0xa1, 0xdc, 0x04, 0x00, 0x00,
}
