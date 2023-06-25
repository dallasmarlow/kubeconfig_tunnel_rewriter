package kubeconfig

type Context struct {
	Cluster string `json:"cluster"`
	User    string `json:"user"`
}

type ContextWithName struct {
	Name    string  `json:"name"`
	Context Context `json:"context"`
}
