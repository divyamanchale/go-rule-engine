package rules

import (
        "github.com/divyamanchale/go-rule-engine/models"
      )



func Run(request models.Request) models.Cart {
  request.Stack = append(request.Stack,CheckIsProductDummy)
  request.Stack = append(request.Stack, CheckIsProductDiscontinued)

  for len(request.Stack) > 0 {
   request.Cart.Product =  request.Stack[0](request.Cart.Product)
    request.Stack = request.Stack[1:]
  }

  return request.Cart
}