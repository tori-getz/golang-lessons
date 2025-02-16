package main

import "fmt"

const VIEW_BOOKMARKS = 1
const ADD_BOOKMARK = 2
const REMOVE_BOOKMARK = 3
const EXIT_APP = 4

func main() {
	fmt.Println("__ bookmarks app __")

	bookmarks := map[string]string{}

	for {
		ShowMenu()

		choice := GetMenuChoice()

		switch choice {
		case VIEW_BOOKMARKS:
			ViewBookmarks(bookmarks)
		case ADD_BOOKMARK:
			bookmarkName, bookmarkUrl := AddBookmark()
			bookmarks[bookmarkName] = bookmarkUrl
			fmt.Println("")
			fmt.Println("Bookmark was added!")
			fmt.Println("")
		case REMOVE_BOOKMARK:
			bookmarkName := RemoveBookmark()
			delete(bookmarks, bookmarkName)
			fmt.Println("")
			fmt.Println("Bookmark was deleted!")
			fmt.Println("")
		case EXIT_APP:
			fmt.Println("Have a nice day!")
			return
		default:
			continue
		}
	}
}

func ViewBookmarks(bookmarks map[string]string) {
	fmt.Println("")
	fmt.Println("Bookmarks list:")
	fmt.Println("")

	index := 1

	for key, value := range bookmarks {
		str := fmt.Sprintf("[%v] %v - %v", index, key, value)
		fmt.Println(str)
		index++
	}
	fmt.Println("")
}

func AddBookmark() (string, string) {
	fmt.Println("")
	fmt.Println("Add bookmark")
	fmt.Println("")

	name := ""
	url := ""

	fmt.Print("Enter bookmark name: ")
	fmt.Scan(&name)

	fmt.Print("Enter bookmark URL: ")
	fmt.Scan(&url)

	return name, url
}

func RemoveBookmark() string {
	fmt.Println("")
	fmt.Println("Remove bookmark")
	fmt.Println("")

	name := ""

	fmt.Print("Enter bookmark name: ")
	fmt.Scan(&name)

	return name
}

func ShowMenu() {
	menu := [4]string{
		"View bookmarks",
		"Add bookmark",
		"Remove bookmark",
		"Exit",
	}

	fmt.Println("")
	fmt.Println("Меню:")
	for index, value := range menu {
		str := fmt.Sprintf("%v. %v", index+1, value)
		fmt.Println(str)
	}
	fmt.Println("")
}

func GetMenuChoice() int {
	var choice int
	fmt.Print("Enter menu number: ")
	fmt.Scan(&choice)
	return choice
}
