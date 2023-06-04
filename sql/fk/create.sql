-- paciente
alter table paciente add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

-- agenda
alter table agenda add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table agenda add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);

-- turno
alter table turno add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table turno add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table turno add constraint nro_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);

-- reprogramacion
alter table reprogramacion add constraint nro_turno_fk foreign key (nro_turno) references turno (nro_turno);

-- error
alter table error add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table error add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table error add constraint nro_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);

-- cobertura
alter table cobertura add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table cobertura add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

-- liquidacion_cabecera
alter table liquidacion_cabecera add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

-- liquidacion_detalle
alter table liquidacion_detalle add constraint nro_liquidacion_fk foreign key (nro_liquidacion) references liquidacion_cabecera (nro_liquidacion);

-- Queda pendiente agregar las fks restantes para liquidacion_detalle. A que tablas referenciamos?