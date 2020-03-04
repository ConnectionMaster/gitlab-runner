package helperimage

import (
	"fmt"
	"strings"
)

const (
	windows1809 = "1809"
	windows1803 = "1803"

	baseImage1809 = "servercore1809"
	baseImage1803 = "servercore1803"

	windowsSupportedArchitecture = "x86_64"
)

var supportedOSVersions = map[string]string{
	windows1803: baseImage1803,
	windows1809: baseImage1809,
}

type unsupportedWindowsVersionError struct {
	version string
}

func newUnsupportedWindowsVersionError(version string) *unsupportedWindowsVersionError {
	return &unsupportedWindowsVersionError{version: version}
}

func (e *unsupportedWindowsVersionError) Error() string {
	return fmt.Sprintf("unsupported Windows version: %s", e.version)
}

var powerShellCmd = []string{"PowerShell", "-NoProfile", "-NoLogo", "-InputFormat", "text", "-OutputFormat", "text", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-Command", "-"}

type windowsInfo struct{}

func (w *windowsInfo) Create(revision string, cfg Config) (Info, error) {
	osVersion, err := w.osVersion(cfg.OperatingSystem)
	if err != nil {
		return Info{}, err
	}

	return Info{
		Architecture:            windowsSupportedArchitecture,
		Name:                    name,
		Tag:                     fmt.Sprintf("%s-%s-%s", windowsSupportedArchitecture, revision, osVersion),
		IsSupportingLocalImport: false,
		Cmd:                     powerShellCmd,
	}, nil

}

func (w *windowsInfo) osVersion(operatingSystem string) (string, error) {
	for osVersion, baseImage := range supportedOSVersions {
		if strings.Contains(operatingSystem, fmt.Sprintf(" %s ", osVersion)) {
			return baseImage, nil
		}
	}

	return "", newUnsupportedWindowsVersionError(operatingSystem)
}
