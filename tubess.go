package main

import "fmt"

//maksimal array
const NMAX int = 123

//tipe data Sampah yang berisi nama, banyak dan daur ulang sampah
type Sampah struct{
	namaSampah	    string	
	banyakSampah	int
	daurUlang 	 	string
}
type tabSampah [NMAX]Sampah //array sampah

//tipe data Sistem yang berisi banyak data sampah yang di input, nama jenis sampah, dan juga data sampah
type Sistem struct{
	nSampah 		int 
	data 			tabSampah
	namaJenis		string
}
type tabSistem [3]Sistem // array sampah yang dimana hanya memiliku 3 jenis


func main(){
	//deklarasi variabel
	var sampah 		tabSistem
	var chooseMenu int

	//nama jenis
	sampah[0].namaJenis = "Organik"
	sampah[1].namaJenis = "Non-Organik"
	sampah[2].namaJenis = "B3"

	//tampilan menu apk
	for {
		menuDisplay()
		fmt.Scan(&chooseMenu)
		fmt.Println("")
		switch chooseMenu{
		case 1 :
			fiturAddData(&sampah)			
		case 2 :
			fiturEditData(&sampah)
		case 3 :
			fiturRemoveData(&sampah)
		case 4 :
			fiturAddRecycleData(&sampah)
		case 5 :
			fiturPrintData(&sampah)
		case 6 :
			fmt.Print("Belum ada")
		case 7 :
			fmt.Print("Belum ada")
		case 8 :
			fiturStatisticData(&sampah)
		case 9 :
			return
		}
	}
}

func fiturAddData(s *tabSistem){
	var a 			int
	for {
		Header("adddata1")
		printData(s)
		Footer("jenis")
		fmt.Scan(&a)

		//menu
		switch a {
		case 1 :
			addData(s, a-1)
		case 2 :
			addData(s, a-1)
		case 3 :
			addData(s, a-1)
		case 4 :
			return
		}
	}
}

func addData(s *tabSistem, a int){
	//deklarasi variabel
	var nama 		string
	var banyak 		int


	//input nama dan banyak sampah	
	Header("adddata2")
	Footer("nama")
	fmt.Scan(&nama)
	Header("adddata2")
	Footer("banyak")
	fmt.Scan(&banyak)

	//jika  n masih kurang dari NMAX, maka data masih bisa di tambahkan
	// CATATAN : masih harus membuat jika namanya sama, maka jumlahnya di tambahkan
	if s[0].nSampah < NMAX {
		s[a].data[s[a].nSampah] = Sampah{namaSampah: nama, banyakSampah: banyak}
		s[a].nSampah++
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Data penuh!")
	}
}

//func check untuk melihat apakah data yang di tambahkan sudah pernah ada atau tidak
//jika sudah pernah ada, maka jumla data di masukan ke array data yang sudah pernah ada
/*func check(){
	var found bool = false
}*/

func fiturEditData(s *tabSistem){
	var index 		int
	var nama 		string
	var banyak 		int
	var confirm 	int
	var jenis 		int

	Header("edit2")
	printData(s)
	Footer("jenis")
	fmt.Scan(&jenis)
	Header("edit2")
	printData(s)
	Footer("index")
	fmt.Scan(&index)

	if index >= 0 && index < s[jenis-1].nSampah {

		Header("edit1")
		Footer("nama")
		fmt.Scan(&nama)
		Header("edit1")
		Footer("banyak")
		fmt.Scan(&banyak)

		for{
			fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Edit Data   ")
			fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
			fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14d  | %-12s        ║\n","     ",s[jenis-1].namaJenis ,s[jenis-1].data[index].namaSampah, s[jenis-1].data[index].banyakSampah, s[jenis-1].data[index].daurUlang)
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
			fmt.Printf("║%45s%-44s║\n"," V","")
			fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
			fmt.Printf("║ %-3s | %-16s | %-20s | %-14d  | %-12s ║\n","     ",s[jenis-1].namaJenis , nama, banyak, "                   ")
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
			fmt.Printf("║%-89s║\n"," [Confirm] / [Not] : (1/0)")
			fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
			fmt.Printf("╚══> ")
			fmt.Scan(&confirm)

			switch confirm{
			case 1 :
				if index >= 0 && index < s[0].nSampah {
					s[jenis-1].data[index].namaSampah = nama 
					s[jenis-1].data[index].banyakSampah = banyak
					fmt.Println("Data berhasil diubah.")
				} else {
					fmt.Println("Index tidak valid.")
				}
				return
			case 0 :
				fmt.Println("Data tidak berubah :D")
				return
			}
		}
	}else {
		fmt.Println("Index tidak valid.")
	}
}

