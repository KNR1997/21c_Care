package report

import (
	"bytes"
	"context"
	"fmt"
	"go-echo-starter/internal/domain"

	"github.com/jung-kurt/gofpdf"
)

type visitRepository interface {
	GetVisitDetails(ctx context.Context, visitID uint) (domain.VisitDetails, error)
}

type Service struct {
	repo visitRepository
}

func NewService(repo visitRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GenerateVisitPDF(ctx context.Context, visitID uint) ([]byte, error) {

	data, err := s.repo.GetVisitDetails(ctx, visitID)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "ABC Health Clinic")

	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Patient: "+data.PatientName)

	pdf.Ln(10)

	// Drugs
	pdf.Cell(40, 10, "Prescription")
	pdf.Ln(8)

	for _, d := range data.Drugs {
		pdf.Cell(40, 8, d.DrugName+" "+d.Dosage)
		pdf.Ln(6)
	}

	pdf.Ln(10)

	// Lab tests
	pdf.Cell(40, 10, "Lab Tests")
	pdf.Ln(8)

	for _, t := range data.LabTests {
		pdf.Cell(40, 8, t.TestName)
		pdf.Ln(6)
	}

	pdf.Ln(10)

	// Notes
	pdf.Cell(40, 10, "Clinical Notes")
	pdf.Ln(8)

	for _, n := range data.Notes {
		pdf.MultiCell(0, 6, n.Note, "", "", false)
	}

	pdf.Ln(10)

	// Billing
	pdf.Cell(40, 10, "Billing")
	pdf.Ln(8)

	pdf.Cell(40, 8, "Drugs Total")
	pdf.Cell(40, 8, fmt.Sprintf("%.2f", data.Bill.DrugsTotal))
	pdf.Ln(6)

	pdf.Cell(40, 8, "Lab Total")
	pdf.Cell(40, 8, fmt.Sprintf("%.2f", data.Bill.LabTestsTotal))
	pdf.Ln(6)

	pdf.Cell(40, 8, "Grand Total")
	pdf.Cell(40, 8, fmt.Sprintf("%.2f", data.Bill.GrandTotal))

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
