package usecase

import (
	"bcc-geazy/internal/controller/delivery"
	"bcc-geazy/internal/entity"
	"bcc-geazy/internal/model"
	"bcc-geazy/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type IKonsultasiUsecase interface {
	KirimPesan(ctx context.Context, buatKonsultasi model.BuatPesan) (*model.KonsultasiResponse, error)
	GetPesan(ctx context.Context, userID, dokterID uuid.UUID) ([]model.KonsultasiResponse, error)
}

type KonsultasiUsecase struct {
	KonsultasiRepository repository.IKonsulRepository
	wsManager            *websocket.WSManager
}

func NewKonsultasiUsecase(korepo repository.IKonsulRepository, ws *websocket.WSManager) *KonsultasiUsecase {
	return &KonsultasiUsecase{
		KonsultasiRepository: korepo,
		wsManager:            ws,
	}
}

func (p *KonsultasiUsecase) KirimPesan(ctx context.Context, buatKonsultasi model.BuatPesan) (*model.KonsultasiResponse, error) {
	konsul := entity.Konsultasi{
		Id:         uuid.New(),
		UserID:     buatKonsultasi.UserID,
		DokterID:   buatKonsultasi.DokterID,
		SenderID:   buatKonsultasi.SenderID,
		Pesan:      buatKonsultasi.Pesan,
		Dibaca:     false,
		WaktuPesan: time.Now(),
	}

	err := p.KonsultasiRepository.CreateKonsultasi(ctx, konsul)
	if err != nil {
		return nil, err
	}

	var receiverID uuid.UUID
	if buatKonsultasi.SenderID == buatKonsultasi.UserID {
		receiverID = buatKonsultasi.DokterID
	} else {
		receiverID = buatKonsultasi.UserID
	}

	if p.wsManager != nil {
		p.wsManager.SendToUser(receiverID.String(), konsul)
	}

	response := model.ToKonsultasiResponse(konsul)
	return &response, nil
}

func (p *KonsultasiUsecase) GetPesan(ctx context.Context, userID, dokterID uuid.UUID) ([]model.KonsultasiResponse, error) {
	data, err := p.KonsultasiRepository.GetByUserdanDokter(ctx, userID, dokterID)
	if err != nil {
		return nil, err
	}

	return model.ToKonsultasiResponses(data), nil
}
