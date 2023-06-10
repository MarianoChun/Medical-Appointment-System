-- paciente
alter table if exists paciente add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

-- agenda
alter table if exists agenda add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table if exists agenda add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);

-- turno
alter table if exists turno add constraint nro_consultorio_fk foreign key (nro_consultorio) references consultorio (nro_consultorio);
alter table if exists turno add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table if exists turno add constraint nro_paciente_fk foreign key (nro_paciente) references paciente (nro_paciente);

-- reprogramacion
alter table if exists reprogramacion add constraint nro_turno_fk foreign key (nro_turno) references turno (nro_turno);

-- cobertura
alter table if exists cobertura add constraint dni_medique_fk foreign key (dni_medique) references medique (dni_medique);
alter table if exists cobertura add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

-- liquidacion_cabecera
alter table if exists liquidacion_cabecera add constraint nro_obra_social_fk foreign key (nro_obra_social) references obra_social (nro_obra_social);

-- liquidacion_detalle
alter table if exists liquidacion_detalle add constraint nro_liquidacion_fk foreign key (nro_liquidacion) references liquidacion_cabecera (nro_liquidacion);

-- Queda pendiente agregar las fks restantes para liquidacion_detalle. A que tablas referenciamos?