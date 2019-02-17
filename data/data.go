package data

// FedTaxData holds data for the federal taxes while we iterate towards getting this data into dynamoDB
var FedTaxData = []byte(`{ 
    "rates": [0.1, 0.12, 0.22, 0.24, 0.32, 0.35, 0.37], 
    "brackets": [0, 9526, 38701, 82501, 157501, 200001, 500001]
}`)

// StateTaxData holds data for each state while we iterate towards getting this into dynamoDB
var StateTaxData = []byte(`{
    "CA":{
        "rates": [0.01, 0.02, 0.04, 0.06, 0.08, 0.093, 0.103, 0.113, 0.123, 0.133],
        "brackets": [0, 8015, 19001, 29989, 41629, 52612, 268750, 322499, 537498, 1000000]
    }
}`)
