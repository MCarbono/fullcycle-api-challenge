<h1 align="center"> Client-Server API Challenge fullcycle </h1>


## Sobre 

Teste desenvolvido para o curso de GO do fullcycle. Consite em duas aplicações(client e server). A aplicação client
solicita a cotação do dólar para a aplicação server através de uma requisição http e salva o retorno em um arquivo .txt.
A aplicação server faz uma requisição http para um serviço externo que lhe retorna a cotação do diária do dólar, salva o registro no banco de dados sqlite e retorna para a client a cotação. 

--- 

Para iniciar a aplicação servidor, vá até a pasta server

```bash
    $ cd server
```

e digite um dos comandos abaixo: 

```bash
    # Makefile command
    $ make run
```

```bash
    # go command
    $ go run main.go
```

---

Para executar a aplicação cliente, vá até a pasta client

```bash
    $ cd client
```

e digite um dos comandos abaixo: 

```bash
    # Makefile command
    $ make run
```

```bash
    # go command
    $ go run main.go
```

é necessário que a aplicação server esteja com o servidor ligado para receber requisições da client.

---