package belajar_golang_goroutines

import (
	"fmt"   // Digunakan untuk menampilkan output ke console
	"sync"  // Digunakan untuk sinkronisasi goroutine (WaitGroup)
	"testing" // Digunakan untuk membuat unit test di Go
	"time"  // Digunakan untuk operasi waktu (Timer, After, AfterFunc)
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second) // Membuat timer yang akan aktif setelah 5 detik
	fmt.Println(time.Now())                 // Menampilkan waktu saat ini sebelum timer selesai

	time := <-timer.C // Menunggu sampai timer selesai dan menerima waktu dari channel C
	fmt.Println(time) // Menampilkan waktu ketika timer selesai
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second) // Mengembalikan channel yang akan mengirim waktu setelah 5 detik
	fmt.Println(time.Now())                // Menampilkan waktu saat ini sebelum menunggu channel

	time := <-channel // Menunggu hingga channel mengirim waktu setelah 5 detik
	fmt.Println(time) // Menampilkan waktu ketika durasi 5 detik telah berlalu
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{} // Membuat WaitGroup untuk menunggu goroutine selesai
	group.Add(1)              // Menandakan ada 1 goroutine yang harus ditunggu

	time.AfterFunc(5*time.Second, func() { // Menjalankan fungsi secara asynchronous setelah 5 detik
		fmt.Println(time.Now()) // Menampilkan waktu saat fungsi dijalankan
		group.Done()            // Memberi tahu WaitGroup bahwa goroutine telah selesai
	})
	fmt.Println(time.Now()) // Menampilkan waktu saat ini (langsung dieksekusi tanpa menunggu)

	group.Wait() // Menunggu hingga semua goroutine yang terdaftar selesai
}

// Kesimpulan:
// Kode ini mendemonstrasikan tiga cara penjadwalan waktu di Go yaitu time.NewTimer,
// time.After, dan time.AfterFunc. time.NewTimer dan time.After sama-sama menggunakan
// channel untuk menunggu durasi tertentu secara blocking, sedangkan time.AfterFunc
// menjalankan fungsi secara asynchronous setelah durasi tertentu sehingga memerlukan
// mekanisme sinkronisasi seperti sync.WaitGroup agar program tidak selesai lebih cepat.
