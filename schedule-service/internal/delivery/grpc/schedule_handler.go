package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/rivaldoyoseps/schedule-service/api/proto_gen/schedulepb"
	"github.com/rivaldoyoseps/schedule-service/internal/domain"
	"github.com/rivaldoyoseps/schedule-service/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScheduleGRPCHandler struct {
	pb.UnimplementedScheduleServiceServer
	usecase usecase.ScheduleUseCase
}

func NewScheduleGRPCHandler(usecase usecase.ScheduleUseCase) *ScheduleGRPCHandler {
	return &ScheduleGRPCHandler{usecase: usecase}
}

func (s *ScheduleGRPCHandler) CreateSlot(ctx context.Context, req *pb.CreateSlotRequest) (*pb.Empty, error) {
	startTime, err := parseTime(req.StartTime)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid start_time: %v", err)
	}
	endTime, err := parseTime(req.EndTime)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid end_time: %v", err)
	}

	slot := &domain.Slot{
		UserID:    req.UserId,
		StartTime: startTime,
		EndTime:   endTime,
		IsBooked:  false,
	}

	err = s.usecase.CreateSlot(slot)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create slot: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *ScheduleGRPCHandler) GetSlotsByUser(ctx context.Context, req *pb.GetSlotsByUserRequest) (*pb.GetSlotsResponse, error) {
	slots, err := s.usecase.GetSlotsByUser(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get slots: %v", err)
	}

	var protoSlots []*pb.Slot
	for _, slot := range slots {
		protoSlots = append(protoSlots, &pb.Slot{
			Id:        slot.ID,
			UserId:    slot.UserID,
			StartTime: slot.StartTime.Format(time.RFC3339),
			EndTime:   slot.EndTime.Format(time.RFC3339),
			IsBooked:  slot.IsBooked,
		})
	}

	return &pb.GetSlotsResponse{Slots: protoSlots}, nil
}



// parseTime mengubah string waktu menjadi time.Time dalam format RFC3339
func parseTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Printf("failed to parse time '%s': %v", timeStr, err)
		return time.Time{}, err
	}
	return t, nil
}
