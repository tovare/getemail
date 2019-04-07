# getemail

Small endpoint to get an email address and persist it to disk using the google Captcha v3 in the front end.

The store endpoint validates the email and persists to disk if ok.

The bot is redirected to denied.html and the user is redirected to thankyou.html

```html
<script src="https://www.google.com/recaptcha/api.js?render=INSERT-API-KEY-HEREonload=onloadCallback"></script>

<form id="emailform" method="POST" action="http://localhost:8080/store">
    <!-- Populated by Captcha v3 -->
    <input type="hidden" id="g-recaptcha-response" name="token"  required>
    <input type="hidden" name="ok" value="https://sportybiz.eu/thankyou.html">
    <input type="hidden" name="failed" value="https://sportybiz.eu/denied.html">
    <input id="email" name="email" type="email" class="validate" required>
    <button class="g-recaptcha btn waves-effect waves-light" 
        type="submit" name="action">Submit
    </button>
</form>
```

```javascript
var onloadCallback = function() {
    grecaptcha.ready(function() {
        console.log("READY")
        grecaptcha.execute('INSERT-API-KEY-HERE', 
            {action: 'homepage'})
            .then(function(token) {
                // add token value to form
                document.getElementById('g-recaptcha-response').value = token;
        });
    });
};
```

## Files not in the repository

To compile you need to add a directory called private which is 
ignored by git containing the file called secrets.json containing

```json
{
    "captchakey" : "<INSERT SECRET API KEY HERE"
}
```

## Compilation

The code is built using pakr which includes the secrets file resulting
in no dependencies outside the binary.

```bash
GOOS=darwin GOARCH=amd64 packr build && mv ./getemail ./releases/darwin-getemail \
  && GOOS=linux GOARCH=amd64 packr build && mv ./getemail ./releases/linux-getemail \
  && GOOS=windows GOARCH=386 packr build && mv ./getemail.exe ./releases/getemail.exe \
  && packr clean
```

## Running

The user running the server needs write access to the current directory
