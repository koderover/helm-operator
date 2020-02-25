package status

import (
	v1 "github.com/fluxcd/helm-operator/pkg/apis/helm.fluxcd.io/v1"
	v1client "github.com/fluxcd/helm-operator/pkg/client/clientset/versioned/typed/helm.fluxcd.io/v1"
	"github.com/fluxcd/helm-operator/pkg/k8sversion"
)

func UpdateStatus(client v1client.HelmReleaseInterface, hr *v1.HelmRelease) (*v1.HelmRelease, error) {
	if k8sversion.GetVersion().LessThan(k8sversion.V112) {
		return client.Update(hr)
	}

	return client.UpdateStatus(hr)
}
