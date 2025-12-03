# Manifest example

Gists can include a `gix.json` to tell gix exactly how to run instead of relying on shebangs or file extensions.

```json
{
  "run": "python app.py --verbose",
  "env": {
    "API_BASE": "https://api.example.com",
    "DEBUG": "1"
  }
}
```

- Place `gix.json` at the root of the gist. Override the filename with `--manifest <name>`.
- `run` is a single string and is executed via the shell (`sh -c` on Unix, `cmd /C` on Windows).
- `env` is optional; keys/values are injected for the run command.
