package models

type Cart struct{

  Product Product


}

type Product struct {
  IsProductDummy bool `json:"isProductDummy"`
  IsProductDiscontinued bool `json:"isProductDiscontinued"`
  Result map[string]bool
}

type Request struct{

  Cart Cart
  Stack []func(Product) Product
}

type Response struct {

  Success bool
  Code    int
}