<!doctype html>
<html lang="ja">
 
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
 
    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <link rel="stylesheet" href="./assets/css/styleClear.css" type="text/css"> 
 
    <title>Top Page</title>
 
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
<p class="display-1 text-center">DataBase Details</p>
</div>

<body>


<div class="tablebox">
<div class="tableborder">
<h1>user Table</h1>
<div class="text-center col-5 ml-3">
        <table class="text-center table table-bordered table-striped">
            <thead>
                <tr>
                    <th class="table-dark">ColumnName</th>
                    <th class="table-dark">id</th>
                    <th class="table-dark">name</th>
                    <th class="table-dark">mail</th>
                    <th class="table-dark">age</th>
                    <th class="table-dark">passwd</th>
                    <th class="table-dark">created_at</th>
                    <th class="table-dark">updated_at</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th class="table-dark">ColumnValue</th>
                    <th>int</th>
                    <th>string</th>
                    <th>string</th>
                    <th>int</th>
                    <th>Encrypted string</th>
                    <th>Timestamp</th>
                    <th>Timestamp</th>
                </tr>
            </thead>
<tbody>
                <tr>
                    <td class="table-dark">Exsample Recode</td>
                    <td>1</td>
                    <td>Amuro Ray</td>
                    <td>Gamdon@renpo.com</td>
                    <td>19</td>
                    <td>PASSWD</td>
                    <td>2019-11-07 09:32:48</td>
                    <td>2019-11-07 10:45:32</td>
                </tr>
            </tbody>
        </table>
</div>

</div>


<div class="tableborder">
<h1>userdetails Table</h1>
<div class="text-center col-5 ml-3">
        <table class="text-center table table-bordered table-striped">
            <thead>
                <tr>
                    <th class="table-dark">ColumnName</th>
                    <th class="table-dark">uid</th>
                    <th class="table-dark">userimage</th>
                    <th class="table-dark">address</th>
                    <th class="table-dark">animal</th>
                    <th class="table-dark">word</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th class="table-dark">ColumnValue</th>
                    <th>int</th>
                    <th>string</th>
                    <th>string</th>
                    <th>string</th>
                    <th>string</th>
                </tr>
            </thead>
<tbody>
                <tr>
                    <td class="table-dark">Exsample Recode</td>
                    <td>1</td>
                    <td>amuro.png</td>
                    <td>SIDE-7</td>
                    <td>RX-78-2</td>
                    <td>見える、見えるぞ！</td>
                </tr>
            </tbody>
        </table>
</div>



<div class="tableborder">
<h1>sessions Table</h1>
<div class="text-center col-5 ml-3">
        <table class="text-center table table-bordered table-striped">
            <thead>
                <tr>
                    <th class="table-dark">ColumnName</th>
                    <th class="table-dark">uid</th>
                    <th class="table-dark">sessionid</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th class="table-dark">ColumnValue</th>
                    <th>int</th>
                    <th>string</th>
                </tr>
            </thead>
<tbody>
                <tr>
                    <td class="table-dark">Exsample Recode</td>
                    <td>1</td>
                    <td>TVMtMDYtU0BaZW9uLmNvbQ==</td>
                </tr>
            </tbody>
        </table>
</div>



<div class="tableborder">
<h1>posts Table</h1>
<div class="text-center col-5 ml-3">
        <table class="text-center table table-bordered table-striped">
            <thead>
                <tr>
                    <th class="table-dark">ColumnName</th>
                    <th class="table-dark">postid</th>
                    <th class="table-dark">uid</th>
                    <th class="table-dark">post</th>
                    <th class="table-dark">created_at</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th class="table-dark">ColumnValue</th>
                    <th>int</th>
                    <th>int</th>
                    <th>string</th>
                    <th>Timestamp</th>
                </tr>
            </thead>
<tbody>
                <tr>
                    <td class="table-dark">Exsample Recode</td>
                    <td>32</td>
                    <td>1</td>
                    <td>僕が一番ガンダムを使えるんだ!!</td>
                    <td>2019-11-07 15:32:56</td>
                </tr>
            </tbody>
        </table>
</div>


<div class="tableborder">
<h1>admins Table</h1>
<div class="text-center col-5 ml-3">
        <table class="text-center table table-bordered table-striped">
            <thead>
                <tr>
                    <th class="table-dark">ColumnName</th>
                    <th class="table-dark">adminid</th>
                    <th class="table-dark">mail</th>
                    <th class="table-dark">passwd</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th class="table-dark">ColumnValue</th>
                    <th>int</th>
                    <th>string</th>
                    <th>string</th>
                </tr>
            </thead>
<tbody>
                <tr>
                    <td class="table-dark">Exsample Recode</td>
                    <td>1</td>
                    <td>admin@admin.com</td>
                    <td>PASSWD</td>
                </tr>
            </tbody>
        </table>
</div>


<div class="tableborder">
<h1>adminsessions Table</h1>
<div class="text-center col-5 ml-3">
        <table class="text-center table table-bordered table-striped">
            <thead>
                <tr>
                    <th class="table-dark">ColumnName</th>
                    <th class="table-dark">adminid</th>
                    <th class="table-dark">adminsessionid</th>
                </tr>
            </thead>
            <thead>
                <tr>
                    <th class="table-dark">ColumnValue</th>
                    <th>int</th>
                    <th>string</th>
                </tr>
            </thead>
<tbody>
                <tr>
                    <td class="table-dark">Exsample Recode</td>
                    <td>1</td>
                    <td>NvfmtqCXPjakVOxdnwLZkHrLCEulspwA</td>
                </tr>
            </tbody>
        </table>
</div>



</div>

</body>
</html>

