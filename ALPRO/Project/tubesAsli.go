package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func acak() string {
	var (
		hasil string
	)
	// Menginisialisasi generator angka acak dengan seed yang unik
	rand.Seed(time.Now().UnixNano())

	// Menghasilkan lima huruf acak
	randomChars := ""
	for i := 0; i < 4; i++ {
		randomRune := rune(rand.Intn(26) + 65) // Menghasilkan huruf kapital acak
		randomChars += string(randomRune)      // Mengonversi rune menjadi string
	}

	hasil = randomChars
	return hasil
}

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func header() {
	CallClear()
	fmt.Println("* --------------------------------------------------- *")
	fmt.Println("*   Selamat Datang di Aplikasi Konsultasi Kesehatan   *")
	fmt.Println("*                  Deore dan Raihan                   *")
	fmt.Println("*           Algoritma Pemrograman IF-46-08            *")
	fmt.Println("* --------------------------------------------------- *")
}

func ending() {
	CallClear()
	fmt.Println("* ------------------------------------------------------------------------ *")
	fmt.Println("*    Terima Kasih Sudah Menggunakan Aplikasi Konsultasi Kesehatan Kami     *")
	fmt.Println("*                             Deore dan Raihan                             *")
	fmt.Println("*                      Algoritma Pemrograman IF-46-08                      *")
	fmt.Println("* ------------------------------------------------------------------------ *")
}

const nmax int = 1024

type pasien struct {
	nama, NIK string
}

type tabPasien struct {
	arrPasien [nmax]pasien
	n         int
}

type dokter struct {
	nama, NIK string
}

type tabDokter struct {
	arrDokter [nmax]dokter
	n         int
}

type pertanyaan struct {
	tag, tanya    string
	kode          string
	tgl, bln, thn int
}

type tabPertanyaan struct {
	arrTanya [nmax]pertanyaan
	n        int
}

type tanggapan struct {
	tanggapan  string
	Penjawab   string
	kode       string
	pertanyaan string
	tag        string
}

type tabTanggapan struct {
	arrTanggapan [nmax]tanggapan
	n            int
}

/*
? 1 dokter
? 2 pasien
*/

func tanyaRole() string {
	var (
		isi string
	)
	CallClear()
	header()
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("ROLE")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("1. Dokter")
	fmt.Println("2. Pasien")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("3. Kembali")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	fmt.Print("Pilihan : ")
	fmt.Scan(&isi)

	for isi != "1" && isi != "3" && isi != "2" {
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("ROLE TIDAK VALID")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		delay := 3 // detik
		time.Sleep(time.Duration(delay) * time.Second)
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("ROLE")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("1. Dokter")
		fmt.Println("2. Pasien")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("3. Kembali")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&isi)
	}
	return isi
}

/*
? Nama dokter harus diakhiri pake spasi .
*/

