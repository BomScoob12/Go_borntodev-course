interface is for polymorphism, it
is a way to define a contract for a class to implement. It is a way to define a
set of methods that a class must implement.

like Interface of Java, C# or TypeScript.

but not like that much because interface not need to be implemented.
just have to define the methods.

```go
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{
        Dog{Name: "Fido"},
        Cat{Name: "Milo"},
    }

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

