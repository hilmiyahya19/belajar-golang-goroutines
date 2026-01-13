package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Kode dieksekusi dengan membuat channel lalu menjalankan goroutine yang sleep 2 detik sebelum mengirim data, sementara goroutine utama langsung terblokir di data := <-channel menunggu kiriman; setelah 2 detik data dikirim sehingga receive terbuka, pesan diterima dan dicetak, lalu Sleep(5 detik) di akhir hanya menahan program agar goroutine sempat mencetak "Selesai mengirim data ke channel", karena operasi receive pada channel bersifat blocking sampai ada data masuk.
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello Channel"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println("Menerima data dari channel:", data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ini adalah data dari channel"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// MENNGIRIM KE CHANNEL
// OnlyIn hanya bisa mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Data hanya bisa dikirim ke channel"
}

// MENERIMA DARI CHANNEL
// OnlyOut hanya bisa menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Menerima:", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel = Channel dengan kapasitas tertentu
// Pada contoh ini, channel dibuat dengan kapasitas 3, sehingga goroutine pengirim dapat mengirim 3 data sebelum goroutine penerima menerima data pertama. Setelah itu, goroutine penerima mulai menerima data satu per satu. Karena kapasitas channel mencukupi, pengirim tidak akan terblokir sampai kapasitas penuh.
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Data 1"
		channel <- "Data 2"
		channel <- "Data 3"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data:", data)
	}

	fmt.Println("Selesai")
}

// SELECT CHANNEL
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// DEFAULT SELECT CHANNEL
// default select digunakan untuk menangani situasi di mana tidak ada case yang siap untuk dieksekusi. Dengan menambahkan blok default, program tidak akan terblokir dan dapat melakukan tindakan lain, seperti mencetak pesan atau menunggu sebentar sebelum mencoba lagi.
// mencegah deadlock pada program ketika tidak ada data yang tersedia di channel.
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2:", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
			time.Sleep(500 * time.Millisecond)
		}

		if counter == 2 {
			break
		}
	}
}