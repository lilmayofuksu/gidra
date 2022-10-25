package utils

import (
	"encoding/binary"
	"errors"

	"github.com/MoonlightPS/Iridium-gidra/gover/gen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var be = binary.BigEndian

type Message = protoreflect.ProtoMessage

type PacketHandler struct {
	key *PacketKey
}

type Packet struct {
	Cmd    int
	Header *gen.PacketHead
	Body   Message
}

func GetCMDDescription(cmd int) string {
	if v, ok := protoMap[cmd]; ok {
		m := v.ProtoReflect().Descriptor().Name()
		return string(m)
	}
	return "Unknown"
}

func (h *PacketHandler) Decode(cmd int, body []byte) (Message, error) {
	if v, ok := protoMap[cmd]; ok {
		m := v.ProtoReflect().New().Interface()
		err := proto.Unmarshal(body, m)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	return nil, errors.New("decode proto packet failed")
}

// only try to parse cmd to ensure proxy performance
func (h *PacketHandler) ParseCmd(data []byte) (int, error) {
	if len(data) < 4 {
		return 0, errors.New("data len error")
	}
	c := make([]byte, 4)
	copy(c, data)
	h.key.Xor(c)
	if be.Uint16(c) != 0x4567 {
		return 0, errors.New("decrypt data failed")
	}
	return int(be.Uint16(c[2:4])), nil
}

func (h *PacketHandler) Parse(data []byte) (*Packet, error) {
	h.key.Xor(data)
	if be.Uint16(data) != 0x4567 {
		return nil, errors.New("decrypt data failed")
	}
	if be.Uint16(data[len(data)-2:]) != 0x89AB {
		return nil, errors.New("decrypt data failed")
	}
	cmd := int(be.Uint16(data[2:4]))
	headerLen := be.Uint16(data[4:6])
	bodyLen := be.Uint32(data[6:10])

	header := &gen.PacketHead{}
	err := proto.Unmarshal(data[10:10+headerLen], header)
	if err != nil {
		return nil, err
	}
	if bodyLen == 0 {
		return &Packet{
			Cmd:    cmd,
			Header: header,
			Body:   nil,
		}, nil
	}

	offset := uint32(headerLen + 10)
	body, err := h.Decode(cmd, data[offset:offset+bodyLen])
	if err != nil {
		return nil, err
	}

	return &Packet{
		Cmd:    cmd,
		Header: header,
		Body:   body,
	}, nil
}

func (h *PacketHandler) Compose(packet *Packet) ([]byte, error) {
	header, err := proto.Marshal(packet.Header)
	if err != nil {
		return nil, err
	}

	var body []byte
	if packet.Body != nil {
		body, err = proto.Marshal(packet.Body)
		if err != nil {
			return nil, err
		}
	}

	data := make([]byte, 12+len(header)+len(body))

	be.PutUint16(data, 0x4567)
	be.PutUint16(data[2:4], uint16(packet.Cmd))
	be.PutUint16(data[4:6], uint16(len(header)))
	be.PutUint32(data[6:10], uint32(len(body)))

	copy(data[10:10+len(header)], header)

	if len(body) > 0 {
		offset := len(header) + 10
		copy(data[offset:offset+len(body)], body)
	}

	be.PutUint16(data[len(data)-2:], 0x89AB)

	h.key.Xor(data)

	return data, nil
}

func (h *PacketHandler) SetKey(key *PacketKey) {
	h.key = key
}

func NewPacketHandler() *PacketHandler {
	return &PacketHandler{}
}
