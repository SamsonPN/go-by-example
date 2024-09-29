package main

import "fmt"

// embedding structs and interfaces allows for composition of types

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// this container struct EMBEDS a base struct
// it looks like a field without a name
type container struct {
	base
	str string
}

func main() {

	// when creating structs with literals
	// we have to initialize the embedding explicitly
	// and the embedded type is the name of the field here
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// can access fields from base directly from co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// or we can use the full path with the embedded type name
	fmt.Println("also num:", co.base.num)

	// since container embeds base
	// the methods of base become the methods of container
	fmt.Println("describe: ", co.describe())

	// base implemented the method describer which is part of this describer interface
	// since container embeds base, container ALSO implements the describer interface
	type describer interface {
		describe() string
	}

	// since container embeds base, which implements the describer interface
	// container also implements the describer interface
	var d describer = co
	fmt.Println("describer: ", d.describe())
}
