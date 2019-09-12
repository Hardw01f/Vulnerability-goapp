create database vulnapp;
create table vulnapp.user (id int not null auto_increment primary key, name varchar(255) not null,mail varchar(255),age int not null, created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp);
insert into vulnapp.user (name,mail,age) values ("Amuro Ray","RX-78-2@E.F.S.F",15),("Char Aznable","MS-06-S@Zeon",20),("Kamille Bidan","MSZ-006@A.E.U.G",17),("Judau Ashta","MSZ-010@A.E.U.G",14),("Banagher Links","RX-0@LONDO.BELL",16);

