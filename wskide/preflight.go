package wskide

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/mitchellh/go-homedir"
)

func preflightEnsureDockerVersion() error {
	vA := semver.New(MinDockerVersion)
	vB := semver.New(strings.TrimSpace(dockerVersion()))
	if vB.Compare(*vA) == -1 {
		return fmt.Errorf("Installed docker version %s is no longer supported", vB)
	}
	return nil
}

func preflightInHomePath(dir string) error {
	// do not check if the directory is empty
	if dir == "" {
		return nil
	}
	homePath, err := homedir.Dir()
	if err != nil {
		return err
	}
	dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(dir, homePath) {
		return fmt.Errorf("work directory %s should be under your home directory; this is required to be accessible by Docker", dir)
	}
	return nil
}

// Preflight perform preflight checks
func Preflight(dir string) error {
	err := preflightEnsureDockerVersion()
	if err != nil {
		return err
	}
	err = preflightInHomePath(dir)
	if err != nil {
		return err
	}
	return nil
}
