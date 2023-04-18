package templates

import (
	"fmt"
	"tssh/defs"
)

// Config template
var Config = fmt.Sprintf(`
# TSSH configuration file
fzf_options: "-i"
%s: ""
%s: "root"
%s: "teleport.domain.com"
%s: "my_username"
%s: false
`,
	defs.ConfigKeyAdminRole,
	defs.ConfigKeyAdminUser,
	defs.ConfigKeyTeleportProxy,
	defs.ConfigKeyTeleportUser,
	defs.ConfigKeyTeleportPasswordless,
)
