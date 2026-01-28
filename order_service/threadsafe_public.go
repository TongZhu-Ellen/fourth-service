package main 

func InitOrder(customerID string, productID string, amount int) *Order {

	
	op := orderOf(customerID, productID, amount)

	mu.Lock()
	defer mu.Unlock()

	orderMap[op.OrderID] = op
	outboxQueue = append(outboxQueue, outboxOf(op, ReserveInvent))
	

	return op
}

func FlushOutbox() {

	for len(outboxQueue) > 0 {

    mu.Lock()
    op := outboxQueue[0]
    outboxQueue = outboxQueue[1:]

    if op.Processed {
        mu.Unlock()
        continue
    }

    if op.TryCount >= MAX_TRY {
        DLQ = append(DLQ, op)
        mu.Unlock()
        continue
    }

    op.advanceOnce()
    if !op.Processed {
        outboxQueue = append(outboxQueue, op) // stick it back!
    }
    mu.Unlock()
}


}