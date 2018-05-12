<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>{{.Title}}</title>
</head>
<body style="text-align:center">
	<h4><a href="/intro">&lt;&lt; Back to intro</a></h4>
	<h2>{{.Title}}</h2>

	{{range .Story}}<p>{{ . }}</p>{{else}}<p>No story :'(</p>{{end}}

	{{range .Options}}<div><a href="/{{ .Arc }}">{{ .Text }}</a></div>{{end}}
</body>
</html>