package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MdTosif/go-deployments/internal/config"
	"github.com/MdTosif/go-deployments/internal/runner"
)


func basicAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Method 1: via Authorization header
        user, pass, ok := r.BasicAuth()
        print(user, pass, ok)  
        if !ok || user != config.Cfg.Auth.Username  || pass != config.Cfg.Auth.Password {
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // (Or Method 2: via URL.User, if you truly want to accept user:pass in the URL)
        // if r.URL.User != nil {
        //     if u := r.URL.User.Username(); u != validUser {
        //         unauthorized(w); return
        //     }
        //     p, hasPass := r.URL.User.Password()
        //     if !hasPass || p != validPass {
        //         unauthorized(w); return
        //     }
        // }

        next.ServeHTTP(w, r)
    })
}

func deployHandler(w http.ResponseWriter, r *http.Request) {
    // r.URL.Path might be "/deploy/foo", "/deploy/bar", etc.
    // Trim the "/deploy/" prefix to get the “:service” value.
    service := strings.TrimPrefix(r.URL.Path, "/deploy/")

    // At this point, `service` will be everything after "/deploy/".
    // You might want to check for empty or validate it:
    if service == "" {
        http.Error(w, "`service` not provided", http.StatusBadRequest)
        return
    }

    runner := runner.New()


    for _, v := range config.Cfg.Services {
        if service == v.Name {
            runner.Run(v.Cmd)
        }
    }

    // Do something with “service”:
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.Write([]byte("You asked to deploy service: " + service))
}

func Start() {
    mux := http.NewServeMux()
    mux.HandleFunc("/deploy/", deployHandler) // "/deploy/:service"

    // wrap your mux (or individual handlers) with basicAuth
    protected := basicAuth(mux)

    srv := &http.Server{
        Addr:    fmt.Sprintf(":%d", config.Cfg.Port),
        Handler: protected,
    }

    log.Printf("Listening on %s …", srv.Addr)
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}