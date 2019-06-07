package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Redirect, Access-Control-Allow-Origin, Referer, Origin, User-Agent, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// https://stackoverflow.com/questions/22972066/how-to-handle-preflight-cors-requests-on-a-go-server
	if r.Method == "OPTIONS" {
		// handle preflight in here
		w.WriteHeader(http.StatusOK)
		return
	}

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(requestDump))
	}
}
