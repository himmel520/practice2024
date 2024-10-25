package httpctrl

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *Handler) webSocketConn(c *gin.Context) {
	ws, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, respErr{err.Error()})
		return
	}
	defer ws.Close()

	for {
		var req reqPayload
		err := ws.ReadJSON(&req)
		if err != nil {
			h.sendError(ws, fmt.Sprintf("failed to read message: %v", err.Error()))
			return
		}

		switch req.Action {
		case "get_mapping":
			mapping := h.uc.Keyword.GetMapping()
			if err := ws.WriteJSON(mapping); err != nil {
				h.log.Error(err)
			}
		case "download":
			if err := h.fetchAndSendContent(req.Url, ws); err != nil {
				h.log.Error(err)
				h.sendError(ws, fmt.Sprintf("failed to download content %v", err.Error()))
				return
			}
		}
	}
}

func (h *Handler) fetchAndSendContent(url string, ws *websocket.Conn) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch content from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d for URL %s", resp.StatusCode, url)
	}

	downloadedData := []byte{}
	totalBytes := resp.ContentLength
	totalLoaded := 0

	buffer := make([]byte, 1024*64) // 64 KB

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			downloadedData = append(downloadedData, buffer[:n]...)
			totalLoaded += int(n)

			loadedMB := float64(totalLoaded) / (1024 * 1024) // Переводим в MB
			totalMB := float64(totalBytes) / (1024 * 1024)   // Переводим в MB
			progressPercent := (float64(totalLoaded) / float64(totalBytes)) * 100

			ws.WriteJSON(gin.H{
				"action":   "progress",
				"loaded":   loadedMB,
				"total":    totalMB,
				"progress": progressPercent,
			})
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error reading body: %w", err)
		}
	}

	ws.WriteJSON(gin.H{
		"action":       "completed",
		"url":          url,
		"content":      downloadedData,
		"content-type": resp.Header.Get("Content-Type"),
	})

	return nil
}
