package ping

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestPing(t *testing.T) {
	LoadRouts(mux.NewRouter())
	if stateInstance == nil || stateInstance.Host == "" ||
		stateInstance.ID == "" || stateInstance.Ts == "" {
		t.Fatal("Uninitialized state")
	}
}
