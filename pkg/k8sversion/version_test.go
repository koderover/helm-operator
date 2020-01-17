package k8sversion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Version(t *testing.T) {
	assert := require.New(t)
	v := SetVersion("v1.10.5-tke.3")
	assert.True(v.LessThan(V112))

	v = SetVersion("v1.14.8-aliyun.1")
	assert.True(V112.LessThan(v))

	v = SetVersion("v1.12.0-alpha")
	assert.False(V112.LessThan(v))

	v = SetVersion("v1.12.0")
	assert.False(V112.LessThan(v))
}
