
# Hacker-Rank-Challenges

Projeto para guardar as implementações dos desafios do https://www.hackerrank.com feitos em Golang.




## Instalação

Certifique-se de ter a versão **^1.21.3** do Golang instalada em sua máquina. 

Após clonar o projeto, acesse a pasta e execute os seguintes comandos para que o Go reconheça o projeto e faça o download das dependências necessárias

```bash
  cd hacker-rank-challenges
  go mod download
  go mod tidy
```
    
## Stack utilizada

**Back-end:** Golang


## Uso/Exemplos

Para executar um desafio, na raiz do projeto, no arquivo `main.go`, realize o import do pacote desejado e faça a chamada da única função existente no pacote. 

Obs: Lembre-se de adicionar os inputs necessários para que a função execute corretamente e também algumas funções não efetuam o print da resposta final, então certifique-se de adicionar a função `fmt.Println()`.

```golang
package main

import (
	"fmt"
	solvemefirst "hacker-rank/solve_me_first"
)

func main() {
	fmt.Println("Hacker Rank Challenges!!!")
	
	inputA := uint32(5)
	inputB := uint32(10)
	fmt.Println(solvemefirst.SumTwoNumbers(inputA, inputB))
}
```

Feito isso, basta executar no terminal na raiz do projeto, o seguinte comando:

```bash
  go run main.go
```

**Output**
```bash
❯ go run main.go
Hacker Rank Challenges!!!
15
```

