package smpl_calc_api

import (
	"errors"
)

type Operand struct {
	Value string `json:"value"`
	Base  int    `json:"base" example:"10"`
}

type Data struct {
	Operand1 Operand
	Operand2 Operand
	ToBase   int `json:"toBase"`
}

func (d Data) CheckData() error {

	if d.ToBase < 2 || d.ToBase > 36 {
		return errors.New("ToBase parameter must be >=2 and <=36]")
	}

	op1 := d.Operand1
	if op1.Value == "" {
		return errors.New("empty first operand")
	}
	//_, err := strconv.Atoi(op1.Value)
	//if err != nil {
	//	return errors.New("can't parse first operand")
	//}
	if op1.Base < 0 {
		return errors.New("wrong base of first operand")
	}

	op2 := d.Operand2
	if op2.Value == "" {
		return errors.New("empty second operand")
	}
	//_, err = strconv.Atoi(op1.Value)
	//if err != nil {
	//	return errors.New("can't parse second operand")
	//}
	if op2.Base < 0 {
		return errors.New("wrong base of second operand")
	}
	return nil
}
