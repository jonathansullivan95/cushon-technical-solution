package main

import (
	"cushon-technical-solution/api"
	"cushon-technical-solution/db"

	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	HTTPServer = &http.Server{
		Handler:     http.DefaultServeMux,
		ReadTimeout: 5 * time.Second,
	}
	port string
)

func main() {
	flag.StringVar(&port, "port", "", "port for http server")
	flag.Parse()

	if port != "" {
		HTTPServer.Addr = fmt.Sprintf(":%s", port)
	}

	err := db.Startup()
	if err != nil {
		fmt.Printf("error starting up db: %v", err)
	}

	http.HandleFunc("/funds", api.GetAvailableFundsHandler)
	http.HandleFunc("/customeraccount/investmentfund/add", api.UpdateCustomerAccountInvestmentFundHandler)
	http.HandleFunc("/customeraccount/get", api.GetCustomerAccountHandler)

	if err := HTTPServer.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("HTTP Server shut down unexpectedly: %v", err)
	}
}
