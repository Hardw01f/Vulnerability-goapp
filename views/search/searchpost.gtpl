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
<p class="display-1 text-center">TimeLine</p>
</div>

</nav>
</header>
<link rel="stylesheet" href="../assets/css/post.css" type="text/css"> 
<body>

<div class="box">
    <div class="box11">
    {{range .StringResults}}
    <div class="box22">
    <h2>{{.}}</h2>
    <br>
    </div>
    {{end}}
    </div>
    <div class="postform">
        <form action="/timeline" method="post">
        <textarea name="post" rows="10" cols="20" wrap="hard" required></textarea>
    </div>
        <input type="submit" value="Post" class="button_postform">
        </form>
</div>

</body>
</html>
