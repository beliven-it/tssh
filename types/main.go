package types

type GoteleportActive struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Logins   []string `json:"logins"`
}

type GoteleportNodeSpec struct {
	Hostname string
}

type GoteleportNode struct {
	Spec GoteleportNodeSpec
}

type GoteleportCMDStatus struct {
	Active GoteleportActive `json:"active"`
}

type TsshConnection struct {
	User string
	Host string
}
