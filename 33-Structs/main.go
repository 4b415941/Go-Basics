package main

import "fmt"

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}

func main() {
	var customer1 = Customer{id: 1, name: "Kadircan", age: 26, address: Address{city: "Eskisehir", district: "Odunpazari"}}
	//var customer2 = Customer{id: 1, name: "Beng", age: 24, address: Address{city: "Eskisehir", district: "Odunpazari"}}

	/* fmt.Println(customer1)
	fmt.Println(customer2) */
	customer1.toString()
	customer1.changeName("Nur")
	customer1.toString()
	//customer2.toString()

	//changeName(&customer1)
	//customer1.toString()

}

func (c *Customer) toString() {
	fmt.Printf("Name: %s, Age: %d Address:%v\n", c.name, c.age, c.address)
}

func (c *Customer) changeName(newName string) {
	c.name = newName
}

/* func changeName(c *Customer) {
	c.name = "Nur"
} */
