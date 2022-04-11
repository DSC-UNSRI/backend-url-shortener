package usecase

import (
	"log"

	"github.com/DSC-UNSRI/backend-url-shortener/model"
	"github.com/DSC-UNSRI/backend-url-shortener/repository"
)

type LinkUsecase struct {
	repository.LinkRepository
}

func NewLinkUsecase(repo repository.LinkRepository) *LinkUsecase {
	return &LinkUsecase{
		LinkRepository: repo,
	}
}

func (uc *LinkUsecase) GetAllLinks() model.Response {
	var response model.Response

	links, err := uc.LinkRepository.GetAllLinks()
	if err != nil {
		response.Status = 500
		response.Data = nil
		response.Message = "Gagal mengambil data links"

		return response
	}

	response.Data = links
	response.Status = 200
	response.Message = "Sukses"
	return response
}

func (uc *LinkUsecase) Redirect(shortenedLink string) string {
	var response model.Response

	if shortenedLink == "" {
		return ""
	}

	link, err := uc.LinkRepository.GetLinkByShortenedLink(shortenedLink)
	if err != nil {
		log.Println(err)
		return "error"
	}

	if link.Id == 0 || link.OriginLink == "" {
		response.Status = 404
		response.Data = nil
		response.Message = "Link not found"
		return ""
	}

	return link.OriginLink
}
