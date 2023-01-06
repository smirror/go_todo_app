package auth

import (
	"bytes"
	"golang.org/x/net/context"
	"testing"
	"todo_app/clock"
	"todo_app/entity"
	"todo_app/testutil/fixture"
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

func TestJWTer_GenerateToken(t *testing.T) {
	ctx := context.Background()
	moq := &StoreMock{}
	wantID := entity.UserID(20)
	u := fixture.User(&entity.User{ID: wantID})
	moq.SaveFunc = func(ctx context.Context, key string, userID entity.UserID) error {
		if userID != wantID {
			t.Errorf("want %d, but got %d", wantID, userID)
		}
		return nil
	}

	sut, err := NewJWTer(moq, clock.RealClocker{})
	if err != nil {
		t.Fatal(err)
	}

	got, err := sut.GenerateToken(ctx, *u)
	if err != nil {
		t.Fatal("not want err:%v", err)
	}

	if len(got) == 0 {
		t.Error("want token, but got empty")
	}
}
