package methods

import "fmt"

type user struct {
	name string
	age  uint8
}

// Método do struct user
func (u user) save() {
	fmt.Println("Conteúdo salvo com sucesso")
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user1 := user{
		"pedro",
		20,
	}
	fmt.Printf("%s possui %d anos\n", user1.name, user1.age)
	user1.save()
	user1.birthday()
	fmt.Printf("%s possui %d anos\n", user1.name, user1.age)

}
