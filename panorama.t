<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta property="og:title" content="{{ . }}">
  <meta property="og:image" content="http://znbang.github.io/s/{{ . | urlquery }}.jpg">
  <title>{{ . }}</title>
</head>
<body style="margin: 0; overflow: hidden">
<iframe style="position: absolute; width: 100%; height: 100%" src="pannellum.htm?autoLoad=true&panorama=images/{{ . }}.jpg"></iframe>
</body>
</html>
