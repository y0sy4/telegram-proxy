package config

import (
	"testing"
)

func TestParseDCIPList_Basic(t *testing.T) {
	input := []string{"2:149.154.167.220", "4:149.154.167.220"}
	regular, media, err := ParseDCIPList(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if regular[2] != "149.154.167.220" {
		t.Errorf("expected DC2=149.154.167.220, got %s", regular[2])
	}
	if regular[4] != "149.154.167.220" {
		t.Errorf("expected DC4=149.154.167.220, got %s", regular[4])
	}
	// Without 'm' suffix, media should mirror regular
	if media[2] != "149.154.167.220" {
		t.Errorf("expected media DC2=149.154.167.220, got %s", media[2])
	}
}

func TestParseDCIPList_MediaSpecific(t *testing.T) {
	input := []string{"2:149.154.167.220", "2m:149.154.167.222", "4:149.154.167.91"}
	regular, media, err := ParseDCIPList(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// DC2 regular
	if regular[2] != "149.154.167.220" {
		t.Errorf("expected regular DC2=149.154.167.220, got %s", regular[2])
	}
	// DC2 media (overridden by 2m:)
	if media[2] != "149.154.167.222" {
		t.Errorf("expected media DC2=149.154.167.222, got %s", media[2])
	}
	// DC4 (no media override)
	if regular[4] != "149.154.167.91" || media[4] != "149.154.167.91" {
		t.Errorf("expected DC4 same for regular and media")
	}
}

func TestParseDCIPList_InvalidFormat(t *testing.T) {
	input := []string{"invalid"}
	_, _, err := ParseDCIPList(input)
	if err == nil {
		t.Error("expected error for invalid format")
	}

	input = []string{"2:not-an-ip"}
	_, _, err = ParseDCIPList(input)
	if err == nil {
		t.Error("expected error for invalid IP")
	}

	input = []string{"abc:149.154.167.220"}
	_, _, err = ParseDCIPList(input)
	if err == nil {
		t.Error("expected error for non-numeric DC")
	}
}

func TestParseDCIPList_Empty(t *testing.T) {
	regular, media, err := ParseDCIPList([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(regular) != 0 || len(media) != 0 {
		t.Error("expected empty maps for empty input")
	}
}
