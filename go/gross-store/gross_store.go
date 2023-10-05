package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}

	// Or...
	// return map[string]int{...}

	// Or...
	// units["quarter_of_a_dozen"] = 3
	// units["half_of_a_dozen"]	= 6
	// units["dozen"] = 12
	// units["small_gross"] = 120
	// units["gross"] = 144
	// units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var bill = map[string]int{}

	// Or...
	// return map[string]int{}
	// return make(map[string]int)

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if value, unitExists := units[unit]; !unitExists {
		return false
	} else {
		if _, itemExists := bill[item]; !itemExists {
			bill[item] = value
		} else {
			bill[item] += value
		}
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, itemExists := bill[item]; !itemExists {
		return false
	}

	if _, unitExists := units[unit]; !unitExists {
		return false
	}

	newQuantity := bill[item] - units[unit]

	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] -= units[unit]
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, itemExists := bill[item]

	if !itemExists {
		return 0, false
	}


	return quantity, true
}
