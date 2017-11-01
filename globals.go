package main

import "fmt"


// convert unit data size
// value is size
// from actual unit
// to new unit
const (
	UNIT_BIT        = "b"
	UNIT_BYTE       = "B"
	UNIT_KILOBYTE   = "KB"
	UNIT_MEGABYTE   = "MB"
	UNIT_GIGABYTE   = "GB"
	UNIT_TERABYTE   = "TB"
	UNIT_PETABYTE   = "PB"
	UNIT_EXABYTE    = "EB"
	UNIT_ZETABYTE   = "ZB"
	UNIT_YOTTABYTE  = "YB"
	UNIT_BASE_VALUE = float64(1024)
	UNIT_BASE_BIT   = float64(8)

	/*BYTE     = 1.0
	KILOBYTE = UNIT_BASE_VALUE * BYTE
	MEGABYTE = UNIT_BASE_VALUE * KILOBYTE
	GIGABYTE = UNIT_BASE_VALUE * MEGABYTE
	TERABYTE = UNIT_BASE_VALUE * GIGABYTE
	PETABYTE = UNIT_BASE_VALUE * TERABYTE
	EXABYTE = UNIT_BASE_VALUE * PETABYTE
	ZETABYTE = UNIT_BASE_VALUE * EXABYTE
	YOTTABYTE = UNIT_BASE_VALUE * ZETABYTE*/
)

var (
	unitList = []string{UNIT_BIT, UNIT_BYTE, UNIT_KILOBYTE, UNIT_MEGABYTE, UNIT_GIGABYTE, UNIT_TERABYTE, UNIT_PETABYTE, UNIT_EXABYTE, UNIT_ZETABYTE, UNIT_YOTTABYTE}
	mapValidUnit = map[string]bool{UNIT_BIT: true, UNIT_BYTE: true, UNIT_KILOBYTE: true, UNIT_MEGABYTE: true, UNIT_GIGABYTE: true, UNIT_TERABYTE: true, UNIT_PETABYTE: true, UNIT_EXABYTE: true, UNIT_ZETABYTE: true, UNIT_YOTTABYTE:true}

)

func ConvertDataSize(value float64, from, to string)(float64, error) {

	if !mapValidUnit[from]{
		return 0, fmt.Errorf("Invalid type [%s] to convert",from)
	}

	if !mapValidUnit[to]{
		return 0, fmt.Errorf("Invalid type [%s] to convert",to)
	}

	if from == to {
		return value, nil
	}

	valueToReturn := value
	idxFrom := 0
	idxTo := 0
	//find units
	for key, v := range unitList {
		if v == from {
			idxFrom = key
		}
		if v == to {
			idxTo = key
		}
	}
	
	base := UNIT_BASE_VALUE
	// crescent
	if idxTo > idxFrom {
		toBit := false
		if from == UNIT_BIT {
			toBit = true
		}	
	
		diff := idxTo-idxFrom
		for diff > 0 {
			if toBit {
				base = UNIT_BASE_BIT
				toBit = false
			}
			valueToReturn = valueToReturn / ( base )
			diff--
			base = UNIT_BASE_VALUE
		}	
	} else {
		// 
		toBit := false
		if to == UNIT_BIT {
			toBit = true
		}	
		diff := idxFrom-idxTo
		for diff > 0 {
			
			valueToReturn = valueToReturn * (base)
			diff--
			if diff == 1 && toBit {
				base = UNIT_BASE_BIT
			}
		}
	}
	return valueToReturn, nil
}

func main() {

	v, e := ConvertDataSize(float64(1), UNIT_GIGABYTE, UNIT_BIT)

	fmt.Printf(" %s - %0.2f%s\n", e, v, UNIT_BIT)


	v, e = ConvertDataSize(float64(8589934592), UNIT_BIT, UNIT_GIGABYTE)

	fmt.Printf(" %s - %0.2f%s\n", e, v, UNIT_GIGABYTE)
}