func isiDokter(t *tabDokter) {
	var (
		isi      dokter
		tanya    string
		yakin    bool = false
		stop     bool = false
		tampung  string
		namaAsli string = ""
	)
	for !yakin {
		yakin = false
		stop = false
		namaAsli = ""
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("REGISTRASI DOKTER")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("NOTE : Akhiri nama anda dengan \" .\"")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan nama anda : ")
		for !stop {
			fmt.Scan(&tampung)
			namaAsli = namaAsli + tampung + " "
			if tampung == "." {
				stop = true
				isi.nama = namaAsli
			}
		}
		fmt.Print("Masukkan NIK anda : ")
		fmt.Scan(&isi.NIK)
		fmt.Print("Sudah yakin ? (Y/N)")
		fmt.Scan(&tanya)
		if tanya == "Y" || tanya == "y" {
			yakin = true
			t.arrDokter[t.n] = isi
			t.n++
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("AKUN BERHASIL DIBUAT")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 2 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("SELAMAT DATANG DOKTER", isi.nama)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay = 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}
}

/*
? Nama pasiennya harus diakhiri pake spasi .
*/
func isiPasien(t *tabPasien) {
	var (
		isi      pasien
		tanya    string
		yakin    bool = false
		stop     bool = false
		tampung  string
		namaAsli string = ""
	)

	for !yakin {
		yakin = false
		stop = false
		namaAsli = ""
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("REGISTRASI PASIEN")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("NOTE : Akhiri nama anda dengan \" .\"")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan nama anda : ")
		for !stop {
			fmt.Scan(&tampung)
			namaAsli = namaAsli + tampung + " "
			if tampung == "." {
				stop = true
				isi.nama = namaAsli
			}
		}
		fmt.Print("Masukkan NIK anda : ")
		fmt.Scan(&isi.NIK)
		fmt.Print("Sudah yakin ? (Y/N) ")
		fmt.Scan(&tanya)
		if tanya == "Y" || tanya == "y" {
			yakin = true
			t.arrPasien[t.n] = isi
			t.n++
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("AKUN BERHASIL DIBUAT")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 2 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("SELAMAT DATANG", isi.nama)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay = 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}
}

/*
? Yang ada di aplikasi
*/
func tampilkanTagYangAda() {
	fmt.Println("Tag yang tersedia : ")
	fmt.Println("1. Penyakit")
	fmt.Println("2. Gejala")
	fmt.Println("3. Obat")
	fmt.Println("4. Tindakan Medis")
	fmt.Println("5. Diet")
	fmt.Println("6. Kegiatan Fisik")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
}

/*
? pasien pilih tag
? pasien buat pertanyaan
? pertanyaan harus diakhiri sama ?
*/
func bertanya(t *tabPertanyaan) {
	var (
		isi            pertanyaan
		pilih          string
		tanya          string
		tampung        string
		stop           bool
		submit         bool
		stop2          bool
		pertanyaannya  string
		kodePertanyaan string
	)
	for !submit {
		pertanyaannya = ""
		stop = false
		submit = false
		stop2 = false
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("TANYA DONG DOK")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("NOTE : Berikan TAG untuk pertanyaan anda (1-6)")
		fmt.Println("NOTE : Penulisan Tanggal (dd mm yyyy)")
		fmt.Println("NOTE : akhiri pertanyaan anda dengan \" ?\"")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		tampilkanTagYangAda()
		fmt.Print("Masukkan TAG untuk pertanyaan anda : ")
		fmt.Scan(&pilih)
		for !stop2 {
			if pilih == "1" {
				t.arrTanya[t.n].tag = "Penyakit"
				stop2 = true
			} else if pilih == "2" {
				t.arrTanya[t.n].tag = "Gejala"
				stop2 = true
			} else if pilih == "3" {
				t.arrTanya[t.n].tag = "Obat"
				stop2 = true
			} else if pilih == "4" {
				t.arrTanya[t.n].tag = "Tindakan Medis"
				stop2 = true
			} else if pilih == "5" {
				t.arrTanya[t.n].tag = "Diet"
				stop2 = true
			} else if pilih == "6" {
				t.arrTanya[t.n].tag = "Kegiatan Fisik"
				stop2 = true
			} else {
				stop2 = false
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("TAG TIDAK VALID")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				delay := 3 // detik
				time.Sleep(time.Duration(delay) * time.Second)
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("TANYA DONG DOK")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("NOTE : Berikan TAG untuk pertanyaan anda (1-6)")
				fmt.Println("NOTE : akhiri pertanyaan anda dengan \" ?\"")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanTagYangAda()
				fmt.Print("Masukkan TAG untuk pertanyaan anda : ")
				fmt.Scan(&pilih)

			}
		}

		fmt.Print("Masukkan pertanyaan anda : ")
		for !stop {
			fmt.Scan(&tampung)
			pertanyaannya = pertanyaannya + tampung + " "
			if tampung == "?" {
				stop = true
			}
		}
		fmt.Print("Masukkan tanggal pertanyaan dibuat : ")
		fmt.Scan(&isi.tgl)
		fmt.Scan(&isi.bln)
		fmt.Scan(&isi.thn)

		for isi.tgl <= 0 || isi.tgl >= 31 || isi.bln <= 0 || isi.bln >= 13 || isi.thn <= 0 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("Tanggal tidak valid")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("TANYA DONG DOK")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("NOTE : Berikan TAG untuk pertanyaan anda (1-6)")
			fmt.Println("NOTE : Penulisan Tanggal (dd mm yyyy)")
			fmt.Println("NOTE : akhiri pertanyaan anda dengan \" ?\"")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			tampilkanTagYangAda()
			fmt.Print("Masukkan tanggal pertanyaan dibuat : ")
			fmt.Scan(&isi.tgl)
			fmt.Scan(&isi.bln)
			fmt.Scan(&isi.thn)
		}

		fmt.Print("Sudah yakin ? (Y/N)")
		fmt.Scan(&tanya)
		t.arrTanya[t.n].tanya = pertanyaannya
		if tanya == "Y" || tanya == "y" {
			kodePertanyaan = acak()
			t.arrTanya[t.n].kode = kodePertanyaan
			t.arrTanya[t.n].tgl = isi.tgl
			t.arrTanya[t.n].bln = isi.bln
			t.arrTanya[t.n].thn = isi.thn
			submit = true
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("PERTANYAAN ANDA SUDAH TERKIRIM")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			t.n++
		}
	}
}

/*
? sorting descending terhadap tanggal selection sort
*/
func sortingDescendingPertanyaan(t *tabPertanyaan) {
	var (
		idx  int
		pass int
		temp pertanyaan
	)
	for pass = 1; pass < t.n; pass++ {
		idx = pass - 1
		for i := pass; i < t.n; i++ {
			if t.arrTanya[idx].thn > t.arrTanya[i].thn ||
				(t.arrTanya[idx].thn == t.arrTanya[i].thn && t.arrTanya[idx].bln > t.arrTanya[i].bln) ||
				(t.arrTanya[idx].thn == t.arrTanya[i].thn && t.arrTanya[idx].bln == t.arrTanya[i].bln && t.arrTanya[idx].tgl > t.arrTanya[i].tgl) {
				idx = i
			}
		}
		temp = t.arrTanya[idx]
		t.arrTanya[idx] = t.arrTanya[pass-1]
		t.arrTanya[pass-1] = temp
	}
}

/*
?sorting Ascendingterhdap tanggal insertion sort
*/
func sortingAscendingPertanyaan(t *tabPertanyaan) {
	var (
		i, j int
		temp pertanyaan
	)

	for i = 1; i < t.n; i++ {
		temp = t.arrTanya[i]
		j = i - 1

		for j >= 0 && (t.arrTanya[j].thn < temp.thn ||
			(t.arrTanya[j].thn == temp.thn && t.arrTanya[j].bln < temp.bln) ||
			(t.arrTanya[j].thn == temp.thn && t.arrTanya[j].bln == temp.bln && t.arrTanya[j].tgl < temp.tgl)) {
			t.arrTanya[j+1] = t.arrTanya[j]
			j = j - 1
		}

		t.arrTanya[j+1] = temp
	}
}

/*
? sorting berdasarkan KODE secara ascending
*/
func sortingKodePertanyaan(t *tabPertanyaan) {
	var (
		idx  int
		pass int
		temp pertanyaan
	)
	for pass = 1; pass < t.n; pass++ {
		idx = pass - 1
		for i := pass; i < t.n; i++ {
			if t.arrTanya[idx].kode > t.arrTanya[i].kode {
				idx = i
			}
		}
		temp = t.arrTanya[idx]
		t.arrTanya[idx] = t.arrTanya[pass-1]
		t.arrTanya[pass-1] = temp
	}
}

/*
? Binary Search berdasarkan KODE pertanyaan
? Ngembaliin indeks KODE pertanyaan
*/
func binarySearch(t tabPertanyaan, target string) int {
	var (
		kiri, kanan, tengah int
		ketemu              int = -1
	)
	sortingKodePertanyaan(&t)
	kiri = 0
	kanan = t.n - 1
	for ketemu == -1 && kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if target > t.arrTanya[tengah].kode {
			kiri = tengah + 1
		} else if target < t.arrTanya[tengah].kode {
			kanan = tengah - 1
		} else {
			ketemu = tengah
		}
	}
	return ketemu
}

