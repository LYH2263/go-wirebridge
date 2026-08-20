package wirebridge_test

import (
	"encoding/binary"
	"errors"
	"testing"

	"example.com/wirebridge"
)

func TestBug05_FrameTooLargeWrapsSentinel(t *testing.T) {
	b := wirebridge.New(wirebridge.WithMaxFrame(32))
	defer b.Close()
	payload := make([]byte, 64)
	bodyLen := 3 + len(payload)
	raw := make([]byte, 4+bodyLen)
	binary.BigEndian.PutUint32(raw[0:4], uint32(bodyLen))
	binary.BigEndian.PutUint16(raw[5:7], 1)
	copy(raw[7:], payload)
	_, err := b.Decode(raw)
	if err == nil {
		t.Fatal("expected too-large error")
	}
	if !errors.Is(err, wirebridge.ErrTooLarge) {
		t.Fatalf("want errors.Is ErrTooLarge, got %v", err)
	}
}
