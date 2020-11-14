package api

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Add(time.Hour * 24)
	currentTime := now.Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}

func TestMe(a int) int {
	return a * a
}
