package notification

import "testing"

func TestSendingNotification(t *testing.T) {
	got := SendNotification("Elvis Salaton", "254719158559")
	want := "Hello there Elvis Salaton Your order has been received"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
