# Nome do Projeto

Este é um projeto de exemplo usando Go e Kafka.

## Estrutura do Projeto

Explicação sobre a estrutura do projeto, destacando os principais diretórios e arquivos.

- **ms-go/**: Diretório principal da aplicação em Go.
  - **app/**: Contém a lógica da aplicação.
    - **controllers/**: Controladores da API.
    - **models/**: Modelos de dados da aplicação.
    - **services/**: Serviços da aplicação, como criação e atualização de produtos.
    - **producers/**: Produtores para enviar mensagens para o Kafka.
    - **consumers/**: Consumidores para receber mensagens do Kafka.
    - **helpers/**: Funções auxiliares.
    - **router/**: Configuração de roteamento da API.
  - **db/**: Configuração e conexão com o banco de dados.
  - **Dockerfile**: Arquivo para a criação da imagem Docker.
  - **go.mod e go.sum**: Arquivos de definição de módulos Go.
  - **main.go**: Arquivo principal da aplicação.
