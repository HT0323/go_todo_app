package auth

import (
	"bytes"
	"context"
	"testing"

	"github.com/HT0323/go_todo_app/clock"
	"github.com/HT0323/go_todo_app/entity"
	"github.com/HT0323/go_todo_app/testutil/fixture"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but go %s", want, rawPubKey)
	}

	want = []byte("-----BEGIN RSA PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but go %s", want, rawPrivKey)
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
		t.Fatalf("now want err: %v", err)
	}
	if len(got) == 0 {
		t.Errorf("token is empty")
	}
}
