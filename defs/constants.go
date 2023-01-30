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

var ConfigKeyAdminUser = "admin_user"

var ConfigKeyAdminRole = "admin_role"

var HomePath = getHomePath()

var ConfigFolderName = "tssh"

var ConfigFileName = "config"

var ConfigFileExtension = "yml"

var ConfigFolderPath = fmt.Sprintf("%s/.config/%s", HomePath, ConfigFolderName)

var ConfigFilePath = fmt.Sprintf("%s/%s.%s", ConfigFolderPath, ConfigFileName, ConfigFileExtension)
