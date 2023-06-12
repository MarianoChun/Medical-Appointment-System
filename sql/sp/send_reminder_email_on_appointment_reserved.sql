create or replace function send_reminder_on_appointment_reserved() returns void as $$
declare
    turno turno%rowtype;
    result record;
    appointment_date_to_remind date := (current_date + interval '2 days')::date;
    email_title text := 'Recordatorio de turno';
    email_body text;
begin
    for turno in select * from turno where estado = 'reservado' and (fecha + interval '2 days')::date = appointment_date_to_remind loop
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
        insert into envio_email (f_generacion, email_paciente, asunto, cuerpo, f_envio, estado)
        values (now(), result.email, email_title, email_body, now(), 'pendiente');

    end loop;
end;
$$ language plpgsql;