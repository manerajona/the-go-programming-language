package main

import (
	"fmt"
	"reflect"
)

// Person is a sample concrete type with a method.
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet(prefix string) string {
	return fmt.Sprintf("%s %s (age %d)", prefix, p.Name, p.Age)
}

func main() {
	// 1) Accessing types and values
	var p Person
	pType := reflect.TypeOf(p)           // reflect.Type for Person
	pValue := reflect.ValueOf(&p).Elem() // reflect.Value for the Person instance
	fmt.Println("Type:", pType)
	fmt.Println("Kind:", pType.Kind())
	fmt.Println("Is addressable/settable?", pValue.CanSet())

	// 2) Creating a new value of a known type at runtime
	newPerson := reflect.New(pType).Elem() // create a zero-Person value
	newPerson.FieldByName("Name").SetString("Alice")
	newPerson.FieldByName("Age").SetInt(30)
	p = newPerson.Interface().(Person)
	fmt.Println("Created via reflection:", p)

	// 3) Making a slice of the type at runtime and appending values
	sliceType := reflect.SliceOf(pType)
	slice := reflect.MakeSlice(sliceType, 0, 0) // []Person
	slice = reflect.Append(slice, newPerson)    // append the Person we made
	// append another literal
	anotherPerson := reflect.ValueOf(Person{Name: "Bob", Age: 25})
	slice = reflect.Append(slice, anotherPerson)
	fmt.Printf("Slice length: %d, contents: %#v\n", slice.Len(), slice.Interface())

	// 4) Making a map of string -> Person and setting entries
	mapType := reflect.MapOf(reflect.TypeOf(""), pType) // map[string]Person
	m := reflect.MakeMap(mapType)
	m.SetMapIndex(reflect.ValueOf("alice"), newPerson)
	m.SetMapIndex(reflect.ValueOf("bob"), anotherPerson)
	fmt.Println("Map contents:", m.Interface())

	// 5) Calling a method by name (runtime method lookup)
	// Use the first element from slice
	elem := slice.Index(0) // reflect.Value of Person
	method := elem.MethodByName("Greet")
	if method.IsValid() {
		// Call expects []reflect.Value arguments and returns []reflect.Value
		results := method.Call([]reflect.Value{reflect.ValueOf("Hello,")})
		fmt.Println("Greet result:", results[0].Interface().(string))
	}

	// 6) Creating a function at runtime with reflect.MakeFunc
	// func(int, int) int  -> returns sum
	intType := reflect.TypeOf(0)
	fnType := reflect.FuncOf([]reflect.Type{intType, intType}, []reflect.Type{intType}, false)
	sumFunc := reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		a := args[0].Int()
		b := args[1].Int()
		res := int(a + b)
		return []reflect.Value{reflect.ValueOf(res)}
	})
	// Call the dynamic function
	out := sumFunc.Call([]reflect.Value{reflect.ValueOf(7), reflect.ValueOf(5)})
	fmt.Println("7 + 5 =", out[0].Interface().(int))

	// 7) (Optional advanced) Create a struct type at runtime
	// NOTE: reflect.StructOf is available on supported Go versions.
	sf := []reflect.StructField{
		{
			Name: "ID",
			Type: reflect.TypeOf(0),
			Tag:  `json:"id"`,
		},
		{
			Name: "Note",
			Type: reflect.TypeOf(""),
			Tag:  `json:"note"`,
		},
	}
	dynType := reflect.StructOf(sf)
	dyn := reflect.New(dynType).Elem()
	dyn.FieldByName("ID").SetInt(123)
	dyn.FieldByName("Note").SetString("dynamic")
	fmt.Printf("Dynamic struct: %#v (type %s)\n", dyn.Interface(), dyn.Type())
}
