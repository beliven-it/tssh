package defs

import (
	"fmt"
	"os"
)

func getHomePath() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		os.Exit(1)
	}

	return homePath
}

var ConfigKeyAdminRole = "admin_role"

var ConfigKeyTeleportProxy = "teleport_proxy"

var ConfigKeyTeleportUser = "teleport_user"

var ConfigKeyTeleportPasswordless = "teleport_passwordless"

var HomePath = getHomePath()

var ConfigFolderName = "tssh"

var ConfigFileName = "config"

var ConfigFileExtension = "yml"

var ConfigSSHAppName = fmt.Sprintf("%s.config", ConfigFolderName)

var ConfigFolderPath = fmt.Sprintf("%s/.config/%s", HomePath, ConfigFolderName)

var ConfigSSHMainPath = fmt.Sprintf("%s/.ssh/config", HomePath)

var ConfigSSHAppPath = fmt.Sprintf("%s/.ssh/%s", HomePath, ConfigSSHAppName)

var ConfigFilePath = fmt.Sprintf("%s/%s.%s", ConfigFolderPath, ConfigFileName, ConfigFileExtension)
