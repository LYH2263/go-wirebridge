package wirebridge_test

import (
	"encoding/binary"
	"errors"
	"testing"

	"example.com/wirebridge"
)

func TestBug04_NilHandlerNoPanic(t *testing.T) {
	b := wirebridge.New()
	defer b.Close()
	// 仅有路由元数据，未注册 Handler
	if err := b.ApplyRoutes([]wirebridge.RouteMeta{{
		Opcode: 99, Name: "ghost", Enabled: true, Tags: []string{"t"},
	}}); err != nil {
		t.Fatal(err)
	}
	raw := make([]byte, 7)
	binary.BigEndian.PutUint32(raw[0:4], 3)
	binary.BigEndian.PutUint16(raw[5:7], 99)
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ServeFrame with nil handler panicked: %v", r)
		}
	}()
	_, err := b.ServeFrame(raw)
	if !errors.Is(err, wirebridge.ErrNilHandler) {
		t.Fatalf("want ErrNilHandler, got %v", err)
	}
}
