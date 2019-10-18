package main

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOauthConfig = &oauth2.Config{
	ClientID:     "616045a39889744799a7",
	ClientSecret: "c04f7a269768cd87bfb0c9a6b329a49a566c3707",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.HandleFunc("/oauth2/receive", completeGithubOauth)
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
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Couldn't login", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), ts)
}
