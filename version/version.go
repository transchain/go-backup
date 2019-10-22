package version

var (
	GitCommit string
	Version   = "master"
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}