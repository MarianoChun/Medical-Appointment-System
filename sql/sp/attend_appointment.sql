create or replace function attend_appointment(appointment_number int) returns boolean as $$
declare
    appointment record;
    appointmentAttended bool;
    existAppointment bool;
begin
    select * from turno where nro_turno = appointment_number into appointment;
    case 
        when not found then
            perform create_attend_appointment_error('?nro de turno no válido.', appointment);
            return false;
        when appointment.fecha::date != now()::date then
            perform create_attend_appointment_error('?turno no corresponde a la fecha del día.', appointment);
            return false;
        when appointment.estado != 'reservado' then
            perform create_attend_appointment_error('?turno no reservado.', appointment);
            return false;
        else
            update turno
            set estado = 'atendido'
            where nro_turno = appointment_number;
            return true;
    end case;    
end;
$$ language plpgsql;

create or replace function create_attend_appointment_error(msg text, appointment record) returns void as $$
declare
 errorNumber int;
begin
    select count(*) into errorNumber from error;
    if appointment.nro_turno is null then
        insert into error(nro_error, operacion, f_error, motivo) 
        values (errorNumber + 1, 'atención', now(), msg);
    else
        insert into error(nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) 
        values (errorNumber + 1, appointment.fecha, appointment.nro_consultorio, appointment.dni_medique, appointment.nro_paciente , 'atención', now(), msg);
    end if;
end;
$$ language plpgsql;
