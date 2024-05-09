package main

import (
	"context"
	funk "h31/31.Working_with_network_in_Go/functions"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	db := funk.DBConnect()
	man := funk.Manager{
		DB: db,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	man.DropLarge_datasetTable(ctx)
	man.CreateLarge_dataset(ctx)
	man.InsertIntoLarge_dataset(ctx)
	info := man.GetFromLarge_dataset(ctx)

	pp.Println(info)

}
