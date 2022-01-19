# go-create

go-create is a tool that helps me create my apps (console apps, back-end web apps, full-stack web apps).

## Build

You need to have Golang v1.17.5+ to be able to build this project.

```bash
cd cmd/go-create
go build -o ../../bin/
```

## Usage

You need to add the executable to your path to use go-create everywhere. For more info, check out these links : <br>
[for windows users](https://medium.com/@kevinmarkvi/how-to-add-executables-to-your-path-in-windows-5ffa4ce61a53)<br>
[for linux users](https://medium.com/codex/adding-executable-program-commands-to-the-path-variable-5e45f1bdf6ce)

```bash
go-create <app-type> <folder-name>
go-create <db or model> <file-name>
go-create service <concerned-model>
```

## Examples

```bash
# creates the project in the current directory
go-create cli .
# creates the project in the directory test
go-create cli test
# creates a web app (you can later choose only back-end or full-stack with react or next or vue)
go-create web test
# creates a database file with postgres config in the current directory
go-create db database
# creates a crud service (articlesService.go) using the model Article
go-create service Article
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
