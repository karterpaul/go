package main

    import (
        "fmt"
        "html/template"
        "net/http"
        "log"
        "strings"
        "github.com/gorilla/csrf"
        "github.com/gorilla/mux"
    )

    var form = `
    <html>
    <head>
    <title>Sign Up!</title>
    </head>
    <body>
    <form method="POST" action="/signup/post" accept-charset="UTF-8">
    <input type="text" name="name">
    <input type="text" name="email">
    <!--
        The default template tag used by the CSRF middleware .
        This will be replaced with a hidden <input> field containing the
        masked CSRF token.
    -->
    {{ .csrfField }}
    <input type="submit" value="Sign up!">
    </form>
    </body>
    </html>
    `

    var t = template.Must(template.New("signup_form.tmpl").Parse(form))

    func main() {
        r := mux.NewRouter()


        /*

         //post and get hanling in mux
         r := mux.NewRouter()
         r.HandleFunc("/members", getMembersHandler).Methods("GET")
         r.HandleFunc("/members", postMembersHandler).Methods("POST")
 
         http.Handle("/", r)
         http.ListenAndServe(":8080", nil)

        */


        r.HandleFunc("/signup", ShowSignupForm)
        // All POST requests without a valid token will return HTTP 403 Forbidden.
        r.HandleFunc("/signup/post", SubmitSignupForm)

        // Add the middleware to your router by wrapping it.
        http.ListenAndServe(":8000",
            csrf.Protect([]byte("32-byte-long-auth-key"),csrf.Secure(false))(r))
        // PS: Don't forget to pass csrf.Secure(false) if you're developing locally
        // over plain HTTP (just don't leave it on in production).
    }

    func ShowSignupForm(w http.ResponseWriter, r *http.Request) {
        // signup_form.tmpl just needs a {{ .csrfField }} template tag for
        // csrf.TemplateField to inject the CSRF token into. Easy!
        t.ExecuteTemplate(w, "signup_form.tmpl", map[string]interface{}{
            csrf.TemplateTag: csrf.TemplateField(r),
        })

        log.Printf("test log");
    }




    func SubmitSignupForm(w http.ResponseWriter, r *http.Request) {
        // We can trust that requests making it this far have satisfied
        // our CSRF protection requirements.
        log.Printf("test log");
        //fmt.Fprintf(w, "%v\n", r.PostForm)
        //fmt.Fprintf(w, "%v\n", r)
        //fmt.Fprintf(w, "%v\n", r.PostForm)
        //str:=r.PostForm
        log.Printf("raw=%s",r.PostForm)
        //log.Printf("str=%s",str[0])
       // a:= str.(map[string]interface{})
        //r:= FormData{}
        //json.Unmarshal(str, &r)
        for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("value:", v[0])
        fmt.Println("val:", strings.Join(v, ""))
    }


        fmt.Fprintf(w,"%s",r.Form["name"])
        fmt.Fprintf(w,"%s",r.Form["name"][0])

       /* for k, v := range str {

        log.Printf("key=%s",k)
        log.Printf("val=%s",v)

        }
*/

    }

