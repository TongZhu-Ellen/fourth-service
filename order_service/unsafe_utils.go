package main

import (

	
	
	"time"
	"context"


	"github.com/google/uuid"

	pb "luckDice_service/pb"
	
	


	
)


// util 函数没有性能安全保障！只是方便我自己复用的工具函数！


func orderOf(customerID string, productID string, amount int) *Order {

	return &Order{
		
		OrderID: uuid.NewString(),
		CustomerID: customerID,
		ProductID:  productID,
		Amount:     amount,

		Status:         Pending,
		
	}
}



func finalizeOrder(op *Order, status OrderStatus, failureType FailureType) {
	op.Status = status
	op.FailureType = failureType
}



func outboxOf(orderp *Order, action Action) *Outbox {

	op := &Outbox{
		Orderp: orderp,
		Action: action,
		Processed: false,
		TryCount: 0,
	}

	return op
}


func prcDice(orderID string, customerID string) (bool, error) {
	// 设置超时时间，防止调用挂住
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 调用 IndemPay
	resp, err := luckDiceClient.IndemPay(ctx, &pb.RequestMsg{
		OrderID:    orderID,
		CustomerID: customerID,
	})
	
	if err != nil { return false, err}

	return resp.Lucky, nil

	
}











