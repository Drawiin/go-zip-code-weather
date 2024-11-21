
# Go Zip Code Temperature

Este projeto é uma aplicação Go que recupera informações de temperatura com base em um código postal fornecido. Ele usa APIs externas para buscar dados meteorológicos e informações de código postal.

## Como Funciona

A aplicação executa os seguintes passos:
1. Recebe um código postal como entrada.
2. Usa o `CEP_SERVICE_URL` para obter informações de localização com base no código postal.
3. Usa o `WEATHER_API_URL` e `WEATHER_API_KEY` para buscar a temperatura atual para a localização.

## Pré-requisitos

- Docker
- Uma chave de API válida para o serviço de clima

## Executando a Aplicação

### Usando Docker

Você pode executar a aplicação usando Docker com o seguinte comando:

```sh
docker run -it --rm -p 8080:8080 \
   -e CEP_SERVICE_URL=https://brasilapi.com.br/api/cep/v2 \
   -e WEATHER_API_URL=https://api.weatherapi.com/v1/current.json \
   -e WEATHER_API_KEY=abc123 \
   -e PORT=8080 \
   vini65599/go-zip-code-temperature:latest
```

### Variáveis de Ambiente

As seguintes variáveis de ambiente são necessárias para executar a aplicação:

- `CEP_SERVICE_URL`: URL para o serviço de código postal.
- `WEATHER_API_URL`: URL para o serviço de clima.
- `WEATHER_API_KEY`: Chave de API para o serviço de clima.
- `PORT`: Porta na qual o servidor web será executado.

### Exemplo

Para executar a aplicação localmente com variáveis de ambiente, use o seguinte comando:

```sh
export CEP_SERVICE_URL=https://brasilapi.com.br/api/cep/v2
export WEATHER_API_URL=https://api.weatherapi.com/v1/current.json
export WEATHER_API_KEY=your_api_key_here
export PORT=8080

go run cmd/main.go
```

## Acessando a Aplicação

A aplicação está disponível em:

```
https://zip-code-temperature-392411933479.us-central1.run.app/temperature/{zipcode}
```

Substitua `{zipcode}` pelo código postal desejado para obter as informações de temperatura.
