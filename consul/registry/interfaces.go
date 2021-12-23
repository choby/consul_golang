package registry

type ServiceRegistry interface {
	Register(serviceInstance ServiceInstance) bool

	Deregister()
}