/*
? menampilkan pertanyaan yang belum diurutkan
*/
func tampilkanPertanyaan(t tabPertanyaan) {
	var j int = 1
	if t.n == 0 {
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("FORUM")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("PERTANYAAN")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Belum ada pertanyaan")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	} else {
		fmt.Println("FORUM")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("PERTANYAAN")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

		fmt.Printf("%-3s %-16s %-5s %-9s  %-50s\n", "NO", "TAG", "KODE", "TANGGAL", "Pertanyaan")
		for i := 0; i < t.n; i++ {
			fmt.Printf("%-3d %-16s %-5s %-2d %-2d %-3d %-50s\n", j, t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
			j++
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	}
}

/*
? menampilkkan tanggapan yang sudah dikirim oleh pasien atau dokter
*/

func tampilkanTanggapan(a *tabDokter, b *tabPasien, c *tabPertanyaan, d *tabTanggapan) {
	if d.n == 0 {
		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("TANGGAPAN")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Belum ada tanggapan")
	} else {
		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("TANGGAPAN")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Printf("%-16s %-5s %-40s %-50s %-50s\n", "TAG ", "KODE", "Penjawab", "Pertanyaan", "Tanggapan")
		for i := 0; i < d.n; i++ {
			fmt.Printf("%-16s %-5s %-40s %-50s %-50s\n", d.arrTanggapan[i].tag, d.arrTanggapan[i].kode, d.arrTanggapan[i].Penjawab, d.arrTanggapan[i].pertanyaan, d.arrTanggapan[i].tanggapan)
		}
	}

}

func tampilkanPasien(t tabPasien) {
	for i := 0; i < t.n; i++ {
		fmt.Println(t.arrPasien[i])
	}
}

/*
? cek akun pasien sudah ada atau belum
? kembaliin index
*/
func cekAkun(t *tabPasien, nik string) int {
	var (
		hasil int = -1
	)
	for i := 0; i < t.n; i++ {
		if nik == t.arrPasien[i].NIK {
			hasil = i
		}
	}
	return hasil
}

/*
? cek akun dokter yang sudah ada atau belum
?Kembaliin index
*/

func cekAkunDokter(t *tabDokter, nik string) int {
	var (
		hasil int = -1
	)
	for i := 0; i < t.n; i++ {
		if nik == t.arrDokter[i].NIK {
			hasil = i
		}
	}

	return hasil
}

/*
? login akun pasien menggunakan nik untuk tanya
? dokter untuk menanggapi
*/
func loginTanya(t *tabPasien, b *tabPertanyaan, c *tabTanggapan, d *tabDokter) {
	var (
		nik   string
		cek   int
		tanya string
		role  string
	)
	CallClear()
	header()
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("LOGIN UNTUK BERTANYA")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	role = tanyaRole()
	if role == "1" {
		// *Dokter
		fmt.Print("Masukkan NIK anda : ")
		fmt.Scan(&nik)
		cek = cekAkunDokter(&*d, nik)
		if cek == -1 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("ANDA BELUM TERDAFTAR DI APLIKASI KAMI ")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Print("Mau membuat akun ? (Y/N)")
			fmt.Scan(&tanya)
			if tanya == "Y" || tanya == "y" {
				role = tanyaRole()
				if role == "1" {
					isiDokter(d)
				} else if role == "2" {
					isiPasien(t)
				}
			}
		} else if cek >= 0 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("LOGIN BERHASIL")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 2 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("SELAMAT DATANG DOKTER", d.arrDokter[cek].nama)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay = 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			bertanya(b)
		}
	} else if role == "2" {
		// *Pasien
		fmt.Print("Masukkan NIK anda : ")
		fmt.Scan(&nik)
		cek = cekAkun(&*t, nik)
		if cek == -1 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("ANDA BELUM TERDAFTAR DI APLIKASI KAMI ")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Print("Mau membuat akun ? (Y/N)")
			fmt.Scan(&tanya)
			if tanya == "Y" || tanya == "y" {
				role = tanyaRole()
				if role == "1" {
					isiDokter(d)
				} else if role == "2" {
					isiPasien(t)
				}
			}
		} else if cek >= 0 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("LOGIN BERHASIL")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 2 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("SELAMAT DATANG ", t.arrPasien[cek].nama)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay = 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			bertanya(b)
		}
	}

}

/*
? login akun pasien dan dokter menggunakan nik untuk menanggapi
*/
func loginMenanggapi(t *tabPasien, b *tabPertanyaan, c *tabTanggapan, d *tabDokter) {
	var (
		nik   string
		cek   int
		tanya string
		role  string
	)
	CallClear()
	header()
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("LOGIN UNTUK MENANGGAPI")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

	// *TANYA ROLENYA DULU
	role = tanyaRole()

	if role == "1" {
		// *LOGIN AKUN DOKTER
		fmt.Print("Masukkan NIK anda : ")
		fmt.Scan(&nik)
		cek = cekAkunDokter(&*d, nik)
		if cek == -1 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("ANDA BELUM TERDAFTAR DI APLIKASI KAMI ")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)

			// *REKOMENDASI BUAT AKUN
			fmt.Print("Mau membuat akun ? (Y/N)")
			fmt.Scan(&tanya)
			if tanya == "Y" || tanya == "y" {
				role = tanyaRole()
				if role == "1" {
					isiDokter(d)
				} else if role == "2" {
					isiPasien(t)
				}
			}
		} else if cek >= 0 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("LOGIN BERHASIL")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 2 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("SELAMAT DATANG DOKTER ", d.arrDokter[cek].nama)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay = 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			menanggapi(t, b, c, d)
		}
	} else {
		// * LOGIN AKUN PASIEN
		fmt.Print("Masukkan NIK anda : ")
		fmt.Scan(&nik)
		cek = cekAkun(&*t, nik)
		if cek == -1 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("ANDA BELUM TERDAFTAR DI APLIKASI KAMI ")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Print("Mau membuat akun ? (Y/N)")
			fmt.Scan(&tanya)
			if tanya == "Y" || tanya == "y" {
				role = tanyaRole()
				if role == "1" {
					isiDokter(d)
				} else if role == "2" {
					isiPasien(t)
				}
			}
		} else if cek >= 0 {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("LOGIN BERHASIL")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay := 2 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("SELAMAT DATANG ", t.arrPasien[cek].nama)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			delay = 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
			menanggapi2(t, b, c, d)
		}
	}
}

/*
? mengembalikan index pertanyaan
*/
func cekKode(t tabPertanyaan, kode string) int {
	var (
		hasil int = -1
	)

	for i := 0; i < t.n; i++ {
		if kode == t.arrTanya[i].kode {
			hasil = i
		}
	}

	return hasil
}

