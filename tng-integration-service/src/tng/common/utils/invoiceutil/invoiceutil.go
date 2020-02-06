package invoiceutil

import "strconv"

func GenBarCodeFromInvoiceCode(invoiceCode string) string {
	l := len(invoiceCode)
	if l < 10 {
		return ""
	}
	arr := []string{}
	for i := 0; i < l; i++ {
		arr = append(arr, string(invoiceCode[i]))
	}

	c := 0
	x := 0
	j := 1
	for i := 0; i < l; i++ {
		temp, _ := strconv.Atoi(arr[i])
		c = (temp + '0') * j
		x = x + c
		j += 3
	}
	x = (x%9 + 1) + '0'
	j = 1
	y := 0
	for i := l - 1; i >= 0; i-- {
		temp, _ := strconv.Atoi(arr[i])
		c = (temp + '0') * j
		y = y + c
		j += 1
	}
	y = (y%9 + 1) + '0'
	arrTemp := []string{}

	for i := 0; i < 6; i++ {
		arrTemp = append(arrTemp, arr[i])
	}
	arrTemp = append(arrTemp, string(x))
	arrTemp = append(arrTemp, string(y))
	for i := 6; i < 10; i++ {
		arrTemp = append(arrTemp, arr[i])
	}
	barCode := Encode(arrTemp, l+2)
	output := ""
	for i := 0; i < len(barCode); i++ {
		output = output + barCode[i]
	}
	return output
}

func GetInvoiceCodeFromBarCode(barCode string) string {
	l := len(barCode)
	temp := []string{}
	if l < 12 {
		return ""
	}
	for i := 0; i < l; i++ {
		temp = append(temp, string(barCode[i]))
	}
	invoiceCode := Decode(temp, l)
	output := ""
	if len(invoiceCode) < 12 {
		return ""
	}
	for i := 0; i < len(invoiceCode); i++ {
		if i == 6 || i == 7 {
			continue
		}
		output = output + invoiceCode[i]
	}
	return output
}
