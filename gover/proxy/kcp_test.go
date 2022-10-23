package proxy

import (
	"testing"

	"github.com/MoonlightPS/Iridium-gidra/gover/kcp_"
)

func TestKCP(t *testing.T) {
	kobj := kcp_.NewKCPWithToken(1, 2, func(buf []byte, size int) {})
	kobj.SetMtu(1200)
	kobj.WndSize(1024, 1024)
	kobj.NoDelay(1, 10, 2, 1)
}
