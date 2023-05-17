package outgoing

import (
	"context"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	authCommon "github.com/MindTickle/content-protos/pb/common"

	// AuditLog "github.com/MindTickle/platform-protos/pb/auditlogsservice"
	// AuditCommon "github.com/MindTickle/platform-protos/pb/common"
	auth "github.com/MindTickle/content-protos/pb/authorisation"
	event "github.com/MindTickle/platform-protos/pb/event"
	tickleDb "github.com/MindTickle/storageprotos/pb/tickleDb"
	sqlStore "github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
)

type TickleDbCreateTableClient interface {
	CreateTable(dbDetail model.TickleDbEnvDetail) (*tickleDb.CreateTableResponse, error)
}

type TickleDbInsertClient interface {
	InsertRow(id string, field model.User, url string, tableName string, reqContext sqlStore.RequestContext, authMeta sqlStore.AuthMeta) (*sqlStore.CreateRowsResponse, error)
}

type AuthorizationClient interface {
	GetCompnanyRolePermission(url string, companyId string, reqMeta authCommon.RequestMeta, recMeta authCommon.RecordMeta) (*auth.GetRoleAndPermissionsResponse, error)
	GetUserRolePermission(url string, userId string, reqMeta authCommon.RequestMeta, recMeta authCommon.RecordMeta) (*auth.GetUserRolesAndPermissionsResponse, error)
}

type EventClient interface {
	CreateEvents(ctx context.Context, url string, channelId int64, eventData model.EventField) (*event.CreateEventsResponse, error)
}

type EventChannelClient interface {
	CreateEventChannel(ctx context.Context, url string, channelData model.EventChannelField) error
}

type EmailClient interface {
	SendEmail(url string, input model.EmailField) (*model.EmailResponse, error)
}

type NotficationClient interface {
	SendNotification(url string, input model.NotificationField) (*model.EmailResponse, error)
}

type AuthenticationClient interface {
	VerifySession(url string, sessionId string) (*model.AuthenticatioResponse, error)
}

// type AuditLogClient interface {
// 	AddLog(url string, reqM AuditCommon.RequestMeta, schemaField []AuditLog.IngestField, field model.AuditField) error
// }

