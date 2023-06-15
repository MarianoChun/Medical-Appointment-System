create or replace function send_reminder_on_appointment_reserved() returns void as $$
declare
    turno turno%rowtype;
    result record;
    appointment_date_to_remind date;
    reminderInterval interval;
    email_title text := 'Recordatorio de turno';
    email_body text;
    has_been_email_sent int := 0;
begin
    case
        when date_part('dow', current_date) = 4 then
            reminderInterval := ('4 days')::interval;
        when date_part('dow', current_date) = 5 then
            reminderInterval :=('3 days')::interval;
        else
            reminderInterval := ('2 days')::interval;
    end case;

    appointment_date_to_remind := (current_date + reminderInterval)::date;

    for turno in select * from turno where estado = 'reservado' and date(fecha) = appointment_date_to_remind and date(fecha) != current_date loop
        select
            p.email,
            concat(p.nombre,' ',p.apellido) as patient_full_name,
            turno.monto_paciente,
            turno.fecha,
            c.nombre as consultory_room_name,
            concat(m.nombre,' ',m.apellido) as medic_full_name
            into result
            from medique m, paciente p, consultorio c
            where
                turno.dni_medique = m.dni_medique and turno.nro_paciente = p.nro_paciente and turno.nro_consultorio = c.nro_consultorio;

        email_body := concat('Estimado ', result.patient_full_name ,',le recordamos que tiene un turno para la fecha ', result.fecha, ' en el consultorio ', result.consultory_room_name,
            ' con el doctor ', result.medic_full_name, '. Recuerde que el monto de la consulta es de ', result.monto_paciente);

        select count(1)
        from envio_email
        where cuerpo = email_body
        and email_paciente = result.email
        and asunto = email_title
        into has_been_email_sent;

        if has_been_email_sent = 0 then
            insert into envio_email (f_generacion, email_paciente, asunto, cuerpo, f_envio, estado)
            values (now(), result.email, email_title, email_body, now(), 'pendiente');
        end if;

    end loop;
end;
$$ language plpgsql;