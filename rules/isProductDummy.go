package rules

import "github.com/divyamanchale/go-rule-engine/models"

func CheckIsProductDummy(product models.Product) models.Product {

  if(product.IsProductDummy == true){
    println("Product is not dummy")
     m := map[string]bool{
        "isProductDummyCheck": true,
    }
    product.Result = m
    println(product.Result["isProductDummyCheck"])
    return product
  }
  return product
}