package main 


import (
	iv "order_service/invent"
   
)


// domino logic, unsafe!

func (op *Outbox) advanceOnce() {
	
    
	orderp := op.Orderp 
	

	switch op.Action {

	case ReserveInvent:
		reserved, err := iv.IndemReserve(orderp.OrderID, orderp.ProductID, orderp.Amount)

		if err != nil { 
			op.TryCount++ 
		} else if !reserved {
			// failed to reserve!
			finalizeOrder(orderp, Failed, FailedNoInvent)
			op.Processed = true
		} else {
			outboxQueue = append(outboxQueue, outboxOf(orderp, PayWithLuck))
			op.Processed = true
		}
		
		





	case PayWithLuck:
		paid, err := prcDice(orderp.OrderID, orderp.CustomerID)

		if err != nil { 
			op.TryCount++ 
		} else if !paid {
			finalizeOrder(orderp, Failed, FailedNoLuck)
			outboxQueue = append(outboxQueue, outboxOf(orderp, ReleaseInvent))
			op.Processed = true
		} else {
			finalizeOrder(orderp, Succeeded, "")
			op.Processed = true
		}







	case ReleaseInvent:
		released, err := iv.IndemRelease(orderp.OrderID, orderp.ProductID, orderp.Amount)

		if err != nil || !released {
			op.TryCount++
		} else {
			op.Processed = true
		}

	
	}


	
}


