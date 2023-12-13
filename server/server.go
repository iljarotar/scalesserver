package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/iljarotar/scalesalgorithm"
	"go.uber.org/zap"
)

type server struct {
	port, host       string
	maxRange, maxNum int
	logger           *zap.SugaredLogger
}

type ServerConfig struct {
	Port, Host       string
	MaxRange, MaxNum int
	Logger           *zap.SugaredLogger
}

func NewServer(c *ServerConfig) *server {
	return &server{
		port:     c.Port,
		host:     c.Host,
		maxRange: c.MaxRange,
		maxNum:   c.MaxNum,
		logger:   c.Logger,
	}
}

func (s *server) Serve() error {
	s.logger.Infow("starting server", "host", s.host, "port", s.port)
	http.HandleFunc("/", s.requestHandler)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.host, s.port), nil)
}

func (s *server) requestHandler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	if err := req.ParseForm(); err != nil {
		s.logger.Errorw("unable to parse form", "error", err)
		fmt.Fprintf(w, "parseForm error: %v\n", err)
		return
	}

	s.logger.Infow("handle request", "request form", req.Form)

	scaleRange, err1 := strconv.Atoi(req.FormValue("range"))
	numNotes, err2 := strconv.Atoi(req.FormValue("notes"))
	if err1 != nil || err2 != nil {
		s.logger.Errorw("unable to parse input", "error", fmt.Errorf("parse errors, %w, %w", err1, err2))
		fmt.Fprintf(w, "unable to parse input, range: %v, notes: %v\n", err1, err2)
		return
	}

	if scaleRange > s.maxRange || scaleRange < 0 {
		s.logger.Errorw("invalid input for range received", "range", scaleRange)
		fmt.Fprintf(w, "invalid input for range, %d, must be between 0 and %d\n", scaleRange, s.maxRange)
		return
	}

	if numNotes > s.maxNum || numNotes < 0 {
		s.logger.Errorw("invalid input for notes received", "notes", numNotes)
		fmt.Fprintf(w, "invalid input for notes, %d, must be between 0 and %d\n", numNotes, s.maxNum)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	scales := scalesalgorithm.GetScales(scaleRange, numNotes, numNotes)
	response := make(map[string][][]int)
	response["scales"] = scales

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		s.logger.Errorw("unable to marshal response", "error", err)
		fmt.Fprintf(w, "json error, %v\n", err)
		return
	}

	w.Write(jsonResponse)
	return
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
