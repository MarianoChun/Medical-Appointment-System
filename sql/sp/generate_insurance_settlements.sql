create or replace function generate_insurance_settlements() returns void as $$
declare
    obraSocial record;
    turnoRecord record;
    fechaMesLiquidacionHasta date := now();
    fechaMesLiquidacionDesde date := fechaMesLiquidacionHasta - interval '1 month';
    montoObraSocial decimal (15, 2);
    nroLiquidacionActual integer;
    dniPacienteActual integer;
    nombrePacienteActual text;
    apellidoPacienteActual text;
    nombreMediqueActual text;
    apellidoMediqueActual text;
    especialidadMediqueActual text;
begin
    for obraSocial in select * from obra_social loop
        insert into liquidacion_cabecera (nro_liquidacion, nro_obra_social, desde, hasta, total) values (default, obraSocial.nro_obra_social, fechaMesLiquidacionDesde, fechaMesLiquidacionHasta, montoObraSocial);

        montoObraSocial := 0.00;
        select into nroLiquidacionActual nro_liquidacion from liquidacion_cabecera where nro_obra_social = obraSocial.nro_obra_social;

        for turnoRecord in select * from turno where turno.nro_obra_social_consulta = obraSocial.nro_obra_social and  turno.estado = 'atendido' and turno.fecha between fechaMesLiquidacionDesde and fechaMesLiquidacionHasta loop
                select dni_paciente, nombre, apellido into dniPacienteActual, nombrePacienteActual, apellidoPacienteActual from paciente where paciente.nro_paciente = turnoRecord.nro_paciente;
                select nombre, apellido, especialidad into nombreMediqueActual, apellidoMediqueActual, especialidadMediqueActual from medique where dni_medique = turnoRecord.dni_medique;
                insert into liquidacion_detalle (nro_linea, nro_liquidacion, f_atencion, nro_afiliade, dni_paciente, nombre_paciente, apellido_paciente, dni_medique, nombre_medique, apellido_medique, especialidad, monto) values (default, nroLiquidacionActual, turnoRecord.fecha, turnoRecord.nro_afiliade_consulta, dniPacienteActual, nombrePacienteActual, apellidoPacienteActual, turnoRecord.dni_medique, nombreMediqueActual, apellidoMediqueActual, especialidadMediqueActual, turnoRecord.monto_obra_social);

                montoObraSocial := montoObraSocial + turnoRecord.monto_obra_social;
                update turno set estado = 'liquidado' where nro_turno = turnoRecord.nro_turno;
        end loop;

        update liquidacion_cabecera set total = montoObraSocial where liquidacion_cabecera.nro_obra_social =  obraSocial.nro_obra_social;
    end loop;
end;
$$ language plpgsql;
