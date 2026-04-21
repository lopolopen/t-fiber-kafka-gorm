package cmd

//go:generate go tool shoot new -json -file=$GOFILE

type Demand struct {
	ProductID uint
	Qty       int
}

type TryReserveCmd struct {
	Origin  string
	Demands []Demand
	Remarks string
}

type CommitCmd struct {
	Origin string
}

type ReleaseCmd struct {
	Origin string
}
