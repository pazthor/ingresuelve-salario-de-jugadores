package service

import (
	"../domain"
	"../repository"
)

//Proceso realiza las operaciones para completar el reto
func Proceso() *domain.Jugadores {
	equipo := repository.GetJugadores()
	equipo.CalcularTotalDeGoles()
	equipo.CalcularSalariosPorJugador()

	return equipo

}
