package services

import (
	"context"
	"net/url"
	"regexp"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/shibafu528/utissue/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var reInvalidTag = regexp.MustCompile("[\\s\\r\\n]")

var defaultUser = pb.User{
	Id:   0,
	Name: "Default user",
}

type server struct {
	pb.UnimplementedCheckinsServer
	checkinMutex sync.Mutex
	checkinSeq   uint64
	checkins     map[uint64]*pb.Checkin
}

func NewCheckinsServer() pb.CheckinsServer {
	return &server{
		checkinSeq: 0,
		checkins:   map[uint64]*pb.Checkin{},
	}
}

func (s *server) Create(ctx context.Context, request *pb.CreateCheckinRequest) (*pb.CreateCheckinResponse, error) {
	if request.CheckedInAt == nil {
		request.CheckedInAt = timestamppb.Now()
	}
	if !request.CheckedInAt.IsValid() {
		return nil, status.Error(codes.InvalidArgument, "checked_in_at is invalid")
	}
	at := request.CheckedInAt.AsTime()
	at = at.Truncate(time.Minute)
	for _, c := range s.checkins {
		if c.CheckedInAt.AsTime() == at {
			return nil, status.Error(codes.AlreadyExists, "checkin already exists in this time")
		}
	}

	if utf8.RuneCountInString(request.Note) > 500 {
		return nil, status.Error(codes.InvalidArgument, "note is too long, maximum is 500 characters")
	}

	if l := utf8.RuneCountInString(request.Link); l > 0 {
		if l > 2000 {
			return nil, status.Error(codes.InvalidArgument, "link is too long, maximum is 2000 characters")
		}

		u, err := url.Parse(request.Link)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "link is invalid URL")
		}
		switch u.Scheme {
		case "http", "https":
			// ok
		default:
			return nil, status.Error(codes.InvalidArgument, "link has invalid scheme, must be http or https")
		}
	}

	for i, t := range request.Tags {
		if utf8.RuneCountInString(t) > 255 {
			return nil, status.Errorf(codes.InvalidArgument, "tags[%d] is too long, maximum is 255 characters", i)
		}
		if reInvalidTag.MatchString(t) {
			return nil, status.Errorf(codes.InvalidArgument, "tags[%d] cannot contain spaces, tabs and newlines", i)
		}
	}

	s.checkinMutex.Lock()
	defer s.checkinMutex.Unlock()

	c := &pb.Checkin{
		Id:          s.checkinSeq,
		User:        &defaultUser,
		CheckedInAt: timestamppb.New(at),
		Note:        request.Note,
		Link:        request.Link,
		Tags:        request.Tags,
	}
	s.checkinSeq++
	s.checkins[c.Id] = c

	return &pb.CreateCheckinResponse{Checkin: c}, nil
}

func (s *server) Get(ctx context.Context, request *pb.GetCheckinRequest) (*pb.GetCheckinResponse, error) {
	c, ok := s.checkins[request.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "checkin not found")
	}

	return &pb.GetCheckinResponse{Checkin: c}, nil
}
