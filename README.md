# GetNinjas DevOps test.


## Sobre o teste
O propósito desse teste não é avaliar certo ou errado. Mas sim formas de **pensar e trabalhar** do candidato. Queremos avaliar também, **qualidade de código e simplicidade** (não confunda simplicidade com relaxo, mas sim o quão **eficiente** é o código/arquitetura proposta)
Sinta-se livre para desenvolver linhas de raciocino, comentar, etc...
E em caso de dúvidas, não tenha vergonha de enviar um email para **devops@getninjas.com.br**

## Cenário
Temos no repo https://github.com/getninjas/devops_test uma aplicação muito simples, uma API rest escrita em Golang, que atualmente só responde à rota `/healthcheck`. Essa aplicação, em compliance com o item III do 12factor app, espera alguns **parametros via ambiente** para rodar corretamente.
Outro ponto importante é que este código tem **cobertura de testes***.

## Objetivos
Dado o Cenário acima queremos que você faça o seguinte:

#### 1. Deploy da aplicação na AWS.
Pra isso, você provavelmente precisará de uma conta free-tier da AWS ou uma já existente, mas não se preocupe, não iremos olhar sua conta ou chamar a api já rodando, queremos que você crie uma forma que possamos recriar toda sua infraestrutura em nossa conta de forma simples.<br>
Mas por favor, leve em conta que seria um ambiente muito próximo à produção, portanto, **escalabilidade**, billing, segurança, **monitoramento** e **logging** tudo isso será levado em consideração na nossa análise.
Sinta-se livre para usar as ferramentas que julgue necessário, mas lembre-se: que na dúvida, vá pelo simples.
Crie um desenho que represente a arquitetura envolvida.

#### 2. Crie uma forma que possamos subir essa aplicação localmente de forma simples.
Mais uma vez, sinta-se livre para usar a ferramenta que julgar necessária.

#### 3. Coloque esta aplicação em um fluxo de CI que realize teste neste código
Pode ser via Travis ou qualquer outro CI. Nos mande a url de um build.

#### 4. Altere o nome da aplicação.
Fazendo um `curl -i http://endpoint-da-aplicacao/healthcheck` a resposta se assemelha a isto:
``` 
HTTP/1.1 200 OK
Date: Wed, 05 Sep 2018 18:07:41 GMT
Content-Length: 24
Content-Type: text/plain; charset=utf-8

Hey Bro, Ninja is Alive!%
```

Como podemos alterar o "nome da aplicação (Ninja)", para outro nome? Nos envie a saída do seu `curl` com a resposta, já com o nome alterado.


#### 5. Discorra qual (ou quais) processos você adotaria para garantir uma entrega contínua desta aplicação, desde o desenvolvimento, até a produção.


## Forma de entrega:
Faça um fork deste repo, e edite o arquivo ANSWERS.md, colocando suas respostas, observações, como rodar o código e o que mais achar necessário de informação.<br>
Faça os commits do seu código/arquivos no seu repo forkado, e quando pronto, nos envie o repo zipado, para que possamos avaliar.
Não esqueça de colocar no ANSWERS.md as instruções. O máximo de detalhes que colocar, fará com que possamos subir mais facilmente a sua infra.

**Boa sorte!**


