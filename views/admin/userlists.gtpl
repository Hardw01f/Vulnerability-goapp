<!doctype html>
<html lang="ja">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

    <title>Admin Page</title>
<ul id="nav">
  <li><div class="username">Admin</div></li>
  <li><a href="/top">Return Home</a></li>
  <ul>
</head>
<div id="header_title">
<p class="display-1 text-center">Admin Page</p>
</div>

</nav>
</header>
<link rel="stylesheet" href="./assets/css/style.css" type="text/css"> 
<body>

<div class="box">
<p>Your uid : {{ .Uid }}</p>
<p>Welcom admin page, {{ .UserName }}</p>
<br>
<p>================================================</p>
<br>
{{ range .UserLists }}
<p>{{ . }}</p>
<br>
{{ end }}
</div>

</body>
</html>

