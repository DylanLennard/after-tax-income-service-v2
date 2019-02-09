package helpers

// This function calculates SSI and medicare related taxes
func CalculateOtherTaxes(income float64, status bool) float64{
	
    medicare_rate := 0.0145
    medicare_upper_rate := 0.0235 // for those making over $200000
	ssi_rate := 0.062 //up to $118,500
	var total_other_tax, ssi_tax, medicare_tax float64

    if !status{
        medicare_rate = medicare_rate * 2
        medicare_upper_rate = medicare_upper_rate * 2
        ssi_rate = ssi_rate * 2
	}

    if income < 118500 {
        ssi_tax = income * ssi_rate
		medicare_tax = income * medicare_rate
	} else {
        ssi_tax = 118500 * ssi_rate
        if income <= 200000{
			medicare_tax = income * medicare_rate
		} else{
			medicare_tax = income * medicare_upper_rate
		}
	}
    total_other_tax = ssi_tax + medicare_tax
    return total_other_tax
}