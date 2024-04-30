package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filename := "Homeworks/25.Uyga_vazifa/example.txt"

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Kontekst bekor qilindi.")
			return
		case <-sig:
			fmt.Println("Tugatish signali qabul qilindi.")
			cancel()
			return
		}
	}()

	data := []byte("Hello, world!")
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("File yaratishdagi xatolik:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Filega ma'lumot yozishdagi xatolik:", err)
		return
	}

	readData := make([]byte, len(data))
	_, err = file.ReadAt(readData, 0)
	if err != nil {
		fmt.Println("Filedan ma'lumot o'qishdagi xatolik:", err)
		return
	}
	fmt.Println("Filedan o'qilgan ma'lumot:", string(readData))

	time.Sleep(2 * time.Second)

	fmt.Println("Datur muvaffaqiyatli yakunlandi!")
}
