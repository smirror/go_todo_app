package auth

import (
	"bytes"
	"testing"
)

func TestEmbedded(t *testing.T) {
	want := []byte("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDgMsu3oTWSLZffs3yRfglFnrPIfx/aAGUE11VkqTutw smirror@LAPTOP-1NQ54DHH")
	if !bytes.Contains(rawPubkey, want) {
		t.Errorf("want %s, but got %s", want, rawPubkey)
	}

	want = []byte("-----BEGIN OPENSSH PRIVATE KEY-----")
	if !bytes.Contains(rawPrivkey, want) {
		t.Errorf("want %s, but got %s", want, rawPrivkey)
	}
}
