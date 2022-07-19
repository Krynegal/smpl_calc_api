package types

type Operand struct {
	Value string `json:"value"`
	Base  int    `json:"base"`
}

type Data struct {
	Operands [2]Operand
	ToBase   int `json:"toBase"`
}
