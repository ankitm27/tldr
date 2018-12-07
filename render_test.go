package tldr

import (
	"os"
	"testing"

	"github.com/mstruebing/tldr/cache"
)

func TestRender(t *testing.T) {
	r, _ := cache.NewRepository(remoteURL, ttl)
	markdown, _ := r.Markdown("linux", "hostname")

	render, err := Render(markdown)

	if err != nil {
		t.Error("Expected to render existing page successfully")
	}

	if len(render) == 0 {
		t.Error("Expected the length of the rendered string to be not empty")
	}
}

func TestWrite(t *testing.T) {
	r, _ := cache.NewRepository(remoteURL, ttl)
	markdown, _ := r.Markdown("linux", "hostname")

	err := Write(markdown, os.Stdout)

	if err != nil {
		t.Error("Expected to write successfully")
	}
}
