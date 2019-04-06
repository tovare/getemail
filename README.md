# getemail

Small endpoint to get an email address and persist it to disk using the google Captcha v3 in the front end.

Work in progress.

 curl -X POST -F "email=mail@tovare.com"  http://127.0.0.1:8080/store

 curl -X POST -F "token=" http://127.0.0.1:8080/verify
