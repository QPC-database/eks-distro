package internal

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

//const docsContentsDirectory = "docs/contents"

var (
	gitRootDirectory = GetGitRootDirectory()
	docsContentsDirectory = filepath.Join(gitRootDirectory, "docs/contents")
)


func GetGitRootDirectory() string {
	gitRootOutput, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(fmt.Sprintf("Unable to get git root directory: %v", err))
	}
	return strings.Join(strings.Fields(string(gitRootOutput)), "")
}

// GetKubernetesReleaseGitTag returns the trimmed value of Kubernetes release GIT_TAG
func GetKubernetesReleaseGitTag(releaseBranch string) (string, error) {
	fileOutput, err := ioutil.ReadFile(fmt.Sprintf("%s/projects/kubernetes/release/%s/GIT_TAG", gitRootDirectory, releaseBranch))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(fileOutput)), nil
}

// FormatDevelopmentReleasePath returns path to RELEASE in development for provided branch.
// Returned path is not guaranteed to exist or be valid.
func FormatDevelopmentReleasePath(branch string) string {
	return formatEnvironmentReleasePath(branch, DevelopmentRelease.String())
}

// FormatProductionReleasePath returns path to RELEASE in production for provided branch.
// Returned path is not guaranteed to exist or be valid.
func FormatProductionReleasePath(branch string) string {
	return formatEnvironmentReleasePath(branch, ProductionRelease.String())
}

// formatEnvironmentReleasePath returns path to RELEASE for provided branch and environment, which is expected to be
// "production" or "development". Returned path is not guaranteed to exist or be valid.
func formatEnvironmentReleasePath(branch, environment string) string {
	return filepath.Join(gitRootDirectory, "release", branch, environment, "RELEASE")
}

// FormatKubeGitVersionFilePath returns path to KUBE_GIT_VERSION_FILE for provided release.
// Expects release.Branch() to return a non-empty value.  Returned path is not guaranteed to exist or be valid.
func FormatKubeGitVersionFilePath(release *Release) string {
	return filepath.Join(gitRootDirectory, "projects/kubernetes/kubernetes", release.Branch(), "KUBE_GIT_VERSION_FILE")
}

// FormatReleaseDocsDirectory returns path to the directory for the docs' directory for provided release.
// Expects release.Branch() and release.Branch() to be non-empty values. Returned path is not guaranteed to exist or be
// valid.
func FormatReleaseDocsDirectory(release *Release) string {
	return filepath.Join(docsContentsDirectory, FormatRelativeReleaseDocsDirectory(release.Branch(), release.Number()))
}

// FormatRelativeReleaseDocsDirectory return relative path to (example: "releases/1-20/1").
// Expects release.Branch() and release.Number() to be non-empty values.
func FormatRelativeReleaseDocsDirectory(branch, number string) string {
	return filepath.Join("releases", branch, number)
}

func GetREADMEPath() string {
	return filepath.Join(gitRootDirectory, "README.md")
}

func GetDocsIndexPath() string {
	return filepath.Join(docsContentsDirectory, "index.md")
}
