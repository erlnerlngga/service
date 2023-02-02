package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	g "github.com/maragudk/gomponents"
	ghttp "github.com/maragudk/gomponents/http"

	"github.com/maragudk/service/html"
	"github.com/maragudk/service/model"
)

type contextKey string

const contextUserKey = contextKey("user")

// getUserFromContext, which may be nil if the user is not authenticated.
func getUserFromContext(ctx context.Context) *model.User {
	user := ctx.Value(contextUserKey)
	if user == nil {
		return nil
	}
	return user.(*model.User)
}

func Signup(mux chi.Router) {
	mux.Get("/signup", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		user := getUserFromContext(r.Context())
		if user != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return nil, nil
		}

		return html.SignupPage(html.PageProps{}), nil
	}))
}

func Login(mux chi.Router) {
	mux.Get("/login", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
		user := getUserFromContext(r.Context())
		if user != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return nil, nil
		}

		return html.LoginPage(html.PageProps{}), nil
	}))
}

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
