package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZAF07/chi-microservice/pkg/common/cmd"
)

func main () {
	log.Println("STARTING PAYMENTS MICROSERVICE")

	defer log.Println("Closing payments microservice")

	ctx := cmd.Context()

	paymentsinterface := createPaymentsMicroservice()
	if err := paymentsinterface.Run(ctx); err != nil {
		panic(err)
	}
}

func createPaymentsMicroservice() amqp.paymentsinterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	
	paymentsService := payments_app.NewPaymentsService(payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")))

	 paymentsInterface, err := amqp.NewPaymentsInterface(
		 fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		 os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		 paymentsService, 
	 )

	 if err != nil {
		 panic(err)
	 }
	 return paymentsInterface
}