/*
? Cek tag ada gk di tabPertanyaan
! parameter tag string (1-6)
*/
func cekTag(t tabPertanyaan, Tag string) int {
	var hasil int
	for i := 0; i < t.n; i++ {
		if Tag == "1" {
			if t.arrTanya[i].tag == "Penyakit" {
				hasil++
			}
		} else if Tag == "2" {
			if t.arrTanya[i].tag == "Gejala" {
				hasil++
			}
		} else if Tag == "3" {
			if t.arrTanya[i].tag == "Obat" {
				hasil++
			}
		} else if Tag == "4" {
			if t.arrTanya[i].tag == "Tindakan Medis" {
				hasil++
			}
		} else if Tag == "5" {
			if t.arrTanya[i].tag == "Diet" {
				hasil++
			}
		} else if Tag == "6" {
			if t.arrTanya[i].tag == "Kegiatan Fisik" {
				hasil++
			}
		}
	}
	return hasil
}

/*
? filter tag pertanyaan
? target sudah valid
*/
func tampilkanTagYangTerfilter(t tabPertanyaan, Tag string) {
	var jmlTag int
	var tagnya string
	if Tag == "1" {
		tagnya = "Penyakit"
	} else if Tag == "2" {
		tagnya = "Gejala"
	} else if Tag == "3" {
		tagnya = "Obat"
	} else if Tag == "4" {
		tagnya = "Tindakan Medis"
	} else if Tag == "5" {
		tagnya = "Diet"
	} else if Tag == "6" {
		tagnya = "Kegiatan Fisik"
	}

	jmlTag = cekTag(t, Tag)
	if jmlTag == 0 {
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("Belum ada pertanyaan dengan tag ", tagnya)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	} else {
		fmt.Printf("%-16s %-5s %-9s  %-50s\n", "TAG", "KODE", "TANGGAL", "Pertanyaan")
		for i := 0; i < t.n; i++ {
			if Tag == "1" {
				if t.arrTanya[i].tag == "Penyakit" {
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
				}
			} else if Tag == "2" {
				if t.arrTanya[i].tag == "Gejala" {
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
				}
			} else if Tag == "3" {
				if t.arrTanya[i].tag == "Obat" {
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
				}
			} else if Tag == "4" {
				if t.arrTanya[i].tag == "Tindakan Medis" {
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
				}
			} else if Tag == "5" {
				if t.arrTanya[i].tag == "Diet" {
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
				}
			} else if Tag == "6" {
				if t.arrTanya[i].tag == "Kegiatan Fisik" {
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", t.arrTanya[i].tag, t.arrTanya[i].kode, t.arrTanya[i].tgl, t.arrTanya[i].bln, t.arrTanya[i].thn, t.arrTanya[i].tanya)
				}
			}
		}
	}
}

