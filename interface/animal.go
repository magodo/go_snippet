package animal

type Animal interface {
	Speak(what string) (out string)
}

// struct implementing interface
type Dog struct {
	accent string
}

func (dog *Dog) Speak(what string) (out string) {
	return (what + dog.accent)
}

// function implement interface
type Cat func() string

func CatFunc() string {
	return " Miewww..."
}

func (cat *Cat) Speak(what string) (out string) {
	return (what + (*cat)())
}

// primary type implement interface
type Cow string

func (cow *Cow) Speak(what string) (out string) {
	return (what + string(*cow))
}

// constructors
func NewDog() *Dog {
	return &Dog{" Awuuu..."}
}

func NewCat() *Cat {
	cat := Cat(CatFunc)
	return &cat
}

func NewCow() *Cow {
	cow := Cow(" Mowww...")
	return &cow
}
