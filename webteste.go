package main

import (
    "fmt"
    "log"
    "net/http"
)

const (
    pHeader = "<!DOCTYPE html><html><body>"
    form    = `<form action="/" method="POST">
<label for="name">Digite seu nome:</label><br />
<input type="text" name="name" />
<input type="submit" />
</form>`
    pFooter = "</body></html>"
)

func indexPage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    fmt.Fprint(writer, pHeader, form)

    if err != nil {
        fmt.Fprint(writer, err)
    } else {
        if name, ok := processRequest(request); ok {
            fmt.Fprint(writer, "Olá, ", name[0])
        } else {
            fmt.Fprint(writer, "Olá, Mundo!")
        }
    }

    fmt.Fprint(writer, pFooter)
}

func processRequest(request *http.Request) ([]string, bool) {
    if name, found := request.Form["name"]; found && len(name) > 0 {
        return name, true
    }
    return nil, false
}

func main() {
    http.HandleFunc("/", indexPage)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("failed to start webserver", err)
    }
}
