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

//SaveToFileJSON writting
func SaveToFileJSON(equipo *domain.Jugadores) {

	j := convertToJSON(*equipo)
	f, err := os.Create("./data/output.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	newBytes := []byte(j)
	_, errW := f.Write(newBytes)
	if errW != nil {
		fmt.Println(errW)
	}

}

func convertToJSON(j domain.Jugadores) (jsonString string) {
	jsonTmp, err := json.Marshal(&j)
	if err != nil {
		fmt.Println(err)
	}
	jsonString = string(jsonTmp)

	return jsonString

}
