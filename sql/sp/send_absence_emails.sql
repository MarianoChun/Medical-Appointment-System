create or replace function send_absence_emails() returns void as $$
declare
    result record;
    has_been_email_sended int;
    email_body varchar;
    email_status varchar := 'enviado';
    email_title varchar := 'Turno cancelado';
begin
    for result in select t.nro_turno, concat(p.nombre, ' ', p.apellido) as nombre_paciente, p.email as email_paciente, concat(m.nombre, ' ', m.apellido) as nombre_medique, t.fecha as fecha_turno into result from turno t, paciente p, medique m where t.nro_paciente = p.nro_paciente and m.dni_medique = t.dni_medique and t.estado = 'reservado' and date(t.fecha) = date(now() - CAST('1 days' AS INTERVAL)) loop
        email_body := concat('¡Hola, ', result.nombre_paciente,'! Su turno con el medico ', result.nombre_medique, ' del día ', result.fecha_turno , ' ha sido cancelado. Pronto el centro de atención se contactará con usted.');

        select count(1) into has_been_email_sended
        from envio_email e
        where e.email_paciente = result.email_paciente
        and e.asunto = email_title
        and cuerpo = email_body;

        if has_been_email_sended == 0 then
            insert into envio_email (f_generacion, email_paciente, asunto, cuerpo, estado)
            values (now(), result.email_paciente, email_title, email_body, email_status);
        end if;
    end loop;
end;
$$ language plpgsql;