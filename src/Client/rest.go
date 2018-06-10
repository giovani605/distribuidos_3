package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type tipo struct {
	id
	data
	origem
	destino
	numero
	tipo
}

type


func main() {
	url := "http://localhost:8080"
	resposta, erro := http.Get(url)
	if erro != nil {
		panic(erro)
	}
	defer resposta.Body.Close()

	body, err := ioutil.ReadAll(resposta.Body)
	if erro != nil {
		panic(erro)
	}
	var p Payload

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

}
