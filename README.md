# tutorial-gowiki (Status: Uncompleted)
Building simplified wiki with Golang

## Current Task

1. Store templates in tmpl/ and page data in data/.
2. Add a handler to make the web root redirect to /view/FrontPage.
3. Spruce up the page templates by making them valid HTML and adding some CSS rules.
4. Implement inter-page linking by converting instances of [PageName] to
5. <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)

## TODO
1. Simple GUI Page.
2. Backend Rest API Calling.
3. Build Simple Simulator.

## commands
```
1. go run wiki.go
```
-- (http://localhost:8080/view/test)
-- (http://localhost:8080/edit/test)

Reference File: https://golang.org/doc/articles/wiki/
