<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>360x180 panorama</title>
</head>
<body>
{{ range $key, $value := . }}
<a href="{{ $value }}.html">{{ $value }}</a><br>
{{ end }}
</body>
</html>
