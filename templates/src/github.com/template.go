package main

import (
        
        "net/http"
        "log"
        "html/template"
        "os"
)

//var templates = template.New("")

type Inventory struct {
	Title string
	Content string
}


func hello(w http.ResponseWriter, r *http.Request) {
        //io.WriteString(w, "web world!")
        log.Printf("reqval:%s",r.RemoteAddr)
        log.Printf("host:%s",r.Host)
        log.Printf("url:%s",r.URL)
        log.Printf("reqval:%s",r.Method)
        //err := templates.ExecuteTemplate(w, tmpl, p)
        sweaters := Inventory{"new title","new content"}
        tmpl, err := template.New("test").Parse("{{.Title}}  {{.Content}}")

        if err != nil { panic(err) }
        err = tmpl.Execute(os.Stdout,sweaters)
        if err != nil { panic(err) }

        tmpl, err = template.ParseFiles("test.html")
        if err != nil {
       panic(err)
       }
       err = tmpl.ExecuteTemplate(w,"test.html",sweaters)
       if err != nil {
       panic(err)
      } 


        //log.Printf("reqval:%s",r.RemoteAddr)
        //log.Printf("reqval:%s",r.RemoteAddr)

}

func main() {
        http.HandleFunc("/", hello)
        http.ListenAndServe(":8000", nil)
}

