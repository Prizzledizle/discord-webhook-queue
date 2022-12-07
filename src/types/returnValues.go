package types

type ReturnBool struct {
	Success bool
	Message string
	Content bool
}

type ReturnInt struct {
	Success bool
	Message string
	Content int
}

type ReturnString struct {
	Success bool
	Message string
	Content string
}
