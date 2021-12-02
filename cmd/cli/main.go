package main

import (
	"fmt"
	"github.com/cmp307/assetman/pkg/storage/sqlite"
)

func main() {
	db, err := sqlite.Connect("C:\\Users\\Jannes\\ScottishGlen\\assets.db")

	if err != nil {
		panic(err)
	}

	repo := sqlite.NewAssetRepository(db)

	ass := repo.GetById(1)

	fmt.Println(ass)
}
