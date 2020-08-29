package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		lmb := num & 1
		lmb_shift := lmb << (31 - i)
		result = lmb_shift | result
		num >>= 1
	}
	return result
}

func main() {
	var i uint32
	fmt.Scanf("%d", &i)
	fmt.Println(i)
	fmt.Println(reverseBits(i))
}
