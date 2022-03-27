package main

import (
    "fmt"
    // "github.com/pkg/profile"
    "github.com/pyroscope-io/client/pyroscope"
    "net/http"
)

func main() {
    // defer profile.Start().Stop()
    pyroscope.Start(pyroscope.Config{
        ApplicationName: "simple.golang.app",
    
        // replace this with the address of pyroscope server
        ServerAddress:   "http://localhost:4040",
    
        // you can disable logging by setting this to nil
        Logger:          pyroscope.StandardLogger,
    
        // optionally, if authentication is enabled, specify the API key:
        // AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN"),
    
        // by default all profilers are enabled,
        // but you can select the ones you want to use:
        ProfileTypes: []pyroscope.ProfileType{
          pyroscope.ProfileCPU,
          pyroscope.ProfileAllocObjects,
          pyroscope.ProfileAllocSpace,
          pyroscope.ProfileInuseObjects,
          pyroscope.ProfileInuseSpace,
        },
    })
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!!!!!", r.URL.Path[1:])
}