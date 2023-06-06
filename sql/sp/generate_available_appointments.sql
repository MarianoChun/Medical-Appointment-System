create or replace function generate_appointments_in_month(year int, month int) returns boolean as $$
declare
    start_of_month timestamp;
    end_of_month timestamp;
    d date;
    medic record;
    any_appointment_in_range boolean;
begin
    start_of_month := make_timestamp(year, month, 1, 0, 0, 0);
    end_of_month := start_of_month + interval '1 month - 1 day';


    select fecha between start_of_month and end_of_month from turno into any_appointment_in_range;
    if any_appointment_in_range then
        return false;
    else
        for d in select generate_series(start_of_month, end_of_month, interval '1 day') loop
            for medic in select * from medique loop
                perform generate_available_slots_in_day(medic.dni_medique, d);
            end loop;
        end loop;

        return true;
    end if;
    
end;
$$ language plpgsql;

create or replace function generate_available_slots_in_day(medic_document_number int, fecha date) returns void as $$
declare
    agenda record;
    tiempo agenda.hora_desde%type;
    appointment_number int; 
    appointment_datetime timestamp;
    dow int; 
begin

    dow := date_part('dow', fecha);
    select * into agenda from agenda where dni_medique = medic_document_number and dia = dow;


    tiempo := agenda.hora_desde;

    while tiempo < agenda.hora_hasta  loop
        select count(*) into appointment_number from turno;
        appointment_datetime := fecha + tiempo;
        insert into turno (nro_turno, fecha, nro_consultorio, dni_medique, estado)
        values (appointment_number + 1, appointment_datetime, agenda.nro_consultorio, medic_document_number, 'disponible');
        tiempo := tiempo + agenda.duracion_turno;
    end loop;
end;
$$ language plpgsql;
