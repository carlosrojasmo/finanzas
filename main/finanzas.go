package main

import (
	"fmt"
)

var registerPedido= []pedido{}
var registerGanancia=[]int{}

type pedido struct{
	idPaquete string
	intentos int
	entregado bool
	tipo string
	valor int
}
func sumSlices(x[]int) int {	

	totalx := 0
	for _, valuex := range x {
		totalx += valuex
	}
	return totalx
}
func newPedido(idPaquete string, intentos int, estregado bool,tipo string,valor int) pedido{
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
		return (ganancia)
	}else {
		if pedido.tipo=="normal"{
			return ganancia
		}else if pedido.tipo=="retail"{
			ganancia=ganancia+pedido.valor
		}else{
			return ganancia 
		}
	}	
}

func main(){
	//escucchar
	//recibir 
	nuevoPedido := newPedido(id,intentos,entragado,tipo,valor)
	registerPedido=append(registerPedido,nuevoPedido)
	ganancia:=calcularGanancia(nuevoPedido)
	registerGanancia=append(registerGanancia,ganancia)
	sum:=sumSlices(registerGanancia)
	fmt.Println("la ganancia es: ",sum )
}
