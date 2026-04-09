package websocket

import (
	"testing"
)

func TestBuildFrame_Binary(t *testing.T) {
	data := []byte("hello")
	frame := BuildFrame(OpBinary, data, true)

	if len(frame) < 2 {
		t.Fatalf("frame too short: %d bytes", len(frame))
	}

	// First byte: FIN (0x80) | OpBinary (0x02) = 0x82
	if frame[0] != 0x82 {
		t.Errorf("expected opcode 0x82, got 0x%02x", frame[0])
	}

	// Second byte: MASK (0x80) | length (5) = 0x85
	if frame[1] != 0x85 {
		t.Errorf("expected masked length byte 0x85, got 0x%02x", frame[1])
	}
}

func TestBuildFrame_Close(t *testing.T) {
	frame := BuildFrame(OpClose, []byte{}, true)

	if len(frame) < 2 {
		t.Fatalf("close frame too short: %d bytes", len(frame))
	}

	// First byte: FIN (0x80) | OpClose (0x08) = 0x88
	if frame[0] != 0x88 {
		t.Errorf("expected close opcode 0x88, got 0x%02x", frame[0])
	}
}

func TestBuildFrame_Ping(t *testing.T) {
	frame := BuildFrame(OpPing, []byte{}, true)

	if frame[0] != 0x89 {
		t.Errorf("expected ping opcode 0x89, got 0x%02x", frame[0])
	}
}

func TestBuildFrame_Pong(t *testing.T) {
	frame := BuildFrame(OpPong, []byte{}, true)

	if frame[0] != 0x8A {
		t.Errorf("expected pong opcode 0x8A, got 0x%02x", frame[0])
	}
}

func TestBuildFrame_Unmasked(t *testing.T) {
	data := []byte("test")
	frame := BuildFrame(OpBinary, data, false)

	// Second byte should NOT have MASK bit set
	if frame[1]&0x80 != 0 {
		t.Error("expected unmasked frame")
	}
}

func TestXORMask(t *testing.T) {
	data := []byte{0x00, 0x11, 0x22, 0x33, 0x44}
	mask := []byte{0xFF, 0xFF, 0xFF, 0xFF}

	result := XORMask(data, mask)

	if result[0] != 0xFF || result[1] != 0xEE || result[2] != 0xDD {
		t.Errorf("unexpected XOR result: %v", result)
	}

	// XORing again should restore original
	restored := XORMask(result, mask)
	if restored[0] != 0x00 || restored[1] != 0x11 || restored[2] != 0x22 {
		t.Errorf("failed to restore original: %v", restored)
	}
}

func TestDefaultPingInterval(t *testing.T) {
	if DefaultPingInterval <= 0 {
		t.Error("DefaultPingInterval should be positive")
	}
}
