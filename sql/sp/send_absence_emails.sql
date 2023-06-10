create or replace function send_absence_emails() returns void as $$
declare
    result record;
    email_title varchar := 'Turno cancelado';
    email_body varchar;
begin
    for result in select t.nro_turno, concat(p.nombre, ' ', p.apellido) as nombre_paciente, p.email as email_paciente, concat(m.nombre, ' ', m.apellido) as nombre_medique, t.fecha as fecha_turno into result from turno t, paciente p, medique m where t.nro_paciente = p.nro_paciente and m.dni_medique = t.dni_medique and t.estado = 'reservado' and t.fecha < now() loop
        email_body := concat('¡Hola, ', result.nombre_paciente,'! Su turno con el medico ', result.nombre_medique, ' del día ', result.fecha_turno , ' ha sido cancelado. Pronto el centro de atención se contactará con usted.');
        select send_email(result.email_paciente, email_title, email_body);
    end loop;
end;
$$ language plpgsql;