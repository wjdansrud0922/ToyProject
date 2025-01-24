package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	list    []string
	str     string
	id      int
	content string
)

func save() {
	doc, _ := json.Marshal(list)
	err := os.WriteFile("todolist/todolist.json", doc, os.FileMode(0644))
	if err != nil {
		return
	}
}

func load() {
	b, err := os.ReadFile("todoList/todolist.json")

	if err != nil {
		return
	}
	json.Unmarshal(b, &list)
}

func main() {
	load()

	for {
		fmt.Println("[ todo List ] \n 1.생성 \n 2.수정 \n 3.삭제 \n 4.리스트")
		fmt.Print("입력 : ")
		fmt.Scanln(&str)
		switch str {
		case "1":

			fmt.Print("내용 입력 : ")
			fmt.Scan(&content)
			list = append(list, content)
			fmt.Println("추가되었습니다.")
			save()
		case "2":

			fmt.Print("수정할 todo 번호 입력 : ")
			fmt.Scanln(&id)
			if len(list) < id || id <= 0 {
				fmt.Println("존재하지 않습니다")
				continue
			}
			fmt.Print("내용 입력 : ")
			fmt.Scanln(&content)
			list[id-1] = strings.TrimSpace(content)
			save()
		case "3":
			fmt.Print("삭제할 todo 번호 입력 : ")
			fmt.Scanln(&id)
			if len(list) < id || id <= 0 {
				fmt.Println("존재하지 않습니다")
				continue
			}
			list = append(list[:id-1], list[(id):]...)
			save()
		case "4":
			fmt.Println("{ 리스트 목록 }")
			for i, _ := range list {
				fmt.Println(i+1, "번 리스트 := ", list[i])
			}
			fmt.Println()
		}
	}
}
