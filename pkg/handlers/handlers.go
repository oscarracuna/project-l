package handlers

import (
	"fmt"
	"net/http"

	"github.com/oscarracuna/project-l/pkg/config"
	"github.com/oscarracuna/project-l/pkg/models"
	"github.com/oscarracuna/project-l/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	in1 := r.Form.Get("input1")
	in2 := r.Form.Get("input2")
	w.Write([]byte(fmt.Sprintf("Input 1 = %s and input 2 = %s", in1, in2)))
}
