package infra

type InfraTestResult struct {
	Bridgenf              error
	Bridgenf6             error
	Ipv4Forward           error
	Ipv6DefaultForwarding error
}
