# parse-import

A tool to help parse imports and all the dependencies associated with a file.

## Usage

### Basic usage

```bash
$ go run main.go -f /path/to/file.ts
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Imports detected: 537
2019/12/25 16:26:45 Writing output to: ./imports.json
2019/12/25 16:26:45 Done
```

### Supports tsconfig `baseUrl` property

```bash
$ go run main.go -tsconfig /path/to/tsconfig.json -f /path/to/file.ts
2019/12/25 18:35:53 Parsing using tsconfig file: /path/to/tsconfig.json
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Imports detected: 598
2019/12/25 16:26:45 Writing output to: ./imports.json
2019/12/25 16:26:45 Done
```

### Supports custom output file path

```bash
$ go run main.go -tsconfig /path/to/tsconfig.json -f /path/to/file.ts -o yoyo.json
2019/12/25 18:35:53 Parsing using tsconfig file: /path/to/tsconfig.json
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Imports detected: 598
2019/12/25 16:26:45 Writing output to: yoyo.json
2019/12/25 16:26:45 Done
```

## Output

Outputs a json file of following format:

1. Now shows details for each importer if something is imported multiple.

    ```json
    "redux-persist": {
        "IsLocal": false,
        "Path": "redux-persist",
        "Info": {
        "Path": "redux-persist",
        "IsDir": false,
        "Importers": [
            {
            "Line": 13,
            "Name": "{ Persistor }",
            "Module": "'redux-persist';",
            "Path": "/path/to/src/app/utils/crossTabSync.ts"
            },
            {
            "Line": 16,
            "Name": "{ persistStore, autoRehydrate }",
            "Module": "'redux-persist';",
            "Path": "/path/to/src/store.ts"
            }
        ]
        }
    }
    ```

2. Local imports json are now consistent with non local imports (No more nested maps)

    ```json
    "/path/to/src/components/home/common/CalendarIcon.tsx": {
        "IsLocal": true,
        "Path": "/path/to/src/components/home/common/CalendarIcon.tsx",
        "Info": {
        "Path": "/path/to/src/components/home/common/CalendarIcon.tsx",
        "IsDir": false,
        "Importers": [
            {
            "Line": 27,
            "Name": "CalendarIcon",
            "Module": "'components/home/common/CalendarIcon';",
            "Path": "/path/to/src/components/home/accountability/time-and-attendance/UpdateTodo.tsx"
            },
            {
            "Line": 10,
            "Name": "CalendarIcon",
            "Module": "'../CalendarIcon';",
            "Path": "/path/to/src/components/home/common/fields/CreateTodo.tsx"
            }
        ]
        }
    }
    ```
