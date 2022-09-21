package handlers

import (
	"net/http"

	"github.com/hamza-starcevic/bookings/pkg/config"
	"github.com/hamza-starcevic/bookings/pkg/models"
	render "github.com/hamza-starcevic/bookings/pkg/renderers"
)

// * Repo the repository used by the handlers
var Repo *Repository

// * Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// * NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// * NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// * Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.Templatedata{})
}

// * About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//*perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//*send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.Templatedata{
		StringMap: stringMap,
	})
}