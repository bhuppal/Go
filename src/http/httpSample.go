package httpSample

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("Successfully performed http request")
}

func httpSample() {
	fmt.Println("Hello World!!")
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}