-- Turnos reservados correctamente
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (1, 38692417, '2023-01-04', '11:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (1, 38692417, '2023-01-25', '11:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (1, 38692417, '2023-01-25', '11:40:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (1, 38692417, '2023-01-18', '11:00:00');

insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (2, 24587963, '2023-01-03', '14:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (2, 24587963, '2023-01-03', '14:20:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (2, 24587963, '2023-01-10', '14:00:00');

insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (3, 28319476, '2023-01-03', '17:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (3, 28319476, '2023-01-03', '17:30:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (3, 28319476, '2023-01-10', '17:00:00');

insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 37894125, '2023-01-02', '13:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 37894125, '2023-01-02', '13:20:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (4, 37894125, '2023-01-09', '13:00:00');

insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (5, 33964752, '2023-01-11', '12:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (5, 33964752, '2023-01-11', '12:40:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (5, 33964752, '2023-01-04', '12:00:00');

insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (21, 41244925, '2023-01-03', '12:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (22, 41244925, '2023-01-03', '13:00:00');
insert into solicitud_reservas (nro_paciente, dni_medique, fecha, hora) values (23, 41244925, now(), '12:00:00');
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
