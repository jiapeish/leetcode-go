package _00517_factorz

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {

}