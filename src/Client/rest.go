package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type passagem struct {
	id      int
	data    string
	origem  string
	destino string
	numero  int
	tipo    int
}

/*
{
	"data" : "oi",
	"origem" : "curitiba",
	"destino" : "longe",
	"numero" : "10",
	"tipo" : "ida"
  }

type
*/

func main() {
	url := "http://localhost/passagens:8080"
	//"http://localhost/hotel:8080"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var p passagem

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
	}

	/* exemplo de envio json

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	*/
}

//Funcoes

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("****Bem Vindo****")
	fmt.Println("")
	fmt.Println("1- Passagens")
	fmt.Println("2- Hoteis")
	fmt.Println("0- Sair do Programa")
}
