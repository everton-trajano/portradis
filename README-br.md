# Portradis

**Portradis (PORtfolio TRAjano DIctionary Server)** é um remake simplificado do Redis, desenvolvido em Go. Este projeto tem como objetivo recriar algumas das funcionalidades básicas do Redis para estudo e demonstração de habilidades em desenvolvimento backend, estruturação de código e uso de boas práticas.
[English version](https://github.com/everton-trajano/portradis/blob/main/README.md)

## Sobre o Projeto

O Redis é um banco de dados de estrutura chave-valor amplamente utilizado devido à sua alta performance e versatilidade. O Portradis foi criado como uma implementação educacional e simplificada, permitindo entender os conceitos fundamentais por trás de um servidor de banco de dados similar ao Redis.

## Funcionalidades

- Integração funcional com redis-cli.
- Servidor HTTP para operações básicas de armazenamento de dados.
- Manipulação de estruturas chave-valor.
- Código modular e organizado para fácil extensão.

## Requisitos

- **Go** (versão 1.20+ recomendada)
- Qualquer sistema operacional compatível com Go.

## Instalação

1. Clone este repositório:
  ```bash

   git clone https://github.com/everton-trajano/portradis.git
```
2. Navegue até o diretório do projeto:
  ```bash

  cd portradis
  ```
3. Instale as dependências:
  ```bash

  go mod tidy
  ```
4. Para iniciar o servidor, use:
  ```bash

  go run main.go
  ```

## Estrutura do Projeto

   * app/: Código principal da aplicação.
   * commands/: Scripts de inicialização e configuração.
   * tools/: Ferramentas e utilitários auxiliares.

## Contribuição

Sinta-se à vontade para abrir issues e pull requests. Feedbacks são sempre bem-vindos!
