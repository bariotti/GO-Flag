# GO-Flag

Projeto exemplificando a utilização da biblioteca FLAG do GO, trabalhando com passagem de parâmetros para aplicações console.

Nesse exemplo, o intuito é vasculhar um arquivo de LOG procurando por logs de um determina tipo.

Os parametros disponiveis são:
path - opcional - valor default: log myapp.log na raiz da aplicação -  arquivo de log que será vasculhado procurando logs de um determinado tipo
logType - tipo do LOG - valor default: ERRO - tipo do log a ser procurado no arquivo

Somente a fim de exemplificar o uso da biblioteca FLAG, o arquivo de LOG utilizado tem um padrão de 1 registro por linha, no formato [data] - [tipo do log] - [mensagem]
Exemplo:

2022-02-22 - SUCESSO - Arquivo lido com sucesso

2022-02-22 - ERRO - Não foi possivel ler um arquivo

Exemplo de como executar a aplicação passando parametros:

lerlogs.exe

lerlogs.exe -logType SUCESSO

lerlogs.exe -logType ERRO

lerlogs.exe -logType SUCESSO -path 

lerlogs.exe -path C:\Logs\myapp.log

lerlogs.exe -path C:\Logs\myapp.log -logType SUCESSO

