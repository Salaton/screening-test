package auth

import "testing"

func TestToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbGFAZ21haWwuY29tIn0.L7YVb0-cTpAPsBR28voYorSCx7I_Bp663bpcsgbX99M"

	t.Run("Creating a new token", func(t *testing.T) {
		got, _ := CreateNewToken("sala@gmail.com")
		want := tokenString
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Parse the token", func(t *testing.T) {
		got, _ := ParseToken(tokenString)
		want := "sala@gmail.com"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
