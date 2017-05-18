package main

type Ejercicio2PostModel struct {
	Origen string `json:"origen"`
}

type Ejercicio2Model struct {
	Nombre string  `json:"nombre"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
}

type Ejercicio2ResponseModel struct {
	ResponseModel []Ejercicio2Model `json:"restaurantes"`
}
