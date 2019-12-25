# parse-import

A tool to help parse imports and all the dependencies associated with a file.

## Usage

### Basic usage

```bash
$ go run main.go -f /path/to/file.ts
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Writing output to: ./imports.json
2019/12/25 16:26:45 Imports detected: 537
2019/12/25 16:26:45 Done
```

### Supports tsconfig `baseUrl` property

```bash
$ go run main.go -tsconfig /path/to/tsconfig.json -f /path/to/file.ts
2019/12/25 18:35:53 Parsing using tsconfig file: /path/to/tsconfig.json
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Writing output to: ./imports.json
2019/12/25 16:26:45 Imports detected: 598
2019/12/25 16:26:45 Done
```

### Supports custom output file path

```bash
$ go run main.go -tsconfig /path/to/tsconfig.json -f /path/to/file.ts -o yoyo.json
2019/12/25 18:35:53 Parsing using tsconfig file: /path/to/tsconfig.json
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Writing output to: yoyo.json
2019/12/25 16:26:45 Imports detected: 598
2019/12/25 16:26:45 Done
```

## Output

Outputs a json file of following format:

```json
{
  "/path/to/src/store.ts": {
    "path": {
      "to": {
        "src": {
          "store.ts": {
            "IsLocal": true,
            "Path": "/path/to/src/store.ts",
            "Info": {
              "Line": 16,
              "Path": "/path/to/src/store.ts",
              "Name": "store",
              "Module": "'../../store';",
              "IsDir": false,
              "ImportedIn": "/path/to/src/services/obs.ts"
            }
          }
        }
      }
    }
  },
  "axios": {
    "IsLocal": false,
    "Path": "axios",
    "Info": {
      "Line": 13,
      "Path": "axios",
      "Name": "{ AxiosPromise }",
      "Module": "'axios';",
      "IsDir": false,
      "ImportedIn": "/path/to/src/app/services/auth.ts"
    }
  }
}
```
