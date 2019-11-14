<!doctype html>
<html lang="ja">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

<ul id="nav">
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
<p class="display-1 text-center">Profile Edit</p>
</div>

</nav>
</header>
<link rel="stylesheet" href="../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="../assets/img/{{.Image}}" width="400" height="300">
    {{ else }}
    <img src="../assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <div class="profileBox">
    <form action="/profile/edit/confirm" method="post">
    <h1>Name : <t> <input type="text" name="username" value="{{.UserName}}"> </h1>
    <h2>Age : <t> <input type="text" name="age" value="{{.Age}}"> </h2>
    <h2>Mail : <t> <input type="text" name="mail" value="{{.Mail}}" disabled="disable"> </h2>
    <h2>Address : <t> <input type="text" name="address" value="{{.Address}}"></h2>
    <h2>Favorite Animal : <t> <input type="text" name="animal" value="{{.Animal}}"> </h2>
    <h2>Word : <t> <input type="text" name="word" value="{{.Word}}"> </h2>
    </div>
    <input type="submit" class="button_edit" value="Confirm">
</div>

</body>
</html>
