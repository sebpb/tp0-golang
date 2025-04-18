package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")

	// validar que la config este cargada correctamente
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}
	// loggeamos el valor de la config
	log.Printf(globals.ClientConfig.Mensaje)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	for true {
		text := utils.LeerConsola()
		if text == "\n" {
			log.Print("Fin del programa")
			break
		}
	}

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
