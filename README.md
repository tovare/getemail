# getemail

Small endpoint to get an email address and persist it to disk using the google Captcha v3 in the front end.

The store endpoint validates the email and persists to disk if ok.

The bot is redirected to denied.html and the user is redirected to thankyou.html

```
<form c id="emailform" method="POST" action="http://localhost:8080/store">
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