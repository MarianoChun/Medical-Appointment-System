create or replace function send_email_on_appointment_reservation() returns trigger as $$
declare
    emailIndex int;
    patient record;
    medic record;
    body text;
begin
    if new.estado = old.estado then
        return new;
    end if;

    if new.estado != 'reservado' then
        return new;
    end if;

    select * from paciente where nro_paciente = new.nro_paciente into patient;
    select * from medique where dni_medique = new.dni_medique into medic;

    select format('Turno reservado para el paciente %s, %s en la fecha de %s a las %s en el consultorio numero %s con el medico %s, %s' 
        ,patient.apellido, patient.nombre, new.fecha::date, new.fecha::time, new.nro_consultorio, medic.apellido, medic.nombre) into body;

    select max(nro_email) from envio_email into emailIndex;

    if emailIndex is null then
        emailIndex = 0;
    end if;

    insert 
    into envio_email (nro_email, f_generacion, email_paciente, asunto, cuerpo, estado)
    values (emailIndex + 1, now(), patient.email, 'Reserva de turno', body, 'pendiente');

    return new;
end;
$$ language plpgsql;

create or replace trigger send_email_on_appointment_reservated after update of estado on turno for each row execute procedure send_email_on_appointment_reservation();