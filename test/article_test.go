package test

import (
	"article-go-gin-semaphore/models"
	"article-go-gin-semaphore/routes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := models.GetAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(models.ArticleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != models.ArticleList[i].Content ||
			v.ID != models.ArticleList[i].ID ||
			v.Title != models.ArticleList[i].Title {

			t.Fail()
			break
		}
	}
}

// Test the function that fetche an Article by its ID
func TestGetArticleByID(t *testing.T) {
	a, err := models.GetArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", routes.ShowIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed test using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a GET request to an article page returns the article page with
// the HTTP code 200 for an unauthenticated user
func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/article/view/:article_id", routes.GetArticle)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Article 1"
		// You can carry out a lot more detailed test using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}