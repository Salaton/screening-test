package tests

import (
	"testing"

	auth "github.com/Salaton/screening-test.git/auth"
	notification "github.com/Salaton/screening-test.git/notification"
	db "github.com/Salaton/screening-test.git/postgres"
)

func TestHashPassword(t *testing.T) {

	// t.Run("test password hash", func(t *testing.T) {
	// 	password := "password12345"
	// 	got := db.HashPassword(password)
	// 	want := db.HashPassword(password)

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}
	// })

	t.Run("compare password and hash", func(t *testing.T) {
		password := "password12345"
		hash := "$2a$14$fNPi4m0o8ooKCUYS4TlU3erQst453fiF.QvtFyKu2EtJGLDPG4kLG"
		got := db.CheckPasswordHash(password, hash)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

}

func TestToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlNhbGF0In0.gvvpBiHU_KRXEw5i5VNpj-bBGweQzQmUDjQ_5bt-KEs"

	t.Run("Creating a new token", func(t *testing.T) {
		got, _ := auth.CreateNewToken("Salat")
		want := tokenString
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Parse the token", func(t *testing.T) {
		got, _ := auth.ParseToken(tokenString)
		want := "Salat"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func TestSendingNotification(t *testing.T) {
	got := notification.SendNotification("Elvis Salaton", "254719158559")
	want := "Hello there Elvis Salaton Your order has been received"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
