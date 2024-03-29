package service

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/christophwitzko/flight-booking-service/pkg/database"
	"github.com/christophwitzko/flight-booking-service/pkg/database/seeder"
	"github.com/christophwitzko/flight-booking-service/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func BenchmarkHandlerGetFlights(b *testing.B) {
	db, _ := database.New()
	_ = seeder.Seed(db, 1000)
	s := New(logger.NewNop(), db)

	resWriter := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.handlerGetFlights(resWriter, req)
	}
}

func BenchmarkHandlerGetFlightsQuery(b *testing.B) {
	db, _ := database.New()
	_ = seeder.Seed(db, 1000)
	s := New(logger.NewNop(), db)

	resWriter := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/?from=AAA", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.handlerGetFlights(resWriter, req)
	}
}

func BenchmarkHandlerGetFlight(b *testing.B) {
	db, _ := database.New()
	_ = seeder.Seed(db, 1000)
	s := New(logger.NewNop(), db)

	req := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	rctx.Reset()
	rctx.URLParams.Add("id", findFlightID(db))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	resWriter := httptest.NewRecorder()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.handlerGetFlight(resWriter, req)
	}
}

func BenchmarkHandlerGetFlightSeats(b *testing.B) {
	db, _ := database.New()
	_ = seeder.Seed(db, 100)
	s := New(logger.NewNop(), db)

	req := httptest.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	rctx.Reset()
	rctx.URLParams.Add("id", findFlightID(db))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	resWriter := httptest.NewRecorder()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.handlerGetFlightSeats(resWriter, req)
	}
}

func BenchmarkHandlerGetDestinations(b *testing.B) {
	db, _ := database.New()
	_ = seeder.Seed(db, 1000)
	s := New(logger.NewNop(), db)

	resWriter := httptest.NewRecorder()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.handlerGetDestinations(resWriter, nil)
	}
}
