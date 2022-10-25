package kcp

/*
#cgo CFLAGS: -O3

#include "ikcp.h"
typedef int (*OutputCallback)(const char *buf, int len, struct IKCPCB *kcp, void *user);
typedef const char const_char;
extern int go_callback(const_char *buf, int len, struct IKCPCB *kcp, void *user);
*/
import "C"
import (
	"sync"
	"time"
	"unsafe"
)

type KCP struct {
	priv *C.struct_IKCPCB
	mu   *sync.Mutex
}
type KCPCallback = func(buf []byte, size int)

type UserInfo struct {
	cb KCPCallback
}

// monotonic reference time point
var refTime time.Time = time.Now()

const IkcpOVERHEAD = 28

// currentMs returns current elapsed monotonic milliseconds since program startup
func currentMs() uint32 { return uint32(time.Since(refTime) / time.Millisecond) }
func CurrentMs() uint32 { return currentMs() }

//export go_callback
func go_callback(buf *C.const_char, len C.int, kcp *C.struct_IKCPCB, arg unsafe.Pointer) C.int {
	userInfo := (ObjectId)(arg).Get().(*UserInfo)
	size := int(len)
	if userInfo.cb != nil {
		data := (*[1 << 31]byte)(unsafe.Pointer(buf))[:size:size]
		userInfo.cb(data, size)
	}
	return C.int(0)
}

func NewKCPWithToken(conv, token uint32, callback KCPCallback) (*KCP, error) {
	userInfo := NewObjectId(&UserInfo{cb: callback})
	h, err := C.ikcp_create(C.uint32_t(conv), C.uint32_t(token), unsafe.Pointer(userInfo))
	h.output = C.OutputCallback(C.go_callback)
	return &KCP{h, &sync.Mutex{}}, err
}

func (k *KCP) Free() {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.priv == nil {
		return
	}
	id := (ObjectId)(k.priv.user)
	(&id).Free()
	C.ikcp_release(k.priv)
	k.priv = nil
}

func (k *KCP) Input(data []byte) int {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.priv == nil {
		return 0
	}
	ptr := (*C.const_char)(unsafe.Pointer(&data[0]))
	size := C.long(len(data))
	ret := C.ikcp_input(k.priv, ptr, size)
	return int(ret)
}

func (k *KCP) Send(data []byte) int {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.priv == nil {
		return 0
	}
	ptr := (*C.const_char)(unsafe.Pointer(&data[0]))
	size := C.int(len(data))
	ret := C.ikcp_send(k.priv, ptr, size)
	return int(ret)
}

func (k *KCP) Recv(data []byte) int {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.priv == nil {
		return 0
	}
	ptr := (*C.const_char)(unsafe.Pointer(&data[0]))
	size := C.int(len(data))
	ret := C.ikcp_recv(k.priv, ptr, size)
	return int(ret)
}

func (k *KCP) Check(current uint32) uint32 {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.priv == nil {
		return 0
	}
	cur := C.uint32_t(current)
	ret := C.ikcp_check(k.priv, cur)
	return uint32(ret)
}

func (k *KCP) Update(current uint32) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.priv == nil {
		return
	}
	cur := C.uint32_t(current)
	C.ikcp_update(k.priv, cur)
}

func (k *KCP) WndSize(sndWnd, rcvWnd int) {
	C.ikcp_wndsize(k.priv, C.int32_t(sndWnd), C.int32_t(rcvWnd))
}

func (k *KCP) NoDelay(nodelay, interval, resend, nc int) {
	C.ikcp_nodelay(k.priv, C.int32_t(nodelay), C.int32_t(interval), C.int32_t(resend), C.int32_t(nc))
}

func (k *KCP) SetMtu(mtu int) {
	C.ikcp_setmtu(k.priv, C.int32_t(mtu))
}
