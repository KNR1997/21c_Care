# AI-Powered Unified Clinic Notes & Billing System

## 📌 Project Overview

The **AI-Powered Unified Clinic Notes & Billing System** is a full-stack application designed to simplify clinic workflows by allowing doctors to enter prescriptions, lab tests, and clinical notes in a single unified interface. The system uses AI to automatically classify unstructured medical text into structured data and generates billing and printable PDF reports.

This project demonstrates practical integration of **AI/NLP with a scalable software architecture**, focusing on usability, structured data storage, billing accuracy, and report generation.

---

## 🚀 Features

* Unified clinic input screen for doctors
* AI-powered medical text classification
* Automatic extraction of:

  * Drugs & dosages
  * Lab tests
  * Clinical notes
* Structured database storage
* Billing calculation based on drugs and lab tests
* Printable PDF medical report
* Clean REST API architecture
* Modular backend design
* Responsive NextJS 13 frontend

---

## 🛠️ Tech Stack

### Backend

* Golang
* Echo Framework
* PostgreSQL
* AI API (LLM-based text classification)
* gofpdf (PDF generation)

### Frontend

* NextJS 13
* TypeScript
* React Query
* Axios

### Database

* PostgreSQL

### Tools

* Docker (optional)
* Postman
* Draw.io (Architecture diagrams)

---

## 🏗️ System Architecture

The system follows a layered architecture with clear separation of concerns.

```
React Frontend
      ↓
Golang API (Echo)
      ↓
Service Layer
      ↓
AI Service
      ↓
Repository Layer
      ↓
PostgreSQL Database
```

### External Services

* AI API for NLP classification
* PDF generator for report printing

---

## 🧠 AI Integration

AI is used to convert unstructured medical text into structured clinical data.

### Workflow

1. Doctor enters medical text
2. Backend sends text to AI service
3. AI extracts:

   * Drugs
   * Lab tests
   * Clinical notes
4. Data is validated
5. Structured data stored in PostgreSQL
6. Billing calculated
7. PDF report generated

### Example

**Input**

```
Patient has fever for 3 days. Prescribe Paracetamol 500mg twice daily. Order CBC test.
```

**AI Output**

```
Drugs:
Paracetamol 500mg

Lab Tests:
CBC

Notes:
Patient has fever for 3 days
```

---

## 🗄️ Database Design

The database is designed using normalized relational tables.

### Main Tables

* patients
* visits
* prescribed_drugs
* lab_tests
* clinical_notes
* drug_catalog
* lab_test_catalog
* billing

### Design Approach

* Catalog tables store official prices
* Visit tables store structured AI output
* Snapshot pricing ensures billing consistency
* Raw AI input is stored for auditing

---

## 💰 Billing Logic

Billing is calculated using:

```
Grand Total = Consultation Fee + Drug Prices + Lab Test Prices
```

### Pricing Strategy

* Drug prices fetched from `drug_catalog`
* Lab test prices fetched from `lab_test_catalog`
* Prices stored in visit tables
* Billing stored in `billing` table

### Benefits

* Consistent billing
* Historical price tracking
* Easy invoice generation

---

## 🖨️ PDF Report Generation

The system generates a printable PDF report for each visit.

### Report Includes

* Clinic information
* Patient details
* Prescription
* Lab tests
* Clinical notes
* Billing summary

### API

```
GET /visits/{id}/report
```

### Output

PDF file ready for printing or download.

---

## 🌐 API Endpoints

### Visits

```
POST /visits
Create new visit and process AI classification

GET /visits
Get all visits

GET /visits/{id}
Get visit details
```

### Billing

```
GET /billing/{visit_id}
Get billing details
```

### Report

```
GET /visits/{id}/report
Download PDF invoice
```

---

## 🖥️ Frontend Workflow

1. Doctor enters medical text
2. AI processes text
3. Structured data displayed
4. Billing calculated
5. PDF invoice generated
6. User downloads invoice

---

## ⚙️ Setup Instructions

### Prerequisites

* Go 1.25+
* Node.js 18+
* PostgreSQL
* AI API Key

---

## 🔧 Backend Setup

### 1. Clone Repository

```
git clone https://github.com/KNR1997/21c_Care.git
cd 21c_Care
```

### 2. Configure Environment

Create `.env`

```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=clinic
AI_API_KEY=your_api_key
```

### 3. Run Database

```
psql -U postgres
CREATE DATABASE clinic;
```

### 4. Run Migrations

```
psql -U postgres -d clinic -f database/schema.sql
```

### 5. Start Server

```
air
```

Server runs on:

```
http://localhost:7788
```

---

## 🎨 Frontend Setup

### 1. Go to frontend

```
cd frontend
```

### 2. Install dependencies

```
yarn
```

### 3. Configure environment

```
copy `.env.template` to `.env`
```

### 4. Run

```
yarn dev
```

App runs on:

```
http://localhost:3002
```

---

## 🧪 Example Workflow

### Step 1

Doctor enters:

```
Patient has headache. Prescribe Ibuprofen 200mg. Order MRI.
```

### Step 2

AI extracts:

* Ibuprofen
* MRI
* headache

### Step 3

Database stores structured data

### Step 4

Billing generated

### Step 5

PDF report downloaded

---

## 📊 Project Structure

```
backend/
    internal/
        handlers/
        services/
        repositories/
        ai/
        report/
    database/
    main.go

frontend/
    src/
        components/
        pages/
        data/
        client/
```

---

## 🧠 Assumptions

* AI classification may not be 100% accurate
* Drug and lab test catalog contains common items
* Consultation fee is fixed
* Single clinic environment
* Internet connection required for AI API

---

## 👨‍💻 Author

Software Engineer Candidate

AI-Powered Clinic System Technical Assessment
