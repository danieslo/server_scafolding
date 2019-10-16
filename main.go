package main

import (
    "log"
    "net/http"
    "time"
    
    "github.com/rs/cors"
    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
)
type Route struct {
    Method string
    Path string
    Handler func(w http.ResponseWriter, r *http.Request)
}
type Handlers struct {
    Middlewares []*negroni.HandlerFunc
    Routes []Route
}

func CreateHandler(h Handlers) (*negroni.Negroni, error) {
    c := cors.AllowAll()
    
    n := negroni.New()
    n.Use(c)
    for _, middleware := range h.Middlewares {
        n.Use(middleware)
    }
    
    r := mux.NewRouter().StrictSlash(false)
    n.UseHandler(r)
    
    for _, route := range h.Routes {
        r.Methods(route.Method).Path(route.Path).HandlerFunc(route.Handler)
    }
    
    return n, nil
}

func CreateServer(n *negroni.Negroni, p string, rT time.Duration, wT time.Duration){
    
    server := &http.Server{
        Addr:              p,
        Handler:           n,
        ReadTimeout:       rT,
        WriteTimeout:      wT,
    }
    
    log.Fatal(server.ListenAndServe().Error())
}