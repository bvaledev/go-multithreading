package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func request[T ApiCep | ViaCep](ch chan<- T, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var generic T
	json.NewDecoder(res.Body).Decode(&generic)

	ch <- generic
}

func main() {
	fmt.Printf("Informe o seu cep (ex: 22470-050): ")
	var cep string
	_, err := fmt.Scanln(&cep)
	if err != nil {
		panic(err)
	}

	chApiCep := make(chan ApiCep)
	chViaCep := make(chan ViaCep)

	go request(chApiCep, fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep))
	go request(chViaCep, fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))

	select {
	case cep := <-chApiCep:
		fmt.Printf("ApiCep: %+v\n", cep)
	case cep := <-chViaCep:
		fmt.Printf("ViaCep: %+v\n", cep)
	case <-time.After(time.Second):
		println("Timeout")
	}
}
