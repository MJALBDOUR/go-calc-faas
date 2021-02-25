package add2

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type TwoNumbers struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func Add2(w http.ResponseWriter, r *http.Request) {
	var twoNumbers TwoNumbers

	err := json.NewDecoder(r.Body).Decode(&twoNumbers)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v", twoNumbers.A + twoNumbers.B)
}
