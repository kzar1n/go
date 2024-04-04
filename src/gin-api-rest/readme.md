Tabela Autorizacao

idAutorizacao
idContaOperacao
dataHoraAutorizacao
quantidadeContas
quantidadeTotal
valorTotal
quantidadeNaoEfetuado
valorNaoEfetuado
quantidadeParcial
valorParcial
quantidadeEfetuado
valorEfetuado
quantidadeAgendado
valorAgendado
situacao

Tabela contasAutorizacao
IdDetalheAutorizacao
idAutorizacao
idContaPagador
quantidadeNaoEfetuado
valorNaoEfetuado
quantidadeParcial
valorParcial
quantidadeEfetuado
valorEfetuado
quantidadeAgendado
valorAgendado

Tabela pagamentosAutorizacao
IdPagamento
idAutorizacao
idContaPagador
idRecebedor
nomeFavorecido
tipoPagamento
dadosConta
dataPagamento
valor
descricaoMotivo




APIs
1 - Listagem das autorizacoes
endpoint: 
    /minhas-autorizacoes
queryParameters:
    id_conta_operacao
    data_inicio
    data_fim


2 - Exibe totais por situacao

3 - Busca por nome ou conta
4 - Lista as contas e os totais por conta
5 - Lista os pagamentos por conta



