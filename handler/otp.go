package handler

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/kyong0612/otp-getter/templates/html"
)

func renderOTPReader(w http.ResponseWriter, result string) error {
	t, err := template.ParseFS(html.File, "otp.gohtml")
	if err != nil {
		return err
	}

	data := struct {
		Result string
	}{
		Result: result,
	}

	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func GetReadOtpPage(w http.ResponseWriter, r *http.Request) {
	err := renderOTPReader(w, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ReadOtpHandler(w http.ResponseWriter, r *http.Request) {
	f, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	fmt.Println(fileHeader.Filename)

	fileBody := make([]byte, fileHeader.Size)
	_, err = f.Read(fileBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp, err := os.CreateTemp("", "temp")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer temp.Close()

	err = os.WriteFile(temp.Name(), fileBody, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(temp.Name())

	result, err := exec.Command("zbarimg", temp.Name()).Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = renderOTPReader(w, string(result))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
