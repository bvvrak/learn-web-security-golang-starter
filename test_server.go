package main

import (
	"fmt"
	"net/http"
	"time"
    "github.com/bootdotdev/learn-web-security/internal/database"
    "github.com/bootdotdev/learn-web-security/internal/httpserver"
    "github.com/bootdotdev/learn-web-security/internal/logging"
    "os"
    "context"
    "path/filepath"
)

func main() {
    wd, _ := os.Getwd()
    db, _ := database.Open(context.Background(), "test.db")
    logger, _ := logging.Open("test.log")
    
    app, err := httpserver.New(db, logger, httpserver.Options{
        DataDirectory:           filepath.Join(wd, "data"),
		FixtureDirectory:        filepath.Join(wd, "data", "fixtures"),
		TemplateDirectory:       filepath.Join(wd, "web", "templates"),
		PublicDirectory:         filepath.Join(wd, "web", "public"),
    })
    if err != nil {
        panic(err)
    }
    
    go http.ListenAndServe(":3032", app.Handler)
    time.Sleep(1 * time.Second)
    
    resp, err := http.Get("http://localhost:3032/health")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Header: %v\n", resp.Header.Get("X-Content-Type-Options"))
}
