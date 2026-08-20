package frame

import (
	"encoding/binary"
	"testing"
)

func TestBug01_DecodePayloadSliceAlias(t *testing.T) {
	payload := []byte("hello-wire")
	bodyLen := 3 + len(payload)
	raw := make([]byte, 4+bodyLen)
	binary.BigEndian.PutUint32(raw[0:4], uint32(bodyLen))
	raw[4] = 0
	binary.BigEndian.PutUint16(raw[5:7], 2)
	copy(raw[7:], payload)

	f, err := Decode(raw, 1<<20, nil)
	if err != nil {
		t.Fatal(err)
	}
	// 污染底层读缓冲
	raw[7] = 'X'
	if f.Payload[0] == 'X' {
		t.Fatal("Decode payload aliases underlying read buffer")
	}
	if string(f.Payload) != "hello-wire" {
		t.Fatalf("payload corrupted: %q", f.Payload)
	}
}
