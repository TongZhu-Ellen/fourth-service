package main

import (
	"sync"
	"fmt"
	
	
    "google.golang.org/grpc"

	iv "order_service/invent"
	pb "luckDice_service/pb"
	
)

var (
	orderMap map[string]*Order = make(map[string]*Order)
	outboxQueue []*Outbox
	DLQ []*Outbox


	mu sync.Mutex

	luckDiceClient pb.LuckDiceServiceClient
	

)

const MAX_TRY = 3




func main() {


	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil { panic(err) }
	defer conn.Close()

	luckDiceClient = pb.NewLuckDiceServiceClient(conn)



	iv.AddInvent("testProduct", 5)

	for i := 1; i <= 10; i++ {
		op := InitOrder("testCustomer", "testProduct", 1)
		FlushOutbox()
		fmt.Printf("%+v\n", op)

	}
	
	

	

	
	


	





}
