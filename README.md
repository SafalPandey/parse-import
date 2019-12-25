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
