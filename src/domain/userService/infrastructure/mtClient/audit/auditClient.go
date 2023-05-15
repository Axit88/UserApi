package audit

// import (
// 	"fmt"

// 	"github.com/Axit88/UserApi/src/constants"
// 	"github.com/Axit88/UserApi/src/domain/userService/core/model"
// 	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
// 	"github.com/MindTickle/mt-go-logger/logger"
// 	pb "github.com/MindTickle/platform-protos/pb/auditlogsservice"
// 	common "github.com/MindTickle/platform-protos/pb/common"
// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// )

// type AuditImpl struct {
// 	AuditService pb.AuditLogsServiceClient
// 	logger       *logger.LoggerImpl
// }

// func NewAuditClient(l *logger.LoggerImpl) outgoing.AuditLogClient {

// 	if constants.IsMock {
// 		return MockAuditClient{}
// 	}

// 	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
// 	conn, err := grpc.Dial(connection, grpc.WithInsecure())
// 	if err != nil {
// 		return nil
// 	}

// 	res := AuditImpl{}
// 	res.AuditService = pb.NewAuditLogsServiceClient(conn)
// 	res.logger = l

// 	return res
// }

// func (client AuditImpl) AddLog(url string, reqM common.RequestMeta, schemaField []pb.IngestField, field model.AuditField) error {

// 	data := pb.AddAuditLogRequest{
// 		RequestMeta: &reqM,
// 		Type:        pb.AuditLogType(field.AuditType),
// 		Timestamp:   field.TimeStamp}

// 	for i, _ := range schemaField {
// 		data.Fields = append(data.Fields, &schemaField[i])
// 	}

// 	_, err := client.AuditService.AddAuditLog(context.Background(), &data)
// 	if err != nil {
// 		client.logger.Errorf(context.Background(), "Audit Client Failed", err)
// 		return err
// 	}

// 	return nil
// }
