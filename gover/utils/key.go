package utils

const DEFAULT_KEY_LEN = 4096

type PacketKey struct {
	key []byte
	len int
}

func (k *PacketKey) SetKey(key []byte) {
	k.key = key
	k.len = len(key)
}

func (k *PacketKey) GenKey(seed uint64) {
	mt := NewMT19937_64()
	mt.Seed(seed)
	mt.Seed(mt.Int64())
	mt.Int64()

	key := make([]byte, 0, DEFAULT_KEY_LEN)
	for i := 0; i < DEFAULT_KEY_LEN; i += 8 {
		key = be.AppendUint64(key, mt.Int64())
	}
	k.key = key
	k.len = DEFAULT_KEY_LEN
}

// In-place xor
func (k *PacketKey) Xor(data []byte) {
	for i := range data {
		data[i] ^= k.key[i%k.len]
	}
}

func (k *PacketKey) Duplicate() *PacketKey {
	key := make([]byte, len(k.key))
	copy(key, k.key)
	return &PacketKey{
		key: key,
		len: k.len,
	}
}

func NewPacketKey() *PacketKey {
	return &PacketKey{}
}
