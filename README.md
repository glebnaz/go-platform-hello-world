# Go Platform

Шаблонный репозиторий для генерации кода.

Шаблон содержит:

- grpc-server
- http-server(grpc-gateway)
- debug-server (metrics,swagger)

Основные команды make:

- make build - билд проекта в папку bin
- make generate - генерация swagger и *.pb файлов из proto из api/
- make lint - запускает golangci lint

При первом запуске нужно запустить [install.sh](https://github.com/glebnaz/go-platform-hello-world/blob/master/install.sh), для установки buf, golangcilint , и остальных нужных инструментов

## Generate your microservice from template:

Для того, что бы сгененрировать сервис используйте эту команду.

```go
cookiecutter -f -o tmp/ template
```

После этого вы можете увидеть в папке tmp ваш новый сервис.

```go
//template/cookiecutter.json
{
	"project_name": "hello-world", -- название проекта
	"github_link": "github.com/glebnaz" -- укажите тут ссылку на ваш репозиторий
}
```

## Roadmap

- [x]  GRPC gateway
- [ ]  live/ready handlers
- [ ]  graceful shutdown
- [ ]  profiler on debug server
- [ ]  trace