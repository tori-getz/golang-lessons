package main

import "fmt"

const VIEW_BOOKMARKS = 1
const ADD_BOOKMARK = 2
const REMOVE_BOOKMARK = 3
const EXIT_APP = 4

type bookmarksMap = map[string]string

func main() {
	fmt.Println("__ bookmarks app __")

	bookmarks := bookmarksMap{}
Menu:
	for {
		ShowMenu()

		choice := GetMenuChoice()

		switch choice {

		case VIEW_BOOKMARKS:
			ViewBookmarks(bookmarks)

		case ADD_BOOKMARK:
			AddBookmark(bookmarks)

		case REMOVE_BOOKMARK:
			RemoveBookmark(bookmarks)

		case EXIT_APP:
			fmt.Println("Have a nice day!")
			break Menu

		default:
			continue

		}
	}
}

func ViewBookmarks(bookmarks bookmarksMap) {
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

func AddBookmark(bookmarks bookmarksMap) {
	fmt.Println("")
	fmt.Println("Add bookmark")
	fmt.Println("")

	bookmarkName := ""
	bookmarkUrl := ""

	fmt.Print("Enter bookmark name: ")
	fmt.Scan(&bookmarkName)

	fmt.Print("Enter bookmark URL: ")
	fmt.Scan(&bookmarkUrl)

	bookmarks[bookmarkName] = bookmarkUrl
	fmt.Println("")
	fmt.Println("Bookmark was added!")
	fmt.Println("")
}

func RemoveBookmark(bookmarks bookmarksMap) {
	fmt.Println("")
	fmt.Println("Remove bookmark")
	fmt.Println("")

	bookmarkName := ""

	fmt.Print("Enter bookmark name: ")
	fmt.Scan(&bookmarkName)

	delete(bookmarks, bookmarkName)

	fmt.Println("")
	fmt.Println("Bookmark was deleted!")
	fmt.Println("")
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
