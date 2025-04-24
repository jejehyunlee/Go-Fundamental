package main

import "fmt"

type alamat struct {
	jalan  string
	kota   string
	negara string
}

// Pointer di Function
func ubahData(alamat2 *alamat) {
	alamat2.negara = "Kamboja"
}

func main() {

	address1 := alamat{
		jalan:  "Bojong menteng",
		kota:   "Bekasi",
		negara: "Indonesia",
	}

	var address2 *alamat = &address1 // pointer
	var address3 *alamat = &address1

	*address2 = alamat{
		jalan:  "Cipendawa",
		kota:   "Jakarta",
		negara: "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *alamat = new(alamat)
	address4.kota, address4.jalan, address4.negara = "Jambi", "Jl. Raya", "Indonesia"
	fmt.Println(address4)

	hasil := &alamat{
		jalan:  "AC Lengkeng",
		kota:   "Jakarta Selatan",
		negara: "",
	}
	ubahData(hasil)

	fmt.Println(hasil)
}
