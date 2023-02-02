package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/honeybadger-io/honeybadger-go"
)

func (s *Server) setupRoutes() {
	s.mux.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer, honeybadger.Handler)
		r.Use(middleware.Compress(5))
		r.Use(middleware.RealIP)
		r.Use(AddMetrics(s.metrics))

		Health(r, s.database, s.log)
		Metrics(r, s.metrics)

		r.Group(func(r chi.Router) {
			r.Use(VersionedAssets)

			Static(r)
		})

		r.Group(func(r chi.Router) {
			r.Use(NoClickjacking, StrictContentSecurityPolicy)
			r.Use(middleware.SetHeader("Content-Type", "text/html; charset=utf-8"))
			r.Use(s.sm.LoadAndSave)

			Signup(r)
			Login(r)
			Logout(r, s.sm, s.log)

			Home(r)
			NotFound(r)
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("admin", map[string]string{"admin": s.adminPassword}))

			Migrate(r, s.database)
		})
	})
}
