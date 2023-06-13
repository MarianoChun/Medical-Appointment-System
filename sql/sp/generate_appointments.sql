create or replace function generate_appointments_in_month(year int, month int) returns boolean as $$
declare
    start_of_month timestamp;
    end_of_month timestamp;
    current_day date;
    item record;
    date_time timestamp;
    is_exists_appointment_in_same_range int;
    count_of_agenda_medics int;
    count_of_appointment_medics int;
begin
    start_of_month := make_timestamp(year, month, 1, 0, 0, 0);
    end_of_month := start_of_month + interval '1 month - 1 day';

    for current_day in select generate_series(start_of_month, end_of_month, interval '1 day') loop
        for item in select nro_consultorio, dni_medique, hora_desde, hora_hasta, duracion_turno from agenda a where a.dia = date_part('dow', current_day) loop
            for date_time in select generate_series(current_day + item.hora_desde, current_day + item.hora_hasta, item.duracion_turno) loop

                select count(1) from turno t where t.fecha = date_time and t.dni_medique = item.dni_medique into is_exists_appointment_in_same_range;
                raise notice 'is_exists_appointment_in_same_range: %', is_exists_appointment_in_same_range;
                if is_exists_appointment_in_same_range <> 0 then
                    raise notice 'is_exists_appointment_in_same_range rollback';
                    ROLLBACK;
                end if;

                insert into turno (fecha, nro_consultorio, dni_medique, estado) values (date_time, item.nro_consultorio, item.dni_medique, 'disponible');
            end loop;
        end loop;
    end loop;

    select count(distinct dni_medique) from agenda a into count_of_agenda_medics;
    select count(distinct dni_medique) from turno t where t.fecha between start_of_month and end_of_month into count_of_appointment_medics;
    raise notice 'count_of_agenda_medics: %', count_of_agenda_medics;
    raise notice 'count_of_appointment_medics: %', count_of_appointment_medics;

    if count_of_agenda_medics <> count_of_appointment_medics then
        raise notice 'count_of_agenda_medics <> count_of_appointment_medics rollback';
        ROLLBACK;
    end if;

    raise notice 'count_of_agenda_medics == count_of_appointment_medics commit';

    return true;
exception
    when others then
        raise notice 'Error in generate_appointments_in_month function. %', SQLERRM;
        return false;
end;
$$ language plpgsql;
