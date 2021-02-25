package sub2

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type TwoNumbers struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func Sub2(w http.ResponseWriter, r *http.Request) {
	var twoNumbers TwoNumbers

	err := json.NewEncoder(r.Body).Decode(&twoNumbers)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "A - B = %v", twoNumbers.A - twoNumbers.B)
}
