package main

import (
	"fmt"
)

/*
func main() {

	fileData, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}

	result, errorr := divide(10, 0)
	if errorr != nil {
		fmt.Println(errorr)
	} else {
		fmt.Println(result)
	}
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divisor cannot be 0")
	}
	return x / y, nil
}
*/

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

type ValidationError struct {
	code string
	text string
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		//return errors.New("Product name cannot be empty!")
		return ValidationError{text: "Product name cannot be empty!", code: "1000"}
	}
	if product.price <= 0 {
		//return errors.New("Price cannot be zero!")
		return ValidationError{text: "Price cannot be zero!", code: "1001"}
	}
	fmt.Println("Product added..")
	return nil
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Error:%s, Code:%s", v.text, v.code)
}

func main() {
	productService := ProductService{}
	err := productService.Add(Product{id: 1, name: "", price: 3000})
	if err != nil {
		fmt.Println(err)
	}

}
