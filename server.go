package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gobuffalo/packr/v2"
)

// Secrets -- keys compiled into the application.
type Secrets struct {
	CaptchaAPIKey string `json:"captchakey"`
}

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
		log.Println(email)
	} else {
		http.Redirect(w, r, failedurl, 301)
	}
}

func main() {
	box := packr.New("My Box", "./private")
	str, _ := box.Find("secrets.json")
	json.Unmarshal(str, &secrets)

	http.HandleFunc("/store", store)
	log.Println("Running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
