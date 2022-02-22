# go-create

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://choosealicense.com/licenses/mit/)
![Build](https://img.shields.io/badge/build-passed-green)
<br>
go-create is a tool that helps me create my apps (console apps, back-end web apps, full-stack web apps).

## Build

You need to have Golang v1.17.5+ to be able to build this project.<br>

```bash
./build.sh
```

## Usage

```bash
go-create web -t=<backend or fullstack> <folder-name>
go-create db <file-name>
go-create model <file-name>
go-create service <model-name>
```

## Examples

```bash
# creates the project in the current directory
go-create cli .
# creates the project in the directory test
go-create cli test
# creates a web app
go-create web -t=backend test
# creates a database file with postgres config in the current directory
go-create db database
# creates a crud service (articlesService.go) using the model Article
go-create service Article
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
