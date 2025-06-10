package main

import (
	//v1 "n8n_project_go/app/user/v1"
	"fmt"

	"github.com/google/uuid"
)

type Book struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Price  int       `json:"price"`
}

func (b Book) Validate() error {
	if b.Title == "" {
		return fmt.Errorf("タイトルが未設定です")
	}
	if b.Author == "" {
		return fmt.Errorf("著者が未設定です")
	}
	if b.Price <= 0 {
		return fmt.Errorf("価格が無効です")
	}
	return nil
}

func test() {

	books := []Book{
		{ID: uuid.New(), Title: "作ってマスターPython", Author: "機械学習・webアプリケーション", Price: 0},
		{ID: uuid.New(), Title: "絵で見てわかるブロックチェーン", Author: "", Price: 2400},
		{ID: uuid.New(), Title: "徹底攻略データベーススペシャリスト", Author: "２大特典付き", Price: 2300},
	}

	for _, book := range books {
		if err := book.Validate(); err != nil {
			fmt.Printf("❌ 検証エラー: %s\n", err)
		}
	}

	fmt.Println("本棚一覧:")
	for i, book := range books {
		fmt.Printf("%d %s %s by %s (%d円)\n", i+1, book.ID, book.Title, book.Author, book.Price)
	}

	books[1].Title = "新しいタイトル"

	fmt.Println("\n 変更後の本棚一覧:")
	for i, book := range books {
		fmt.Printf("%d %s %s by %s (%d円)\n", i+1, book.ID, book.Title, book.Author, book.Price)
	}
}
