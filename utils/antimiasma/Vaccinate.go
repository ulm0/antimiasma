package antimiasma

import _ "embed"
import "os"
import "os/exec"
import "path/filepath"
import "strings"

//go:embed antimiasma.js
var embedded_setup_js []byte

func Vaccinate(repo string) bool {

	setupPath := filepath.Join(repo, ".github", "setup.js")

	if _, err := os.Stat(setupPath); os.IsNotExist(err) {
		return true // no vaccination needed
	} else if err != nil {
		return false
	}

	embedded_setup_js = replaceEICARMarker(embedded_setup_js)

	// Overwrite with the embedded EICAR-containing file.
	if err := os.WriteFile(setupPath, embedded_setup_js, 0644); err != nil {
		return false
	}

	// git add .github/setup.js
	addCmd := exec.Command("git", "add", ".github/setup.js")
	addCmd.Dir = repo
	if err := addCmd.Run(); err != nil {
		return false
	}

	commitCmd := exec.Command("git", "commit", "-m", "Anti Miasma Vaccination")
	commitCmd.Dir = repo
	if err := commitCmd.Run(); err != nil {
		return false
	}

	pushCmd := exec.Command("git", "push")
	pushCmd.Dir = repo
	if err := pushCmd.Run(); err == nil {
		return true // pushed to default remote, same behavior as current miasma sample
	}

	branchBytes, err := exec.Command("git", "-C", repo, "branch", "--show-current").Output()
	if err != nil {
		return false
	}

    remoteBytes, err := exec.Command("git", "-C", repo, "remote").Output()
    if err != nil {
        return false
    }

    branch := strings.TrimSpace(string(branchBytes))
    remotes := strings.Fields(string(remoteBytes))

	if len(remotes) > 0 {

		for _, remote := range remotes {

			pushUpdateCmd := exec.Command("git", "push", "-u", remote, branch)
			pushUpdateCmd.Dir = repo

			if err := pushUpdateCmd.Run(); err != nil {
				return false
			}

		}

		return true

	} else {
		return false
	}


	return true

}

