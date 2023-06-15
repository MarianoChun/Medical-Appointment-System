create or replace function send_email_on_appointment_reservation() returns trigger as $$
declare
    patient record;
    medic record;
    title text := 'Reserva de turno';
    body text;
    estado text := 'pendiente';
begin
    if new.estado = old.estado then
        return new;
    end if;

    if new.estado != 'reservado' then
        return new;
    end if;

    select * from paciente where nro_paciente = new.nro_paciente into patient;
    select * from medique where dni_medique = new.dni_medique into medic;

    select format('Turno reservado para el paciente %s, %s en la fecha de %s a las %s en el consultorio numero %s con el medico %s, %s', patient.apellido, patient.nombre, new.fecha::date, new.fecha::time, new.nro_consultorio, medic.apellido, medic.nombre) into body;

    insert into envio_email (f_generacion, email_paciente, asunto, cuerpo, estado) values (now(), patient.email, title, body, estado);

    return new;
end;
$$ language plpgsql;

create or replace trigger send_email_on_appointment_reserved after update of estado on turno for each row execute procedure send_email_on_appointment_reservation();