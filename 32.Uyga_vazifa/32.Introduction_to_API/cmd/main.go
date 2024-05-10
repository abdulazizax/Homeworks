package main

import (
	"context"
	"fmt"
	funk "h32/32.Introduction_to_API/functions"

	"golang.org/x/sync/errgroup"
)

func main() {
	db := funk.DBConnect()
	man := funk.Manager{
		DB: db,
	}

	ctx := context.Background()

	group, ctx := errgroup.WithContext(ctx)

	man.DropLarge_datasetTable(ctx)
	man.CreateLarge_datasetTable(ctx)
	for i := 0; i < 10; i++ {
		man.InsertIntoLarge_dataset(ctx)
	}

	group.Go(func() error {
		man.InsertIntoLarge_dataset(ctx)
		return nil
	})

	group.Go(func() error {
		man.UpdateLarge_datasetTable(ctx)
		return nil
	})

	group.Go(func() error {
		man.SelectFromLarge_dataset(ctx)
		return nil
	})

	err := group.Wait()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All functions completed successfully")
	}
}
