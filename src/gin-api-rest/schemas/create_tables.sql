

create table if not exists minhas_autorizacoes.autorizacao (
    id_autorizacao              varchar(36) default(uuid()) primary key,
    id_conta_operacao           varchar(36) not null,
    data_hora_autorizacao       datetime not null,
    quantidade_contas           int,
    quantidade_total            int,
    valor_total                 decimal(15,2),
    quantidade_nao_efetuado     int,
    valor_nao_efetuado          decimal(15,2),
    quantidade_parcial          int,
    valor_parcial               decimal(15,2),
    quantidade_efetuado         int,
    valor_efetuado              decimal(15,2),
    quantidade_agendado         int,
    valor_agendado              decimal(15,2),
    situacao                    varchar(3),
    data_hora_criacao           datetime not null default '0001-01-01 00:00:00',
    data_hora_atualizacao       datetime not null default '0001-01-01 00:00:00',
    data_hora_delecao           datetime not null default '0001-01-01 00:00:00'
);

create table if not exists minhas_autorizacoes.contas_autorizacao (
    id_detalhe_autorizacao      varchar(36) default(uuid()) primary key,
    id_autorizacao              varchar(36) not null,
    id_conta_pagador            varchar(36) not null,
    quantidade_nao_efetuado     int,
    valor_nao_efetuado          decimal(15,2),
    quantidade_parcial          int,
    valor_parcial               decimal(15,2),
    quantidade_efetuado         int,
    valor_efetuado              decimal(15,2),
    quantidade_agendado         int,
    valor_agendado              decimal(15,2),
    foreign key (id_autorizacao) references autorizacao(id_autorizacao)
);

create table if not exists minhas_autorizacoes.pagamentos_autorizacao (
    id_pagamento                varchar(36) primary key,
    id_autorizacao              varchar(36) not null,
    id_conta_pagador            varchar(36) not null,
    id_recebedor                varchar(36) not null,
    nome_recebedor              varchar(50),
    tipo_pagamento              varchar(3),
    dados_conta_recebedor       varchar(30),
    data_pagamento              date not null,
    valor                       decimal(15,2) not null,
    descricao_motivo            varchar(250),
    foreign key (id_autorizacao) references contas_autorizacao(id_autorizacao)
);
