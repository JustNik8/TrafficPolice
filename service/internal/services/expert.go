package services

import (
	"TrafficPolice/errs"
	"TrafficPolice/internal/domain"
	"TrafficPolice/internal/repository"
	"errors"
	"github.com/google/uuid"
)

type ExpertService interface {
	GetExpertByUserID(userID string) (domain.Expert, error)
	GetCase(userID string) (domain.Case, error)
	SetCaseDecision(decision domain.Decision) (domain.CaseDecisionInfo, error)
	GetCaseWithPersonInfo(caseID string) (domain.Case, error)
}

type expertService struct {
	expertRepo repository.ExpertRepo
	caseRepo   repository.CaseRepo
	consensus  int
}

func NewExpertService(
	expertRepo repository.ExpertRepo,
	caseRepo repository.CaseRepo,
	consensus int,
) ExpertService {
	return &expertService{
		expertRepo: expertRepo,
		caseRepo:   caseRepo,
		consensus:  consensus,
	}
}

func (s *expertService) GetCase(userID string) (domain.Case, error) {
	expert, err := s.expertRepo.GetExpertByUserID(userID)
	if err != nil {
		return domain.Case{}, err
	}

	caseID, err := s.expertRepo.GetLastNotSolvedCaseID(expert.ID)
	if err == nil {
		return s.caseRepo.GetCaseByID(caseID)
	}
	if !errors.Is(err, errs.ErrNoLastNotSolvedCase) {
		return domain.Case{}, err
	}

	notSolvedCase, err := s.expertRepo.GetNotSolvedCase(expert)
	if err != nil {
		return domain.Case{}, err
	}

	err = s.expertRepo.InsertNotSolvedCase(domain.SolvedCase{
		SolvedCaseID:  uuid.New().String(),
		ExpertID:      expert.ID,
		CaseID:        notSolvedCase.ID,
		IsExpertSolve: false,
		FineDecision:  false,
	})
	if err != nil {
		return domain.Case{}, err
	}

	return notSolvedCase, err
}

func (s *expertService) GetExpertByUserID(userID string) (domain.Expert, error) {
	return s.expertRepo.GetExpertByUserID(userID)
}

func (s *expertService) SetCaseDecision(decision domain.Decision) (domain.CaseDecisionInfo, error) {
	err := s.expertRepo.SetCaseDecision(decision)
	if err != nil {
		return domain.CaseDecisionInfo{}, err
	}

	caseDecisions, err := s.expertRepo.GetCaseFineDecisions(decision.CaseID)
	if err != nil {
		return domain.CaseDecisionInfo{}, err
	}

	if caseDecisions.PositiveDecisions >= s.consensus {
		err = s.caseRepo.SetCaseFineDecision(decision.CaseID, true)
		if err != nil {
			return domain.CaseDecisionInfo{}, err
		}
		return domain.CaseDecisionInfo{
			CaseID:         decision.CaseID,
			ShouldSendFine: true,
			IsSolved:       true,
		}, err
	}
	if caseDecisions.NegativeDecisions >= s.consensus {
		err = s.caseRepo.SetCaseFineDecision(decision.CaseID, false)
		if err != nil {
			return domain.CaseDecisionInfo{}, err
		}
		return domain.CaseDecisionInfo{
			CaseID:         decision.CaseID,
			ShouldSendFine: false,
			IsSolved:       true,
		}, nil
	}

	expertsCnt, err := s.expertRepo.GetExpertsCountBySkill(decision.Expert.CompetenceSkill)
	if err != nil {
		return domain.CaseDecisionInfo{}, err
	}

	totalDecisions := caseDecisions.PositiveDecisions + caseDecisions.NegativeDecisions

	leftExperts := expertsCnt - totalDecisions
	leftDecisions := s.consensus - max(caseDecisions.PositiveDecisions, caseDecisions.NegativeDecisions)
	if leftExperts < leftDecisions {
		return domain.CaseDecisionInfo{
			CaseID:         decision.CaseID,
			ShouldSendFine: false,
			IsSolved:       false,
		}, s.caseRepo.UpdateCaseRequiredSkill(decision.CaseID, decision.Expert.CompetenceSkill+1)
	}

	return domain.CaseDecisionInfo{
		CaseID:         decision.CaseID,
		ShouldSendFine: false,
		IsSolved:       false,
	}, nil
}

func (s *expertService) GetCaseWithPersonInfo(caseID string) (domain.Case, error) {
	return s.caseRepo.GetCaseWithPersonInfo(caseID)
}
