package proxy

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/MoonlightPS/Iridium-gidra/gover/kcp"
)

func BenchmarkSyncMap(b *testing.B) {
	m := &sync.Map{}
	for i := 0; i < 10000; i++ {
		k, err := kcp.NewKCPWithToken(1, 2, func(buf []byte, size int) {})
		if err != nil {
			panic(err)
		}
		m.Store(i, k)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Range(func(key, value interface{}) bool {
			obj := value.(*kcp.KCP)
			obj.Input(make([]byte, 1024))
			return true
		})
	}
}

func TestUDP(t *testing.T) {
	l, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		b := make([]byte, 1024)
		for {
			n, a, e := l.ReadFrom(b)
			fmt.Println("recv", n, a, e)
			l.WriteTo(b[:n], a)
		}
	}()

	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 9981}
	c1, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	dstAddr = &net.UDPAddr{IP: sip, Port: 9981}
	c2, err := net.DialUDP("udp", nil, dstAddr)
	if err != nil {
		fmt.Println(err)
	}

	n, err := c1.Write([]byte("Halo world"))
	fmt.Println("write", n, err)
	n, err = c2.Write([]byte("Halo world1"))
	fmt.Println("write", n, err)

	b := make([]byte, 1024)
	n, err = c1.Read(b)
	fmt.Println(n, err)
	n, err = c2.Read(b)
	fmt.Println(n, err)

	time.Sleep(time.Second * 2)
}
