package handler

import (
	"Valter/service"
)

type FeatHandler struct {
	featServ service.FeatService
}

func NewFeatHandler(featService service.FeatService) *FeatHandler {
	return &FeatHandler{featService}
}

func (fh *FeatHandler) FeatDummy() {
	_, err := fh.featServ.FeatDummy()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy2()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy3()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy4()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy5()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy6()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy7()
	if err != nil {
		return
	}
	_, err = fh.featServ.FeatDummy8()
	if err != nil {
		return
	}
}
