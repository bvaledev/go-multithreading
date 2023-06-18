# GO Multithreading

Este projeto tem como objetivo estudo de multithreading, retornando o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

- `https://cdn.apicep.com/file/apicep/" + cep + ".json`

- `http://viacep.com.br/ws/" + cep + "/json/`

Os requisitos são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.


## Features

- Channels
- Select
- Go Routines