package model

type WebInfo struct {
	WebName     string `db:"web_name"`
	Keywords    string `db:"keywords"`
	Description string `db:"description"`
	Url         string `db:"url"`
}
