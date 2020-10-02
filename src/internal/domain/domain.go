package domain

import "fmt"

// Jugadores struct which contains
// an array of users
type Jugadores struct {
	Jugadores           []Jugador `json:"jugadores"`
	GolesAlMes          int
	GolesAnotados       int
	totalGolesPorNivel  int
	PromedioGoles       float32
	TotalGolesMetidos   int
	TotalGolesEstimados int
}

// Jugador representation of Jugador into  data struct
type Jugador struct {
	Nombre         string  `json:"nombre"`
	Nivel          string  `json:"nivel"`
	Goles          int     `json:"goles"`
	SueldoBase     float32 `json:"sueldo"`
	Bono           float32 `json:"bono"`
	SueldoCompleto float32 `json:"sueldo_completo"`
	Equipo         string  `json:"equipo"`
	GolesMinimo    int     `json:"goles_minimos"`
}

//CalcularTotalDeGoles show
func (jugadores *Jugadores) CalcularTotalDeGoles() {

	var totalGolesEstimados = 0
	var totalGolesMetidos = 0

	for _, v := range jugadores.Jugadores {
		totalGolesMetidos += v.Goles
		golesEstimados := getNivelDeGoles(v.Nivel)
		totalGolesEstimados += golesEstimados
	}

	jugadores.TotalGolesEstimados = totalGolesEstimados
	jugadores.TotalGolesMetidos = totalGolesMetidos
}

//CalcularSalariosPorJugador actualiza los salarios correspondientes
// para cada jugador
func (jugadores *Jugadores) CalcularSalariosPorJugador() {

	totalDeGolesEstimados := float32(jugadores.TotalGolesEstimados)

	tam := len(jugadores.Jugadores)
	for i := 0; i < tam; i++ {
		bonoPorEquipo := jugadores.Jugadores[i].CalcularBonoPorEquipo(float32(jugadores.TotalGolesMetidos), totalDeGolesEstimados, jugadores.getPesoBonoPorEquipo())
		jugadores.Jugadores[i].SueldoCompleto = jugadores.Jugadores[i].calcularSueldoCompleto(bonoPorEquipo)

	}
}

//ShowSalarios muestra  el sueldo completo de cada jugador
func (jugadores *Jugadores) ShowSalarios() {

	for i, v := range jugadores.Jugadores {
		fmt.Printf("%d: El sueldo completo de %s  es %.2f\n", i, v.Nombre, v.SueldoCompleto)
	}
}

//CalcularBonoPorEquipo return float
func (jugador *Jugador) CalcularBonoPorEquipo(golesMetidos, golesEstimados, peso float32) (result float32) {

	bono := peso * jugador.Bono

	promedio := golesMetidos / golesEstimados
	result = bono * promedio
	return
}
func (jugador *Jugador) calcularSueldoCompleto(bonoPorEquipo float32) float32 {
	bonoIndividual := jugador.BonoIndividual()
	bonoTotal := bonoIndividual + bonoPorEquipo
	sueldoTotal := jugador.SueldoBase + bonoTotal

	return sueldoTotal
}

//BonoIndividual return float
func (jugador *Jugador) BonoIndividual() (result float32) {

	golesEstimados := getNivelDeGoles(jugador.Nivel)
	porcentajeGoles := float32(jugador.Goles) / float32(golesEstimados)

	bono := jugador.getPesoBonoIndividual() * jugador.Bono

	result = bono * porcentajeGoles

	return
}

func (jugadores *Jugadores) getPesoBonoPorEquipo() float32 {
	return 0.5
}

func (jugador *Jugador) getPesoBonoIndividual() float32 {
	return 0.5
}
func getNivelDeGoles(s string) int {
	var nivelJugador = map[string]int{"A": 5, "B": 10, "C": 15, "Cuauh": 20}
	return nivelJugador[s]
}
