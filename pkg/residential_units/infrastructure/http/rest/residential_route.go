package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	mt_app "github.com/mitribu/mt-residentials/pkg/residential_units/application"
	mt_domain "github.com/mitribu/mt-residentials/pkg/residential_units/domain"
	mth_rest "github.com/mitribu/mt-residentials/pkg/shared/infrastructure/http/rest"
)

type ResidentialRequest struct {
	Name      string "json:name"
	CreatedBy string `json:"created_by,omitempty"`
}

func (a *ResidentialRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("missing required fields.")
	}

	a.CreatedBy = "123"              // It can take from auth_user
	a.Name = strings.ToLower(a.Name) // as an example, we down-case
	return nil
}

type ResidentialResponse struct {
	*mt_domain.Residential

	CreatedBy string `json:"created_by,omitempty"`
}

func NewResidentialResponse(r *mt_domain.Residential) *ResidentialResponse {
	resp := &ResidentialResponse{Residential: r}
	if resp.CreatedBy == "" {
		resp.CreatedBy = "123" // set creator
	}

	return resp
}

func (rd *ResidentialResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Server struct {
	app    mt_app.ResidentialApp
	router chi.Router
}

func NewServer(app mt_app.ResidentialApp) *Server {
	server := &Server{app: app}
	server.router = route(server)
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func route(rr *Server) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Post("/", rr.createResidential)
	return r
}

func (rr *Server) createResidential(w http.ResponseWriter, r *http.Request) {
	data := &ResidentialRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, mth_rest.ErrInvalidRequest(err))
		return
	}
	residential, _ := rr.app.Add(data.Name, data.CreatedBy)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewResidentialResponse(residential))
}
