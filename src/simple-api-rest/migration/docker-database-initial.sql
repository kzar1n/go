create database db;

use db;


create table db.pagamentos (
    id                  varchar(36) default(uuid()) primary key,
    data                date,
    tipo                varchar(10),
    valor               decimal(15,2),
    id_pagador          varchar(36),
    id_conta_pagador  varchar(36),
    id_recebedor        varchar(36),
    banco_recebedor     varchar(3),
    agencia_recebedor   varchar(4),
    conta_recebedor     varchar(20)
);

insert into db.pagamentos  (data, tipo, valor, id_pagador, id_conta_pagador, id_recebedor, banco_recebedor, agencia_recebedor, conta_recebedor) 
                    values ('2024-12-31', "TED", 132, uuid(), uuid(), uuid(), "001", "0001", "000001230943949"),
                           ('2024-12-31', "PIX", 5453, uuid(), uuid(), uuid(), "123", "4356", "000001323339-X"),
                           ('2024-12-31', "TEF", 133, uuid(), uuid(), uuid(), "341", "1234", "0005201-2");

