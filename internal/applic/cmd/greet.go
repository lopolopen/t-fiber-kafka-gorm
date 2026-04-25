package cmd

//go:generate go tool shoot new -json -file=$GOFILE

type Greet struct {
	Name string
}
