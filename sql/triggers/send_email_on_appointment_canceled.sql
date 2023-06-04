create or replace function send_email_on_appointment_canceled() returns trigger as $$
declare
    result record;
    email_body varchar;
begin
    if new.estado == old.estado and new.estado == 'cancelado' then
        return new;
    end if;

    if new.estado != 'cancelado' then
        return new;
    end if;

    select
        t.nro_turno,
        concat(p.nombre, ' ', p.apellido) as nombre_paciente,
        p.email as email_paciente,
        concat(m.nombre, ' ', m.apellido) as nombre_medique
    into result
    from turno t, paciente p, medique m
    where t.nro_paciente = p.nro_paciente
    and m.dni_medique = t.dni_medique;

    email_body := concat('Hola, ',  result.nombre_paciente,'! Su turno nro', result.nro_turno, 'ha sido cancelado con el medico ', result.nombre_medique);

    insert into envio_email (f_generacion, email_paciente, asunto, cuerpo, f_envio, estado)
    values (now(), result.email_paciente, 'Cancelación de turno', email_body, now(), 'enviado');

    return new;
end;
$$ language plpgsql;

create trigger send_email_on_appointment_canceled
after update on turno
for each row
execute procedure send_email_on_appointment_canceled();