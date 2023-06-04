create function generate_available_slots_in_day(medic_document_number int, day int) returns void as $$
declare
 tiempo agenda.hora_desde%type;
 start_time agenda.hora_desde%type;
 end_time agenda.hora_hasta%type;
 duration agenda.duracion_turno%type;
begin
    select hora_desde into start_time from agenda as a where a.dni_medique = medic_document_number and a.dia = day;
    select duracion_turno into duration from agenda as a where a.dni_medique = medic_document_number and a.dia = day;
    select hora_hasta into end_time from agenda as a where a.dni_medique = medic_document_number and a.dia = day;

    tiempo := start_time;

    while tiempo < end_time  loop
        raise notice 'Turno disponible a las %', tiempo;
        tiempo := tiempo + duration;
    end loop;
end;
$$ language plpgsql;