package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NoTabaco/nomadcoin/blockchain"
	"github.com/NoTabaco/nomadcoin/utils"
)

var port string

type url string

// type TextMarshaler interface
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See a Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")

	/*
		1.	rw.Header().Add("Content-Type", "application/json")
			b, err := json.Marshal(data)
			utils.HandleErr(err)
			fmt.Fprintf(rw, "%s", b)
	*/

	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
		}
	case "POST":
		{
			var addBlockBody addBlockBody
			utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
			blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
			rw.WriteHeader(http.StatusCreated)
		}
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
