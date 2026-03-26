package usecase

import (
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"

	"github.com/google/uuid"
)

type IAnakUsecase interface {
	CreateDataAnak(ctx context.Context, buatAnak model.TambahDataAnak, userID uuid.UUID) (*model.AnakResponse, error)
	GetDataAnak(ctx context.Context, pagination model.Pagination) ([]model.AnakResponse, error)
	DeleteDataAnak(ctx context.Context, id uuid.UUID) error
	EditDataAnak(ctx context.Context, id uuid.UUID, edit model.EditDataAnak) error
}

type AnakUsecase struct {
	AnakRepository repository.IAnakRepository
}

func NewAnakUsecase(anakRepository repository.IAnakRepository) *AnakUsecase {
	return &AnakUsecase{anakRepository}
}

func (p *AnakUsecase) CreateDataAnak(ctx context.Context, buatAnak model.TambahDataAnak, userID uuid.UUID) (*model.AnakResponse, error) {

	bmi, status := HitungBMI(buatAnak.BeratBadan, buatAnak.Tinggi)
	umur := HitungUmur(buatAnak.TahunLahir)

	anak := entity.Anak{
		Id:              uuid.New(),
		UserID:          userID,
		Nama:            buatAnak.Nama,
		TanggalLahir:    buatAnak.TanggalLahir,
		BulanLahir:      buatAnak.BulanLahir,
		TahunLahir:      buatAnak.TahunLahir,
		Umur:            umur,
		Tinggi:          buatAnak.Tinggi,
		BeratBadan:      buatAnak.BeratBadan,
		Gender:          buatAnak.Gender,
		LingkarLengan:   buatAnak.LingkarLengan,
		LingkarKepala:   buatAnak.LingkarKepala,
		GolonganDarah:   buatAnak.GolonganDarah,
		Alergi:          buatAnak.Alergi,
		RiwayatPenyakit: buatAnak.RiwayatPenyakit,
		BMI:             bmi,
		StatusGizi:      status,
	}

	err := p.AnakRepository.CreateDataAnak(ctx, anak)
	if err != nil {
		return nil, err
	}

	response := model.ToAnakResponse(anak)
	return &response, nil
}

func (p *AnakUsecase) GetDataAnak(ctx context.Context, pagination model.Pagination) ([]model.AnakResponse, error) {
	anak, err := p.AnakRepository.GetDataAnak(ctx, pagination)
	if err != nil {
		return nil, err
	}

	responses := model.ToAnakResponses(anak)
	return responses, nil
}

func (p *AnakUsecase) DeleteDataAnak(ctx context.Context, id uuid.UUID) error {
	return p.AnakRepository.DeleteDataAnak(ctx, id)
}

func (p *AnakUsecase) EditDataAnak(ctx context.Context, id uuid.UUID, edit model.EditDataAnak) error {
	return p.AnakRepository.EditDataAnak(ctx, id, edit)
}
