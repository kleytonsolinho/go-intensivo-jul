package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/devfullcycle/go-intensivo-jul/internal/infra/database"
	usecase "github.com/devfullcycle/go-intensivo-jul/internal/useCase"
	"github.com/devfullcycle/go-intensivo-jul/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
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
	// car := Car{ // declarando e atribuindo valores ao struct Car
	// 	Model: "Ferrari",
	// 	Color: "Red",
	// }
	// car.Model = "Lamborghini" // alterando o valor do atributo Model
	// car.Start()               // chamando o método Start
	// car.ChangeColor("Blue")   // chamando o método ChangeColor
	// total := soma(10, 20)     // chamando a função soma
	// println(car.Model, car.Color)
	// println(total)

	// --------------------

	// a := 10
	// b := &a // copiou o valor de A mas criou o seu próprio espaço na memória
	// *b = 20 // alterou o valor de B mas também alterou o valor de A

	// println("a:", a)
	// println("b:", b)

	// --------------------

	// order, err := entity.NewOrder("1", 10, 1)

	// if err != nil {
	// 	println("error:", err.Error())
	// } else {
	// 	println("order:", order.ID)
	// }

	// --------------------

	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // fica escutando a fila do rabbitmq // trava // T2

	rabbitmqWorker(msgRabbitmqChannel, uc)

	// input := usecase.OrderInput{
	// 	ID:    "4",
	// 	Price: 10,
	// 	Tax:   1,
	// }

	// output, err := uc.Execute(input)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(*output)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalcilateFinalPrice) {
	fmt.Println("Starting rabbitMQ")

	for msg := range msgChan {
		var input usecase.OrderInput

		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco: ", output)
	}
}
