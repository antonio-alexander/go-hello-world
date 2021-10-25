package cli

// These variables are populated at build time
// REFERENCE: https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
var (
	Version   string = "latest"
	GitCommit string
	GitBranch string
)
