# Live-Code-3-Phase-1

## RULES

1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Clone repo ini kemudian buatlah **branch dengan nama kalian**.
4. Kerjakan pada file JavaScript (\*.js) yang telah disediakan.
5. Waktu pengerjaan: **90 menit** untuk **2 soal**.
6. **Pada text editor hanya ada file yang terdapat pada repository ini**.
7. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
8. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).
9. Lakukan `git push origin <branch-name>` dan create pull request **hanya jika waktu Live Code telah usai (bukan ketika kalian sudah selesai
mengerjakan)**. Tuliskan nama lengkap kalian saat membuat pull request dan assign buddy.
10. **Penilaian berbasis logika dan hasil akhir**. Pastikan keduanya sudah benar.

## Notes

Live code ini memiliki bobot nilai sebagai berikut:

### KETENTUAN

Here are some dos and don'ts to consider when working on this livecode:

Do's:

- Clearly define the problem statement and requirements for the challenge.
- Use realistic and representative data for the challenge.
- Provide a database schema or ER diagram for the challenge.
- Specify the SQL dialect and version to be used for the challenge.
- Provide a clear submission format and deadline for the challenge.
- Evaluate the challenge solutions based on clear and transparent criteria.
- Provide feedback and follow-up for the candidates after evaluating their solutions.

Don'ts:

- Use overly complex or unrealistic scenarios for the challenge.
- Include ambiguous or poorly defined requirements for the challenge.
- Use inconsistent or illogical data for the challenge.
- Assume the candidates have access to proprietary or confidential information for the challenge.
- Ignore the importance of data security and privacy in the challenge.
- Use vague or subjective evaluation criteria for the challenge.
- Provide no or insufficient feedback to the candidates after evaluating their solutions.

SET 1

## NUMBER 1 LIVE CODE 3

## HealthTrack - Reinventing Hospital Management Systems

### Restrictions

- Don't use global variables unnecessarily. Instead, use local variables and pass them as parameters to functions as needed.
- Don't forget to handle errors returned by functions, using either an if err != nil statement or a panic statement where appropriate.
- Don't use fmt.Println excessively for debugging purposes. Instead, use Go's built-in testing framework or a separate debugging tool.
- Don't forget to use the correct type assertions when working with interfaces.
- Incorrect type assertions can result in runtime errors and unexpected behavior.
- Don't use unidiomatic or non-idiomatic Go code, such as using while loops instead of for loops, or using := instead of var for variable declarations when appropriate.

## Objectives

- Mampu menyelesaikan masalah yang diberikan.
- Memahami konsep `DDL` dan `DML`
- Memahami konsep `Query SQL`
- Memahami penggunaan logika kondisi
- Memahami penggunaan logika perulangan
- Memahami penggunaan konkurensi dengan goroutines, channel hingga wg
- Memahami penggunaan Golang Basic CLI

## Directions

### Title: HealthTrack - Reinventing Hospital Management Systems

In the sprawling metropolis of BioTech City, an ambitious healthcare professional named Dr. Jane Foster, decided to revolutionize the traditional hospital management system. Her vision, a state-of-the-art hospital named "HealthTrack," aimed to offer unparalleled patient care and service. But Dr. Foster realized that to operate HealthTrack seamlessly and enhance the patient experience, she needed a technologically advanced backend system.

Recognizing the importance, Dr. Foster hired the most sought-after backend developer in town - you, known for your expertise in electronic health record (EHR) systems and integrations.

### RELEASE 1 - Database Design

Your first mission is to design a robust and secure EHR database for `HealthTrack`. Your preliminary analysis identifies the critical entities of the database: `patients`, `doctors`, `appointments`, `treatments`, and `billing`.

Field Requirements:

- `Patients`: This table would store details with fields - `PatientId`, `Name`, `DateOfBirth`, and `MedicalHistory`.
- `Doctors`: This table would comprise details like - `DoctorId`, `Name`, `Specialty`, and `Availability`.
- `Appointments`: This table captures appointment data with fields - `AppointmentId`, `PatientId`, `DoctorId`, `Date`, and `Time`.
- `Treatments`: Detailed information about patient treatments - `TreatmentId`, `PatientId`, `DoctorId`, `Diagnosis`, and `Medication`.
- `Billing`: This table focuses on the financials with fields - `BillId`, `PatientId`, `TreatmentId`, and `Amount`.

### Your task, as the backend developer, Dr. Foster expects you to create an Entity Relationship Diagram (ERD) showing the interconnections between these entities. The relationships will likely include many-to-many and one-to-many connections.

### Please Notes

