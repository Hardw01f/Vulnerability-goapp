<!doctype html>
<html lang="ja">
 
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
 
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="./assets/css/style.css" type="text/css">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
 
    <title>Hello, world!</title>
<link rel="stylesheet" href="./assets/css/style.css" type="text/css">
<title></title>
</head>
<body>
<div class="title">
<p class="display-1">Vulnerability Apps</p>
</div>
<div class="login_form">
<h2>Sign in</h2>
<form action="/login" method="post">
    <p>Mail-Address:<input type="email" name="mail" required></p>
    <p>Password:<input type="password" name="passwd" required></p>
    <input type="submit" value="Login">
</form>
</div>
<div class="register_form">
<h2>New User Register</h2>
<form action="/new" method="post">
    <p>Mail-Address:<input type="email" name="mail" required></p>
    <p>Name:<input type="text" name="name" onsubmit="return validateForm()" required></p>
    <p>Age:<input type="number" name="age" min="0" max="125" pattern=" 0+\.[0-9]*[1-9][0-9]*$" required></p>
    <br>
    <p>Password:<input type="password" name="passwd" id="passwd" required></p>
    <p>Confirm Password:<input type="password" name="confirm" oninput="CheckPassword(this)" required></p>
    <input type="submit" value="Register">
</form>
</div>
</body>
</html>
