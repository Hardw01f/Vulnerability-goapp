<!doctype html>
<html lang="ja">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

    <title>Admin Login Page</title>
<link rel="stylesheet" href="./assets/css/style.css" type="text/css">
</head>
<body>
<div class="title">
<p class="display-1 text-center">Admin User SignIn</p>
</div>
<div class="login_form">
<h2>Admin Sign in</h2>
<form action="/adminconfirm" method="post">
    <p>Mail-Address:<input type="text" name="adminmail" required></p>
    <p>Password:<input type="password" name="adminpasswd" required></p>
    <input type="submit" value="Admin Login">
</form>
</div>
</body>
</html>
