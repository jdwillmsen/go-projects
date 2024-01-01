# Discord Ping Bot
A simple discord bot that will respond when pinged.

## Run App
### Local
To run the app you can either run the executable or build and run. (Or some combination depending on your workflow)
```shell
./main.exe
```
or
```shell
go build main.go && go run main.go
```
### Docker
To run the app you can also use the docker image.
```shell
docker run -d jdwillmsen/go-discord-ping-bot
```

### Docker Compose
To run the app you can also use the docker compose file.
```shell
docker compose up -d
```

## Config
This bot will require a config setup to run.
```json
{
  "Token": "<<Insert Token Here>>",
  "BotPrefix": "<<Insert Prefix Here | ! >>"
}
```

## Docker
### Build
#### Docker Hub
To be able to build the image you can run the following command from this directory.
```shell
docker build . -t jdwillmsen/go-discord-ping-bot:latest
```
i.e. (With additional version tag)
```shell
docker build . -t jdwillmsen/go-discord-ping-bot:latest -t jdwillmsen/go-discord-ping-bot:1.0.0
```

#### GitHub Containers
To be able to build the image for GitHub containers you can run the following command.
```shell
docker build . -t ghcr.io/jdwillmsen/go-discord-ping-bot:latest
```
i.e. (With additional version tag)
```shell
docker build . -t ghcr.io/jdwillmsen/go-discord-ping-bot:latest -t ghcr.io/jdwillmsen/go-discord-ping-bot:1.0.0
```

### Tag
After the image is built you can add additional tags with the following command.
```shell
docker tag jdwillmsen/go-discord-ping-bot:latest jdwillmsen/go-discord-ping-bot:<tag>
```

### Push
#### Docker Hub
To be able to push the image you can run the following command. (-a for all tags)
```shell
docker push -a jdwillmsen/go-discord-ping-bot
```

#### GitHub Containers
To be able to push the image you can run the following command. (-a for all tags)
```shell
docker push -a ghcr.io/jdwillmsen/go-discord-ping-bot
```

## Examples
![img.png](example.png)