package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	/**
	Na alinha abaixo, o Go usa o DefaultServeMux como o multiplexer de solicitações.

	Então, o que é um multiplexer (também frequentemente chamado de mux ou router)?

	No contexto de servidores web, um multiplexer de solicitações é uma ferramenta que distribui as
	solicitações HTTP recebidas para seus respectivos manipuladores com base em critérios como o método HTTP (GET, POST, etc.)
	e o caminho da URL. Isso permite que seu aplicativo responda de maneira diferente a diferentes URLs ou diferentes tipos
	de solicitações.
	*/
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := req.URL.Query().Get("cep")
	if cepParam == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	viaCEP, err := buscaCepAPI(cepParam)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	// O código comentado abaixo converte a struct ViaCep para um json e retorna na linha w.Write(result)
	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	// Na linha abaixo, fazemos o mesmo que o código comentando acima, porém, a função json json.NewEncoder(rw).Encode converte a struct ViaCep para json e coloca esse json dentro do nosso Response !
	json.NewEncoder(rw).Encode(viaCEP)
}

func buscaCepAPI(cep string) (*ViaCEP, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	return readResponseBody(res)
}

func readResponseBody(res *http.Response) (*ViaCEP, error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var data ViaCEP
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
