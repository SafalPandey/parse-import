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

## Supports multiple languages

Currently, **parse-import** supports `ts` and `python` and is open to contributions for other languages. :)

## Recursively parses through the local imports

### Use `entrypoint` flag to recursively parse local imports in Python

```bash
go run main.go -f /path/to/main.py -l py -entryPoint /path/to/main.py
2020/01/25 17:11:21 Parsing using entry point file: /path/to/main.py
2020/01/25 17:11:21 Parsing imports for: [/path/to/main.py]
2020/01/25 17:11:21 Imports detected: 5
2020/01/25 17:11:21 Writing output to: ./imports.json
2020/01/25 17:11:21 Done
```

<details>
<summary>Here is how the `./imports.json` looks</summary>

```json
{
  "/path/to/baseDir/ablah/abc.py": {
    "IsLocal": true,
    "Path": "/path/to/baseDir/ablah/abc.py",
    "Info": {
      "Path": "/path/to/baseDir/ablah/abc.py",
      "IsDir": false,
      "Importers": [
        {
          "Line": 2,
          "Name": "some2",
          "Module": "ablah.abc",
          "Path": "/path/to/baseDir/main.py"
        }
      ]
    }
  },
  "/path/to/baseDir/utils/blahhah/something/": {
    "IsLocal": true,
    "Path": "/path/to/baseDir/utils/blahhah/something/",
    "Info": {
      "Path": "/path/to/baseDir/utils/blahhah/something/",
      "IsDir": true,
      "Importers": [
        {
          "Line": 4,
          "Name": "ess",
          "Module": "utils.blahhah.something",
          "Path": "/path/to/baseDir/main.py"
        }
      ]
    }
  },
  "/path/to/baseDir/utils/blahhah/uss.py": {
    "IsLocal": true,
    "Path": "/path/to/baseDir/utils/blahhah/uss.py",
    "Info": {
      "Path": "/path/to/baseDir/utils/blahhah/uss.py",
      "IsDir": false,
      "Importers": [
        {
          "Line": 3,
          "Name": "some3",
          "Module": "utils.blahhah.uss",
          "Path": "/path/to/baseDir/main.py"
        }
      ]
    }
  },
  "/path/to/baseDir/utils/utils.py": {
    "IsLocal": true,
    "Path": "/path/to/baseDir/utils/utils.py",
    "Info": {
      "Path": "/path/to/baseDir/utils/utils.py",
      "IsDir": false,
      "Importers": [
        {
          "Line": 1,
          "Name": "some",
          "Module": "utils.utils",
          "Path": "/path/to/baseDir/ablah/abc.py"
        },
        {
          "Line": 3,
          "Name": "some",
          "Module": "utils.utils",
          "Path": "/path/to/baseDir/utils/blahhah/uss.py"
        },
        {
          "Line": 1,
          "Name": "some",
          "Module": "utils.utils",
          "Path": "/path/to/baseDir/main.py"
        }
      ]
    }
  },
  "datetime": {
    "IsLocal": false,
    "Path": "datetime",
    "Info": {
      "Path": "datetime",
      "IsDir": false,
      "Importers": [
        {
          "Line": 1,
          "Name": "datetime",
          "Module": "datetime",
          "Path": "/path/to/baseDir/utils/blahhah/uss.py"
        }
      ]
    }
  }
}
```

</details>

**Note:** Supports relative imports in typescript/javascript by default.

### Supports tsconfig `baseUrl` property in Typescript

```bash
$ go run main.go -tsconfig /path/to/tsconfig.json -f /path/to/src/index.tsx
2020/01/25 17:12:55 Parsing using tsconfig file: /path/to/tsconfig.json
2020/01/25 17:12:55 Parsing imports for: [/path/to/src/index.tsx]
2020/01/25 17:12:55 Imports detected: 1295
2020/01/25 17:12:55 Writing output to: ./imports.json
2020/01/25 17:12:55 Done
```

## Supports custom output file path

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

1. Now shows details for each importer if something is imported multiple times.

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

2. Local imports json are now consistent with non local imports. No more nested maps.

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
