// File generated by Gopher Sauce
// DO NOT EDIT!!
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	gosweb "github.com/cheikhshift/gos/web"
	"github.com/gorilla/sessions"
	opentracing "github.com/opentracing/opentracing-go"
	templates "github.com/thestrukture/dap/api/templates"

	sessionStore "github.com/thestrukture/dap/api/sessions"
)

// Access you .gxml's end tags with
// this http.HandlerFunc.
// Use MakeHandler(http.HandlerFunc) to serve your web
// directory from memory.
func MakeHandler(fn func(http.ResponseWriter, *http.Request, opentracing.Span)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		span := opentracing.StartSpan(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		defer span.Finish()
		carrier := opentracing.HTTPHeadersCarrier(r.Header)
		if err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier); err != nil {
			log.Fatalf("Could not inject span context into header: %v", err)
		}

		if attmpt := ApiAttempt(w, r, span); !attmpt {
			fn(w, r, span)
		}

	}
}

func Handler(w http.ResponseWriter, r *http.Request, span opentracing.Span) {
	var p *gosweb.Page
	p, err := templates.LoadPage(r.URL.Path)
	var session *sessions.Session
	var er error
	if session, er = sessionStore.Store.Get(r, "session-"); er != nil {
		session, _ = sessionStore.Store.New(r, "session-")
	}

	var sp opentracing.Span
	opName := fmt.Sprintf(fmt.Sprintf("Web:/%s", r.URL.Path))

	if true {
		carrier := opentracing.HTTPHeadersCarrier(r.Header)
		wireContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, carrier)
		if err != nil {
			sp = opentracing.StartSpan(opName)
		} else {
			sp = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
		}
	}
	defer sp.Finish()

	if err != nil {
		log.Println(err.Error())

		w.WriteHeader(http.StatusNotFound)
		span.SetTag("error", true)
		span.LogEvent(fmt.Sprintf("%s request at %s, reason : %s ", r.Method, r.URL.Path, err))
		pag, err := templates.LoadPage("/your-404-page")

		if err != nil {
			log.Println(err.Error())
			//
			return
		}
		pag.R = r
		pag.Session = session
		if p != nil {
			p.Session = nil
			p.Body = nil
			p.R = nil
			p = nil
		}

		if pag.IsResource {
			w.Write(pag.Body)
		} else {
			templates.RenderTemplate(w, pag, span) //"/your-500-page"
		}
		session = nil

		return
	}

	if !p.IsResource {
		w.Header().Set("Content-Type", "text/html")
		p.Session = session
		p.R = r
		templates.RenderTemplate(w, p, span) //fmt.Sprintf("web%s", r.URL.Path)
		session.Save(r, w)
		// log.Println(w)
	} else {
		w.Header().Set("Cache-Control", "public")
		if strings.Contains(r.URL.Path, ".css") {
			w.Header().Add("Content-Type", "text/css")
		} else if strings.Contains(r.URL.Path, ".js") {
			w.Header().Add("Content-Type", "application/javascript")
		} else {
			w.Header().Add("Content-Type", http.DetectContentType(p.Body))
		}

		w.Write(p.Body)
	}

	p.Session = nil
	p.Body = nil
	p.R = nil
	p = nil
	session = nil

	return
}
