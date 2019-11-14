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
<p class="display-1 text-center">User Image</p>
</div>

</nav>
</header>
<link rel="stylesheet" href="../../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="showUpdateImage">
    {{ if .Image }}
    <img src="../../assets/img/{{.Image}}" width="500" height="400">
    {{ else }}
    <img src="../../assets/img/noimage.png" width="500" height="400">
    {{end}}
    </div>
    <form enctype="multipart/form-data" action="/profile/edit/upload" method="post">
  <input type="file" name="uploadfile" multiple="multiple" />
  <input type="submit" class="button_changeImage" value="Change Image" />
</div>

</body>
</html>
