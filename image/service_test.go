package image

import (
	"fmt"
	"testing"

	"github.com/gorilla/mux"
)

func TestImageRoutes(t *testing.T) {
	LoadRouts(mux.NewRouter())
}

func TestRequestStatusType(t *testing.T) {

	if fmt.Sprintf("%v", StatusProcessed) != "Processed" {
		t.Fatal("Status not parsing as string")
	}

	var status RequestStatusType

	if status.IsSet() {
		t.Fatal("Unset status shouldn't be valid")
	}

	status = StatusProcessed

	if !status.IsSet() {
		t.Fatal("Set status shouldn't be invalid")
	}

}
