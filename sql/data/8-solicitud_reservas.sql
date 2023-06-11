-- Turnos reservados correctamente
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (1, 1, 38692417, '2023-01-04', '11:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (2, 1, 38692417, '2023-01-25', '11:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (3, 1, 38692417, '2023-01-25', '11:40:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (4, 1, 38692417, '2023-01-18', '11:00:00');

insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (5, 2, 24587963, '2023-01-03', '14:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (6, 2, 24587963, '2023-01-03', '14:20:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (7, 2, 24587963, '2023-01-10', '14:00:00');

insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (8, 3, 28319476, '2023-01-03', '17:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (9, 3, 28319476, '2023-01-03', '17:30:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (10, 3, 28319476, '2023-01-10', '17:00:00');

insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (11, 4, 37894125, '2023-01-02', '13:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (12, 4, 37894125, '2023-01-02', '13:20:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (13, 4, 37894125, '2023-01-09', '13:00:00');

insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (14, 5, 33964752, '2023-01-11', '12:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (15, 5, 33964752, '2023-01-11', '12:40:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (16, 5, 33964752, '2023-01-04', '12:00:00');

insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (17, 21, 41244925, '2023-01-03', '12:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (18, 22, 41244925, '2023-01-03', '13:00:00');
insert into solicitud_reservas (nro_orden, nro_paciente, dni_medique, fecha, hora) values (19, 23, 41244925, '2023-01-10', '12:00:00');
-- Turnos reservados incorrectamente, generan error

-- nro de paciente invalido (no existe)

-- dni medique invalido (no existe)

-- el medique no atiende la obra social del paciente

-- turno no disponible

-- turno inexistente

-- el paciente supero el limite de turnos que es maximo 5
