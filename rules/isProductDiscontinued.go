package rules

import "github.com/divyamanchale/go-rule-engine/models"

func CheckIsProductDiscontinued(product models.Product) models.Product {
  if(product.IsProductDiscontinued == true){
    println("Product is available")
    return product
  }
  return product
}