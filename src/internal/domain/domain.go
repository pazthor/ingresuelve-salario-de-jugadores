package domain

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
