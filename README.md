# 📦 go-json-server

Simple server that replies json from file.

## How to use

1. create `db` directory
2. add json file to ./db directory. name of the file will be the endpoint and content of the file will be the response.
3. just run the following command:

```shell
$ curl -L -o app https://github.com/chaewonkong/go-json-server/raw/main/bin/go-json-server
$ chmod 777 app
$ ./app
```

This server will automatically finds files in ./db directory and serves it.
your server will be at `localhost:8080`

👨‍💻 Happy Hacking~
