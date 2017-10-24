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
	
	BYTE     = 1.0
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE
	PETABYTE = 1024 * TERABYTE
	EXABYTE = 1024 * PETABYTE
	ZETABYTE = 1024 * EXABYTE
	YOTTABYTE = 1024 * ZETABYTE
)

var (
	unitList = []string{UNIT_BIT, UNIT_BYTE, UNIT_KILOBYTE, UNIT_MEGABYTE, UNIT_GIGABYTE, UNIT_TERABYTE, UNIT_PETABYTE, UNIT_EXABYTE, UNIT_ZETABYTE, UNIT_YOTTABYTE}
)

func ConvertDataSize(value float64, from, to string) float64 {

	if from == to {
		return value
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

	if idxTo > idxFrom {
		diff := idxTo-idxFrom
		for diff > 0{
			valueToReturn = valueToReturn / ( UNIT_BASE_VALUE )
			diff--
		}	
	} else {
		diff := idxFrom-idxTo
		for diff > 0{
			valueToReturn = valueToReturn * (UNIT_BASE_VALUE)
			diff--
		}
	}
	return valueToReturn
}

/*func main() {
	fmt.Printf(" %0.2f%s\n", ConvertDataSize(float64(6831826696), UNIT_BYTE, UNIT_GIGABYTE), UNIT_GIGABYTE)
	fmt.Printf(" %0.2f%s\n", ConvertDataSize(float64(6831826696), UNIT_BYTE, UNIT_MEGABYTE), UNIT_MEGABYTE)
	fmt.Printf(" %0.2f%s\n", ConvertDataSize(float64(6831826696), UNIT_BYTE, UNIT_KILOBYTE), UNIT_KILOBYTE)

	fmt.Printf(" %0.2f%s\n", ConvertDataSize(float64(1), UNIT_GIGABYTE, UNIT_KILOBYTE), UNIT_KILOBYTE)
	fmt.Printf(" %0.2f%s\n", ConvertDataSize(float64(1), UNIT_MEGABYTE, UNIT_KILOBYTE), UNIT_KILOBYTE)
	
	fmt.Printf(" %0.2f%s\n", ConvertDataSize(float64(1), UNIT_BYTE, UNIT_BIT), UNIT_BYTE)

}*/

