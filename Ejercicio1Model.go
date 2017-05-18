package main

type Ejercicio1PostModel struct {
	Origen  string `json:"origen"`
	Destino string `json:"destino"`
}

type Ejercicio1Model struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Ejercicio1ResponseModel struct {
	ResponseModel []Ejercicio1Model `json:"ruta"`
}
