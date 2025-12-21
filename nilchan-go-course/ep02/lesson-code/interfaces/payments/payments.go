package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

// PaymentModule struct to store:
// - map with payment description
// - payment method (e.g. paypal, bank, crypto)
type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

// Constructor for PaymentModule struct to prevent hacking
func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Pay() method
// Recieves:
// 1. Desctiption
// 2. Payment
// Returns:
// 1. ID
func (p *PaymentModule) Pay(description string, usd int) int {
	// 1. Make a payment
	// 2. Get ID
	id := p.paymentMethod.Pay(usd)

	// 3. Save payment info
	// - description
	// - cost
	// - is it canceled?
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Canceled:    false,
	}

	p.paymentsInfo[id] = info

	// 4. Return ID
	return id
}

// Cancel() method
// Recieves:
// 1. ID
// Returns:
// -
func (p *PaymentModule) Cancel(id int) {
	// info -> just a copy of p.paymentsInfo[id]
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	// copy changed
	info.Canceled = true

	// change a real value with our copy
	p.paymentsInfo[id] = info
}

// Info() method
// Recieves:
// 1. ID
// Returns:
// Info
func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// AllInfo() method
// Recieves:
// -
// Returns:
// All info
func (p PaymentModule) AllInfo() map[int]PaymentInfo {
	// use a temp map for security reasons
	// user can steal and rewrite our original map if we return it
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))

	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}

	return tempMap
}
