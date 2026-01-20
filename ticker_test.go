package belajar_golang_goroutines

import (
	"fmt"    // Digunakan untuk menampilkan output ke console
	"testing" // Digunakan untuk membuat unit test di Go
	"time"   // Digunakan untuk fitur waktu seperti Ticker dan Tick
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // Membuat ticker yang mengirim event setiap 1 detik

	go func() {
		time.Sleep(5 * time.Second) // Menunggu selama 5 detik
		ticker.Stop()               // Menghentikan ticker agar tidak mengirim event lagi
	}()

	for time := range ticker.C { // Melakukan iterasi setiap kali ticker mengirim waktu
		fmt.Println("Ticker at", time) // Menampilkan waktu dari ticker
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second) // Membuat ticker sederhana yang mengirim waktu tiap 1 detik

	for time := range channel { // Melakukan iterasi setiap kali channel mengirim waktu
		fmt.Println("Tick at", time) // Menampilkan waktu dari tick
	}
}

// Kesimpulan:
// Kode ini menunjukkan dua cara penggunaan ticker di Go, yaitu time.NewTicker dan time.Tick.
// time.NewTicker memberikan kontrol penuh karena ticker dapat dihentikan menggunakan Stop(),
// sehingga lebih aman untuk penggunaan jangka panjang. Sementara itu, time.Tick lebih sederhana
// namun tidak bisa dihentikan, sehingga berpotensi menyebabkan kebocoran resource jika digunakan
// tanpa mekanisme penghentian yang jelas.
