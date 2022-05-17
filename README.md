# api-numerais-romanos
API criada para receber uma palavra normal e encontrar números romanos dentro dela e escolher o maior numeral.

<h3> Descrição </h3>
<p> Recebi o desafio de criar uma API Rest que recebe uma palavra contínua e que indentifique onde há numerais romanos nesta palavra, e,
  ao analisá-los, identifique o maior. <br>
 A API possui uma única rota /search que recebe uma requisição REST em formato JSON
contendo a lista de números no corpo da requisição. Por exemplo, para a entrada apresentada, a requisição seria:<br>

{ "text": "AXXBLX" }<br>
  A resposta está feita também em formato JSON, retornando o maior número romano encontrado e
também o seu respectivo valor em formato decimal. No exemplo anteriormente apresentado, o resultado
seria:<br>
  {
"number": "LX",
"value": 60
}<br>
</p>
  ![](https://github.com/Natanael-devops/ilustracoes/blob/main/programa-romano.gif)

  <h3>Stacks</h3>
  <p>Desenvolvido utilizando:
<ul>
  <li>GO</li>
  <li>Docker</li>
  <li>Gin</li>
  <li>GORM</li>
  </ul></p>
  
  <h3>Abrir e rodar o projeto</h3>
  <p>É necessário possuir docker e go instalados no computador.
  Eu utilizei o Postman para analisar a aplicação, recomendo sua uilização.
  Após baixar o projeto, para abrir e rodar o projeto execute um comando em cada terminal:<br>
  <code>docker-compose up</code><br>
  <code>go run main.go</code><br>
  Para acessar a imagem docker do PgAdmin4 é necessário usar as configurações do arquivo "docker-compose.yml", ou, se necessário, alterá-las.
  
  
  
