package visit

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
	"go-echo-starter/internal/services/ai"
	"go-echo-starter/internal/services/clinicalnote"
	"go-echo-starter/internal/services/labtest"
	"go-echo-starter/internal/services/prescribeddrug"
)

type Repository interface {
	Create(ctx context.Context, visit *models.Visit) error
	List(ctx context.Context) ([]models.Visit, error)
	ListPaginated(pagination repositories.Pagination[models.Visit]) (*repositories.Pagination[models.Visit], error)
	Get(ctx context.Context, id uint) (models.Visit, error)
}

type Service struct {
	repo               Repository
	ai                 *ai.Service
	labTests           *labtest.Service
	notes              *clinicalnote.Service
	drugCatalogRepo    *repositories.DrugCatalogRepository
	labTestCatalogRepo *repositories.LabTestCatalogRepository
	drugService        *prescribeddrug.Service
}

func NewService(
	repo Repository,
	ai *ai.Service,
	labTests *labtest.Service,
	notes *clinicalnote.Service,
	drugCatalogRepo *repositories.DrugCatalogRepository,
	labTestCatalogRepo *repositories.LabTestCatalogRepository,
	drugService *prescribeddrug.Service) *Service {
	return &Service{
		repo:               repo,
		ai:                 ai,
		labTests:           labTests,
		notes:              notes,
		drugCatalogRepo:    drugCatalogRepo,
		labTestCatalogRepo: labTestCatalogRepo,
		drugService:        drugService,
	}
}

func (s *Service) Preview(ctx context.Context, rawInput string) (*domain.AIResponse, error) {
	return s.ai.ParseMedicalText(ctx, rawInput)
}

func (s *Service) Create(ctx context.Context, visit *models.Visit, aiResult *domain.AIResponse) error {

	if err := s.repo.Create(ctx, visit); err != nil {
		return fmt.Errorf("create visit in repository: %w", err)
	}

	for _, d := range aiResult.Drugs {

		catalogDrug, err := s.drugCatalogRepo.FindBestMatch(ctx, d.Name)

		price := 0.0
		drugName := d.Name

		if err == nil {
			price = catalogDrug.DefaultPrice
			drugName = catalogDrug.Name
		}

		drug := models.PrescribedDrug{
			VisitID:  visit.ID,
			DrugName: drugName,
			Price:    price,
		}

		if err := s.drugService.Create(ctx, &drug); err != nil {
			return err
		}
	}

	for _, name := range aiResult.LabTests {

		labTestCatalog, err := s.labTestCatalogRepo.FindBestMatch(ctx, name)

		price := 0.0
		labTestName := name

		if err == nil {
			price = labTestCatalog.DefaultPrice
			labTestName = labTestCatalog.Name
		}

		labTest := models.LabTest{
			VisitID:  visit.ID,
			TestName: labTestName,
			Price:    price,
		}

		if err := s.labTests.Create(ctx, &labTest); err != nil {
			return err
		}

	}

	note := models.ClinicalNote{
		VisitID: visit.ID,
		Note:    aiResult.Notes,
	}
	s.notes.Create(ctx, &note)

	return nil
}

// func (s *Service) Create(ctx context.Context, visit *models.Visit) error {
// 	aiResult, err := s.ai.ParseMedicalText(ctx, visit.RawInput)
// 	if err != nil {
// 		return fmt.Errorf("ai parsing failed: %w", err)
// 	}

// 	if err := s.repo.Create(ctx, visit); err != nil {
// 		return fmt.Errorf("create visit in repository: %w", err)
// 	}

// 	// Drugs
// 	for _, d := range aiResult.Drugs {
// 		drug := models.PrescribedDrug{
// 			VisitID:   visit.ID,
// 			DrugName:  d.Name,
// 			Dosage:    d.Dosage,
// 			Frequency: d.Frequency,
// 			Duration:  d.Duration,
// 		}
// 		s.drugs.Create(ctx, &drug)
// 	}

// 	// Lab tests
// 	for _, t := range aiResult.LabTests {
// 		labTest := models.LabTest{
// 			VisitID:  visit.ID,
// 			TestName: t,
// 		}
// 		s.labTests.Create(ctx, &labTest)
// 	}

// 	// Notes
// 	note := models.ClinicalNote{
// 		VisitID: visit.ID,
// 		Note:    aiResult.Notes,
// 	}
// 	s.notes.Create(ctx, &note)

// 	return nil
// }

func (s *Service) List(ctx context.Context) ([]models.Visit, error) {
	visits, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list visits: %w", err)
	}

	return visits, nil
}

func (s *Service) ListPaginated(pagination repositories.Pagination[models.Visit]) (*repositories.Pagination[models.Visit], error) {
	paginatedResult, err := s.repo.ListPaginated(pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated visits from repository: %w", err)
	}
	return paginatedResult, nil
}

func (s *Service) Get(ctx context.Context, id uint) (models.Visit, error) {
	visit, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.Visit{}, fmt.Errorf("get visit from repository: %w", err)
	}

	return visit, nil
}
