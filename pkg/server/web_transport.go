package server

import (
	"TrainTracking/internal/features/repository"
	"TrainTracking/internal/features/service"
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

func StartWebTransportServer(DB *gorm.DB) {
	cert, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/key.pem")
	if err != nil {
		log.Fatal("Failed to load cert:", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"h3"},
	}

	server := webtransport.Server{
		H3: http3.Server{
			Addr:      ":4433",
			TLSConfig: tlsConfig,
		},
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/v1/wt/train/", func(w http.ResponseWriter, r *http.Request) {
		session, err := server.Upgrade(w, r)
		if err != nil {
			log.Println("WebTransport upgrade error:", err)
			return
		}
		go handleWebTransportSession(session, r, DB)
	})

	go func() {
		log.Println("Running WebTransport server on :4433")
		if err := server.ListenAndServeTLS("certs/cert.pem", "certs/key.pem"); err != nil {
			log.Fatal("WebTransport server error:", err)
		}
	}()
}

func handleWebTransportSession(session *webtransport.Session, r *http.Request, DB *gorm.DB) {
	defer func(session *webtransport.Session, code webtransport.SessionErrorCode, msg string) {
		err := session.CloseWithError(code, msg)
		if err != nil {
			return
		}
	}(session, 0, "closed")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		log.Println("Invalid path")
		return
	}
	id := parts[4]

	stream, err := session.AcceptStream(context.Background())
	if err != nil {
		log.Println("Failed to accept stream:", err)
		return
	}

	encoder := json.NewEncoder(stream)
	trainRepo := repository.NewTrainRepository(DB)
	trainService := service.NewTrainService(trainRepo)

	for {
		resp, err := trainService.GetTrainPosition(id)
		if err != nil {
			log.Println("GetTrainPosition error:", err)
			return
		}

		if err := encoder.Encode(resp); err != nil {
			log.Println("WebTransport write error:", err)
			return
		}
		time.Sleep(2 * time.Second)
	}
}
