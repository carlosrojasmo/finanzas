package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
)

var registerPedido= []pedido{}
var registerGanancia=[]int{}


type paquete struct {
	IDPaquete string
	Tipo string
	Valor int
	Seguimiento int
	Intentos int
	Estado string
}
type pedido struct{
	idPaquete string
	tipo string
	valor int
	intentos int
	entregado bool
	
	
}
func sumSlices(x[]int) int {	

	totalx := 0
	for _, valuex := range x {
		totalx += valuex
	}
	return totalx
}
func newPedido(idPaquete string, intentos int, entregado bool,tipo string,valor int) pedido{
	pedidoNuevo := pedido{idPaquete:idPaquete,intentos:intentos,entregado: entregado, tipo:tipo,valor:valor}
	return pedidoNuevo
}

func calcularGanancia(pedido pedido) int{
	var ganancia int = -(pedido.intentos*10)
	if pedido.tipo=="prioritario"{
		var prio float32=0.3*float32(pedido.valor)
		ganancia=ganancia+int(prio)
	}
	if pedido.entregado==true{
		ganancia=ganancia+pedido.valor
	}else {
		if pedido.tipo=="normal"{
		}else if pedido.tipo=="retail"{
			ganancia=ganancia+pedido.valor
		}
	}
	return ganancia	
}
func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
  }

func main(){
	//escucchar
	//recibir 
	conn, err := amqp.Dial("amqp://logistica:logistica@10.10.28.102:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
	    "hello", // name
  		false,   // durable
    	false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )
	  failOnError(err, "Failed to register a consumer")
	  
	  forever := make(chan bool)
	  
	  go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var paqueteCalcular paquete
			var entregado bool
			err=json.Unmarshal(d.Body,&paqueteCalcular)
			fmt.Print("Paquete a calcular ganancia: ", paqueteCalcular.IDPaquete)
			if paqueteCalcular.Estado=="Recibido"{
				entregado=true
			}else{
				entregado=false
			}
			nuevoPedido := newPedido(paqueteCalcular.IDPaquete,paqueteCalcular.Intentos,entregado,paqueteCalcular.Tipo,paqueteCalcular.Valor)
			registerPedido=append(registerPedido,nuevoPedido)
			ganancia:=calcularGanancia(nuevoPedido)
			registerGanancia=append(registerGanancia,ganancia)
			sum:=sumSlices(registerGanancia)
			fmt.Println(", la ganancia acumulada es: ",sum )
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever

}
