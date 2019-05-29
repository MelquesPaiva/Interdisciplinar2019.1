# Arquivo de orientações do projeto


## Obs:

- Não há necessidade de instalação de qualquer tipo de servidor (XAMP, WAMP, etc). Programa compilado
- Necessário realizar instalação do GO e do MongoDB em sua maquina

## Instalação do MongoDb

- Windows: https://medium.com/@LondonAppBrewery/how-to-download-install-mongodb-on-windows-4ee4b3493514
- Linux: https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/

## Instalação do GO

- Windows: https://www.freecodecamp.org/news/setting-up-go-programming-language-on-windows-f02c8c14e2f/
- Linux: https://tecadmin.net/install-go-on-ubuntu/

## Inicializando api

- cd /go/src/go-restapi
- go run main.go (Executando programa)

## Descrição de rotas

- interdisciplinar.yaml (Copia e cola em editor.swagger.io)

## Pacotes Utilizados

- (gopkg.in/mgo.v2) **Driver de conexão com o mongoDb**
- (github.com/gorilla/mux) **Pacote utilizado para trabalhar as rotas**
- (github.com/BurntSushi/toml) **Utilizado para criar parser e encoder de dados**


## Arquivos do projeto


### Configuração

- (config/config.go) **Dados de configuração com o banco de dados**
- (config/dao/sensor-dao.go) **Onde são feitas todas as nossas operações de CRUD**

### Model

- (models/senso.go) **Model que representa todos os dados enviados pelo ESP** 

### Rotas

- (router/sensor-router.go) **Onde serão criadas todas as rotas de nosso software**