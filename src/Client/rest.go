package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Passagem struct {
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
	url := "http://localhost:8080/passagens"
	//"http://localhost/hotel:8080"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("vai dar pau 1if")

		panic(err)
	}

	//defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("vai dar pau 2if")
		panic(err)
	}
	fmt.Println(body)

	var dadosJson Passagem
	erro := json.Unmarshal(body, &dadosJson) // aqui acontece a magica!

	bodye := []byte(body)
	fmt.Println(string(bodye))

	if erro != nil {
		fmt.Println("vai dar pau 3if")
		fmt.Println(erro)

		panic(erro)
	}
	fmt.Println(dadosJson)

	//err = json.Unmarshal(body, &p)
	//if err != nil {
	//	panic(err)
	//}
	/*
		slcD := []string{"data", "origem", "destino", "numero", "tipo"}
		slcB, _ := json.Marshal(slcD)
		fmt.Println(string(slcB))
	*/

	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Passagens")
	case 2:
		fmt.Println("Hoteis")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
	}
	/*
		res1D := &response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))
	*/
	var p Passagem

	fmt.Println("Data:")
	fmt.Scan(&p.data)
	fmt.Println("Origem:")
	fmt.Scan(&p.origem)
	fmt.Println("Destino:")
	fmt.Scan(&p.destino)
	fmt.Println("Poltrona:")
	fmt.Scan(&p.numero)
	fmt.Println("Ida (1), Volta(2)")
	fmt.Scan(&p.tipo)
	/*
			res := Passagem{}
			res1D := &Passagem{
				data:    p.data,
				origem:  p.origem,
				destino: p.destino,
				numero:  p.numero,
				tipo:    p.tipo}
			res1B, _ := json.Marshal(res1D)
			fmt.Println(string(res1B))
			json.Unmarshal([]byte(res1B), &res)

			mapD := map[string]int{"apple": 5, "lettuce": 7}
		    mapB, _ := json.Marshal(mapD)
		    fmt.Println(string(mapB))

			fmt.Println(res1B)

			pass, _ := json.Marshal(p)
			fmt.Println(string(pass))
	*/
	//simpleObject := Passagem {id, data, origem,	destino, numero, tipo}
	/*simpleObject := Passagem{p.id, p.data, p.origem, p.destino, p.numero, p.tipo}
	simpleObjectJson, err := json.Marshal(simpleObject)
	if err != nil {
		fmt.Println("Erro!")
	}
	fmt.Print(string(simpleObjectJson))
	var passag Passagem
	err := json.Unmarshal(simpleObjectJson, &passag)
	*/
	/*res1D := &passagem{x
	        data: ,
	        origem: ,
			destino: ,
			numero: ,
			tipo: ,
		}*/
	//pass, _ := json.Marshal(p)
	//fmt.Println(string(pass))

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
	}
	*/
}

//Funcoes

func leComando() int {
	var comandoLido int
	fmt.Println("")
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

func compra() {

}
