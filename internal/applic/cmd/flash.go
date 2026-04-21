package cmd

//go:generate go tool shoot new -json -file=$GOFILE

type FlashTryCmd struct {
	SO        string
	ProductID uint
}