/*
? fungsi menanggapi khusus dokter
*/
func menanggapi(a *tabPasien, b *tabPertanyaan, c *tabTanggapan, d *tabDokter) {
	var (
		isi                     tanggapan
		tanya                   string
		yakin                   bool = false
		stop                    bool = false
		tampung                 string
		tanggapanAsli           string = ""
		kodePertanyaan          string
		cek, cek2               int
		kembali                 string
		tanyaIdentitas          string
		NIKPenjawab             string
		indexDataPenjawab       int
		ada                     bool = false
		pilihanMenu             string
		pilihanSorting          string
		pilihanCari             string
		menanggapiKode          string
		pilihanFilter           string
		pilihanFilterMenanggapi string
		jmlTag                  int
	)

	if b.n == 0 {
		// *KALO BELUM ADA PERTANYAAN
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		tampilkanPertanyaan(*b)
		tampilkanTanggapan(&*d, &*a, &*b, &*c)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("kembali (Input apa ae )")
		fmt.Scan(&kembali)
	} else {
		// *KALO UDH ADA PERTANYAAN
		pilihanMenu = ""
		for pilihanMenu != "5" {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			tampilkanPertanyaan(*b)
			tampilkanTanggapan(&*d, &*a, &*b, &*c)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("1. Sorting Pertanyaan ")
			fmt.Println("2. Menanggapi")
			fmt.Println("3. Cari Pertanyaan")
			fmt.Println("4. Filter Pertanyaan")
			fmt.Println("5. Kembali")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihanMenu)

			for pilihanMenu != "1" && pilihanMenu != "2" && pilihanMenu != "3" && pilihanMenu != "4" && pilihanMenu != "5" {
				// *KALO PILIHAN USER ASAL ULANG SAMPE PILIHAN 1-5
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("PILIHAN TIDAK VALID")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				delay := 3 // detik
				time.Sleep(time.Duration(delay) * time.Second)
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				tampilkanTanggapan(&*d, &*a, &*b, &*c)
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("1. Sorting Pertanyaan ")
				fmt.Println("2. Menanggapi")
				fmt.Println("3. Cari Pertanyaan")
				fmt.Println("4. Filter Pertanyaan")
				fmt.Println("5. Kembali")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Pilihan anda : ")
				fmt.Scan(&pilihanMenu)
			}

			// * KALO PILIHAN 1 MASUK KE SORTING PERTANYAAN
			if pilihanMenu == "1" {
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("1. Terbaru - Terlama")
				fmt.Println("2. Terlama - Terbaru")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Pilihan anda : ")
				fmt.Scan(&pilihanSorting)

				for pilihanSorting != "1" && pilihanSorting != "2" {
					// * KALO PILIHAN USER ASAL ULANG TERUS SAMPE 1-2
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("PILIHAN TIDAK VALID")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					delay := 3 // detik
					time.Sleep(time.Duration(delay) * time.Second)
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					tampilkanPertanyaan(*b)
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("1. Terbaru - Terlama ")
					fmt.Println("2. Terlama - Terbaru ")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Pilihan anda : ")
					fmt.Scan(&pilihanSorting)
				}
				if pilihanSorting == "1" {
					// *KALO USER PILIH 1 LAKUIN ASCENDING
					sortingAscendingPertanyaan(&*b)
				} else {
					// *KALO USER PILIIH 2 LAKUIN DESCENDING
					sortingDescendingPertanyaan(&*b)
				}
			} else if pilihanMenu == "2" {
				// * USER MENANGGAPI PERTANYAAN
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan kode pertanyaan : ")
				fmt.Scan(&kodePertanyaan)
				cek = cekKode(*b, kodePertanyaan)
				for cek == -1 {
					//* KALO USER INPUT KODE YANG GK ADA
					// * ULANG TERUS SAMPE INPUT KODE YANG ADA
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("KODE TIDAK TERSEDIA")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					delay := 3 // detik
					time.Sleep(time.Duration(delay) * time.Second)
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					tampilkanPertanyaan(*b)
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Masukkan kode pertanyaan : ")
					fmt.Scan(&kodePertanyaan)
					cek = cekKode(*b, kodePertanyaan)
				}

				yakin = false
				stop = false
				for !yakin {
					yakin = false
					stop = false
					tanggapanAsli = ""
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("PERTANYAAN : ", b.arrTanya[cek].tanya)
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("NOTE : Tanggapan di akhiri dengan \" .\"")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("Silakan menanggapi : ")
					for !stop {
						fmt.Scan(&tampung)
						tanggapanAsli = tanggapanAsli + tampung + " "
						if tampung == "." {
							stop = true
							isi.tanggapan = tanggapanAsli
							isi.pertanyaan = b.arrTanya[cek].tanya
							isi.kode = b.arrTanya[cek].kode
							isi.tag = b.arrTanya[cek].tag
						}
					}
					// *TANYA MAU CANTUMIN IDENTITAS ?
					fmt.Print("Mau mencantumkan nama anda ? (Y/N)")
					fmt.Scan(&tanyaIdentitas)
					if tanyaIdentitas == "Y" || tanyaIdentitas == "y" {
						// *TANYA NIK
						fmt.Print("Masukkan NIK anda :")
						fmt.Scan(&NIKPenjawab)
						for !ada {
							indexDataPenjawab = cekAkun(a, NIKPenjawab)
							if indexDataPenjawab == -1 {
								indexDataPenjawab = cekAkunDokter(d, NIKPenjawab)
								if indexDataPenjawab == -1 {
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Println("Anda Salah memasukkan NIK anda")
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									delay := 3 // detik
									time.Sleep(time.Duration(delay) * time.Second)
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Print("Masukkan NIK anda lagi :")
									fmt.Scan(&NIKPenjawab)
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								} else {
									isi.Penjawab = d.arrDokter[indexDataPenjawab].nama
									ada = true
								}
							} else {
								isi.Penjawab = a.arrPasien[indexDataPenjawab].nama
								ada = true
							}
						}
					} else {
						isi.Penjawab = "Nama Disamarkan"
					}
					fmt.Print("Sudah yakin ? (Y/N)")
					fmt.Scan(&tanya)
					if tanya == "Y" || tanya == "y" {
						yakin = true
						c.arrTanggapan[c.n] = isi
						c.n++
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("TANGGAPAN ANDA SUDAH TERKIRIM")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						delay := 3 // detik
						time.Sleep(time.Duration(delay) * time.Second)
					}
				}
			} else if pilihanMenu == "3" {
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan Kode Pertanyaan : ")
				fmt.Scan(&pilihanCari)
				cek2 = binarySearch(*b, pilihanCari)
				for cek2 == -1 {
					//* KALO USER INPUT KODE YANG GK ADA
					// * ULANG TERUS SAMPE INPUT KODE YANG ADA
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("KODE TIDAK TERSEDIA")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					delay := 3 // detik
					time.Sleep(time.Duration(delay) * time.Second)
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					tampilkanPertanyaan(*b)
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Masukkan Kode Pertanyaan : ")
					fmt.Scan(&pilihanCari)
					cek2 = binarySearch(*b, pilihanCari)
				}
				menanggapiKode = ""
				for menanggapiKode != "2" {
					CallClear()
					header()
					sortingKodePertanyaan(&*b)
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Printf("%-16s %-5s %-9s  %-50s\n", "TAG", "KODE", "TANGGAL", "Pertanyaan")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", b.arrTanya[cek2].tag, b.arrTanya[cek2].kode, b.arrTanya[cek2].tgl, b.arrTanya[cek2].bln, b.arrTanya[cek2].thn, b.arrTanya[cek2].tanya)
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("1. Menanggapi")
					fmt.Println("2. Kembali ")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Pilihan anda : ")
					fmt.Scan(&menanggapiKode)
					for menanggapiKode != "2" && menanggapiKode != "1" {
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("PILIHAN TIDAK VALID")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						delay := 3 // detik
						time.Sleep(time.Duration(delay) * time.Second)
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Printf(" %-16s %-5s %-9s  %-50s\n", "TAG", "KODE", "TANGGAL", "Pertanyaan")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", b.arrTanya[cek2].tag, b.arrTanya[cek2].kode, b.arrTanya[cek2].tgl, b.arrTanya[cek2].bln, b.arrTanya[cek2].thn, b.arrTanya[cek2].tanya)
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("1. Menanggapi")
						fmt.Println("2. Kembali ")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Print("Pilihan anda : ")
						fmt.Scan(&menanggapiKode)
					}

					if menanggapiKode == "1" {
						yakin = false
						stop = false
						for !yakin {
							yakin = false
							stop = false
							tanggapanAsli = ""
							CallClear()
							header()
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("PERTANYAAN : ", b.arrTanya[cek2].tanya)
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("NOTE : Tanggapan di akhiri dengan \" .\"")
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("Silakan menanggapi : ")
							for !stop {
								fmt.Scan(&tampung)
								tanggapanAsli = tanggapanAsli + tampung + " "
								if tampung == "." {
									stop = true
									isi.tanggapan = tanggapanAsli
									isi.pertanyaan = b.arrTanya[cek2].tanya
									isi.kode = b.arrTanya[cek2].kode
									isi.tag = b.arrTanya[cek2].tag
								}
							}

							// *TANYA MAU CANTUMIN IDENTITAS ?
							fmt.Print("Mau mencantumkan nama anda ? (Y/N)")
							fmt.Scan(&tanyaIdentitas)
							if tanyaIdentitas == "Y" || tanyaIdentitas == "y" {
								// *TANYA NIK
								fmt.Print("Masukkan NIK anda :")
								fmt.Scan(&NIKPenjawab)
								for !ada {
									indexDataPenjawab = cekAkun(a, NIKPenjawab)
									if indexDataPenjawab == -1 {
										indexDataPenjawab = cekAkunDokter(d, NIKPenjawab)
										if indexDataPenjawab == -1 {
											CallClear()
											header()
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
											fmt.Println("Anda Salah memasukkan NIK anda")
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
											delay := 3 // detik
											time.Sleep(time.Duration(delay) * time.Second)
											CallClear()
											header()
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
											fmt.Print("Masukkan NIK anda lagi :")
											fmt.Scan(&NIKPenjawab)
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
										} else {
											isi.Penjawab = d.arrDokter[indexDataPenjawab].nama
											ada = true
										}
									} else {
										isi.Penjawab = a.arrPasien[indexDataPenjawab].nama
										ada = true
									}
								}
							} else {
								isi.Penjawab = "Nama Disamarkan"
							}
							fmt.Print("Sudah yakin ? (Y/N)")
							fmt.Scan(&tanya)
							if tanya == "Y" || tanya == "y" {
								yakin = true
								c.arrTanggapan[c.n] = isi
								c.n++
								CallClear()
								header()
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								fmt.Println("TANGGAPAN ANDA SUDAH TERKIRIM")
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								delay := 3 // detik
								time.Sleep(time.Duration(delay) * time.Second)
								menanggapiKode = "2"
							}
						}
					}
				}
			} else if pilihanMenu == "4" {
				pilihanFilter = ""
				for pilihanFilter != "7" {
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					tampilkanPertanyaan(*b)
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					tampilkanTagYangAda()
					fmt.Println("7. Kembali")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Pilihan anda : ")
					fmt.Scan(&pilihanFilter)
					for pilihanFilter != "1" && pilihanFilter != "2" && pilihanFilter != "3" && pilihanFilter != "4" && pilihanFilter != "5" && pilihanFilter != "6" && pilihanFilter != "7" {
						// * kalo pilihannya bkn (1-7)
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("PILIHAN TIDAK VALID")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						delay := 3 // detik
						time.Sleep(time.Duration(delay) * time.Second)
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						tampilkanPertanyaan(*b)
						fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
						tampilkanTagYangAda()
						fmt.Println("7. Kembali")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Print("Pilihan anda : ")
						fmt.Scan(&pilihanFilter)
					}
					if pilihanFilter != "7" {
						jmlTag = cekTag(*b, pilihanFilter)
						if jmlTag == 0 {
							CallClear()
							header()
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							tampilkanTagYangTerfilter(*b, pilihanFilter)
							fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
							fmt.Print("Kembali (input apa ae ): ")
							fmt.Scan(&pilihanFilterMenanggapi)
						} else {
							CallClear()
							header()
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							tampilkanTagYangTerfilter(*b, pilihanFilter)
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("1. Memanggapi")
							fmt.Println("2. Kembali")
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Print("Pilihan anda : ")
							fmt.Scan(&pilihanFilterMenanggapi)
							for pilihanFilterMenanggapi != "1" && pilihanFilterMenanggapi != "2" {
								CallClear()
								header()
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								fmt.Println("PILIHAN TIDAK VALID")
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								delay := 3 // detik
								time.Sleep(time.Duration(delay) * time.Second)
								CallClear()
								header()
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								tampilkanTagYangTerfilter(*b, pilihanFilter)
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
								fmt.Println("1. Memanggapi")
								fmt.Println("2. Kembali")
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								fmt.Print("Pilihan anda : ")
								fmt.Scan(&pilihanFilterMenanggapi)
							}
							if pilihanFilterMenanggapi == "1" {
								// * USER MENANGGAPI PERTANYAAN
								CallClear()
								header()
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								tampilkanTagYangTerfilter(*b, pilihanFilter)
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								fmt.Print("Masukkan kode pertanyaan : ")
								fmt.Scan(&kodePertanyaan)
								cek = cekKode(*b, kodePertanyaan)
								for cek == -1 {
									//* KALO USER INPUT KODE YANG GK ADA
									// * ULANG TERUS SAMPE INPUT KODE YANG ADA
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Println("KODE TIDAK TERSEDIA")
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									delay := 3 // detik
									time.Sleep(time.Duration(delay) * time.Second)
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									tampilkanTagYangTerfilter(*b, pilihanFilter)
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Print("Masukkan kode pertanyaan : ")
									fmt.Scan(&kodePertanyaan)
									cek = cekKode(*b, kodePertanyaan)
								}

								yakin = false
								stop = false
								for !yakin {
									yakin = false
									stop = false
									tanggapanAsli = ""
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Println("PERTANYAAN : ", b.arrTanya[cek].tanya)
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Println("NOTE : Tanggapan di akhiri dengan \" .\"")
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Println("Silakan menanggapi : ")
									for !stop {
										fmt.Scan(&tampung)
										tanggapanAsli = tanggapanAsli + tampung + " "
										if tampung == "." {
											stop = true
											isi.tanggapan = tanggapanAsli
											isi.pertanyaan = b.arrTanya[cek].tanya
											isi.kode = b.arrTanya[cek].kode
											isi.tag = b.arrTanya[cek].tag
										}
									}
									// *TANYA MAU CANTUMIN IDENTITAS ?
									fmt.Print("Mau mencantumkan nama anda ? (Y/N)")
									fmt.Scan(&tanyaIdentitas)
									if tanyaIdentitas == "Y" || tanyaIdentitas == "y" {
										// *TANYA NIK
										fmt.Print("Masukkan NIK anda :")
										fmt.Scan(&NIKPenjawab)
										for !ada {
											indexDataPenjawab = cekAkun(a, NIKPenjawab)
											if indexDataPenjawab == -1 {
												indexDataPenjawab = cekAkunDokter(d, NIKPenjawab)
												if indexDataPenjawab == -1 {
													CallClear()
													header()
													fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
													fmt.Println("Anda Salah memasukkan NIK anda")
													fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
													delay := 3 // detik
													time.Sleep(time.Duration(delay) * time.Second)
													CallClear()
													header()
													fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
													fmt.Print("Masukkan NIK anda lagi :")
													fmt.Scan(&NIKPenjawab)
													fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
												} else {
													isi.Penjawab = d.arrDokter[indexDataPenjawab].nama
													ada = true
												}
											} else {
												isi.Penjawab = a.arrPasien[indexDataPenjawab].nama
												ada = true
											}
										}
									} else {
										isi.Penjawab = "Nama Disamarkan"
									}
									fmt.Print("Sudah yakin ? (Y/N)")
									fmt.Scan(&tanya)
									if tanya == "Y" || tanya == "y" {
										yakin = true
										c.arrTanggapan[c.n] = isi
										c.n++
										CallClear()
										header()
										fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
										fmt.Println("TANGGAPAN ANDA SUDAH TERKIRIM")
										fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
										delay := 3 // detik
										time.Sleep(time.Duration(delay) * time.Second)
										pilihanFilter = "7"
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

/*
? Fungsi menanggapi khusus pasien
*/
func menanggapi2(a *tabPasien, b *tabPertanyaan, c *tabTanggapan, d *tabDokter) {
	var (
		kembali           string
		pilihanMenu       string
		cek               int
		kodePertanyaan    string
		yakin, stop       bool
		tanggapanAsli     string
		tampung           string
		isi               tanggapan
		tanyaIdentitas    string
		NIKPenjawab       string
		indexDataPenjawab int
		ada               bool
		tanya             string
		pilihanCari       string
		cek2              int
		menanggapiKode    string
	)

	if b.n == 0 {
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		tampilkanPertanyaan(*b)
		tampilkanTanggapan(&*d, &*a, &*b, &*c)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("kembali (Input apa ae )")
		fmt.Scan(&kembali)
	} else {
		pilihanMenu = ""
		for pilihanMenu != "3" {
			CallClear()
			header()
			sortingDescendingPertanyaan(&*b)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			tampilkanPertanyaan(*b)
			tampilkanTanggapan(&*d, &*a, &*b, &*c)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("1. Menanggapi")
			fmt.Println("2. Cari Pertanyaan")
			fmt.Println("3. Kembali")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihanMenu)

			for pilihanMenu != "1" && pilihanMenu != "2" && pilihanMenu != "3" {
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("PILIHAN TIDAK VALID")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				delay := 3 // detik
				time.Sleep(time.Duration(delay) * time.Second)
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				tampilkanTanggapan(&*d, &*a, &*b, &*c)
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println("1. Menanggapi")
				fmt.Println("2. Cari Pertanyaan")
				fmt.Println("3. Kembali")
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Pilihan anda : ")
				fmt.Scan(&pilihanMenu)
			}

			if pilihanMenu == "1" {
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan kode pertanyaan : ")
				fmt.Scan(&kodePertanyaan)
				cek = cekKode(*b, kodePertanyaan)
				for cek == -1 {
					//* KALO USER INPUT KODE YANG GK ADA
					// * ULANG TERUS SAMPE INPUT KODE YANG ADA
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("KODE TIDAK TERSEDIA")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					delay := 3 // detik
					time.Sleep(time.Duration(delay) * time.Second)
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					tampilkanPertanyaan(*b)
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Masukkan kode pertanyaan : ")
					fmt.Scan(&kodePertanyaan)
					cek = cekKode(*b, kodePertanyaan)
				}
				yakin = false
				stop = false
				for !yakin {
					yakin = false
					stop = false
					tanggapanAsli = ""
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("PERTANYAAN : ", b.arrTanya[cek].tanya)
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("NOTE : Tanggapan di akhiri dengan \" .\"")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("Silakan menanggapi : ")
					for !stop {
						fmt.Scan(&tampung)
						tanggapanAsli = tanggapanAsli + tampung + " "
						if tampung == "." {
							stop = true
							isi.tanggapan = tanggapanAsli
							isi.pertanyaan = b.arrTanya[cek].tanya
							isi.kode = b.arrTanya[cek].kode
							isi.tag = b.arrTanya[cek].tag
						}
					}
					// *TANYA MAU CANTUMIN IDENTITAS ?
					fmt.Print("Mau mencantumkan nama anda ? (Y/N)")
					fmt.Scan(&tanyaIdentitas)
					if tanyaIdentitas == "Y" || tanyaIdentitas == "y" {
						// *TANYA NIK
						fmt.Print("Masukkan NIK anda :")
						fmt.Scan(&NIKPenjawab)
						for !ada {
							indexDataPenjawab = cekAkun(a, NIKPenjawab)
							if indexDataPenjawab == -1 {
								indexDataPenjawab = cekAkunDokter(d, NIKPenjawab)
								if indexDataPenjawab == -1 {
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Println("Anda Salah memasukkan NIK anda")
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									delay := 3 // detik
									time.Sleep(time.Duration(delay) * time.Second)
									CallClear()
									header()
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
									fmt.Print("Masukkan NIK anda lagi :")
									fmt.Scan(&NIKPenjawab)
									fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								} else {
									isi.Penjawab = d.arrDokter[indexDataPenjawab].nama
									ada = true
								}
							} else {
								isi.Penjawab = a.arrPasien[indexDataPenjawab].nama
								ada = true
							}
						}
					} else {
						isi.Penjawab = "Nama Disamarkan"
					}
					fmt.Print("Sudah yakin ? (Y/N)")
					fmt.Scan(&tanya)
					if tanya == "Y" || tanya == "y" {
						yakin = true
						c.arrTanggapan[c.n] = isi
						c.n++
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("TANGGAPAN ANDA SUDAH TERKIRIM")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						delay := 3 // detik
						time.Sleep(time.Duration(delay) * time.Second)
					}
				}
			} else if pilihanMenu == "2" {
				CallClear()
				header()
				fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
				tampilkanPertanyaan(*b)
				fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan Kode Pertanyaan : ")
				fmt.Scan(&pilihanCari)
				cek2 = binarySearch(*b, pilihanCari)
				for cek2 == -1 {
					//* KALO USER INPUT KODE YANG GK ADA
					// * ULANG TERUS SAMPE INPUT KODE YANG ADA
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("KODE TIDAK TERSEDIA")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					delay := 3 // detik
					time.Sleep(time.Duration(delay) * time.Second)
					CallClear()
					header()
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					tampilkanPertanyaan(*b)
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Masukkan Kode Pertanyaan : ")
					fmt.Scan(&pilihanCari)
					cek2 = binarySearch(*b, pilihanCari)
				}
				menanggapiKode = ""
				for menanggapiKode != "2" {
					CallClear()
					header()
					sortingKodePertanyaan(&*b)
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Printf("%-16s %-5s %-9s  %-50s\n", "TAG", "KODE", "TANGGAL", "Pertanyaan")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", b.arrTanya[cek2].tag, b.arrTanya[cek2].kode, b.arrTanya[cek2].tgl, b.arrTanya[cek2].bln, b.arrTanya[cek2].thn, b.arrTanya[cek2].tanya)
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
					fmt.Println("1. Menanggapi")
					fmt.Println("2. Kembali ")
					fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
					fmt.Print("Pilihan anda : ")
					fmt.Scan(&menanggapiKode)
					for menanggapiKode != "2" && menanggapiKode != "1" {
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("PILIHAN TIDAK VALID")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						delay := 3 // detik
						time.Sleep(time.Duration(delay) * time.Second)
						CallClear()
						header()
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Printf(" %-16s %-5s %-9s  %-50s\n", "TAG", "KODE", "TANGGAL", "Pertanyaan")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Printf("%-16s %-5s %-2d %-2d %-3d %-50s\n", b.arrTanya[cek2].tag, b.arrTanya[cek2].kode, b.arrTanya[cek2].tgl, b.arrTanya[cek2].bln, b.arrTanya[cek2].thn, b.arrTanya[cek2].tanya)
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
						fmt.Println("1. Menanggapi")
						fmt.Println("2. Kembali ")
						fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
						fmt.Print("Pilihan anda : ")
						fmt.Scan(&menanggapiKode)
					}

					if menanggapiKode == "1" {
						yakin = false
						stop = false
						for !yakin {
							yakin = false
							stop = false
							tanggapanAsli = ""
							CallClear()
							header()
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("PERTANYAAN : ", b.arrTanya[cek2].tanya)
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("NOTE : Tanggapan di akhiri dengan \" .\"")
							fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("Silakan menanggapi : ")
							for !stop {
								fmt.Scan(&tampung)
								tanggapanAsli = tanggapanAsli + tampung + " "
								if tampung == "." {
									stop = true
									isi.tanggapan = tanggapanAsli
									isi.pertanyaan = b.arrTanya[cek2].tanya
									isi.kode = b.arrTanya[cek2].kode
									isi.tag = b.arrTanya[cek2].tag
								}
							}

							// *TANYA MAU CANTUMIN IDENTITAS ?
							fmt.Print("Mau mencantumkan nama anda ? (Y/N)")
							fmt.Scan(&tanyaIdentitas)
							if tanyaIdentitas == "Y" || tanyaIdentitas == "y" {
								// *TANYA NIK
								fmt.Print("Masukkan NIK anda :")
								fmt.Scan(&NIKPenjawab)
								for !ada {
									indexDataPenjawab = cekAkun(a, NIKPenjawab)
									if indexDataPenjawab == -1 {
										indexDataPenjawab = cekAkunDokter(d, NIKPenjawab)
										if indexDataPenjawab == -1 {
											CallClear()
											header()
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
											fmt.Println("Anda Salah memasukkan NIK anda")
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
											delay := 3 // detik
											time.Sleep(time.Duration(delay) * time.Second)
											CallClear()
											header()
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
											fmt.Print("Masukkan NIK anda lagi :")
											fmt.Scan(&NIKPenjawab)
											fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
										} else {
											isi.Penjawab = d.arrDokter[indexDataPenjawab].nama
											ada = true
										}
									} else {
										isi.Penjawab = a.arrPasien[indexDataPenjawab].nama
										ada = true
									}
								}
							} else {
								isi.Penjawab = "Nama Disamarkan"
							}
							fmt.Print("Sudah yakin ? (Y/N)")
							fmt.Scan(&tanya)
							if tanya == "Y" || tanya == "y" {
								yakin = true
								c.arrTanggapan[c.n] = isi
								c.n++
								CallClear()
								header()
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								fmt.Println("TANGGAPAN ANDA SUDAH TERKIRIM")
								fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
								delay := 3 // detik
								time.Sleep(time.Duration(delay) * time.Second)
								menanggapiKode = "2"
							}
						}
					}
				}
			}
		}
	}
}

func tampilanUtama(A *tabPasien, B *tabPertanyaan, C *tabTanggapan, D *tabDokter) {
	var (
		pilihan       string = "1"
		tanya, tanya2 string
		role          string
		kembali       string
	)

	for pilihan >= "1" && pilihan <= "4" {
		CallClear()
		header()
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("MENU UTAMA")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Forum")
		fmt.Println("3. Tanya")
		fmt.Println("4. Diskusi")
		fmt.Println("5. Exit ")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == "1" {
			role = tanyaRole()
			if role == "1" {
				isiDokter(D)
			} else if role == "2" {
				isiPasien(A)
			}
		} else if pilihan == "2" {
			CallClear()
			header()
			sortingDescendingPertanyaan(&*B)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			tampilkanPertanyaan(*B)
			tampilkanTanggapan(D, A, B, C)
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------")
			fmt.Print("Kembali (Input apa ae) : ")
			fmt.Scan(&kembali)
		} else if pilihan == "3" {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("Anda harus membuat akun terlebih dahulu !!!")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Print("Sudah punya akun ? (Y/N)")
			fmt.Scan(&tanya)
			if tanya == "Y" || tanya == "y" {
				loginTanya(A, B, C, D)
			} else {
				fmt.Print("Mau buat akun ? (Y/N)")
				fmt.Scan(&tanya2)
				if tanya2 == "Y" || tanya2 == "y" {
					role = tanyaRole()
					if role == "1" {
						isiDokter(D)
					} else if role == "2" {
						isiPasien(A)
					}
				}
			}
		} else if pilihan == "4" {
			CallClear()
			header()
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Println("Anda harus membuat akun terlebih dahulu !!!")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
			fmt.Print("Sudah punya akun ? (Y/N)")
			fmt.Scan(&tanya)
			if tanya == "Y" || tanya == "y" {
				loginMenanggapi(A, B, C, D)
			} else {
				fmt.Print("Mau buat akun ? (Y/N)")
				fmt.Scan(&tanya2)
				if tanya2 == "Y" || tanya2 == "y" {
					role = tanyaRole()
					if role == "1" {
						isiDokter(D)
					} else if role == "2" {
						isiPasien(A)
					}
				}
			}
		} else {
			ending()
			delay := 3 // detik
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}
}

func main() {
	var (
		a tabPasien
		b tabPertanyaan
		c tabTanggapan
		d tabDokter
	)
	tampilanUtama(&a, &b, &c, &d)

}
