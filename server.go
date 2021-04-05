package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const port = ":5500"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/products/{fetchCountPercentage}", products).Methods("GET")

	fmt.Println("Serving @ http://127.0.0.1" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root page"))
}

func products(w http.ResponseWriter, r *http.Request) {
	fetchCountPercentage, errInput := strconv.ParseFloat(mux.Vars(r)["fetchCountPercentage"], 64)

	fetchCount := 0

	if errInput != nil {
		fmt.Println(errInput.Error())
	} else {
		fetchCount = int(float64(len(productList)) * fetchCountPercentage / 100)
		if fetchCount > len(productList) {
			fetchCount = len(productList)
		}
	}

	// write response
	jsonList, err := json.Marshal(productList[0:fetchCount])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(jsonList)
	}
}

type product struct {
	Name  string
	Price float64
	Count int
}

var productList = []product{
	product{"p1", 15000, 30},
	product{"p2", 25000, 10},
	product{"p3", 35000, 320},
	product{"p4", 45000, 730},
	product{"p5", 25000, 340},
	product{"p6", 55000, 300},
	product{"p7", 35000, 230},
	product{"p8", 25000, 130},
	product{"p9", 65000, 10},
	product{"p10", 15000, 20},
}
