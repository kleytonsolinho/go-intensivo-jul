package main

import (
	"database/sql"
	"fmt"

	"github.com/devfullcycle/go-intensivo-jul/internal/entity"
	"github.com/devfullcycle/go-intensivo-jul/internal/infra/database"
	usecase "github.com/devfullcycle/go-intensivo-jul/internal/useCase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

func (c Car) Start() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color // duplicando o valor de c.color na memória - copia do color original
	println("New color: " + c.Color)
}

func soma(x, y int) int {
	return x + y
}

func main() {
	car := Car{ // declarando e atribuindo valores ao struct Car
		Model: "Ferrari",
		Color: "Red",
	}
	car.Model = "Lamborghini" // alterando o valor do atributo Model
	car.Start()               // chamando o método Start
	car.ChangeColor("Blue")   // chamando o método ChangeColor
	total := soma(10, 20)     // chamando a função soma
	println(car.Model, car.Color)
	println(total)

	// --------------------

	a := 10
	b := &a // copiou o valor de A mas criou o seu próprio espaço na memória
	*b = 20 // alterou o valor de B mas também alterou o valor de A

	println("a:", a)
	println("b:", b)

	// --------------------

	order, err := entity.NewOrder("1", 10, 1)

	if err != nil {
		println("error:", err.Error())
	} else {
		println("order:", order.ID)
	}

	// --------------------

	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "2",
		Price: 10,
		Tax:   1,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(*output)
}
