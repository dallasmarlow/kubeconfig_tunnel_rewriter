package kubeconfig

type User struct {
	ClientCertificateData []byte   `json:"client-certificate-data,omitempty"`
	ClientKeyData         []byte   `json:"client-key-data,omitempty"`
	Exec                  UserExec `json:"exec,omitempty"`
	Password              string   `json:"password,omitempty"`
	Username              string   `json:"username,omitempty"`
	Token                 string   `json:"token,omitempty"`
}

type UserExec struct {
	ApiVersion string              `json:"apiVersion"`
	Args       []string            `json:"args"`
	Command    string              `json:"command"`
	Env        []map[string]string `json:"env,omitempty"`
}

type UserWithName struct {
	Name string `json:"name"`
	User User   `json:"user"`
}
