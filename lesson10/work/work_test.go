package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_fetchPosts(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		args         args
		wantStatus   int
		wantResponse string
	}{
		{
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/posts", nil),
			},
			wantStatus:   http.StatusOK,
			wantResponse: `{"posts":[{"id":1,"title":"First Post"},{"id":2,"title":"Second Post"}]}`,
		},
		{
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/posts", nil),
			},
			wantStatus:   http.StatusMethodNotAllowed,
			wantResponse: `method not allowed`,
		},
	}
	result := fetchPosts()
	assertEqual(t, fetchPosts.Result(),[{"id":1,"title":"First Post"},{"id":2,"title":"Second Post"}])
}
