package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	t.Run("happy path -- game is finished", func(t *testing.T) {
		stories := Story(map[string]Arc{
			"intro": {
				Title:   "intro page",
				Story:   []string{"a", "b", "c"},
				Options: nil,
			},
		})
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handlePages(stories).ServeHTTP(w, r)

		assert.Contains(t, w.Body.String(), " <title>intro page</title>")
		assert.Contains(t, w.Body.String(), "<div>a\n\nb\n\nc</div>")
		assert.Contains(t, w.Body.String(), "Thank you for playing the game!")
		assert.Contains(t, w.Body.String(), "Restart")
	})

	t.Run("happy path -- game has options", func(t *testing.T) {
		stories := Story(map[string]Arc{
			"intro": {
				Title:   "intro page",
				Story:   []string{"a", "b", "c"},
				Options: nil,
			},
			"other": {
				Title: "other-page",
				Story: []string{"great", "story"},
				Options: []*option{
					{
						Text: "continue here",
						Arc:  "intro",
					},
				},
			},
		})
		r := httptest.NewRequest(http.MethodGet, "/?page=other", nil)
		w := httptest.NewRecorder()
		handlePages(stories).ServeHTTP(w, r)

		assert.Contains(t, w.Body.String(), " <title>other-page</title>")
		assert.Contains(t, w.Body.String(), "<div>great\n\nstory</div>")
		assert.Contains(t, w.Body.String(), "What next?")
		assert.Contains(t, w.Body.String(), "<a href=\"/?page=intro\">continue here</a><br/>")
	})

}
