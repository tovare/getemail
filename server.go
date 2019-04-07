package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gobuffalo/packr/v2"
)

// Secrets -- keys compiled into the application.
type Secrets struct {
	CaptchaAPIKey string `json:"captchakey"`
}

// API secrets initialized by main
var secrets Secrets

// CaptchaResponse -- structure returned by google
type CaptchaResponse struct {
	Success     bool      `json:"success"`
	ErrorCodes  []string  `json:"error-codes"` // On error
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
}

func store(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	token := r.PostFormValue("token")
	email := r.PostFormValue("email")
	okurl := r.PostFormValue("ok")
	failedurl := r.PostFormValue("failed")
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"response": {token}, "secret": {secrets.CaptchaAPIKey}})
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	var captcha CaptchaResponse
	json.Unmarshal(result, &captcha)
	if captcha.Success {
		http.Redirect(w, r, okurl, 301)
		log.Printf("%v, ok", email)
	} else {
		http.Redirect(w, r, failedurl, 301)
		log.Printf("%v, %v, key=%v", email,
			strings.Join(captcha.ErrorCodes, " "), secrets.CaptchaAPIKey)
	}
}

func main() {
	// Get secret API key from secrets file. This will be embedded
	// into the executable.
	box := packr.New("My Box", "./private")
	str, _ := box.Find("secrets.json")
	json.Unmarshal(str, &secrets)
	// Open the output file.
	fh, err := os.OpenFile("epost.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(fh)
	defer fh.Close()
	http.HandleFunc("/store", store)
	log.Println("Running on 8020")
	log.Fatal(http.ListenAndServe(":8020", nil))
}
