CREATE DATABASE IF NOT EXISTS healthtrack;
USE healthtrack;

CREATE TABLE Patients (
  PatientId INT AUTO_INCREMENT PRIMARY KEY,
  Name VARCHAR(100) NOT NULL,
  DateOfBirth DATE NOT NULL,
  MedicalHistory TEXT NOT NULL
);

CREATE TABLE Doctors (
  DoctorId INT AUTO_INCREMENT PRIMARY KEY,
  Name VARCHAR(100) NOT NULL,
  Specialty VARCHAR(100) NOT NULL,
  Availability VARCHAR(100) NOT NULL
);

CREATE TABLE Appointments (
  AppointmentId INT AUTO_INCREMENT PRIMARY KEY,
  PatientId INT NOT NULL,
  DoctorId INT NOT NULL,
  Date DATE NOT NULL,
  Time TIME NOT NULL,
  FOREIGN KEY (PatientId) REFERENCES Patients(PatientId) ON DELETE CASCADE,
  FOREIGN KEY (DoctorId) REFERENCES Doctors(DoctorId) ON DELETE CASCADE
)

CREATE TABLE Treatments (
  TreatmentId INT AUTO_INCREMENT PRIMARY KEY,
  PatientId INT NOT NULL,
  DoctorId INT NOT NULL,
  Diagnosis TEXT NOT NULL,
  Medication TEXT NOT NULL,
  FOREIGN KEY (PatientId) REFERENCES Patients(PatientId) ON DELETE CASCADE,
  FOREIGN KEY (DoctorId) REFERENCES Doctors(DoctorId) ON DELETE CASCADE
)

CREATE TABLE Billing (
  BillId INT AUTO_INCREMENT PRIMARY KEY,
  PatientId INT NOT NULL,
  TreatmentId INT NOT NULL,
  Amount DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (PatientId) REFERENCES Patients(PatientId) ON DELETE CASCADE,
  FOREIGN KEY (TreatmentId) REFERENCES Treatments(TreatmentId) ON DELETE CASCADE
)

-- Data Dummy
-- Patients
INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('John Doe', '1990-05-15', 'No known allergies. Had chickenpox in 1995.');
INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('Jane Smith', '1985-08-12', 'Allergic to penicillin. Broke left arm in 2005.');
INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('Alice Johnson', '1999-12-03', 'Has asthma. No other known conditions.');

-- Doctors
INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Peter Parker', 'Cardiologist', 'Monday-Wednesday');
INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Bruce Wayne', 'Orthopedic', 'Thursday-Saturday');
INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Clark Kent', 'General Practitioner', 'Monday-Friday');

-- Appointments
INSERT INTO Appointments (PatientId, DoctorId, Date, Time) VALUES (1, 1, '2023-10-10', '10:00:00');
INSERT INTO Appointments (PatientId, DoctorId, Date, Time) VALUES (2, 2, '2023-10-11', '11:00:00');
INSERT INTO Appointments (PatientId, DoctorId, Date, Time) VALUES (3, 3, '2023-10-12', '14:00:00');

-- Treatments
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (1, 1, 'Minor chest pain due to stress.', 'Prescribed Ibuprofen.');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (2, 2, 'Mild joint pain.', 'Prescribed Paracetamol.');
INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (3, 3, 'Flu', 'Prescribed general flu medication.');

-- Billing
INSERT INTO Billing (PatientId, TreatmentId, Amount) VALUES (1, 1, 150.00);
INSERT INTO Billing (PatientId, TreatmentId, Amount) VALUES (2, 2, 100.50);
INSERT INTO Billing (PatientId, TreatmentId, Amount) VALUES (3, 3, 80.25);
