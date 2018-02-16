package prime

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestDefaultPrimeCalc(t *testing.T) {
	LoadRouts(mux.NewRouter())
	r := getPrimeResp(defaultMaxNumber)

	if defaultMaxNumber != r.Max {
		t.Fatalf("Max value not equal to max constant (%d) %d ",
			defaultMaxNumber, r.Max)
	}

}

func TestArcPrimeCalc(t *testing.T) {
	LoadRouts(mux.NewRouter())

	n := 9000

	r := getPrimeResp(n)

	if n != r.Max {
		t.Fatalf("Max value not equal to max constant (%d) %d ",
			n, r.Max)
	}

	if r.Value > r.Max {
		t.Fatalf("Calculated prime higher than the max (%d) %d ",
			r.Max, r.Value)
	}
}
