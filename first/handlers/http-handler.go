package handlers

import (
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/denstepanov/gobonacci/shared/services"
)

func InitCount(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(chi.URLParam(r, "num"))
	if err != nil {
		panic("Param should be a number!")
	}

	if num > 0 {
		for i := 0; i < num; i++ {
			go func() {
				prev := big.NewInt(0).Bytes()
				curr := big.NewInt(1).Bytes()
				for {
					result, err := Send(services.Count(prev, curr))
					if err != nil {
						log.Fatal(err)
					}
					prev = result.Prev
					curr = result.Curr
				}
			}()
		}
	}
}
