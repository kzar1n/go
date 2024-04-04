drop procedure if exists minhas_autorizacoes.dowhile;
delimiter //  
create procedure minhas_autorizacoes.dowhile()   
begin
declare i int default 0; 
while (i <= 100) do

    set @id_autorizacao = uuid();

    insert into minhas_autorizacoes.autorizacao 
    (id_autorizacao, id_conta_operacao, data_hora_autorizacao, quantidade_contas, quantidade_total, valor_total, quantidade_nao_efetuado, valor_nao_efetuado, quantidade_parcial, valor_parcial, quantidade_efetuado, valor_efetuado, quantidade_agendado, valor_agendado, situacao, data_hora_criacao)
    values
    (@id_autorizacao, uuid(), current_timestamp(), 10, 10, 100.0, 0, 0, 0, 0, 5, 50.0, 5, 50.0, "ok", current_timestamp());

    set @id_conta_pagador = uuid();

    insert into minhas_autorizacoes.contas_autorizacao 
    (id_autorizacao, id_conta_pagador, quantidade_nao_efetuado, valor_nao_efetuado, quantidade_parcial, valor_parcial, quantidade_efetuado, valor_efetuado, quantidade_agendado, valor_agendado)
    values
    (@id_autorizacao, @id_conta_pagador, 1, 10.2, 1, 22.2, 1, 33.3, 1, 44.4);

    insert into minhas_autorizacoes.pagamentos_autorizacao 
    (id_pagamento, id_autorizacao, id_conta_pagador, id_recebedor, nome_recebedor, tipo_pagamento, dados_conta_recebedor, data_pagamento, valor, descricao_motivo)
    values
    (uuid(), @id_autorizacao, @id_conta_pagador, uuid(), "nome_recebedor", "pix", "4052/9384762-9", current_date(), 23.7, "descricao_motivo"),
    (uuid(), @id_autorizacao, @id_conta_pagador, uuid(), "nome_recebedor", "pix", "4052/9384762-9", current_date(), 132.7, "descricao_motivo"),
    (uuid(), @id_autorizacao, @id_conta_pagador, uuid(), "nome_recebedor", "pix", "4052/9384762-9", current_date(), 225.6, "descricao_motivo"),
    (uuid(), @id_autorizacao, @id_conta_pagador, uuid(), "nome_recebedor", "pix", "4052/9384762-9", current_date(), 5423.7, "descricao_motivo");

    set i = i+1;
end while;
end;
//  

call minhas_autorizacoes.dowhile(); 






