package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// Jadwalkan tugas setiap menit
	c.AddFunc("* * * * *", func() {
		fmt.Println("Tugas dijalankan setiap menit:", time.Now())
	})

	// Jadwalkan tugas dengan interval kustom
	c.AddFunc("@hourly", func() {
		fmt.Println("Tugas dijalankan setiap jam:", time.Now())
	})

	// Mulai jadwal
	c.Start()

	// Biarkan program berjalan selama beberapa saat
	time.Sleep(5 * time.Minute)

	// Stop jadwal
	c.Stop()
}
