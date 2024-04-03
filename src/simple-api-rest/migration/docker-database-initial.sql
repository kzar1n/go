create database db;

use db;

create table db.contas (
    id_conta         varchar(36) default(uuid()) primary key,
    banco           varchar(3),
    agencia         varchar(4),
    conta           varchar(7),
    dac             varchar(1),
    tipo_conta       varchar(20)
);

create table db.pessoas (
    id_pessoa        varchar(36) default(uuid()) primary key,
    nome            varchar(30),
    documento       varchar(14),
    tipo_documento   varchar(4),
    id_conta         varchar(36),
    foreign key (id_conta) references db.contas(id_conta)
);

create table db.pagamentos (
    id              varchar(36) default(uuid()) primary key,
    data            date,
    tipo            varchar(10),
    valor           decimal(15,2),
    id_recebedor     varchar(36),
    id_pagador       varchar(36),
    foreign key (id_recebedor) references db.pessoas(id_pessoa),
    foreign key (id_pagador) references db.pessoas(id_pessoa)
);

SET @id_conta = uuid();
SET @id_pessoa = uuid();
insert into db.contas (id_conta, banco, agencia, conta, dac, tipo_conta) values (@id_conta, "341", "1500", "0005201", "2", "conta corrente");
insert into db.pessoas (id_pessoa, nome, documento, tipo_documento, id_conta) values (@id_pessoa, "Carol M. Young", "00096654119002", "CPF", @id_conta);

SET @id_conta2 = uuid();
insert into db.contas (id_conta, banco, agencia, conta, dac, tipo_conta) values (@id_conta2, "001", "0001", "1234567", "X", "conta digital");
SET @id_pessoa2 = uuid();
insert into db.pessoas (id_pessoa, nome, documento, tipo_documento, id_conta) values (@id_pessoa2, "Gary D. Fick", "00055025669006", "CPF", @id_conta2);
insert into db.pagamentos (data, tipo, valor, id_recebedor, id_pagador) values ('2024-12-31', "TED", 132, @id_pessoa, @id_pessoa2);
insert into db.pagamentos (data, tipo, valor, id_recebedor, id_pagador) values ('2024-03-02', "PIX", 12, @id_pessoa, @id_pessoa2);
insert into db.pagamentos (data, tipo, valor, id_recebedor, id_pagador) values ('2025-01-23', "TEF", 200, @id_pessoa, @id_pessoa2);

