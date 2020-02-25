package k8sversion

import "github.com/Masterminds/semver"

var version *semver.Version
var V112 = semver.MustParse("v1.12.0")

func SetVersion(ver string) *semver.Version {
	var err error
	if version, err = semver.NewVersion(ver); err != nil {
		panic(err)
	}

	return version
}

func GetVersion() *semver.Version {
	if version != nil {
		return version
	}

	return V112
}
