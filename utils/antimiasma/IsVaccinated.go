package antimiasma

import "os"
import "path/filepath"

func IsVaccinated(repo string) bool {

	setupPath := filepath.Join(repo, ".github", "setup.js")

	data, err := os.ReadFile(setupPath)
	if err != nil {
		return true
	}

	return containsEICARMarker(string(data))

}
