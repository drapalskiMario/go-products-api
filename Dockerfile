# Use a imagem oficial do Golang
FROM golang:latest

# Defina o diretório de trabalho no container
WORKDIR /go/src/app

# Copie o código para o diretório de trabalho no container
COPY . .

# Instale as dependências
RUN go mod tidy

# Exponha a porta 3000
EXPOSE 3000

# Comando a ser executado quando o contêiner for iniciado
CMD ["tail", "-f", "/dev/null"]
