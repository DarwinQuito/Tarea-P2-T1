package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Multa struct {
	Id       string `json:id`
	Nombre   string `json:nombre`
	Sancion  string `json:Sancion`
	Cantidad int    `json:cantidad`
}

// Array global de Multas:
var Multas []Multa

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h2>Servidor funcionando by Darwin Quito!</h2></body></html>")
	fmt.Println("Solicitud atendida: homePage")
}

func findAllMultas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findAllMultas")
	json.NewEncoder(w).Encode(Multas)
}

func findMultaById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findMultaById")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, Multa := range Multas {
		if Multa.Id == key {
			json.NewEncoder(w).Encode(Multa)
		}
	}
}

func createNewMulta(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el body desde el request y
	// se deserializa en una variable Multa:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Multa Multa
	json.Unmarshal(reqBody, &Multa)
	// adicionamos en el array el nuevo Multa:
	Multas = append(Multas, Multa)
	json.NewEncoder(w).Encode(Multa)
}

func deleteMulta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: deleteMulta")
	vars := mux.Vars(r)
	key := vars["id"]
	// buscar el Multa a eliminar:
	for index, Multa := range Multas {
		if Multa.Id == key {
			// borrar del array:
			Multas = append(Multas[:index], Multas[index+1:]...)
		}
	}
}

func updateMulta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: updateMulta")
	// Se obtiene el body desde el request y
	// se deserializa en una variable Multa:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Multa Multa
	json.Unmarshal(reqBody, &Multa)
	key := Multa.Id
	// buscar el Multa a actualizar:
	for index, p := range Multas {
		if p.Id == key {
			// actualizar el array:
			Multas[index] = Multa
			break
		}
	}
	json.NewEncoder(w).Encode(Multa)
}

func iniciarServidor() {
	fmt.Println("API REST simple con lenguaje go.")
	ruteador := mux.NewRouter().StrictSlash(true)
	ruteador.HandleFunc("/", homePage)
	ruteador.HandleFunc("/Multas", findAllMultas)
	//el orden de definicion es importante en el manejo de rutas:
	ruteador.HandleFunc("/Multa", createNewMulta).Methods("POST")
	ruteador.HandleFunc("/Multa/{id}", deleteMulta).Methods("DELETE")
	ruteador.HandleFunc("/Multa", updateMulta).Methods("PUT")
	ruteador.HandleFunc("/Multa/{id}", findMultaById)

	log.Fatal(http.ListenAndServe(":8080", ruteador))
}

func main() {
	Multas = []Multa{
		Multa{Id: "1", Nombre: "Exceso de velocidad", Sancion: "Primer grado", Cantidad: 150},
		Multa{Id: "2", Nombre: "Estacionado en una ciclovia", Sancion: "Segundo grado", Cantidad: 200},
	}
	iniciarServidor()
}
