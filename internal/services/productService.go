package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"taobao/internal/dto"
)

type ProductService struct {
	baseURL string
	client  *http.Client
}

func NewProductService(baseUrl string) *ProductService {
	return &ProductService{
		baseURL: baseUrl,
		client:  &http.Client{},
	}
}

func (p ProductService) RecommendedProducts(ctx context.Context, pageNo, pageSize int, language string) (*dto.HotRecommendedResponse, error) {
	jsonData := []byte(fmt.Sprintf(`{"pageNo":%d,"pageSize":%d,"language":"%s"}`, pageNo, pageSize, language))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/open/product/recommend", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("access-key", "eb4e1d5755e246cf986caa3b6cce28d7")
	req.Header.Add("nonce-str", "U91NHJA5CJ8UYAWE6NT1WOH80452V9RJ")
	req.Header.Add("timestamp", "1733474179")
	req.Header.Add("user-id", "2100008175167")
	req.Header.Add("signature", "ADB3C101BAEAB8B67B0AF793CD595CE3")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hot recommended products returned status: %d", resp.StatusCode)
	}

	var result *dto.HotRecommendedResponse

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p ProductService) ProductDetails(ctx context.Context, itemId int64) (*dto.ProductDetailsResponse, error) {
	u, err := url.Parse(p.baseURL + "/open/product/get")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	q := u.Query()
	q.Set("itemId", strconv.FormatInt(itemId, 10))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("access-key", "eb4e1d5755e246cf986caa3b6cce28d7")
	req.Header.Add("nonce-str", "U91NHJA5CJ8UYAWE6NT1WOH80452V9RJ")
	req.Header.Add("timestamp", "1733474179")
	req.Header.Add("user-id", "2100008175167")
	req.Header.Add("signature", "ADB3C101BAEAB8B67B0AF793CD595CE3")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hot recommended products returned status: %d", resp.StatusCode)
	}

	var result *dto.ProductDetailsResponse

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
