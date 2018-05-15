package common

type BaseStrategy interface {
	Handle(data map[string]string)
	BuildData()
}
