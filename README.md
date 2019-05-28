# Arquivo de orientações do projeto


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