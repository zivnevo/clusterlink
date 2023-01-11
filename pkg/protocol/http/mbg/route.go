package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

type MbgHandler struct{}

func (m MbgHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", m.mbgWelcome)

	r.Route("/peer", func(r chi.Router) {
		r.Get("/{mbgID}", m.peerGet)   //GET /peer/{mbgID}  - Get MBG peer Id
		r.Post("/{mbgID}", m.peerPost) // Post /peer/{mbgID} - Add MBG peer Id to MBg peers list
	})

	r.Route("/hello", func(r chi.Router) {
		//r.Use(PostCtx)
		r.Post("/{mbgID}", m.sendHello) // send Hello to MBG peer
		r.Post("/", m.sendHello2All)    // send Hello to MBG peer
	})

	r.Route("/service", func(r chi.Router) {
		r.Post("/", m.addServicePost)   // Post /service  - Expose mbg service
		r.Get("/", m.allServicesGet)    // Get  /service  - Expose mbg service
		r.Get("/{svcId}", m.serviceGet) // Get  /service  - Expose mbg service
	})

	r.Route("/expose", func(r chi.Router) {
		r.Post("/", m.exposePost) // Post /expose  - Expose mbg service
	})

	r.Route("/connect", func(r chi.Router) {
		r.Post("/", m.connectPost)      // Post /connect  - Connect mbg service
		r.Connect("/", m.handleConnect) // Connect /connect  - Connect mbg service
		r.Delete("/", m.connectDelete)  // Disconnect /connect  - Disconnect mbg service

	})

	return r
}

func (m MbgHandler) mbgWelcome(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to Multi-cloud Border Gateway"))
	if err != nil {
		log.Println(err)
	}
}