create or replace function send_email_on_appointment_canceled() returns trigger as $$
declare
    result record;
    last_id integer;
    email_id integer;
    email_title varchar;
    email_body varchar;
    email_status varchar;
begin
    if new.estado = old.estado then
        return new;
    end if;

    if new.estado <> 'cancelado' then
        return new;
    end if;

    select
        t.nro_turno,
        concat(p.nombre, ' ', p.apellido) as nombre_paciente,
        p.email as email_paciente,
        concat(m.nombre, ' ', m.apellido) as nombre_medique,
        t.fecha as fecha_turno
    into result
    from turno t, paciente p, medique m
    where t.nro_paciente = p.nro_paciente
    and m.dni_medique = t.dni_medique
    and t.nro_turno = old.nro_turno;

    select max(nro_email) into last_id from envio_email;

    if last_id is null then
        last_id = 0;
    end if;

    email_id := last_id + 1;
    email_title := 'Cancelación de turno';
    email_body := concat('¡Hola, ', result.nombre_paciente,'! Su turno con el medico ', result.nombre_medique, ' del día ', result.fecha_turno , ' ha sido cancelado. Pronto el centro de atención se contactará con usted.');
    email_status := 'enviado';

    insert into envio_email (nro_email, f_generacion, email_paciente, asunto, cuerpo, f_envio, estado)
    values (email_id, now(), result.email_paciente, email_title, email_body, now(), email_status);

    return new;
end;
$$ language plpgsql;

create trigger send_email_on_appointment_canceled after update on turno for each row execute procedure send_email_on_appointment_canceled();