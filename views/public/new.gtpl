<!doctype html>
<html lang="ja">
 
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
 
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
 
    <title>New User register</title>
<link rel="stylesheet" href="./assets/css/style.css" type="text/css">
</head>
<body>
<div class="title">
<p class="display-1 text-center">Vulnerability Apps</p>
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

<style type="text/css">
input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
     -webkit-appearance: none;
     margin: 0;
}

input[type="number"] {
     -moz-appearance:textfield;
}
</style>

<script>
	function CheckPassword(confirm){
		var input1 = passwd.value;
		var input2 = confirm.value;
		if(input1 != input2){
			confirm.setCustomValidity("Password not matched!!");
		}else{
			confirm.setCustomValidity('');
		}
	}
</script>
