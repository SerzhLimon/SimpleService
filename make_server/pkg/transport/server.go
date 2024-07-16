package transport

import (
	"SimpleService/pkg/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const (
	Main          = "pages/main.html"
	Article       = "pages/article.html"
	Admin         = "pages/admin.html"
	Authorization = "pages/authorization.html"
)

type Storage interface {
	GetTotalCount() (int, error)
	Search(limit, offset int) ([]model.Tribe, error)
	SearchById(id int) (model.Tribe, error)
	PublishArticle(data []string) error
	Authorization(login, passwoed string) (string, error)
}

type Server struct {
	limit   int
	storage Storage
	token   string
}

func New(storage Storage, limit int) *Server {
	return &Server{storage: storage, limit: limit}
}

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	
	pageParam := r.URL.Query().Get("page")
	offset, err := strconv.Atoi(pageParam)
	if err != nil || offset < 1 {
		w.Write([]byte("Incorrect value"))
		return
	}

	total, err := s.storage.GetTotalCount()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var next int
	if s.limit*offset < total {
		next = offset + 1
	}

	result, err := s.storage.Search(s.limit, offset-1)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := model.Response{
		Name:   "Tribes",
		Tribes: result,
		Page:   offset,
		Prev:   offset - 1,
		Next:   next,
	}
	tmpl := template.Must(template.ParseFiles(Main))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tmpl.Execute(w, res)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (s *Server) SearchHandlerById(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.Write([]byte("Incorrect value"))
		return
	}

	data, err := s.storage.SearchById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var number_page int
	for i := 0; i < id; i = i + 3 {
		number_page++
	}

	var tribes []model.Tribe
	tribes = append(tribes, data)

	res := model.Response{
		Name:   data.Content,
		Tribes: tribes,
		Page:   number_page,
		Prev:   number_page - 1,
	}

	tmpl := template.Must(template.ParseFiles(Article))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tmpl.Execute(w, res)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func (s *Server) PublishHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("Name")
	description := r.FormValue("Description")
	content := r.FormValue("Content")

	if len(name) < 1 || len(description) < 1 || len(content) < 1 {
		log.Println("empty name, description or content")
	} else {
		data := []string{name, description, content}
		err := s.storage.PublishArticle(data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	http.Redirect(w, r, "/?page=1", http.StatusSeeOther)
}

func (s *Server) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	if s.token != "" {
		s.CreateHandler(w, r)
	} else {
		res := model.Response{}
		tmpl := template.Must(template.ParseFiles(Authorization))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err := tmpl.Execute(w, res)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (s *Server) CheckDataHandler(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("Login")
	password := r.FormValue("Password")

	var err error
	s.token, err = s.storage.Authorization(login, password)
	if err != nil {
		fmt.Println(err.Error())
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (s *Server) CreateHandler(w http.ResponseWriter, r *http.Request) {

	res := model.Response{}
	tmpl := template.Must(template.ParseFiles(Admin))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.Execute(w, res)
	if err != nil {
		fmt.Println(err.Error())
	}
}
