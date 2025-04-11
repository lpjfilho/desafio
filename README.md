# Desafio Técnico - Novo Portal Parceiros

## Descrição do Desafio

Neste desafio, você deverá construir um sistema para automatizar o cálculo de comissões que a empresa deve cobrar de seus parceiros.  
A empresa ajuda os parceiros a receberem dos consumidores suas dívidas e cobra uma comissão sobre cada pagamento.  

### O sistema deve ser capaz de:
- Calcular o valor da comissão que a empresa vai cobrar de cada parceiro.
- Retornar as dívidas pagas e o valor da comissão.

## Requisitos

### Funcionalidades do Sistema:
- **Cadastro de Parceiros**: A empresa tem vários parceiros, cada um com uma tabela de preço associado.
- **Tabelas de Preço**:
  - Faixas de dias de atraso da dívida.
- **Comissão**: 
  - O sistema deve calcular a comissão de acordo com a tabela de preço configurada.
- **Base de Dados de Dívidas Pagas**: 
  - O sistema deve acessar um array de dívidas pagas para calcular a comissão.
- **Endpoint**:
  - Deve receber um id do parceiro e retornar os valores cálculados. 

## Regras de Negócio:
- O cálculo da comissão é baseado na tabela associada ao parceiro.

## Tecnologias Sugeridas
- Linguagem: Golang.
- Framework Gin.

## O que será Avaliado
- Estruturação do Código: Organização de pastas, modularização, e clareza de código.
- Boas Práticas: Uso de Go idiomático, manuseio de erros, e concisão.
- Testes: Inclua testes unitários e, se possível, testes de integração.

## Duração do Teste
Este teste será realizado em uma sessão de live coding com duração aproximada de 40 minutos.

## Instruções Finais
- Antes de iniciar o teste, sinta-se à vontade para fazer perguntas e esclarecer quaisquer dúvidas que tenha.
- O tempo será cronometrado a partir do momento em que você começar a codificar.

Estamos ansiosos para ver como você abordará este desafio. Boa sorte! 🚀