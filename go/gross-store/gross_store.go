package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	storeUnits := make(map[string]int)
	storeUnits["quarter_of_a_dozen"] = 3
	storeUnits["half_of_a_dozen"] = 6
	storeUnits["dozen"] = 12
	storeUnits["small_gross"] = 120
	storeUnits["gross"] = 144
	storeUnits["great_gross"] = 1728

	return storeUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billValue, billExists := bill[item]
	unitValue, unitExists := units[unit]

	if !billExists || !unitExists {
		return false
	}

	valueChange := billValue - unitValue

	if valueChange < 0 {
		return false
	}

	if valueChange == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = valueChange
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
