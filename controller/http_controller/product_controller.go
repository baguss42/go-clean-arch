package http_controller

import (
	"context"
	"encoding/json"
	"github.com/baguss42/go-clean-arch/entity/dto"
	_interface "github.com/baguss42/go-clean-arch/interface"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Service _interface.ProductServiceInterface
	Ctx     context.Context
}

func (p Product) All(w http.ResponseWriter, r *http.Request) (int, error) {
	// TODO: remove this
	resp, err := http.Get("http://localhost:9090/product")
	if err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	type Products struct {
		Code int    `json:"code"`
		Info string `json:"info"`
		Data []struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Price       string `json:"price"`
			Qty         int    `json:"qty"`
		} `json:"data"`
	}

	var result Products
	if err = json.Unmarshal(body, &result); err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	return WriteSuccess(w, http.StatusOK, result.Data)
}

func (p Product) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	var request dto.ProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return WriteError(w, http.StatusBadRequest, err)
	}

	response, err := p.Service.Create(p.Ctx, request)
	if err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	return WriteSuccess(w, http.StatusCreated, response)
}

func (p Product) Read(w http.ResponseWriter, r *http.Request) (int, error) {
	response, err := p.Service.List(p.Ctx, dto.ListOption{})
	if err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	return WriteSuccess(w, http.StatusOK, response)
}

func (p Product) Update(w http.ResponseWriter, r *http.Request) (int, error) {
	response, err := p.Service.Update(p.Ctx, 0, dto.ProductRequest{})
	if err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	return WriteSuccess(w, http.StatusOK, response)
}

func (p Product) Delete(w http.ResponseWriter, r *http.Request) (int, error) {
	err := p.Service.Delete(context.Background(), 1)
	if err != nil {
		return WriteError(w, http.StatusInternalServerError, err)
	}

	return WriteSuccess(w, http.StatusOK, "deleted")
}
