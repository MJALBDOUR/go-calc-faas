package div2

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

	if B == 0.0 {
		http.Error(w, "Division by zero", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "A / B = %v", twoNumbers.A / twoNumbers.B)
}
