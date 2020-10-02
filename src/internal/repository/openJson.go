package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../domain"
)

func openFileJSON() []byte {
	jsonFile, errfile := os.Open("./data/jugadores.json")
	if errfile != nil {
		fmt.Println(errfile)
	}

	defer jsonFile.Close()

	byteValue, errbyteValue := ioutil.ReadAll(jsonFile)
	if errfile != nil {
		fmt.Println(errbyteValue)
	}

	return byteValue

}

//GetJugadores obtiene la lista de jugadores
func GetJugadores() *domain.Jugadores {
	var jugadores *domain.Jugadores

	byteValue := openFileJSON()
	json.Unmarshal(byteValue, &jugadores)

	return jugadores
}
