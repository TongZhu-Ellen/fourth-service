package invent

import (
	"sync"
)

var (
	reservedp map[string]int = make(map[string]int)
	remainedp map[string]int = make(map[string]int)
	

	reservation map[string]bool = make(map[string]bool)

	mu sync.Mutex
)


func AddInvent(productID string, amount int) {

	remainedp[productID] += amount

}




// return reserved
func IndemReserve(orderID string, productID string, amount int) (bool, error) {

	

	mu.Lock()
	defer mu.Unlock()
	

	if reservation[orderID] {
		// reserved
		return true, nil
	}

	if remainedp[productID] < amount {
		// no enough amount
		return false, nil
	}

	reservedp[productID] += amount
	remainedp[productID] -= amount

	reservation[orderID] = true

	return true, nil

}

// return reserved
func IndemRelease(orderID string, productID string, amount int) (bool, error) {

	mu.Lock()
	defer mu.Unlock()

	if !reservation[orderID] {
		return true, nil
	}

	if reservedp[productID] < amount {
		panic("Reservation logic broke!!!")
	}

	reservedp[productID] -= amount
	remainedp[productID] += amount

	reservation[orderID] = false

	return true, nil

}


