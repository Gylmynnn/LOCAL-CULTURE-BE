package main

import (
	"local-culture-backend/db"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	return nil
}
