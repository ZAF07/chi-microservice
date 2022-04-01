package cmd

// WAIT
import (
	"log"
	"net"
	"time"
)

// Takes in RabbitMQ host string and connects to RabbitMQ, returns RabbitMQ Connection instance
func WaitForService(host string) {
	log.Println("Waiting for host %s", host)

	for {
		log.Println("testing cinnection to %s", host)
		conn, err := net.Dial("tcp", host)
		if err == nil {
			_ = conn.Close()
			log.Printf("%s is up!", host)
			return
		}
		time.Sleep(time.Millisecond *500)
	}
}