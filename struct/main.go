package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Person struct {
	Name    string
	Age     int
	balance int
}

type Address struct {
	City    string
	Street  string
	ZipCode string
}

type Person2 struct {
	Name    string
	Age     int
	Address Address // вложенная структура
}

type Person3 struct {
	Person2 // заносят сюда поля
	Address
	flag bool
}

func (p *Person) setAge(newAge int) {
	p.Age = newAge
}

func (p *Person) setBalance(bal int) error {
	if bal > 10_000_000 {
		return errors.New("too many")
	}
	p.balance = bal
	return nil
}

func main() {
	p1 := Person{"1234", 1, 1000}
	p2 := Person{Name: "name", Age: 20}

	var p3 Person
	p3.Name = "12344"
	p3.Age = 1

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	p4 := new(Person)
	fmt.Println(p4)

	p5 := Person2{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "Moscow",
			Street:  "Tverskaya",
			ZipCode: "101000",
		},
	}
	fmt.Println(p5)

	fmt.Println(p1)
	p1.setAge(10000)
	fmt.Println(p1)

	point := struct{ x, y int }{10, 10}
	fmt.Println(point)

	type User struct {
		ID        int       `json:"id" db:"user_id"`
		Name      string    `json:"name" validate:"required"`
		Email     string    `json:"email" validate:"required,email"`
		CreatedAt time.Time `json:"created_at"`
		Password  string    `json:"-"` // игнорировать в JSON
	}

	// Использование с JSON
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))
	// {"id":1,"name":"Alice","email":"alice@example.com","created_at":"0001-01-01T00:00:00Z"}
}
