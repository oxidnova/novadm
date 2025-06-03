package n8n

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/novadm/backend/internal/config"
)

type Proxy interface {
	CallWebhookForGenConsultation(string) error
}

func NewProxy(d dependencies) Proxy {
	return &defaultProxy{d: d}
}

type dependencies interface {
	Logger() logx.Logger
	Config() *config.Config
}

type defaultProxy struct {
	d dependencies
}

type GenConsultationRequest struct {
	Prompt string `json:"prompt"`
}

type GenConsultationResponse struct {
	Message string `json:"message"`
}

func (p *defaultProxy) CallWebhookForGenConsultation(prompt string) error {
	ctx := context.Background()

	req := &GenConsultationRequest{Prompt: prompt}
	httpReq, err := p.newHttpRequest(ctx, req)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp := &GenConsultationResponse{}
	if err := p.send(ctx, httpReq, resp); err != nil {
		return err
	}

	return nil
}

func (p *defaultProxy) newHttpRequest(ctx context.Context, pl any) (*http.Request, error) {
	var payload bytes.Buffer
	if err := json.NewEncoder(&payload).Encode(pl); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		strings.ToUpper(p.d.Config().N8N.Webhook.Method),
		p.d.Config().N8N.Webhook.Endpoint,
		&payload,
	)
	if err != nil {
		return nil, err
	}

	return httpReq, nil
}

func (p *defaultProxy) send(ctx context.Context, httpReq *http.Request, resp interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("status: %d, resp: %s", httpResp.StatusCode, string(data))
	}

	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return err
	}

	return nil
}
