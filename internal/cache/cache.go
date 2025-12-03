package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Manifest struct {
	GistID      string    `json:"gist_id"`
	SHA         string    `json:"sha"`
	Description string    `json:"description"`
	Owner       string    `json:"owner"`
	Files       []string  `json:"files"`
	Source      string    `json:"source_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

var cleaner = regexp.MustCompile(`[^a-zA-Z0-9._-]`)

func Dir(cacheRoot, gistID, sha string) string {
	cleanID := cleaner.ReplaceAllString(gistID, "-")
	cleanSHA := cleaner.ReplaceAllString(sha, "-")
	return filepath.Join(cacheRoot, cleanID, cleanSHA)
}

func ManifestPath(cacheDir string) string {
	return filepath.Join(cacheDir, "manifest.json")
}

func LoadManifest(path string) (Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, err
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return Manifest{}, fmt.Errorf("parse manifest: %w", err)
	}
	return m, nil
}

func SaveManifest(path string, m Manifest) error {
	buf, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("encode manifest: %w", err)
	}
	if err := os.WriteFile(path, buf, 0o644); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}
	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func EnsureDir(path string) error {
	return os.MkdirAll(path, 0o755)
}

func PresentFiles(dir string, files []string) bool {
	for _, f := range files {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			return false
		}
	}
	return true
}

func Shorten(id string) string {
	if len(id) <= 8 {
		return id
	}
	return id[:8]
}

func JoinPath(base string, elems ...string) string {
	return filepath.Join(append([]string{base}, elems...)...)
}

func IsEmptyDir(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return len(entries) == 0
}
