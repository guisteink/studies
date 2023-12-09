### Informações do Docker:
```docker version``` Exibe informações sobre a versão do Docker instalada.
```docker info``` Mostra informações detalhadas sobre o sistema Docker.

### Imagens e Contêineres:
```docker images``` ou ```docker image ls``` Lista as imagens Docker no sistema.
```docker ps``` Lista os contêineres em execução.
```docker ps -a``` Lista todos os contêineres (em execução e parados).
```docker run <imagem>``` Cria e inicia um novo contêiner a partir de uma imagem.
```docker stop <contêiner>``` Para um contêiner em execução.
```docker start <contêiner>``` Inicia um contêiner parado.

### Construção de Imagens:
```docker build -t <nome-imagem> <caminho-contexto>``` Constrói uma imagem Docker a partir de um Dockerfile.
```docker rmi <nome-imagem>``` Remove uma imagem Docker.
```docker build -t <nome-imagem> .``` Constrói uma imagem Docker usando o Dockerfile no diretório atual.

### Remoção de Recursos:
```docker rm <contêiner>``` Remove um contêiner.
```docker rmi <imagem>``` Remove uma imagem.
```docker volume rm <volume>``` Remove um volume.

### Logs e Execução
```docker logs <contêiner>``` Exibe os logs de um contêiner.
```docker exec -it <contêiner> <comando>``` Executa um comando dentro de um contêiner em execução.

### Rede Docker
```docker network ls``` Lista as redes Docker.
```docker network inspect <rede>``` Exibe detalhes de uma rede Docker.

### Comandos Gerais:
```docker-compose up``` Inicia serviços definidos em um arquivo docker-compose.yml.
```docker-compose down``` Para e remove serviços definidos em um arquivo docker-compose.yml.
```docker-compose ps``` Lista os serviços em execução.

### Limpeza:
```docker system prune``` Remove todos os contêineres, redes e imagens não utilizados.

### Autenticação Docker Hub:
```docker login``` Faz login no Docker Hub para realizar push de imagens.
