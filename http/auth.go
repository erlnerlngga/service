package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	g "github.com/maragudk/gomponents"
	ghttp "github.com/maragudk/gomponents/http"

	"github.com/maragudk/service/html"
)

type sessionDestroyer interface {
	Destroy(ctx context.Context) error
}

// Logout creates an http.Handler for logging out.
// It just destroys the current user session.
func Logout(mux chi.Router, s sessionDestroyer, log *log.Logger) {
	mux.Post("/logout", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		if err := s.Destroy(r.Context()); err != nil {
			log.Println("Error logging out:", err)
			return html.ErrorPage(), err
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return nil, nil
	}))
}
