package handler

import (
	"Event-Management-System-Go-PSQL/storage"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type eventTypeData struct {
	EventType []storage.EventType
}

/*
type EventTypeForm struct {
	CSRFToken     string
	EventTypeName string
} */

type EventFormData struct {
	CSRFField  template.HTML
	Form       storage.EventType
	FormErrors map[string]string
}

func (s *Server) getEventType(w http.ResponseWriter, r *http.Request) {
	tmp := s.templates.Lookup("event_type_list.html")
	if tmp == nil {
		log.Println("Unable to look event type.html")
		return
	}
	et, err := s.store.GetEventType()
	fmt.Printf("%+v", et)
	if err != nil {
		log.Println("Unable to get event type.  ", err)
	}
	tempData := eventTypeData{
		EventType: et,
	}
	err = tmp.Execute(w, tempData)
	if err != nil {
		log.Println("Error executing tempalte:", err)
		return
	}
}

func (s *Server) createEventType(w http.ResponseWriter, r *http.Request) {
	log.Println("Method : createEventType")

	data := EventFormData{
		CSRFField: csrf.TemplateField(r),
	}

	s.loadCreateEventTypeTemplate(w, r, data)

}
func (s *Server) saveEventType(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalln("Parsing error")
	}

	var form storage.EventType
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln("Decoding error")
	}
	fmt.Printf("%#v", form)
	if form.EventTypeName == "" {
		data := EventFormData{
			CSRFField: csrf.TemplateField(r),
			Form:      form,
			FormErrors: map[string]string{
				"EventTypeName": "Event Type name is required.",
			},
		}
		s.loadCreateEventTypeTemplate(w, r, data)
	}
	id, err := s.store.CreateEventType(form)
	if err != nil {
		log.Fatalln("Unable to save data :", err)

	}
	fmt.Printf("%#v", id)
}

func (s *Server) loadCreateEventTypeTemplate(w http.ResponseWriter, r *http.Request, form EventFormData) {
	tmpl := s.templates.Lookup("event-type-form.html")
	if err := tmpl.Execute(w, form); err != nil {
		log.Println("Error executing tempalte : ", err)
		return
	}
}
