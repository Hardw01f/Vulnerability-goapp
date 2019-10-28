<html lang="ja">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Matrix Code</title>
    <style TYPE="text/css">
      body {
        width: 100%;
        margin: 0;
        padding: 0;
      }
      body #wrapper {
        width: 100%;
        height: 100%;
        min-height: 100%;
        overflow: hidden;
        position: fixed;
        transform: rotateY( 180deg );
      }
      #container {
        position: absolute;
        top: 0;
        right: 20px;
        left: 20px;
      }
    </style>
  </head>
  <body style=margin:0 onload="for (s = window.screen, w = q.width = s.width, h = q.height = s.height, m = Math.random, p = [], i = 0; i < 256; p[i++] = 1); setInterval('9Style=\'rgba(0,0,0,.05)\'9Rect(0,0,w,h)9Style=\'#0F0\';p.map(function(v,i){9Text(String.fromCharCode(5e2+m()*33),i*10,v);p[i]=v>758+m()*1e4?0:v+10})'.split(9).join(';q.getContext(\'2d\').fill'), 33)">
    <!-- matrix code background area -->
    <div id="wrapper">
      <canvas id="q"></canvas>
    </div>
    <!-- contents area -->
    <div id="container">
    </div>
  </body>
</html>
