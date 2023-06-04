-- paciente
alter table paciente drop constraint nro_obra_social_fk;

-- agenda
alter table agenda drop constraint dni_medique_fk;
alter table agenda drop constraint nro_consultorio_fk;

-- turno
alter table turno drop constraint nro_consultorio_fk;
alter table turno drop constraint dni_medique_fk;
alter table turno drop constraint nro_paciente_fk;

-- reprogramacion
alter table reprogramacion drop constraint nro_turno_fk;

-- error
alter table error drop constraint nro_consultorio_fk;
alter table error drop constraint dni_medique_fk;
alter table error drop constraint nro_paciente_fk;

-- cobertura
alter table cobertura drop constraint dni_medique_fk;
alter table cobertura drop constraint nro_obra_social_fk;

-- liquidacion_cabecera
alter table liquidacion_cabecera drop constraint nro_obra_social_fk;

-- liquidacion_detalle
alter table liquidacion_detalle drop constraint nro_liquidacion_fk;

-- Queda pendiente agregar las fks restantes para liquidacion_detalle. A que tablas referenciamos?