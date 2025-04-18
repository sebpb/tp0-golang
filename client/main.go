package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// validar que la config este cargada correctamente
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// generamos un paquete y lo enviamos al servidor
	paquete := utils.LeerYGuardarPaquete()
	utils.EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
}
