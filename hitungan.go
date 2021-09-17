package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type segitiga struct {
	alas   float64
	tinggi float64
}

type persegi struct {
	panjang float64
	lebar   float64
}

func (s segitiga) luas() float64 {
	hasil := s.alas * s.tinggi / 2
	return hasil
}

func (s segitiga) keliling() float64 {
	a := (s.alas / 2) * (s.alas / 2)
	b := s.tinggi * s.tinggi
	c := math.Sqrt(a + b)
	return c * 3
}

func (p persegi) luas() float64 {
	hasil := p.panjang * p.lebar
	return hasil
}

func (p persegi) keliling() float64 {
	hasil := 2 * (p.panjang + p.lebar)
	return hasil
}

func main() {
	var x = true
	var input int
	var s segitiga
	var p persegi
	var hasil hitung
	for x {
		fmt.Println("===MENU===")
		fmt.Println("1. Luas Segitiga")
		fmt.Println("2. Keliling Segitiga")
		fmt.Println("3. Luas Persegi/ Persegi Panjang")
		fmt.Println("4. Keliling Persegi/ Persegi Panjang")
		fmt.Println("5. Keluar Program")
		fmt.Print("Masukkan Pilihan Anda: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Print("Masukkan Alas Segitiga: ")
			fmt.Scan(&s.alas)
			fmt.Print("Masukkan Tinggi Segitiga: ")
			fmt.Scan(&s.tinggi)
			hasil = segitiga{s.alas, s.tinggi}
			fmt.Printf("Luas Segitiga: %.2f\n\n", hasil.luas())
		case 2:
			fmt.Print("Masukkan Alas Segitiga: ")
			fmt.Scan(&s.alas)
			fmt.Print("Masukkan Tinggi Segitiga: ")
			fmt.Scan(&s.tinggi)
			hasil = segitiga{s.alas, s.tinggi}
			fmt.Printf("Sisi Miring = %.2f\n", hasil.keliling()/3)
			fmt.Printf("Keliling Segitiga (3 x sisi Miring): %.2f\n\n", hasil.keliling())
		case 3:
			fmt.Print("Masukkan Panjang Persegi: ")
			fmt.Scan(&p.panjang)
			fmt.Print("Masukkan Lebar Persegi: ")
			fmt.Scan(&p.lebar)
			hasil = persegi{p.panjang, p.lebar}
			fmt.Printf("Luas Persegi: %.2f\n\n", hasil.luas())
		case 4:
			fmt.Print("Masukkan Panjang Persegi: ")
			fmt.Scan(&p.panjang)
			fmt.Print("Masukkan Lebar Persegi: ")
			fmt.Scan(&p.lebar)
			hasil = persegi{p.panjang, p.lebar}
			fmt.Printf("Luas Persegi: %.2f\n\n", hasil.keliling())
		case 5:
			fmt.Println("PROGRAM AKAN DITUTUP..")
			x = false
		}

	}

}
