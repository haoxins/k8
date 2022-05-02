package k8

import (
	"github.com/banzaicloud/k8s-objectmatcher/patch"
	"k8s.io/apimachinery/pkg/runtime"
)

func Equals(a, b runtime.Object) (bool, error) {
	result, err := patch.DefaultPatchMaker.Calculate(a, b)
	if err != nil {
		return false, err
	}

	return result.IsEmpty(), nil
}
