package kubeconfig

type Cluster struct {
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty"`
	Server                   string `json:"server,omitempty"`
	TlsServerName            string `json:"tls-server-name,omitempty"`
}

type ClusterWithName struct {
	Cluster Cluster `json:"cluster"`
	Name    string  `json:"name"`
}