Following the ERD, your responsibility extends to transforming this design into a real-world database using a tool of your choice, perhaps XAMPP. You'd also create SQL commands for CRUD operations and form views for intuitive data access.

Hint **Release 1** :

- **Do Analyze Data** based on requirement on Field Requirements above
- **Create your ERD** based on your analyze on `point a`
- **Provide your SQL Query** to create db on localhost

### Expected Output

- Analyze Data

## sample

![image](https://github.com/H8-FTGO-P1/FTGO-P1-V2-LC3/blob/main/Release1A-AnalyzeERD.png)

- ERD
  
## sample

![image](https://github.com/H8-FTGO-P1/FTGO-P1-V2-LC3/blob/main/Release1B-ERD.png)

- SQL Query

## sample

![image](https://github.com/H8-FTGO-P1/FTGO-P1-V2-LC3/blob/main/Release1C-Query.png)

### RELEASE 2 - Patient Management Interface

Dr. Foster requires a Simple Command Line Interface (CLI) program using Golang. The CLI will allow staff to register and manage patient details, track appointments, and handle treatments.

- Patient Registration: Capture and store patients' details.
- Appointment Scheduling: Ability to assign a patient to a doctor based on the doctor's specialty and availability
- Treatment Details: Store diagnosis and medication details post an appointment.
- Billing Integration: Automatically calculate the bill amount based on the treatment provided.

## Your task is to design and implement this system based on these requirements. The CLI application should interact with a SQL database (as established in release 1) to handle data storage and retrieval.

Here the sample data that Dr. Jane Foster share to you :

```sql

`Patients`

`INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('John Doe', '1990-05-15', 'No known allergies. Had chickenpox in 1995.');`
`INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('Jane Smith', '1985-08-12', 'Allergic to penicillin. Broke left arm in 2005.');`
`INSERT INTO Patients (Name, DateOfBirth, MedicalHistory) VALUES ('Alice Johnson', '1999-12-03', 'Has asthma. No other known conditions.');`

`Doctors`

`INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Peter Parker', 'Cardiologist', 'Monday-Wednesday');`
`INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Bruce Wayne', 'Orthopedic', 'Thursday-Saturday');`
`INSERT INTO Doctors (Name, Specialty, Availability) VALUES ('Dr. Clark Kent', 'General Practitioner', 'Monday-Friday');`

`Appointments`

`INSERT INTO Appointments (PatientId, DoctorId, Date, Time) VALUES (1, 1, '2023-10-10', '10:00:00');`
`INSERT INTO Appointments (PatientId, DoctorId, Date, Time) VALUES (2, 2, '2023-10-11', '11:00:00');`
`INSERT INTO Appointments (PatientId, DoctorId, Date, Time) VALUES (3, 3, '2023-10-12', '14:00:00');`

`Treatments`

`INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (1, 1, 'Minor chest pain due to stress.', 'Prescribed Ibuprofen.');`
`INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (2, 2, 'Mild joint pain.', 'Prescribed Paracetamol.');`
`INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (3, 3, 'Flu', 'Prescribed general flu medication.');`

`Billing`

`INSERT INTO Billing (PatientId, TreatmentId, Amount) VALUES (1, 1, 150.00);`
`INSERT INTO Billing (PatientId, TreatmentId, Amount) VALUES (2, 2, 100.50);`
`INSERT INTO Billing (PatientId, TreatmentId, Amount) VALUES (3, 3, 80.25);`

```

Hint **Release 2**:

Review patient management flow and requirements.

Expected Output :

![image](https://github.com/H8-FTGO-P1/FTGO-P1-V2-LC3/blob/main/Release2.gif)

Ensure to test the application thoroughly to make sure it handles all the scenarios outlined in the user journey.

### RELEASE 3 - Admin Report

Your challenge here is to offer Dr. Jane Foster a comprehensive dashboard directly through the CLI:

- a. Patient Reports: Display all patient details, including historical medical records.
- b. Doctor Activity: Show the number of appointments and treatments handled by each doctor.
- c. Billing Overview: Aggregate data to display total revenue, outstanding bills, and daily income.
- d. Appointment Stats: Reflect the number of appointments scheduled, completed, or missed.

To ensure a streamlined approach, please create these functionalities in a separate Python file named dashboard.py. This file will contain the functions for all reports, and they should be callable from the primary application.

Ensure robust testing to guarantee accuracy and security of patient data.

Please make sure to test the reports functionality thoroughly to ensure accurate and reliable data retrieval.

Hint **Release 3** :

- **Maximize** SQL Query and display on *CLI Program*

Expected Output :

![image](https://github.com/H8-FTGO-P1/FTGO-P1-V2-LC3/blob/main/Release3.gif)
