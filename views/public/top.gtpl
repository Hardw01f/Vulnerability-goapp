<!doctype html>
<html lang="ja">
 
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
 
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <link rel="stylesheet" href="./assets/css/style.css" type="text/css"> 
 
    <title>Top Page</title>
 
<ul id="nav">
  <li><div class="username">{{.UserName}}</div></li>
  <li><a href="/top">Home</a></li>
  <li><a href="/profile">Profile</a></li>
  <li><a href="/timeline">TimeLine</a></li>
  <li><a href="/post">Post</a></li>
  <li><a href="/hints">Hints</a></li>
  <li><a href="/db">DB</a></li>
  <li><a href="/logout">Logout</a></li>
<ul>
</head>


<div id="header_title">
<p class="display-1 text-center">Vulnerability Apps</p>
</div>

<body>


<div class="box">
<div class="cp_cont">
<div class="cp_anime14">
	<div class="text1">Good morning , Mr.{{.UserName}} !</div>
	<div class="text2">I will tell about this mission</div>
	<div class="text2">This Web application was attacked and tampered to defective app by Rookie Hacker, HardWolf</div>
	<div class="text2">Therefore, We need to investigate it <br>for prevent more damage</div>
	<div class="text2">Your mission is investigating and detect <br>this Web Application vulnerabilities</div>
	<div class="text2">I pray for your consideration</div>
	<div class="text2">This message will not be eliminate automatically.</div>
</div>
</div>
</div>

</body>
</html>

