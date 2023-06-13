begin transaction;
-- Turnos reservados correctamente
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (9, 26387951, now(), '12:30:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (14, 26387951, now(), '13:00:00');
-- Turnos reservados incorrectamente, generan error

-- nro de paciente invalido (no existe)
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4444, 12341234, now(), '12:00:00');

-- dni medique invalido (no existe)
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (1, 4444, now(), '12:00:00');

-- el medique no atiende la obra social del paciente
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (1, 12341234, now(), '12:00:00');

-- turno no disponible
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '09:30:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '09:30:00');

-- turno inexistente
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now() - interval '1 year', '12:00:00');

-- el paciente supero el limite de turnos que es maximo 5
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '10:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '10:30:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '11:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '11:30:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 12341234, now(), '12:00:00');
commit;