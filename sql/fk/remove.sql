begin transaction;
alter table paciente drop constraint if exists nro_obra_social_fk;
alter table agenda drop constraint if exists dni_medique_fk;
alter table agenda drop constraint if exists nro_consultorio_fk;
alter table turno drop constraint if exists nro_consultorio_fk;
alter table turno drop constraint if exists dni_medique_fk;
alter table turno drop constraint if exists nro_paciente_fk;
alter table turno drop constraint if exists nro_obra_social_consulta_fk;
alter table reprogramacion drop constraint if exists nro_turno_fk;
alter table cobertura drop constraint if exists dni_medique_fk;
alter table cobertura drop constraint if exists nro_obra_social_fk;
alter table liquidacion_cabecera drop constraint if exists nro_obra_social_fk;
alter table liquidacion_detalle drop constraint if exists nro_liquidacion_fk;
alter table liquidacion_detalle drop constraint if exists dni_medique_fk;
commit;