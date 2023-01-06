package auth

import (
	"bytes"
	"testing"
)

func TestEmbedded(t *testing.T) {
	want := []byte("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDgMsu3oTWSLZffs3yRfglFnrPIfx/aAGUE11VkqTutw smirror@LAPTOP-1NQ54DHH")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}

	want = []byte("-----BEGIN OPENSSH PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rawPrivKey)
	}
}
