package utils

import (
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	tokenReq := protoMap[GetPlayerTokenReq]
	fmt.Println(tokenReq, tokenReq.ProtoReflect().Descriptor().Name())
}
