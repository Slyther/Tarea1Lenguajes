package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sieg Heil!")
}

func Ejercicio1(w http.ResponseWriter, r *http.Request) {
	var ejercicio1PostModel Ejercicio1PostModel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &ejercicio1PostModel); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	if len(ejercicio1PostModel.Origen) == 0 || len(ejercicio1PostModel.Destino) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err = json.NewEncoder(w).Encode("Formato no correcto."); err != nil {
			panic(err)
		}
		return
	}
	response, err := GetDirections(ejercicio1PostModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func Ejercicio2(w http.ResponseWriter, r *http.Request) {
	var ejercicio2PostModel Ejercicio2PostModel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &ejercicio2PostModel); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	if len(ejercicio2PostModel.Origen) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err = json.NewEncoder(w).Encode("Formato no correcto."); err != nil {
			panic(err)
		}
		return
	}
	response, err := GetNearbyRestaurants(ejercicio2PostModel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func Ejercicio3(w http.ResponseWriter, r *http.Request) {
	var ejercicio3PostModel Ejercicio3PostModel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10485760))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &ejercicio3PostModel); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode("Formato no correcto."); err != nil {
			panic(err)
		}
		return
	}
	if len(ejercicio3PostModel.Nombre) == 0 || len(ejercicio3PostModel.Data) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err := json.NewEncoder(w).Encode("Formato no correcto."); err != nil {
			panic(err)
		}
		return
	}
	response := GrayScale(ejercicio3PostModel)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func Ejercicio4(w http.ResponseWriter, r *http.Request) {
	var ejercicio4PostModel Ejercicio4PostModel
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10485760))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &ejercicio4PostModel); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	if len(ejercicio4PostModel.Nombre) == 0 || len(ejercicio4PostModel.Data) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err := json.NewEncoder(w).Encode("Formato no correcto."); err != nil {
			panic(err)
		}
		return
	}
	response := Resize(ejercicio4PostModel)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
