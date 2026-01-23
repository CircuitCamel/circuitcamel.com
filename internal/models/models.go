package models

type Config struct {
	ENV  string
	CRT  string
	KEY  string
	PORT string
}

type Page struct {
	Title string
	Body  string
}

type BlogPost struct {
	Slug  string
	Title string
	Body  string
	Date  string
}

type BlogListPage struct {
	Title string
	Posts []BlogPost
}
