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
<p class="display-1 text-center">Hints</h1>
</div>

<link rel="stylesheet" href="./assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
<div class="whitebox">
<h1><u>レンダリング</u></h1>
<br>
<h2>golangの場合はレンダリングの処理は以下のようになる</h2>
<br>
<h2><pre>&lt;h1&gt;&#123;&#123;.UserName&#125;&#125;&lt;/h1&gt;</pre></h2>
<h2>&#123;&#123; &#125;&#125;に囲われた部分にレンダリングします</h2>
<br><br><br>
<h1><u>どこかで呼ばれているSQL文</u></h1>
<br>
<h2>"select post,created_at from vulnapp.posts where post like \"%" + searchWord + "%\";"</h2>
<br>
<h2>"select adminid from vulnapp.admins where mail=\"" + requestMail + "\" and passwd=\"" + requestPasswd + "\";"</h2>
</div>
</div>

</body>
</html>
