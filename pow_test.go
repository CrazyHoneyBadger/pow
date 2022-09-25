package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_POWServer_create(t *testing.T) {
	lenUniqKey := 30
	complexity := 6
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.NotNil(t, ps)
	assert.Equal(t, lenUniqKey, ps.lengthUniqueKey)
	assert.Equal(t, complexity, ps.complexity)
}
func Test_POWServer_create_negative_uniq_key(t *testing.T) {
	lenUniqKey := -1
	complexity := 6
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.NotNil(t, ps)
	assert.Equal(t, 0, ps.lengthUniqueKey)
	assert.Equal(t, 0, ps.complexity)
}
func Test_POWServer_create_negative_complexity(t *testing.T) {
	lenUniqKey := 30
	complexity := -1
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.NotNil(t, ps)
	assert.Equal(t, 0, ps.lengthUniqueKey)
	assert.Equal(t, 0, ps.complexity)
}
func Test_POWServer_create_negative_uniq_key_and_complexity(t *testing.T) {
	ps := NewPOWServer(-1, -1)
	assert.NotNil(t, ps)
	assert.Equal(t, 0, ps.lengthUniqueKey)
	assert.Equal(t, 0, ps.complexity)
}

func Test_POWServer_generateUniqKey(t *testing.T) {
	lenUniqKey := 30
	complexity := 6
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.NotEmpty(t, ps.GenerateUniqKey())
	assert.Len(t, ps.GenerateUniqKey(), lenUniqKey)

}
func Test_POWServer_generateUniqKey_negative(t *testing.T) {
	ps := NewPOWServer(-1, 6)
	assert.Empty(t, ps.GenerateUniqKey())
}

func Test_POWServer_validateMessage_not_active(t *testing.T) {
	lenUniqKey := 0
	complexity := 6
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.NoError(t, ps.ValidateMessage(VERSION, ps.GenerateUniqKey()))
}
func Test_POWServer_validateMessage_version_error(t *testing.T) {
	lenUniqKey := 5
	complexity := 6
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.ErrorIs(t, ps.ValidateMessage("1.0.0", ps.GenerateUniqKey()), ErrVersion)
}
func Test_POWServer_validateMessage_complexity_error(t *testing.T) {
	lenUniqKey := 5
	complexity := 6
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.ErrorIs(t, ps.ValidateMessage(VERSION, "test:100"), ErrHash)
}

func Test_POWServer_validateMessage(t *testing.T) {
	lenUniqKey := 3
	complexity := 2
	ps := NewPOWServer(lenUniqKey, complexity)
	assert.NoError(t, ps.ValidateMessage(VERSION, "test:103"))
}

func Test_POWClient_create(t *testing.T) {
	maxIterationCount := uint64(100)
	pc := NewPOWClient(maxIterationCount)
	assert.NotNil(t, pc)
}

func Test_POWClient_signMessage_VersionErr(t *testing.T) {
	maxIterationCount := uint64(100)
	pc := NewPOWClient(maxIterationCount)
	_, err := pc.SignMessage("avs", "test", 0)
	assert.ErrorIs(t, err, ErrVersion)

}
func Test_POWClient_signMessage_MaxIterationErr(t *testing.T) {
	maxIterationCount := uint64(0)
	pc := NewPOWClient(maxIterationCount)
	_, err := pc.SignMessage(VERSION, "test", 6)
	assert.ErrorIs(t, err, ErrMaxIterationOverflow)

}
func Test_POWClient_signMessage(t *testing.T) {
	maxIterationCount := uint64(1 << 23)
	pc := NewPOWClient(maxIterationCount)
	data, err := pc.SignMessage(VERSION, "test", 2)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}
