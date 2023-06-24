package handler

import "net/http"

func GetReadOtpPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sup"))
}

func ReadOtpHandler(w http.ResponseWriter, r *http.Request) {

}
