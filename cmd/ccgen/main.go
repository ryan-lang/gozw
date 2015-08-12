package main

import "github.com/helioslabs/gozw/ccgen"

func main() {

	gen, err := ccgen.NewGenerator()
	if err != nil {
		panic(err)
	}

	// devices, err := gen.GenDevices()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(devices)

	err = gen.GenCommandClasses()
	if err != nil {
		panic(err)
	}
}
