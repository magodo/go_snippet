package animal

import (
	"fmt"
	"testing"
)

func TestAnimal(t *testing.T) {
	dog := NewDog()
	cat := NewCat()
	cow := NewCow()

	fmt.Println(dog.Speak("Dog say: "))
	fmt.Println(cat.Speak("Cat say: "))
	fmt.Println(cow.Speak("Cow say: "))
}
