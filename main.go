package main

import "fmt"

func main() {
	var conferenceName = "KITA-KITA"
	const conferenceTicket = 70
	var remainingTicket = 70
	var name string
	var bookings []string

	fmt.Println("Selamat Datang Ke Konferensi", conferenceName)
	fmt.Println("Kita memiliki", conferenceTicket, "tiket dan tersisa", remainingTicket, "tiket")
	fmt.Println("Ambil Tiketmu Disini!!!")

	fmt.Println(conferenceName)

	fmt.Println("Siapa Namamu?")
	fmt.Scan(&name)
	bookings = append(bookings, name)
	fmt.Printf("Hallo selamat datang %v", name)
	fmt.Printf("Konfirmasi nama kamu adalah %v", bookings[0])
}
