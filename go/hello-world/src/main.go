package main

import (
  "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello from GO"))
}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":80", nil); err != nil {
    panic(err)
  }
}
