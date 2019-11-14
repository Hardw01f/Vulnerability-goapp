<!doctype html>
<html lang="ja">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <link rel="stylesheet" href="./assets/css/styleClear.css" type="text/css"> 

    <title>Hello, world!</title>
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
<p class="display-1 text-center">Profile</p>
</div>

</nav>
</header>
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="./assets/img/{{.Image}}" width="400" height="300">
    <h2>{{.Word}}</h2>
    {{ else }}
    <img src="./assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <div class="profileBox">
    <h1>Name : <t> <i>{{.UserName}}</i></h1>
    <h2>Age : <t> {{.Age}}</h2>
    <h2>Mail : <t> {{.Mail}}</h2>
    <h2>Address : <t> {{.Address}}</h2>
    <h2>Favorite Animal : <t> {{.Animal}}</h2>
    </div>
    <a href="/profile/edit"><button type="button" class="button">Edit</button></a>
    <a href="/profile/edit/image"><button type="button" class="button_changeImage">Change Image</button></a>
    <a href="/profile/changepasswd"><button type="button" class="button_changepasswd">Change Password</button></a>
</div>

</body>
</html>
