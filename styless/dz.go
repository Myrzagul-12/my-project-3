func main (
	"fmt"
	"string"
func CreateRequest(title, description, location string) {

if len(strings.TrimSpace(title)) == 0 
   len(strings.TrimSpace(description)) == 0 
   len(strings.TrimSpace(location)) == 0 {
	fmt.Println("Ошибка: все поля  должны быть заполнены!")
	return 
}
fmt.Println("Заявка успешно создана:")
}

func main() {
	CreateRequest("Починить кран", " ", "Кухня")
}