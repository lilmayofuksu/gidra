package proxy

import (
	"testing"

	"github.com/MoonlightPS/Iridium-gidra/gover/kcp"
)

func TestKCP(t *testing.T) {
	kobj, err := kcp.NewKCPWithToken(1, 2, func(buf []byte, size int) {})
	if err != nil {
		panic(err)
	}
	kobj.SetMtu(1200)
	kobj.WndSize(1024, 1024)
	kobj.NoDelay(1, 10, 2, 1)
}
