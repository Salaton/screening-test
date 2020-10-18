package tests

import (
	"testing"

	auth "github.com/Salaton/screening-test/auth"
	notification "github.com/Salaton/screening-test/notification"
	db "github.com/Salaton/screening-test/postgres"
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
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbGFAZ21haWwuY29tIn0.L7YVb0-cTpAPsBR28voYorSCx7I_Bp663bpcsgbX99M"

	t.Run("Creating a new token", func(t *testing.T) {
		got, _ := auth.CreateNewToken("sala@gmail.com")
		want := tokenString
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Parse the token", func(t *testing.T) {
		got, _ := auth.ParseToken(tokenString)
		want := "sala@gmail.com"
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
