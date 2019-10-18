package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<form action="/oauth/github" method="post">
		<input type="submit" value="Login with Github">
	</form>
</body>
</html>`)
}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
}
