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
`,
	defs.ConfigKeyAdminRole,
	defs.ConfigKeyAdminUser,
)
