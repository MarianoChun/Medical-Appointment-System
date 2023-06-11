create or replace function reserve_appointment(nro_historia_clinica integer, dni_medique_reserva integer, fechaHora timestamp) returns boolean as $$
declare
    result record;
    turnoAReservar record;
    nroObraSocialPaciente integer;
    nroAfiliadePaciente integer := null;
    turnosReservadosPaciente integer;
    montoPaciente decimal(12,2);
    montoObraSocial decimal(12,2);
    timeStampTurnoSolicitado timestamp := fechaHora;
begin
    select * into result from medique where medique.dni_medique = dni_medique_reserva;
    if not found then
        insert into error (nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) values (default, null, null, dni_medique_reserva, nro_historia_clinica, 'reserva', now(), '?dni de médique no válido');
        raise notice 'No existe un medique con dicho dni, ingrese un dni existente';
        return false;
    end if;

    select * into result from paciente where paciente.nro_paciente = nro_historia_clinica;
    if not found then
        insert into error (nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) values (default, null, null, dni_medique_reserva, nro_historia_clinica, 'reserva', now(), '?nro de historia clínica no válido');
        raise notice 'No existe un paciente con dicho nro de historia clinica, ingrese uno existente';
        return false;
    end if;

    select into nroObraSocialPaciente nro_obra_social from paciente where paciente.nro_paciente = nro_historia_clinica;
    if nroObraSocialPaciente is not null then
        select * into result from cobertura where cobertura.dni_medique = dni_medique_reserva and cobertura.nro_obra_social = nroObraSocialPaciente;
        if not found then
            insert into error (nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) values (default, null, null, dni_medique_reserva, nro_historia_clinica, 'reserva', now(), '?obra social de paciente no atendida por le médique');
            raise notice 'La obra social del paciente no es atendida por le médique';
            return false;
        end if;

        select into nroAfiliadePaciente nro_afiliade from paciente where paciente.nro_paciente = nro_historia_clinica;
    end if;


    select * into turnoAReservar from turno where date_trunc('hour', turno.fecha) = timeStampTurnoSolicitado and turno.dni_medique = dni_medique_reserva and turno.estado = 'disponible' limit 1;
    if not found then
        insert into error (nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) values (default, null, null, dni_medique_reserva, nro_historia_clinica, 'reserva', now(), '?turno inexistente ó no disponible');
        raise notice 'El turno es inexistente ó no esta disponible';
        return false;
    end if;

    turnosReservadosPaciente := count_reserved_appointments_for_patient(nro_historia_clinica);
    if turnosReservadosPaciente = 5 then
        insert into error(nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) values (default, null, null, dni_medique_reserva, nro_historia_clinica, 'reserva', now(), '?supera límite de reserva de turnos');
        raise notice 'El turno a reservar supera el límite de reserva de turnos';
        return false;
    end if;

    if nroObraSocialpaciente is null then
        select monto_consulta_privada into montoPaciente from medique where dni_medique = dni_medique_reserva;
    else
        select monto_paciente into montoPaciente from cobertura where dni_medique = dni_medique_reserva and nro_obra_social = nroObraSocialpaciente;
        select monto_obra_social into montoObraSocial from cobertura where dni_medique = dni_medique_reserva and nro_obra_social = nroObraSocialpaciente;
    end if;

    update turno set nro_paciente = nro_historia_clinica, nro_obra_social_consulta = nroObraSocialPaciente,
                     nro_afiliade_consulta = nroAfiliadePaciente, monto_paciente = montoPaciente, monto_obra_social = montoObraSocial, f_reserva = now(),
                     estado = 'reservado', fecha = turnoAReservar.fecha where nro_turno = turnoAReservar.nro_turno;

    return true;
end;
$$ language plpgsql;

create or replace function create_reserve_appointment_error(fecha_turno timestamp, nro_consultorio_error integer, nro_historia_clinica integer, dni_medique_error integer, msg text) returns void as $$
begin
    insert into error (nro_error, f_turno, nro_consultorio, dni_medique, nro_paciente, operacion, f_error, motivo) values (default, fecha_turno, nro_consultorio_error, dni_medique_error, nro_historia_clinica, 'reserva', now(), msg);
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

create or replace function calculate_patient_amount(nro_paciente_consulta integer, dni_medique_consulta integer) returns decimal(12,2) as $$
declare
    nroObraSocialpaciente integer;
    patientAmount decimal(12,2);
begin
    select nro_obra_social into nroObraSocialpaciente from paciente where paciente.nro_paciente = nro_paciente_consulta;
    if nroObraSocialpaciente is null then
        select monto_consulta_privada into patientAmount from medique where dni_medique = dni_medique_consulta;
    else
        select monto_paciente into patientAmount from cobertura where dni_medique = dni_medique_consulta and nro_obra_social = nroObraSocialpaciente;
    end if;

    return patientAmount;
end;
$$ language plpgsql;