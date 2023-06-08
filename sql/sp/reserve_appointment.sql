create or replace function reserve_appointment(nro_historia_clinica integer, dni_medique_reserva integer, fecha date, hora time) returns boolean as $$
-- validaciones
/*
- Los turnos sean asignados correctamente —e.g. un mismo turno no puede
ser reservado dos veces, el proceso de reserva de un turno no puede quedar a medio hacer (Transacciones),
etc.
- Que el DNI de médique exista. En caso de que no cumpla, se debe cargar un error
con el mensaje ?dni de médique no válido.
– Que el número de historia clínica exista. En caso de que no cumpla, se debe cargar
un error con el mensaje ?nro de historia clínica no válido.
– Si le paciente tiene una obra social, que le médique trabaje con esa obra social. En
caso de que no cumpla, se debe cargar un error con el mensaje ?obra social de
paciente no atendida por le médique.
– Que exista el turno de le médique para la fecha y la hora solicitadas, y que se
encuentre disponible. En caso de que no cumpla, se debe cargar un error con el
mensaje ?turno inexistente ó no disponible.
– Que le paciente no haya llegado al límite de 5 turnos en estado reservado. En caso
de que no cumpla, se debe cargar un error con el mensaje ?supera límite de
reserva de turnos.
*/
declare
    result record;
    nroObraSocialPaciente integer;
    turnosReservadosPaciente integer;
    montoConsulta decimal(12,2);
    timestampTurno timestamp;
begin
    select * into result from medique where medique.dni_medique = dni_medique_reserva;
    if not found then 
        perform create_reserve_appointment_error('?dni de médique no válido.', nro_historia_clinica, dni_medique_reserva, fecha, hora);
        raise notice 'No existe un medique con dicho dni, ingrese un dni existente';
        return false;
    end if;

    select * into result from paciente where paciente.nro_paciente = nro_historia_clinica;
    if not found then 
        perform create_reserve_appointment_error('?nro de historia clínica no válido.', nro_historia_clinica, dni_medique_reserva, fecha, hora);
        raise notice 'No existe un paciente con dicho nro de historia clinica, ingrese uno existente';
        return false;
    end if;

    select into nroObraSocialPaciente nro_obra_social from paciente where paciente.nro_paciente = nro_historia_clinica;
    select * into result from cobertura where cobertura.dni_medique = dni_medique_reserva and cobertura.nro_obra_social = nroObraSocialPaciente;
    if not found then
        perform create_reserve_appointment_error('?obra social de paciente no atendida por le médique.', nro_historia_clinica, dni_medique_reserva, fecha, hora);
        raise notice 'La obra social del paciente no es atendida por le médique';
        return false;
    end if; 

    timestampTurno := fecha + hora;
    select * into result from turno where turno.fecha = timestampTurno and turno.dni_medique = dni_medique_reserva;
    if not found or result.estado != 'disponible' then
        perform create_reserve_appointment_error('?turno inexistente ó no disponible.', nro_historia_clinica, dni_medique_reserva, fecha, hora);
        raise notice 'El turno es inexistente ó no esta disponible';
        return false;
    end if;

    turnosReservadosPaciente := count_reserved_appointments_for_patient(nro_historia_clinica);
    if turnosReservadosPaciente == 5 then
        perform create_reserve_appointment_error('?supera límite de reserva de turnos.', nro_historia_clinica, dni_medique_reserva, fecha, hora);
        raise notice 'El turno a reservar supera el límite de reserva de turnos';
        return false;
    end if;

    montoConsulta := calculate_consultation_amount(nro_historia_clinica, dni_medique_reserva);

    /*
    TODO: Actualizar la fila correspondiente en la tabla turno con los datos de le paciente y los montos de la consulta a abonar por elle y por la obra social, marcando el estado como reservado.
     */

    commit;
end;
$$ language plpgsql;

create or replace function create_reserve_appointment_error(msg text, nro_historia_clinica integer, dni_medique integer, fecha date, hora time) returns void as $$
declare
    errorCount int;
begin
    select count(*) into errorCount from error;
    insert into error (nro_error, dni_medique, nro_paciente, operacion, f_error, motivo)
    values (errorCount + 1, dni_medique, nro_historia_clinica, 'reserva', now(), msg); 
end;
$$ language plpgsql;

create or replace function count_reserved_appointments_for_patient(nro_paciente_consultado integer) returns integer as $$
    declare
        appointmentCount integer;
begin
    select count(*) into appointmentCount from turno where turno.nro_paciente = nro_paciente_consultado and turno.estado = 'reservado';
    return appointmentCount;
end;
$$ language plpgsql;

create or replace function calculate_consultation_amount(nro_paciente_consulta integer, dni_medique_consulta integer) returns decimal(12,2) as $$
declare
    nro_obra_social_paciente integer;
    consultationAmount decimal(12,2);
begin
    select nro_obra_social into nro_obra_social_paciente from paciente where paciente.nro_paciente = nro_paciente_consulta;
    if nro_obra_social_paciente is null then
        select monto_consulta_privada into consultationAmount from medique where dni_medique = dni_medique_consulta;
    else
        select (monto_paciente + monto_obra_social) into consultationAmount from cobertura where dni_medique = dni_medique_consulta and nro_obra_social = nro_obra_social_paciente;
    end if;

    return consultationAmount;
end;
$$ language plpgsql;