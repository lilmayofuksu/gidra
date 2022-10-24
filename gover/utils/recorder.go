package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/yezihack/colorlog"
)

type messagePack struct {
	ProtoName string  `json:"protoName"`
	Object    Message `json:"object"`
}

type Recorder struct {
	q *BytesQueue
	h *PacketHandler
	d []*messagePack
}

func (r *Recorder) Stop() {
	r.q.Close()
	content, err := json.Marshal(r.d)
	if err != nil {
		colorlog.Warn("marshaling records failed, err: %+v", err)
		return
	}
	fileName := fmt.Sprintf("log_%s_%d.json", time.Now().Format("2006_01_02_15_04_05"), rand.Int()%10000)
	os.WriteFile(fileName, content, 0644)
	colorlog.Info("log write to %s, with %d packet(s)", fileName, len(r.d))
}

func (r *Recorder) Record(packet []byte) {
	data := make([]byte, len(packet))
	copy(data, packet)
	r.q.Enqueue(data)
}

func (r *Recorder) Start() {
	go func() {
		parser := r.h
		for {
			data := r.q.Dequeue()
			if data == nil {
				break
			}
			// process data
			d, err := parser.Parse(data)
			if err != nil {
				colorlog.Warn("parse packet data failed")
				continue
			}
			if d.Body == nil {
				continue
			}
			pack := &messagePack{
				ProtoName: string(d.Body.ProtoReflect().Descriptor().Name()),
				Object:    d.Body,
			}
			colorlog.Info("Record proto: %s", pack.ProtoName)
			r.d = append(r.d, pack)
		}
		colorlog.Warn("recorder quit")
	}()
}

func NewRecorder(cap int, handler *PacketHandler) *Recorder {
	return &Recorder{
		q: NewBytesQueue(cap),
		h: handler,
	}
}
