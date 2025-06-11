package grpc

import (
	"context"
	"time"

	"github.com/rivaldoyoseps/booking-service/api/proto_gen/bookingpb"
	"github.com/rivaldoyoseps/booking-service/internal/domain"
	"github.com/rivaldoyoseps/booking-service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingGRPCHandler struct {
	bookingpb.UnimplementedBookingServiceServer
	bookingUC usecase.BookingUseCase
}

func NewBookingGrpcHandler(usecase usecase.BookingUseCase) *BookingGRPCHandler {
	return &BookingGRPCHandler{
		bookingUC: usecase,
	}
}

func (h *BookingGRPCHandler) CreateBooking(ctx context.Context, req *bookingpb.CreateBookingRequest) (*bookingpb.BookingResponse, error) {
	booking := &domain.Booking{
		UserID: req.UserId,
		ScheduleId: req.ScheduleId,
		Note: req.Note,
	}

	err := h.bookingUC.CreateBooking(booking)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create booking %v", err)
	}

	return &bookingpb.BookingResponse{
		Id: booking.ID,
		UserId: booking.UserID,
		ScheduleId: booking.ScheduleId,
		Note: booking.Note ,
		CreatedAt: booking.CreatetAt.Format(time.RFC3339),
	}, nil
}

func (h *BookingGRPCHandler) GetBookingByUser(ctx context.Context, req *bookingpb.GetBookingByUserRequest) (*bookingpb.GetBookingsResponse, error) {
	bookings, err := h.bookingUC.GetBookingsByUser(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get bookings: %v", err)
	}

	var protoBookings []*bookingpb.BookingResponse
	for _, b := range bookings {
		protoBookings = append(protoBookings, &bookingpb.BookingResponse{
			Id:         b.ID,
			UserId:     b.UserID,
			ScheduleId: b.ScheduleId,
			Note:       b.Note,
			CreatedAt:  b.CreatetAt.Format(time.RFC3339),
		})
	}

	return &bookingpb.GetBookingsResponse{Bookings: protoBookings}, nil
}

func (h *BookingGRPCHandler) GetBookingDetail(ctx context.Context, req *bookingpb.GetBookingByDetailRequest) (*bookingpb.BookingResponse, error) {
	booking, err := h.bookingUC.GetBookingDetail(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "booking not found: %v", err)
	}

	return &bookingpb.BookingResponse{
		Id:         booking.ID,
		UserId:     booking.UserID,
		ScheduleId: booking.ScheduleId,
		Note:       booking.Note,
		CreatedAt:  booking.CreatetAt.Format(time.RFC3339),
	}, nil
}