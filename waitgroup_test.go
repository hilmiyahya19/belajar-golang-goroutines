package belajar_golang_goroutines // Nama package untuk contoh pembelajaran goroutine

import (
	"fmt"  // Digunakan untuk menampilkan output ke console
	"sync" // Menyediakan WaitGroup untuk sinkronisasi goroutine
	"testing" // Digunakan agar fungsi bisa dijalankan sebagai unit test
	"time" // Digunakan untuk memberi jeda waktu (Sleep)
)

func RunAsynchronous(group *sync.WaitGroup) { // Fungsi yang akan dijalankan secara goroutine
	// defer adalah keyword di Go yang digunakan untuk menunda eksekusi sebuah fungsi sampai fungsi yang membungkusnya selesai (return).
	defer group.Done() // Menandai bahwa goroutine ini selesai saat fungsi berakhir

	group.Add(1) // Menambah counter WaitGroup (INI SEBENARNYA SALAH PENEMPATANNYA)

	fmt.Println("Hello") // Menampilkan teks ke console
	time.Sleep(1 * time.Second) // Mensimulasikan proses dengan delay 1 detik
}

func TestWaitGroup(t *testing.T) { // Fungsi test untuk mendemonstrasikan penggunaan WaitGroup
	group := &sync.WaitGroup{} // Membuat instance WaitGroup

	for i := 0; i < 100; i++ { // Loop untuk menjalankan 100 goroutine
		go RunAsynchronous(group) // Menjalankan fungsi secara asynchronous
	}

	group.Wait() // Menunggu sampai counter WaitGroup kembali ke 0
	fmt.Println("Selesai") // Dieksekusi setelah semua goroutine selesai
}

// Kesimpulan:
// Kode ini mendemonstrasikan penggunaan sync.WaitGroup untuk menunggu banyak goroutine
// agar selesai sebelum program melanjutkan eksekusi. Namun terdapat kesalahan logika,
// yaitu pemanggilan group.Add(1) yang dilakukan di dalam goroutine, sehingga berpotensi
// menyebabkan race condition atau panic. Praktik yang benar adalah memanggil Add()
// sebelum goroutine dijalankan, lalu Done() di dalam goroutine untuk menandai selesai.
