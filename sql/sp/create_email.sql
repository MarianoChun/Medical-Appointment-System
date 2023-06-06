create or replace function send_email(email varchar, email_title varchar, email_body varchar) returns void as $$
declare
    last_id integer;
    email_id integer;
    email_status varchar;
begin
    select max(nro_email) into last_id from envio_email;

    if last_id is null then
        last_id = 0;
    end if;

    email_id := last_id + 1;
    email_status := 'enviado';

    insert into envio_email (nro_email, f_generacion, email_paciente, asunto, cuerpo, f_envio, estado)
    values (email_id, now(), email, email_title, email_body, now(), email_status);
end;
$$ language plpgsql;