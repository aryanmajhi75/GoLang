package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type post struct {
	uid      int
	hashtags string
	content  string
	likes    int
	shares   int
	score    float64
}

func popular(posts []post, counter int) {
	var pposts [10]int
	var index int

	fmt.Println("The Popular posts are:")
	for i := 0; i < counter; i++ {
		if posts[i].likes >= 1000 {
			fmt.Printf("\n Unique ID: \n\t-> %d", posts[i].uid)
			fmt.Printf("\n Content : \n\t-> %s", posts[i].content)
			fmt.Printf("\n Hashtags : \n\t-> %s", posts[i].hashtags)
			fmt.Printf("\n Likes : \n\t-> %d", posts[i].likes)
			fmt.Printf("\n Shares : \n\t-> %d", posts[i].shares)
			fmt.Printf("\n Popularity Score :\n\t-> %0.2f", posts[i].score)
			pposts[index] = posts[i].uid
			index++
		}
	}

	fmt.Printf("%d", pposts)
}

func popularityscore(posts post) (score float64) {

	score = float64(posts.likes) + float64(posts.shares)
	score = posts.score / 112
	return score
}

func display(posts []post, length int) {
	for i := 0; i < length; i++ {
		fmt.Println("\n +-------------------------------")
		fmt.Printf("\n | Unique ID \n\t-> %d", posts[i].uid)
		fmt.Printf("\n | Content \n\t-> %s", posts[i].content)
		fmt.Printf("\n | Hashtags \n\t-> %s", posts[i].hashtags)
		fmt.Printf("\n | Likes \n\t-> %d", posts[i].likes)
		fmt.Printf("\n | Shares \n\t-> %d", posts[i].shares)
		fmt.Printf("\n |Popularity Score \n\t-> %0.2f", posts[i].score)
		fmt.Println("\n +-------------------------------")
	}
}

func search(posts []post, counter int) {
	var hashs string
	fmt.Println("Enter the search tag")
	find := readInput()
	for i := 0; i < counter; i++ {
		hashs = posts[i].hashtags
		if strings.Contains(hashs, find) {
			fmt.Printf("\n Unique ID: \n\t-> %d", posts[i].uid)
			fmt.Printf("\n Content : \n\t-> %s", posts[i].content)
			fmt.Printf("\n Hashtags : \n\t-> %s", posts[i].hashtags)
			fmt.Printf("\n Likes : \n\t-> %d", posts[i].likes)
			fmt.Printf("\n Shares : \n\t-> %d", posts[i].shares)
			fmt.Printf("\n Popularity Score :\n\t-> %0.2f", posts[i].score)
		}
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	var posts = make([]post, 10)
	var choice, counter int
	// var p[10] int
	// var err error
	for {
		fmt.Printf("\n")
		fmt.Println(" [1] Add Post ")
		fmt.Println(" [2] Display Posts ")
		fmt.Println(" [3] Check Popularity ")
		fmt.Println(" [4] Find Post by Hashtag ")
		fmt.Println(" [5] Popular Post ")
		fmt.Println(" [6] Exit ")
		fmt.Printf("\nEnter choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if counter < 10 {
				fmt.Println("Enter content :\n")
				posts[counter].content = readInput()
				fmt.Println("Enter hashtags :\n")
				posts[counter].hashtags = readInput()
				fmt.Println("Enter likes :\n")
				fmt.Scan(&posts[counter].likes)
				fmt.Println("Enter shares :\n")
				fmt.Scan(&posts[counter].shares)
				posts[counter].score = float64((posts[counter].likes + posts[counter].shares) / 112)
				counter = counter + 1
			} else {
				fmt.Println("Maximum limit reached")
			}

		case 2:
			if counter == 0 {
				fmt.Println("No posts available!")
			} else {
				display(posts, counter)

			}
		case 3:
			if counter == 0 {
				fmt.Println("No posts available!")
			} else {
				popular(posts, counter)
			}
		case 4:
			if counter == 0 {
				fmt.Println("No posts available!")
			} else {
				search(posts, counter)
			}
		case 5:
			if counter == 0 {
				fmt.Println("No posts available!")
			} else {
				popular(posts, counter)
			}
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Enter a valid choice!")
		}

	}
}
