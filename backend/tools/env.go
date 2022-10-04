package tools

import "os"

// GetEnv returns environment variable by name
func GetEnv(name, val string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return val
}
