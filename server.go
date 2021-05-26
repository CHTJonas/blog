package blog

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/CHTJonas/blog/assets"
	"github.com/gorilla/mux"
)

type Server struct {
	b   *Blog
	r   *mux.Router
	srv *http.Server
}

func NewServer() *Server {
	s := new(Server)
	s.r = mux.NewRouter().StrictSlash(true)
	s.r.PathPrefix("/static/").Handler(assets.Server())
	s.r.HandleFunc("/", s.index)
	s.r.HandleFunc("/robots.txt", s.robots)
	s.r.HandleFunc("/rss.xml", s.feed)
	s.r.HandleFunc("/sitemap.xml", s.sitemap)
	s.r.HandleFunc("/{slug}", s.show)
	return s
}

func (serv *Server) Start(addr string) error {
	serv.srv = &http.Server{
		Addr:    addr,
		Handler: serv.r,
	}
	return serv.srv.ListenAndServe()
}

func (serv *Server) Stop(ctx context.Context) error {
	serv.srv.SetKeepAlivesEnabled(false)
	return serv.srv.Shutdown(ctx)
}

func (serv *Server) index(w http.ResponseWriter, r *http.Request) {
	partial := string(assets.EnsureReadFile("index.html"))
	layout := string(assets.EnsureReadFile("layout.html"))
	t, err := template.New("index").Parse(partial + layout)
	if err != nil {
		panic(err)
	}
	data := serv.getTemplateData()
	data.ContentData = serv.b.GetPosts()
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (serv *Server) robots(w http.ResponseWriter, r *http.Request) {
	tpl := assets.EnsureReadFile("robots.txt")
	t, err := template.New("robots").Parse(string(tpl))
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, serv.b)
	if err != nil {
		panic(err)
	}
}

func (serv *Server) feed(w http.ResponseWriter, r *http.Request) {
	rss, err := serv.b.GetRSS()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, rss)
}

func (serv *Server) sitemap(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, serv.b.GetSitemap())
}

func (serv *Server) show(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	partial := string(assets.EnsureReadFile("show.html"))
	layout := string(assets.EnsureReadFile("layout.html"))
	t, err := template.New("show").Parse(partial + layout)
	if err != nil {
		panic(err)
	}
	for _, p := range serv.b.GetPosts() {
		if p.Slug == slug {
			data := serv.getTemplateData()
			data.ContentData = p
			err = t.Execute(w, data)
			if err != nil {
				panic(err)
			}
			return
		}
	}
}
