package main

import (
	"fmt"
	"github.com/cmp307/assetman/pkg/storage/sqlite"
)

func main() {
	db, err := sqlite.Connect()

	if err != nil {
		panic(err)
	}

	repo := sqlite.NewAssetRepository(db)

	ass, _ := repo.GetAll()

	fmt.Println(ass)
}
