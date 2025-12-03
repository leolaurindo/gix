package alias

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

func Load(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return map[string]string{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read aliases: %w", err)
	}
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parse aliases: %w", err)
	}
	if m == nil {
		m = map[string]string{}
	}
	return m, nil
}

func Save(path string, aliases map[string]string) error {
	buf, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return fmt.Errorf("encode aliases: %w", err)
	}
	if err := os.WriteFile(path, buf, 0o644); err != nil {
		return fmt.Errorf("write aliases: %w", err)
	}
	return nil
}

func Sorted(aliases map[string]string) []string {
	out := make([]string, 0, len(aliases))
	for k := range aliases {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func CopyMap(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func PrintList(w io.Writer, aliases map[string]string) {
	names := Sorted(aliases)
	for _, n := range names {
		fmt.Fprintf(w, "%s -> %s\n", n, aliases[n])
	}
}
