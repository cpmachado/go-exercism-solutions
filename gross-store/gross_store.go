package gross

// Units stores the Gross Store unit measurement.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to the customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if exists {
		bill[item] += value
	}
	return exists
}

// RemoveItem removes an item from the customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	bill_value, bill_exists := bill[item]
	unit_value, unit_exists := units[unit]
	if !bill_exists || !unit_exists || bill_value < unit_value {
		return false
	}
	if bill_value == unit_value {
		delete(bill, item)
	} else {
		bill[item] = bill_value - unit_value
	}
	return true
}

// GetItem returns the quantity of item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
