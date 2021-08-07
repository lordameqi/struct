package main

import "fmt"

type person struct {
	nama string
	hobi string
	school
	teachers map[string]string
}
type school struct {
	nama       string
	tahunMulai int
	tahunAkhir int
}
type teachers struct {
	namaKelas string
	namaGuru  string
}

func main() {

	//huy  := 3
	// var depan, belakang, gelar string = "Rachmad", "Fadillah", "Ganteng"
	// fmt.Println(depan, belakang, gelar)
	// var decimal = 2.62

	// fmt.Printf("bilangan desimal: %f\n", decimal)
	// fmt.Printf("bilangan desimal: %.3f\n", decimal)
	// fmt.Println(decimal)

	// var exist bool = true
	// fmt.Printf("exist? %t \n", exist)

	// var nilai = 4 + 7/5
	// var hasil string
	// if nilai%2 == 0 {
	// 	hasil = "Genap"
	// } else {
	// 	hasil = "Ganjil"
	// }
	// fmt.Println(hasil)

	// nilai := 3

	// if nilai == 10 {
	// 	fmt.Println("Nilai anda adalah: ", nilai, " Nilai Sempurna")
	// } else if nilai > 8 && nilai < 10 {
	// 	fmt.Print("Nilai anda adalah: ", nilai, " Memuaskan")
	// } else if nilai > 4 && nilai < 8 {
	// 	fmt.Print("Nilai anda adalah: ", nilai, " Tingkatkan lagi")
	// } else {
	// 	fmt.Print("Nilai anda adalah: ", nilai, " Gagal")
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// var arraykan [3]string
	// arraykan[0] = "Aku"
	// arraykan[1] = "Cinta"
	// arraykan[2] = "Kamu"
	// fmt.Println("Isi elemen adalah ", arraykan)
	// fmt.Println("jumlah elemen adalah ", len(arraykan))
	// i := 42
	// f := 3.142
	// g := 0.867 + 0.5i

	// fmt.Println("tipe datanya", reflect.TypeOf(i))
	// fmt.Println("tipe datanya", reflect.TypeOf(f))
	// fmt.Println("tipe datanya", reflect.TypeOf(g))

	s2 := person{
		nama: "Rachmad Fadillah",
		hobi: "Tidur",
		school: school{
			nama:       "UIN SUSKA RIAU",
			tahunMulai: 2015,
			tahunAkhir: 2020,
		},
		teachers: map[string]string{
			"namaKelas": "Teknik Informatika",
			"namaGuru":  "Febiyanto",
		},
	}

	huy(s2)
}
func huy(s person) {
	fmt.Printf("%+v", s)
}
