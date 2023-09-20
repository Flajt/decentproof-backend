package decentproof_backend

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func HandleWithCors(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%q", dump)

	// Sets the response headers to allow CORS requests.

	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)

	_, err = io.WriteString(w, "This function is allowing most CORS requests")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
}
