package handler

import (
	"Event-Management-System-Go-PSQL/storage"
	"fmt"
	"log"
	"net/http"
)

type (
	events struct {
		Events []storage.Events
	}
)

func (s *Server) getEvents(w http.ResponseWriter, r *http.Request) {

	tmp := s.templates.Lookup("events.html")

	if tmp == nil {
		log.Println("Unable to look event ")
		return
	}
	et, err := s.store.GetEvent()

	fmt.Printf("%+v", et)

	if err != nil {
		log.Println("Unable to get event type.  ", err)
	}

	tempData := events{
		Events: et,
	}

	err = tmp.Execute(w, tempData)
	if err != nil {
		log.Println("Error executing tempalte:", err)
		return
	}
}