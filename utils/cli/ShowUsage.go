package cli

import "fmt"
import "os"

func ShowUsage() {

	fmt.Fprintf(os.Stderr, "Antimiasma\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "GitHub: https://github.com/Team-Rockstars-Security/antimiasma\n")
	fmt.Fprintf(os.Stderr, "Usage:  Antimiasma scans for Miasma infected repositories inside the <path> folder.\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  antimiasma discover <path>\n")
	fmt.Fprintf(os.Stderr, "  antimiasma mitigate <path>\n")
	fmt.Fprintf(os.Stderr, "  antimiasma vaccinate <path>\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "IMPORTANT:\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "  The vaccinate action will implant this program and spread it as an anti-worm\n")
	fmt.Fprintf(os.Stderr, "  vaccine using the same spreading technique to fix infiltrated machones.\n")
	fmt.Fprintf(os.Stderr, "  It will self-delete if after 7 days no newly infected repository was discovered.\n")
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)

}
