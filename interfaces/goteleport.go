package interfaces

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"tssh/defs"
	"tssh/types"
	"tssh/utils"
)

type goteleport struct {
	status       types.GoteleportActive
	user         string
	proxy        string
	passwordless bool
}

type Goteleport interface {
	ListRoles() ([]string, error)
	ListHosts() ([]string, error)
	ListLogins() ([]string, error)
	Login() error
	Logout() error
	CreateSshConfig() error
	Connect(string) error
}

func (t *goteleport) includeSSHConfig() error {
	// Check the ssh/config file exists
	_, err := os.Stat(defs.ConfigSSHMainPath)
	if errors.Is(err, os.ErrNotExist) {
		_, e := os.Create(defs.ConfigSSHMainPath)
		if e != nil {
			return e
		}
	}

	file, err := os.OpenFile(defs.ConfigSSHMainPath, os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	defer file.Close()

	oldContent, err := os.ReadFile(defs.ConfigSSHMainPath)
	if err != nil {
		return err
	}

	oldContentToString := string(oldContent)

	delimiterStart := "# TSSH start managed"
	delimiterEnd := "# TSSH end managed"
	includeString := "Include " + defs.ConfigSSHAppName

	checker := regexp.MustCompile(defs.ConfigSSHAppName)
	hasReference := checker.MatchString(oldContentToString)

	var row = delimiterStart + "\n" + includeString + "\n" + delimiterEnd + "\n\n"
	if hasReference {
		deleteRegex := regexp.MustCompile("(?ms)" + delimiterStart + ".*" + delimiterEnd + "\n\n")
		oldContentToString = deleteRegex.ReplaceAllString(oldContentToString, "")
	}

	_, err = file.WriteString(row + oldContentToString)
	if err != nil {
		return err
	}

	return nil
}

func (t *goteleport) getStatus() error {
	output, err := utils.Exec("tsh", "status", "--format=json")
	if err != nil {
		return err
	}

	var response types.GoteleportCMDStatus
	err = json.Unmarshal(output, &response)
	if err != nil {
		return err
	}

	t.status = response.Active

	return nil
}

func (t *goteleport) ListHosts() ([]string, error) {
	output, err := utils.Exec("tsh", "ls", "--format=json")
	if err != nil {
		return nil, err
	}

	var response []types.GoteleportNode
	err = json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}

	hostnames := []string{}
	for _, node := range response {
		hostnames = append(hostnames, node.Spec.Hostname)
	}

	return hostnames, nil
}

func (t *goteleport) ListRoles() ([]string, error) {
	return t.status.Roles, nil
}

func (t *goteleport) ListLogins() ([]string, error) {
	return t.status.Logins, nil
}

func (t *goteleport) Connect(connection string) error {
	err := utils.ExecStdout("tsh", "ssh", connection)
	if err != nil {
		return err
	}

	return nil
}

func (t *goteleport) CreateSshConfig() error {
	output, err := utils.Exec("tsh", "config")
	if err != nil {
		return err
	}

	outputAsString := string(output)

	// Replace this snippet. In particular the %h must become
	// %n because %h doesn't handle uppercase characters
	replaceRule := regexp.MustCompile("%r@%h:%p")
	outputAsString = replaceRule.ReplaceAllString(outputAsString, "%r@%n:%p")

	err = os.WriteFile(defs.ConfigSSHAppPath, []byte(outputAsString), 0600)
	if err != nil {
		return err
	}

	// Add the refernce of the file created in main ssh config
	return t.includeSSHConfig()
}

func (t *goteleport) Login() error {
	args := []string{
		"login",
		"--proxy",
		t.proxy,
		"--user",
		t.user,
		"--auth",
	}

	if t.passwordless {
		args = append(args, "passwordless")
	} else {
		args = append(args, "local")
	}

	err := utils.ExecDevNull("tsh", args...)
	if err != nil {
		return err
	}

	return t.getStatus()
}

func (t *goteleport) Logout() error {
	return utils.ExecDevNull("tsh", "logout")
}

func NewGoteleportNotAuthInterface() Goteleport {
	i := goteleport{}

	return &i
}

func NewGoteleportInterface(user, proxy string, passwordless bool) (Goteleport, error) {
	i := goteleport{
		user:         user,
		proxy:        proxy,
		passwordless: passwordless,
	}

	// Check the status of current account
	err := i.getStatus()
	if err != nil {
		// If the user is not connected, must perform a login first
		err = i.Login()
		if err != nil {
			return &i, err
		}
	}

	return &i, err
}