func fiturRemoveData(s *tabSistem){
	var index 		int
	var a 			int

	fmt.Print("Jenis data yang ingin dihapus: ")
	fmt.Scan(&a)
	fmt.Print("Index data yang ingin dihapus: ")
	fmt.Scan(&index)

	if index >= 0 && index < s[a].nSampah {
		for i := index; i < s[a].nSampah-1; i++ {
			s[a].data[i] = s[a].data[i+1]
		}
		s[a].nSampah--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func fiturAddRecycleData(s *tabSistem){
	//deklarasi variabel
	var index		int
	var metode 		string
	var jenis 			int

	//input
	fmt.Print("Jenis data : ")
	fmt.Scan(&jenis)
	fmt.Print("Index data sampah: ")
	fmt.Scan(&index)
	fmt.Print("Metode daur ulang: ")
	fmt.Scan(&metode)
	
	//
	if index >= 0 && index < s[jenis-1].nSampah {
		s[jenis-1].data[index].daurUlang = metode
		fmt.Println("Metode daur ulang dicatat.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func fiturStatisticData(s *tabSistem){
	total := 0
	didaurUlang := 0
	for i := 0; i < 3; i++{
		for j := 0; j < s[i].nSampah; j++ {
			total += s[i].data[j].banyakSampah
			if s[i].data[j].daurUlang != "" {
				didaurUlang += s[i].data[j].banyakSampah
			}
		}
	}
	fmt.Println("Total Sampah:", total)
	fmt.Println("Total Didaur Ulang:", didaurUlang)
}


func menuDisplay(){
	fmt.Printf("\n\n╔╡%-10s╞══════════════════════════════════╗\n","   MENU")
	fmt.Printf("║══════════════════════════════════════════════║\n")
	fmt.Printf("║%-46s║\n"," 1. Add Data")
	fmt.Printf("║%-46s║\n"," 2. Edit Data")
	fmt.Printf("║%-46s║\n"," 3. Remove Data")
	fmt.Printf("║%-46s║\n"," 4. Add Recycle Data")
	fmt.Printf("║%-46s║\n"," 5. Display Data")
	fmt.Printf("║%-46s║\n"," 6. Sort Data")
	fmt.Printf("║%-46s║\n"," 7. Search Data")
	fmt.Printf("║%-46s║\n"," 8. Display Statistic Data")
	fmt.Printf("║%-46s║\n"," 9. Exit")
	fmt.Printf("║──────────────────────────────────────────────║\n")
	fmt.Printf("║%-46s║\n","Choose (1/2/3/4/5/6/7/8/9) :")
	fmt.Printf("║══════════════════════════════════════════════╝\n")
	fmt.Printf("╚══> ")
}

func fiturPrintData(s *tabSistem){
	Header("display")
	printData(s)
	Footer("")
}

func printData(s *tabSistem){
	for i := 0; i < len(s); i++ {
		for j := 0; j < s[i].nSampah; j++ {
				fmt.Printf("║  %d.%d  | %-16s | %-20s | %-14d  |  %-12s       ║\n",i+1,j,s[i].namaJenis ,s[i].data[j].namaSampah, s[i].data[j].banyakSampah, s[i].data[j].daurUlang)
		}
	}
}
func Header(title string){
	switch title {
	case "adddata1" :
		fmt.Printf("\n\n╔╡%-10s╞═════════════════════════════════════════════════════════════════════════╗\n","   Add Data   ")
		fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")

	case "adddata2" :
		fmt.Printf("\n\n╔╡%-10s╞═════════════════════════════════════════════════════════════════════════╗\n","   Add Data   ")

	case "display" :
		fmt.Printf("\n\n╔╡%-10s╞═════════════════════════════════════════════════════════════════════╗\n","   Display Data   ")
		fmt.Printf("║═══════╤══════════════════╤══════════════════════╤═════════════════╤═════════════════════║\n")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")

	case "edit1" :
		fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Edit Data   ")

	case "edit2" :
		fmt.Printf("\n\n╔╡%-10s╞════════════════════════════════════════════════════════════════════════╗\n","   Edit Data   ")
		fmt.Printf("║ %-3s | %-16s | %-20s | %-14s | %-12s ║\n"," No. ","   Jenis Sampah", "     Nama Sampah", " Banyak Sampah ", " Metode Daur Ulang ")
		fmt.Printf("║═══════╪══════════════════╪══════════════════════╪═════════════════╪═════════════════════║\n")
	}
}
func Footer(title string){
	switch title {

	case "" :
		fmt.Printf("╚═════════════════════════════════════════════════════════════════════════════════════════╝\n")

	case "jenis" :
		fmt.Printf("║═══════╧══════════════════╧══════════════════════╧═════════════════╧═════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Pilih jenis sampah : (1/2/3/4)")
		fmt.Printf("║%-89s║\n"," 1. Organik")
		fmt.Printf("║%-89s║\n"," 2. Non-Organik")
		fmt.Printf("║%-89s║\n"," 3. B3")
		fmt.Printf("║%-89s║\n"," 4. Exit")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")
	
	case "nama" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan nama sampah :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")

	case "banyak" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan banyak sampah :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")
	
	case "index" :
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════║\n")
		fmt.Printf("║%-89s║\n"," Masukan index sampah :")
		fmt.Printf("║═════════════════════════════════════════════════════════════════════════════════════════╝\n")
		fmt.Printf("╚══> ")
	}

}