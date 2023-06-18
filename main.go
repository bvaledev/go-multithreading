package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func request(ch chan<- []byte, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	ch <- resBody
}

func main() {
	fmt.Printf("Informe o seu cep (ex: 22470-050): ")
	var cep string
	_, err := fmt.Scanln(&cep)
	if err != nil {
		panic(err)
	}

	chApiCep := make(chan []byte)
	chViaCep := make(chan []byte)

	go request(chApiCep, fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep))
	go request(chViaCep, fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))

	select {
	case cep := <-chApiCep:
		fmt.Printf("%s\n", cep)
	case cep := <-chViaCep:
		fmt.Printf("%s\n", cep)
	case <-time.After(time.Second):
		println("Timeout")
	}
}
