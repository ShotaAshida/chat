package main

import(
  "log"
  "net/http"
  "sync"
  "text/template"
  "path/filepath"
  "fmt"
)

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  t.templ.Execute(w, nil)
}


func main() {
    var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
    flag.Parse() //フラグを解釈
    r := newRoom()
    http.Handle("/", &templateHandler{filename: "chat.html"})
    http.Handle("/room", r)
    //ちゃっとるーむ開始
    fmt.Print("run前")
    go r.run()
    fmt.Print("run後")

    // Webサーバを開始します
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListernAndServe:", err)
    }
}
