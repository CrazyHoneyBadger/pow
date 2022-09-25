package pow

import "fmt"

var (
	ErrMaxIterationOverflow = fmt.Errorf("max iteration count overflow")
)

type POWClient struct {
	maxIteration uint64
}

func NewPOWClient(maxIterationCount uint64) *POWClient {
	return &POWClient{
		maxIteration: maxIterationCount,
	}
}
func (c POWClient) SignMessage(version, message string, complexity int) (string, error) {
	if VERSION != version {
		return "", ErrVersion
	}
	var i uint64
	for i = 0; i < c.maxIteration; i++ {
		data := fmt.Sprintf("%s:%d", message, i)
		hash := generate3Hash(data)
		if hashValidate(hash, complexity) {
			return data, nil
		}
	}
	return "", ErrMaxIterationOverflow
}
