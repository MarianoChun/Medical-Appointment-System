create or replace function cancel_appointment(dni integer, date_from date, date_to date) returns int as $$
declare
    canceled_appointment_count int;
    result record;
begin
    canceled_appointment_count := 0;

    for result in select t.nro_turno as nro_turno, t.nro_paciente, p.nombre as nombre_paciente, p.apellido as apellido_paciente, p.telefono as telefono_paciente, p.email as email_paciente, m.nombre as nombre_medique, m.apellido as apellido_medique from turno t, paciente p, medique m where p.nro_paciente = t.nro_paciente and m.dni_medique = t.dni_medique and t.dni_medique = dni and t.estado in ('reservado') and date(t.fecha) between date_from and date_to
                  union
                  select t.nro_turno as nro_turno, t.nro_paciente, null as nombre_paciente, null as apellido_paciente, null as telefono_paciente, null as email_paciente, m.nombre as nombre_medique, m.apellido as apellido_medique from turno t, medique m where m.dni_medique = t.dni_medique and t.dni_medique = dni and t.estado in ('disponible') and date(t.fecha) between date_from and date_to loop
        update turno set estado = 'cancelado' where nro_turno = result.nro_turno;

        if result.nro_paciente is not null then
            insert into reprogramacion (nro_turno, nombre_paciente, apellido_paciente, telefono_paciente, email_paciente, nombre_medique, apellido_medique, estado) values (result.nro_turno, result.nombre_paciente, result.apellido_paciente, result.telefono_paciente, result.email_paciente, result.nombre_medique, result.apellido_medique, 'pendiente');
        end if;

        canceled_appointment_count := canceled_appointment_count + 1;
    end loop;

    return canceled_appointment_count;
end;
$$ language plpgsql;