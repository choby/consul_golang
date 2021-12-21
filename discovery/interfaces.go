package discovery

import "github.com/choby/consul_go/registry"

type DiscoveryClient interface {

	/**
	 * Gets all ServiceInstances associated with a particular serviceId.
	 * @param serviceId The serviceId to query.
	 * @return A List of ServiceInstance.
	 */
	GetInstances(serviceId string) ([]registry.ServiceInstance, error)

	/**
	 * @return All known service IDs.
	 */
	GetServices() ([]string, error)
}
