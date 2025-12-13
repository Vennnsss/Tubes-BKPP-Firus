package main

import "fmt"

func main() {
	var tipeKendaraan string
	var jam1Motor, jam1Mobil, jam2Motor, jam2Mobil int
	var jam, menit, tarif, pilihan int
	var stop bool
	stop = false
	jam1Motor = 3000
	jam2Motor = 2000
	jam1Mobil = 5000
	jam2Mobil = 3000
	for !stop {
		fmt.Print("Masukkan tipe kendaraan : ")
		fmt.Scan(&tipeKendaraan)
		for tipeKendaraan != "motor" && tipeKendaraan != "mobil" {
			fmt.Println("Input Invalid, kendaraan hanya motor atau mobil")
			fmt.Print("Masukkan tipe kendaraan : ")
			fmt.Scan(&tipeKendaraan)
		}
		fmt.Print("Masukkan Jam Parkir: ")
		fmt.Scan(&jam)
		for jam < 0 {
			fmt.Println("Input Invalid, jam hanya bisa bernilai 0 atau lebih")
			fmt.Print("Masukkan Jam Parkir: ")
			fmt.Scan(&jam)
		}
		fmt.Print("Masukkan Menit Parkir: ")
		fmt.Scan(&menit)
		for menit < 0 {
			fmt.Println("Input Invalid, menit hanya bisa bernilai 0 atau lebih")
			fmt.Print("Masukkan Menit Parkir: ")
			fmt.Scan(&menit)
		}
		if menit == 0 && jam == 0 {
			fmt.Println("Kendaraan tidak parkir")
		} else {
			if tipeKendaraan == "motor" {
				if menit != 0 && jam == 0 {
					tarif = jam1Motor
				} else {
					tarif = ((jam - 1) * jam2Motor) + jam1Motor
					if menit != 0 {
						tarif += jam2Motor
					}
				}
			} else if tipeKendaraan == "mobil" {
				if menit != 0 && jam == 0 {
					tarif = jam1Mobil
				} else {
					tarif = ((jam - 1) * jam2Mobil) + jam1Mobil
					if menit != 0 {
						tarif += jam2Mobil
					}
				}
			}
			fmt.Printf("Tarif parkir yang harus dibayar adalah Rp. %d dengan tipe kendaraan %s \n", tarif, tipeKendaraan)
		}
		fmt.Printf("\n Apakah ingin memasukkan tarif parkir lagi? \n")
		fmt.Println("1. Iya")
		fmt.Println("2. Tidak")
		fmt.Print("Pilihan (1/2): ")
		fmt.Scan(&pilihan)
		stop = pilihan == 2
	}
	fmt.Println("Program Selesai")
}
