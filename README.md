# Pat Life Care API

## cross compilation

Para generar el archivo compilado del API se tiene que ejecutar los siguientes comandos

* Para SO basados en linux

`````bash
GOOS=linux GOARCH=amd64 go build
`````

* Para SO basados en Windows

`````bash
GOOS=windows GOARCH=amd64 go build
`````

> _[NOTE]_ Para más información sobre cross compilation consultar la siguiente
> página [GO COMPILATION](https://golang.org/doc/install/source#environment)
