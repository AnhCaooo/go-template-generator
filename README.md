# Go Web Server generator
A CLI tool to generate web server in [Go](https://go.dev/). The purpose of this tool is to save time for developers when they initialize the web server which is using Go. 

## General info
This CLI helps generate a web server based on my preference template which I have learned through self-referencing and the time I have been working with Go. This template is also inspired from these GitHub repositories: 
- [project-layout](https://github.com/golang-standards/project-layout)
- [go-nuts](https://github.com/francisco3ferraz/go-nuts)

This CLI tool will not only generate a web server, it will also provide a default Dockerfile configuration which you could use right way to build your own image. 


## Quick start 
Clone the project. Navigate to the project and run the following command
```bash
go run main.go
```
This will ask for where your project is located and your Go Module path.

### Tech stacks 
- Web server framework: [gorilla/mux](https://github.com/gorilla/mux)
- Logging framework: [zap](https://github.com/uber-go/zap)


## Changelog
- 1.0.0 - The first version. All basic features are working. However, currently, after the project is initialized, you, the consumer have to manually remove all 'template'import dependencies. In next version, maybe this will be handled automatically
