package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var intValue = [4]int{1, 2, 3, 4}
	fmt.Println("before", intValue)

	unsafePtr := unsafe.Pointer(&intValue[0])
	uintPtr := uintptr(unsafePtr) + 2*unsafe.Sizeof(intValue[0])
	fmt.Println("size", unsafe.Sizeof(intValue[0]))
	*(*int)(unsafe.Pointer(uintPtr)) = 6

	//unsafePtr := unsafe.Pointer(&intValue[0])
	//uintPtr := uintptr(unsafePtr) + 2
	//*(*int)(unsafe.Pointer(uintPtr)) = 6

	//unsafePtr := unsafe.Pointer(&intValue[0])
	//*(*int)(uintptr(unsafePtr) + 2) = 6

	//unsafePtr := unsafe.Pointer(&intValue[0])
	//*(*int)(uintptr(unsafePtr) + 2*unsafe.Sizeof(intValue[0])) = 6

	fmt.Println("after", intValue)

	a := struct {}{}
	fmt.Println("a size", unsafe.Sizeof(a))


}
