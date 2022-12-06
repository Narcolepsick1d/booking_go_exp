package handlers

import (
	"github.com/Narcolepsick1d/bookings/pkg/config"
	"github.com/Narcolepsick1d/bookings/pkg/models"
	"github.com/Narcolepsick1d/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	StringMap := make(map[string]string)
	StringMap["test"] = "hello again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: StringMap,
	})

}
