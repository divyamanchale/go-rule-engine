package main

import (
        "net/http"
        "github.com/divyamanchale/go-rule-engine/models"
        "github.com/divyamanchale/go-rule-engine/rules"
        "encoding/json"
      )

const ( defaultport = "8080" )

func main() {
  http.HandleFunc("/validate", requestHandler)
  http.ListenAndServe(":"+defaultport, nil)
}

func requestHandler(w http.ResponseWriter,r *http.Request){
  println("I got request")
  var p models.Product
  err := json.NewDecoder(r.Body).Decode(&p)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  var cart models.Cart
  cart.Product = p
  var request models.Request
  request.Cart = cart
  cart = rules.Run(request)
  println("I got a response")
  json.NewEncoder(w).Encode(cart)
}