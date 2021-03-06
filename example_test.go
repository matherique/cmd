package cmd_test

import (
	"fmt"

	cmd "github.com/matherique/cmd"
)

type foo struct{}
func (foo) Handler(a []string) error {
	fmt.Println("foo", a)
  return nil
}

type bar struct{}
func (bar) Handler(a []string) error {
	fmt.Println("bar", a)
  return nil
}

func ExampleNew() {
	c := cmd.New("foo")

	fmt.Println(c.Name())

	// Output:
	// foo
}

func ExampleCommand_AddSub() {
	c := cmd.New("foo")
	sc := cmd.New("bar")

	c.AddSub(sc)
	fmt.Println(len(c.Sub()))

	// Output:
	// 1
}

func ExampleCommand_SetDesc() {
	c := cmd.New("foo")

	c.SetDesc("foo foo")

	fmt.Println(c.Desc())
	// Output:
	// foo foo
}

func ExampleCommand_SetLongDesc() {
	c := cmd.New("foo")
	c.SetLongDesc("long foo")

	fmt.Println(c.LongDesc())

	// Output:
	// long foo
}

func ExampleCommand_SetHandler() {
	c := cmd.New("foo")

	c.SetHandler(foo{})

	h := c.Handler()
	arg := make([]string, 0)
	h(arg)

	// Output:
	// foo []
}

func ExampleCommand_HasSub() {
	c := cmd.New("foo")
	sc := cmd.New("bar")

	c.AddSub(sc)

	s := c.HasSub("bar")

	fmt.Println(s.Name())

	s = c.HasSub("baz")

	if s == nil {
		fmt.Println("nil")
	}

	// Output:
	// bar
	// nil
}

func ExampleCommand_HasAlias() {
	c := cmd.New("foo", "f")

	alias := c.HasAlias("f")

	if alias {
		fmt.Println(alias)
	}

	alias = c.HasAlias("b")

	if !alias {
		fmt.Println(alias)
	}

	// Output:
	// true
	// false
}

func ExampleCommand_SetAlias() {
	c := cmd.New("foo")
	c.SetAlias("f")

	fmt.Println(c.HasAlias("f"))

	// Output:
	// true
}

func ExampleCommand_AddAlias() {
	c := cmd.New("foo")
	c.AddAlias("f")

	fmt.Println(c.HasAlias("f"))

	// Output:
	// true
}

func ExampleCommand_Run() {
	c := cmd.New("foo")
	c.SetLongDesc("foo long desc")
	c.SetHandler(foo{})

	sc := cmd.New("bar")
	sc.SetLongDesc("bar long desc")
	sc.SetHandler(bar{})

	c.AddSub(sc)

	var a []string
	c.Run(a)

	a = []string{"bar"}
	c.Run(a)

	a = []string{"help"}
	c.Run(a)

	a = []string{"bar", "help"}
	c.Run(a)

	a = []string{"arg1", "arg2"}
	c.Run(a)

	// Output:
	// foo []
	// bar []
	// foo long desc
	// bar long desc
	// foo [arg1 arg2]
}
