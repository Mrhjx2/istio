// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/router/v2/router.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Router struct {
	// Whether the router generates dynamic cluster statistics. Defaults to
	// true. Can be disabled in high performance scenarios.
	DynamicStats *types.BoolValue `protobuf:"bytes,1,opt,name=dynamic_stats,json=dynamicStats,proto3" json:"dynamic_stats,omitempty"`
	// Whether to start a child span for egress routed calls. This can be
	// useful in scenarios where other filters (auth, ratelimit, etc.) make
	// outbound calls and have child spans rooted at the same ingress
	// parent. Defaults to false.
	StartChildSpan bool `protobuf:"varint,2,opt,name=start_child_span,json=startChildSpan,proto3" json:"start_child_span,omitempty"`
	// Configuration for HTTP upstream logs emitted by the router. Upstream logs
	// are configured in the same way as access logs, but each log entry represents
	// an upstream request. Presuming retries are configured, multiple upstream
	// requests may be made for each downstream (inbound) request.
	UpstreamLog []*v2.AccessLog `protobuf:"bytes,3,rep,name=upstream_log,json=upstreamLog,proto3" json:"upstream_log,omitempty"`
	// Do not add any additional *x-envoy-* headers to requests or responses. This
	// only affects the :ref:`router filter generated *x-envoy-* headers
	// <config_http_filters_router_headers_set>`, other Envoy filters and the HTTP
	// connection manager may continue to set *x-envoy-* headers.
	SuppressEnvoyHeaders bool     `protobuf:"varint,4,opt,name=suppress_envoy_headers,json=suppressEnvoyHeaders,proto3" json:"suppress_envoy_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_822295bdbc4788d1, []int{0}
}
func (m *Router) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Router.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(dst, src)
}
func (m *Router) XXX_Size() int {
	return m.Size()
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetDynamicStats() *types.BoolValue {
	if m != nil {
		return m.DynamicStats
	}
	return nil
}

func (m *Router) GetStartChildSpan() bool {
	if m != nil {
		return m.StartChildSpan
	}
	return false
}

func (m *Router) GetUpstreamLog() []*v2.AccessLog {
	if m != nil {
		return m.UpstreamLog
	}
	return nil
}

func (m *Router) GetSuppressEnvoyHeaders() bool {
	if m != nil {
		return m.SuppressEnvoyHeaders
	}
	return false
}

func init() {
	proto.RegisterType((*Router)(nil), "envoy.config.filter.http.router.v2.Router")
}
func (m *Router) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Router) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DynamicStats != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRouter(dAtA, i, uint64(m.DynamicStats.Size()))
		n1, err := m.DynamicStats.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.StartChildSpan {
		dAtA[i] = 0x10
		i++
		if m.StartChildSpan {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.UpstreamLog) > 0 {
		for _, msg := range m.UpstreamLog {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRouter(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.SuppressEnvoyHeaders {
		dAtA[i] = 0x20
		i++
		if m.SuppressEnvoyHeaders {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRouter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Router) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DynamicStats != nil {
		l = m.DynamicStats.Size()
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.StartChildSpan {
		n += 2
	}
	if len(m.UpstreamLog) > 0 {
		for _, e := range m.UpstreamLog {
			l = e.Size()
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	if m.SuppressEnvoyHeaders {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRouter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRouter(x uint64) (n int) {
	return sovRouter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Router) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRouter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Router: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Router: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DynamicStats == nil {
				m.DynamicStats = &types.BoolValue{}
			}
			if err := m.DynamicStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartChildSpan", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StartChildSpan = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpstreamLog", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpstreamLog = append(m.UpstreamLog, &v2.AccessLog{})
			if err := m.UpstreamLog[len(m.UpstreamLog)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuppressEnvoyHeaders", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SuppressEnvoyHeaders = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRouter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRouter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRouter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRouter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRouter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRouter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRouter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRouter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRouter   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/router/v2/router.proto", fileDescriptor_router_822295bdbc4788d1)
}

var fileDescriptor_router_822295bdbc4788d1 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x26, 0x43, 0xb2, 0x29, 0x52, 0x44, 0xca, 0x0e, 0x65, 0xec, 0x54, 0x10, 0x92,
	0x51, 0xbd, 0x8b, 0x53, 0xc1, 0xc3, 0x10, 0xe9, 0xc0, 0x83, 0x97, 0x92, 0x75, 0x59, 0x56, 0xc8,
	0xfa, 0x42, 0x92, 0x56, 0xf7, 0x0d, 0x3d, 0xfa, 0x11, 0x64, 0xdf, 0xc2, 0x9b, 0x34, 0x59, 0xf5,
	0x32, 0xf0, 0xf6, 0xfe, 0xef, 0xff, 0xfe, 0xbf, 0x47, 0x5e, 0x30, 0xe5, 0x65, 0x0d, 0x5b, 0x9a,
	0x43, 0xb9, 0x2a, 0x04, 0x5d, 0x15, 0xd2, 0x72, 0x4d, 0xd7, 0xd6, 0x2a, 0xaa, 0xa1, 0x6a, 0xea,
	0x3a, 0xd9, 0x57, 0x44, 0x69, 0xb0, 0x10, 0x8c, 0x5d, 0x80, 0xf8, 0x00, 0xf1, 0x01, 0xd2, 0x04,
	0xc8, 0x7e, 0xac, 0x4e, 0x86, 0x93, 0x43, 0x50, 0x96, 0xe7, 0xdc, 0x18, 0x09, 0xa2, 0x41, 0xfe,
	0x0a, 0x4f, 0x1d, 0x46, 0x02, 0x40, 0x48, 0x4e, 0x9d, 0x5a, 0x54, 0x2b, 0xfa, 0xa6, 0x99, 0x52,
	0x5c, 0x1b, 0xef, 0x8f, 0xbf, 0x11, 0xee, 0xa5, 0x8e, 0x1f, 0xdc, 0xe0, 0x93, 0xe5, 0xb6, 0x64,
	0x9b, 0x22, 0xcf, 0x8c, 0x65, 0xd6, 0x84, 0x68, 0x84, 0xe2, 0x7e, 0x32, 0x24, 0x1e, 0x41, 0x5a,
	0x04, 0x99, 0x02, 0xc8, 0x17, 0x26, 0x2b, 0x9e, 0x0e, 0xf6, 0x81, 0x79, 0x33, 0x1f, 0xc4, 0xf8,
	0xcc, 0x58, 0xa6, 0x6d, 0x96, 0xaf, 0x0b, 0xb9, 0xcc, 0x8c, 0x62, 0x65, 0xd8, 0x19, 0xa1, 0xf8,
	0x38, 0x3d, 0x75, 0xfd, 0xbb, 0xa6, 0x3d, 0x57, 0xac, 0x0c, 0x9e, 0xf0, 0xa0, 0x52, 0xc6, 0x6a,
	0xce, 0x36, 0x99, 0x04, 0x11, 0x76, 0x47, 0xdd, 0xb8, 0x9f, 0x5c, 0x92, 0x43, 0x27, 0xf8, 0x7b,
	0x51, 0x9d, 0x90, 0x5b, 0x27, 0x66, 0x20, 0xd2, 0x7e, 0x0b, 0x98, 0x81, 0x08, 0xae, 0xf1, 0x85,
	0xa9, 0x94, 0xd2, 0xdc, 0x98, 0xcc, 0x31, 0xb2, 0x35, 0x67, 0x4b, 0xae, 0x4d, 0x78, 0xe4, 0xf6,
	0x9f, 0xb7, 0xee, 0x43, 0x63, 0x3e, 0x7a, 0x6f, 0x7a, 0xff, 0xb1, 0x8b, 0xd0, 0xe7, 0x2e, 0x42,
	0x5f, 0xbb, 0x08, 0xe1, 0x49, 0x01, 0x7e, 0xbf, 0xd2, 0xf0, 0xbe, 0x25, 0xff, 0xff, 0xc6, 0x33,
	0x7a, 0xed, 0xd4, 0xc9, 0xa2, 0xe7, 0xee, 0x72, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xf5,
	0x46, 0x07, 0xf1, 0x01, 0x00, 0x00,
}
