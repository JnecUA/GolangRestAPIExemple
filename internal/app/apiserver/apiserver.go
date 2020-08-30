package apiserver

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//APIserver ... is type of the main server structure
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

type PageOptions struct {
	Title string
}

//Init ... Initialize default server
func Init(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ... Configure and start server
func (s *APIserver) Start() error {
	//Configuring and start logger
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Logger start work, you can change logger log level in configure file")

	//Configureing router
	s.configureRouter()
	s.logger.Info("API server start work")

	//Connect sytatic files
	staticPath, _ := filepath.Abs("../public")
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fs)

	//Start router
	addr := s.config.Ip + ":" + strconv.Itoa(s.config.Port)
	s.logger.Info("Router start on http://", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Panic(err)
	}

	return nil
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/", s.handleHome())
}

func (s *APIserver) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./internal/app/public/basictemplate.html")
		p := PageOptions{Title: "Bitcoin Curensi"}
		t.Execute(w, p)
	}
}
