create or replace function generate_insurance_settlements() returns void as $$
declare
    obraSocial record;
    turnoRecord record;
    result record;
    fechaMesLiquidacionHasta date := current_date;
    fechaMesLiquidacionDesde date := fechaMesLiquidacionHasta - interval '1 month';
    montoObraSocial decimal (15, 2);
    nroLiquidacionActual integer;
begin
    for obraSocial in select * from obra_social loop
        montoObraSocial := 0.00;

        insert into liquidacion_cabecera (nro_liquidacion, nro_obra_social, desde, hasta, total) values (default, obraSocial.nro_obra_social, fechaMesLiquidacionDesde, fechaMesLiquidacionHasta, montoObraSocial);

        select into nroLiquidacionActual nro_liquidacion from liquidacion_cabecera where nro_obra_social = obraSocial.nro_obra_social;

        for turnoRecord in select * from turno where turno.nro_obra_social_consulta = obraSocial.nro_obra_social and turno.estado = 'atendido' and date(turno.fecha) between fechaMesLiquidacionDesde and fechaMesLiquidacionHasta loop
                select
                    p.dni_paciente,
                    p.nombre as nombre_paciente,
                    p.apellido as apellido_paciente,
                    m.nombre as nombre_medique,
                    m.apellido as apellido_medique,
                    m.especialidad as especialidad_medique
                    into result from paciente p, medique m
                    where turnoRecord.nro_paciente = p.nro_paciente and turnoRecord.dni_medique = m.dni_medique;

                insert into liquidacion_detalle (nro_linea, nro_liquidacion, f_atencion, nro_afiliade, dni_paciente, nombre_paciente, apellido_paciente, dni_medique, nombre_medique, apellido_medique, especialidad, monto) values (default, nroLiquidacionActual, turnoRecord.fecha, turnoRecord.nro_afiliade_consulta, result.dni_paciente, result.nombre_paciente, result.apellido_paciente, turnoRecord.dni_medique, result.nombre_medique, result.apellido_medique, result.especialidad_medique, turnoRecord.monto_obra_social);

                montoObraSocial := montoObraSocial + turnoRecord.monto_obra_social;
                update turno set estado = 'liquidado' where nro_turno = turnoRecord.nro_turno;
        end loop;

        update liquidacion_cabecera set total = montoObraSocial where liquidacion_cabecera.nro_obra_social =  obraSocial.nro_obra_social;
    end loop;
end;
$$ language plpgsql;
