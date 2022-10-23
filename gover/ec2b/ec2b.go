package ec2b

import (
	"encoding/binary"
	"errors"

	"github.com/MoonlightPS/Iridium-gidra/gover/utils"
)

var le = binary.LittleEndian

// UnityPlayer:$26EA90
func key_scramble(key []byte) {
	round_keys := make([]byte, 11*16)
	for round := 0; round <= 10; round++ {
		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				idx := (round << 8) + (i * 16) + j
				round_keys[round*16+i] ^= aes_xorpad_table[1][idx] ^ aes_xorpad_table[0][idx]
			}
		}
	}
	var chip [16]byte
	oqs_mhy128_enc_c(key, round_keys, chip[:])
	copy(key, chip[:])
}

// UnityPlayer:$19DA40
func get_decrypt_vector(key, crypt []byte) ([]byte, error) {
	var val uint64 = 0xFFFFFFFFFFFFFFFF
	for i := 0; i < len(crypt)>>3; i++ {
		val = le.Uint64(crypt[i<<3:]) ^ val
	}

	key_hi, key_lo := le.Uint64(key), le.Uint64(key[8:])
	mt := utils.NewMT19937_64()
	mt.Seed(key_lo ^ 0xCEAC3B5A867837AC ^ val ^ key_hi)

	out := make([]byte, 0, utils.DEFAULT_KEY_LEN)
	for i := 0; i < utils.DEFAULT_KEY_LEN; i += 8 {
		out = le.AppendUint64(out, mt.Int64())
	}
	return out, nil
}

func Derive(ec2b []byte) ([]byte, error) {
	if len(ec2b) != 2076 {
		return nil, errors.New("ec2b size must be 2076")
	}

	var key [16]byte
	var data [2048]byte

	copy(key[:], ec2b[8:])
	copy(data[:], ec2b[28:])

	key_scramble(key[:])

	for i := 0; i < 16; i++ {
		key[i] ^= key_xorpad_table[i]
	}

	return get_decrypt_vector(key[:], data[:])
}
