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
	User  string
	Host  string
	Refer string
}

type TctlRoleSpecAllowNodeLabels struct {
	Hostname string `json:"hostname"`
}

type TctlRoleSpecAllow struct {
	Logins     []string                    `json:"logins"`
	NodeLabels TctlRoleSpecAllowNodeLabels `json:"node_labels"`
}

type TctlRoleSpec struct {
	Allow TctlRoleSpecAllow `json:"allow"`
}

type TctlRoleMetadata struct {
	Name string `json:"name"`
}

type TctlRole struct {
	Kind     string           `json:"kind"`
	Metadata TctlRoleMetadata `json:"metadata"`
	Spec     TctlRoleSpec     `json:"spec"`
}
