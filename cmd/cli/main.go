package main

import (
	"fmt"
	"github.com/cmp307/assetman/pkg/storage"
	"github.com/cmp307/assetman/pkg/storage/sqlite"
)

func main() {
	db, err := sqlite.Connect()

	if err != nil {
		panic(err)
	}

	repo := sqlite.NewRepository(db)
	asset, _ := repo.PaginateByName("Off", storage.QueryOptions{
		Limit:  10,
		Offset: 0,
	})

	fmt.Println(asset)
}
