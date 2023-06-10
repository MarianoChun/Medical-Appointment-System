create or replace function generate_appointments_in_month(year int, month int) returns boolean as $$
declare
    start_of_month timestamp;
    end_of_month timestamp;
    current_day date;
    medic record;
    any_appointment_in_range boolean;
    medic_agenda record;
    agenda_time timestamp;
begin
    start_of_month := make_timestamp(year, month, 1, 0, 0, 0);
    end_of_month := start_of_month + interval '1 month - 1 day';

    select exists(select 1 from turno  where fecha between start_of_month and end_of_month) into any_appointment_in_range;
    if any_appointment_in_range then
        return false;
    else
            for current_day in select generate_series(start_of_month, end_of_month, interval '1 day') loop
                for medic in select * from medique loop
                    select  * from agenda where dni_medique = medic.dni_medique and dia = date_part('dow', current_day) into medic_agenda;
                    for agenda_time in select generate_series(current_day + medic_agenda.hora_desde, current_day + medic_agenda.hora_hasta, medic_agenda.duracion_turno) loop
                        insert into turno (fecha, nro_consultorio, dni_medique, estado)
                        values (agenda_time, medic_agenda.nro_consultorio, medic.dni_medique, 'disponible');
                    end loop;
                end loop;
            end loop;   
        return true;
    end if;
    
end;
$$ language plpgsql;
