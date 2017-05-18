package main

type Ejercicio4PostModel struct {
	Nombre string     `json:"nombre"`
	Data   string     `json:"data"`
	Tamano SizeParams `json:"tamaño"`
}

type SizeParams struct {
	Alto  int `json:"alto"`
	Ancho int `json:"ancho"`
}

type Ejercicio4ResponseModel struct {
	Nombre string `json:"nombre"`
	Data   string `json:"data"`
}
