package main

import (
	"ethereum-wallet/address"
	"fmt"
)

func main() {
	address1, err := address.CreateEthAddress()
	if err != nil {
		fmt.Println("get address from privateKey error: ", err)
		return
	}
	fmt.Println(address1.Private, address1.Address, address1.Address)

	address2, err2 := address.PublicKeytoAddress("022505a03b55f896c0948f35a2c63b46f6a4cdb8221164bc27bb9980617dacbce7")
	if err2 != nil {
		fmt.Println("get address from publicKey error: ", err)
		return
	}
	fmt.Println(address2)

	// fmt.Println(1 << 1)
	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(unsafe.Sizeof(arr))

	// p := new(int)
	// fmt.Println(reflect.TypeOf(p))

	// var slice1 []int = make([]int, 0)
	// for i := 0; i < 10000; i++ {
	// 	slice1 = append(slice1, 1)
	// 	fmt.Println("slice1[", i, "]的值为：", slice1[i], "-len为:", len(slice1), "-cap为:", cap(slice1))
	// 	if (i%1000 == 0) && (i != 0) {
	// 		fmt.Println("Press Enter to continue...")
	// 		fmt.Scanln()
	// 	}
	// }

	var slic []int
	fmt.Println(slic)

	var a *A = createA()
	a.SetName("xx")
	fmt.Printf("Address of a: %p\n", &a.Name)
	fmt.Println(a.Name)
	updateName(*a, "yy")
	fmt.Println(a.Name)
}

func updateName(a A, name string) {
	a.Name = name
}

type A struct {
	Name string
}

func createA() *A {
	return &A{}
}

func (a *A) SetName(name string) {
	a.Name = name
}

func reverseArray(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
