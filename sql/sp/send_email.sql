create or replace function send_email(email varchar, email_title varchar, email_body varchar) returns void as $$
declare
    email_status varchar := 'enviado';
begin
    insert into envio_email (f_generacion, email_paciente, asunto, cuerpo, f_envio, estado)
    values (now(), email, email_title, email_body, now(), email_status);
end;
$$ language plpgsql;