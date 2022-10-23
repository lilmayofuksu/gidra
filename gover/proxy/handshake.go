package proxy

import (
	"encoding/binary"
	"errors"
)

var le = binary.LittleEndian

const (
	HANDSHAKE_SYN = 0xff
	HANDSHAKE_ACK = 0x145
	HANDSHAKE_FIN = 0x194
)

var ErrInvalidHandshakePacket = errors.New("invalid handshake packet")

type Handshake struct {
	m1    uint32
	conv  uint32
	token uint32
	enet  uint32
	m2    uint32
	lid   uint64
}

func (h *Handshake) Compose() []byte {
	data := make([]byte, 20)
	be.PutUint32(data, h.m1)
	be.PutUint32(data[4:], h.conv)
	be.PutUint32(data[8:], h.token)
	be.PutUint32(data[12:], h.enet)
	be.PutUint32(data[16:], h.m2)
	return data
}

func ParseHandshakePacket(data []byte) (*Handshake, error) {
	if len(data) != 20 {
		return nil, ErrInvalidHandshakePacket
	}
	m1, conv, token, enet, m2 := be.Uint32(data), be.Uint32(data[4:]), be.Uint32(data[8:]), be.Uint32(data[12:]), be.Uint32(data[16:])
	switch m1 {
	case HANDSHAKE_SYN:
		if m2 != 0xffffffff || enet != 1234567890 {
			return nil, ErrInvalidHandshakePacket
		}
	case HANDSHAKE_ACK:
		if m2 != 0x14514545 || enet != 1234567890 {
			return nil, ErrInvalidHandshakePacket
		}
	case HANDSHAKE_FIN:
		if m2 != 0x19419494 || (enet != 1 && enet != 0) {
			return nil, ErrInvalidHandshakePacket
		}
	default:
		return nil, ErrInvalidHandshakePacket
	}
	lid := be.Uint64(data[4:])
	return &Handshake{m1, conv, token, enet, m2, lid}, nil
}

// fast check if handshake packet
func IsHandshakePacket(data []byte) bool {
	if len(data) != 20 {
		return false
	}
	switch be.Uint32(data) {
	case HANDSHAKE_SYN:
		return true
	case HANDSHAKE_ACK:
		return true
	case HANDSHAKE_FIN:
		return true
	}
	return false
}

func ToLID(conv, token uint32) uint64 { return (uint64(conv) << 32) | uint64(token) }

func NewHandshakePacket(m1, conv, token, enet, m2 uint32) *Handshake {
	return &Handshake{m1, conv, token, enet, m2, ToLID(conv, token)}
}
