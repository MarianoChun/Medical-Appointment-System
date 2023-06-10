create or replace function attend_appointment(appointment_number int) returns boolean as $$
declare
    appointment record;
begin
    select * from turno where nro_turno = appointment_number into appointment;
    case 
        when not found then
            insert into error(operacion, f_error, motivo) values ('atención', now(), '?nro de turno no válido.');
            return false;
        when appointment.fecha::date != now()::date then
            insert into error(operacion, f_error, motivo) values ('atención', now(), '?turno no corresponde a la fecha del día.');
            return false;
        when appointment.estado != 'reservado' then
            insert into error(operacion, f_error, motivo) values ('atención', now(), '?turno no reservado.');
            return false;
        else
            update turno
            set estado = 'atendido'
            where nro_turno = appointment_number;
            return true;
    end case;    
end;
$$ language plpgsql;
