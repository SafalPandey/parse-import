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

## Features

### Supports multiple languages

Currently, **parse-import** supports `ts` and `python` and is open to contributions for other languages. :)

### Recursively parses through the local imports

#### Use `entrypoint` flag to recursively parse local imports in Python

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
          "Name": "some",
          "Module": "utils.utils",
          "Path": "/path/to/baseDir/ablah/abc.py"
        },
        {
          "Name": "some",
          "Module": "utils.utils",
          "Path": "/path/to/baseDir/utils/blahhah/uss.py"
        },
        {
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

### Supports custom output file path

```bash
$ go run main.go -tsconfig /path/to/tsconfig.json -f /path/to/file.ts -o yoyo.json
2019/12/25 18:35:53 Parsing using tsconfig file: /path/to/tsconfig.json
2019/12/25 16:26:45 Parsing imports for: [/path/to/file.ts]
2019/12/25 16:26:45 Imports detected: 598
2019/12/25 16:26:45 Writing output to: yoyo.json
2019/12/25 16:26:45 Done
```

### You can now install parse-import in your TS/JS project with `package.json` :tada:

Simply run the install command to get the latest released version.

```bash
# With yarn
yarn add SafalPandey/parse-import

# With npm
npm install SafalPandey/parse-import
```

Note: You can run `yarn add SafalPandey/parse-import#<version-tag>` to get a specific version.

Test if installation is successful.

```bash
$ yarn run parse-import -h

Usage of /path/to/node_modules/.bin/parse-import:
  -entry-point string
        Path to main.py
  -f string
        File to parse
  -h    Show usage
  -l string
        Language to parse (default "ts")
  -o string
        Output file path (default "./imports.json")
  -tsconfig string
        Path to tsconfig file
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
                "Name": "{ Persistor }",
                "Module": "'redux-persist';",
                "Path": "/path/to/src/app/utils/crossTabSync.ts"
                },
                {
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
                "Name": "CalendarIcon",
                "Module": "'components/home/common/CalendarIcon';",
                "Path": "/path/to/src/components/home/accountability/time-and-attendance/UpdateTodo.tsx"
                },
                {
                "Name": "CalendarIcon",
                "Module": "'../CalendarIcon';",
                "Path": "/path/to/src/components/home/common/fields/CreateTodo.tsx"
                }
            ]
        }
    }
    ```
