package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"regexp"
	"strings"
	"time"

	"googlemaps.github.io/maps"

	resizer "github.com/nfnt/resize"
	gini "golang.org/x/image/bmp"
)

func GetDirections(model Ejercicio1PostModel) (response Ejercicio1ResponseModel, err error) {
	response.ResponseModel = make([]Ejercicio1Model, 0)
	if !ValidateAddress(model.Destino) || !ValidateAddress(model.Origen) {
		err = errors.New("Formato de direccion incorrecto.")
		return response, err
	}
	if err != nil {
		return response, err
	}
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyAcd1jNXWTE1U0fBI-At5hstqEV0NvoWfo"))
	if err != nil {
		return response, err
	}
	r := &maps.DirectionsRequest{
		Origin:      model.Origen,
		Destination: model.Destino,
	}
	resp, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return response, err
	}
	for _, element := range resp {
		lat := element.Bounds.NorthEast.Lat
		lon := element.Bounds.NorthEast.Lng
		response.ResponseModel = append(response.ResponseModel, Ejercicio1Model{Lat: lat, Lon: lon})
	}
	return response, nil
}

func GetNearbyRestaurants(model Ejercicio2PostModel) (response Ejercicio2ResponseModel, err error) {
	response.ResponseModel = make([]Ejercicio2Model, 0)
	if !ValidateAddress(model.Origen) {
		err = errors.New("Formato de direccion incorrecto.")
		return response, err
	}
	geocodingApi, err := maps.NewClient(maps.WithAPIKey("AIzaSyDd02xLxxzBlCnBxFmu4dw-zkJnSfRlPEg"))
	if err != nil {
		return response, err
	}
	req := &maps.GeocodingRequest{
		Address: model.Origen,
	}
	res, err := geocodingApi.Geocode(context.Background(), req)
	if err != nil {
		return response, err
	}
	placesApi, err := maps.NewClient(maps.WithAPIKey("AIzaSyAbkP1a62We_YU-Znw5KHfsw8n21nKxJ_Y"))
	if err != nil {
		return response, err
	}
	r := &maps.NearbySearchRequest{
		Location: &(res[0].Geometry.Location),
		Radius:   1500,
		Type:     maps.PlaceTypeRestaurant,
	}
	resp, err := placesApi.NearbySearch(context.Background(), r)
	if err != nil {
		return response, err
	}
	for _, element := range resp.Results {
		req := &maps.GeocodingRequest{
			Address: element.Name,
		}
		res, err := geocodingApi.Geocode(context.Background(), req)
		time.Sleep(time.Second)
		if err != nil {
			return response, err
		}

		lat := res[0].Geometry.Bounds.NorthEast.Lat
		lon := res[0].Geometry.Bounds.NorthEast.Lng
		response.ResponseModel = append(response.ResponseModel, Ejercicio2Model{Nombre: element.Name, Lat: lat, Lon: lon})
	}
	return response, nil
}

func GrayScale(model Ejercicio3PostModel) (response Ejercicio3ResponseModel) {
	bitmap := DecodeBitmap(model.Data)
	appendStr := "(blanco y negro)."
	response.Nombre = ModifyImageName(model.Nombre, appendStr)
	bounds := bitmap.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewRGBA64(bitmap.Bounds())
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := bitmap.At(x, y)
			r, g, b, a := oldColor.RGBA()
			gs := uint16((0.2126 * float32(r)) + (0.7152 * float32(g)) + (0.0722 * float32(b)))
			gray.Set(x, y, color.RGBA64{gs, gs, gs, uint16(a)})
		}
	}
	response.Data = EncodeBitmap(gray)
	return response
}

func Resize(model Ejercicio4PostModel) (response Ejercicio4ResponseModel) {
	bitmap := DecodeBitmap(model.Data)
	appendStr := "(reducida)."
	rect := bitmap.Bounds()
	if rect.Max.X <= model.Tamano.Ancho && rect.Max.Y <= model.Tamano.Alto {
		response.Nombre = ModifyImageName(model.Nombre, "(No Modificada).")
		response.Data = model.Data
		return response
	}
	response.Nombre = ModifyImageName(model.Nombre, appendStr)
	//demasiada cosa para hacer resizing, no quiero fatality en el examen D:
	modifiedBitmap := resizer.Resize(uint(model.Tamano.Ancho), uint(model.Tamano.Alto), bitmap, resizer.NearestNeighbor)
	response.Data = EncodeBitmap(modifiedBitmap)
	return response
}

func ModifyImageName(name string, modification string) (modifiedName string) {
	var buffer bytes.Buffer
	stringsArray := strings.Split(name, ".")
	buffer.WriteString(stringsArray[0])
	buffer.WriteString(modification)
	buffer.WriteString(stringsArray[1])
	modifiedName = buffer.String()
	return modifiedName
}
func DecodeBitmap(b64String string) (bitmap image.Image) {
	bytesarray, err := base64.StdEncoding.DecodeString(b64String)
	if err != nil {
		fmt.Println(err)
	}
	bytesReader := bytes.NewReader(bytesarray)
	bitmap, err = gini.Decode(bytesReader)
	if err != nil {
		fmt.Println(err)
	}
	return bitmap
}

func EncodeBitmap(bitmap image.Image) (b64String string) {
	var bytesBuffer bytes.Buffer
	bytesWriter := bufio.NewWriter(&bytesBuffer)
	gini.Encode(bytesWriter, bitmap)
	err := bytesWriter.Flush()
	if err != nil {
		fmt.Println(err)
	}
	resizedBytes := bytesBuffer.Bytes()
	b64String = base64.StdEncoding.EncodeToString(resizedBytes)
	return b64String
}

func ValidateAddress(address string) (isValid bool) {
	//http://stackoverflow.com/questions/9397485/regex-street-address-match
	//¯\_(ツ)_/¯
	isValid, _ = regexp.MatchString("..*,..*", address)
	return isValid
}
