package middleware

import (
	"net/http"

	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/project/projectservice"
	"github.com/WeiXinao/msProject/task/taskservice"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
)

const PrivateProject = 1;

type ProjectAuthMiddleware struct {
	encrypter     encrypts.Encrypter
	projectClient projectservice.ProjectService
	taskClient    taskservice.TaskService
}

func NewProjectAuthMiddleware(
	encrypter encrypts.Encrypter,
	projectClient projectservice.ProjectService,
	taskClient taskservice.TaskService,
) *ProjectAuthMiddleware {
	return &ProjectAuthMiddleware{
		encrypter:     encrypter,
		projectClient: projectClient,
		taskClient:    taskClient,
	}
}

func (m *ProjectAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logx.WithContext(r.Context())
		encoder := sonic.ConfigDefault.NewEncoder(w)
		projectCode := r.FormValue("projectCode")
		taskCode := r.FormValue("taskCode")
		if projectCode == "" && taskCode == "" {
			next(w, r)
			return
		}

		var (
			projectId = int64(0)
			err       error
		)
		if projectCode != "" {
			projectId, err = m.encrypter.DecryptInt64(projectCode)
			if err != nil {
				logger.Errorf("[ProjectAuthMiddleware] %#v", err)
				encoder.Encode(respx.Fail(respx.ErrInternalServer))
				return
			}
		}

		if projectCode == "" && taskCode != "" {
			taskId, err := m.encrypter.DecryptInt64(taskCode)
			if err != nil {
				logger.Errorf("[ProjectAuthMiddleware] %#v", err)
				encoder.Encode(respx.Fail(respx.ErrInternalServer))
				return
			}
			task, err := m.taskClient.FindTaskById(r.Context(), &taskservice.FindTaskByIdRequest{
				TaskId: taskId,
			})
			if err != nil {
				logger.Errorf("[ProjectAuthMiddleware] %#v", err)
				encoder.Encode(respx.Fail(respx.ErrInternalServer))
				return
			}
			if task.ProjectCode == "" {
				next(w, r)
				return
			}
			projectId, err = m.encrypter.DecryptInt64(projectCode)
			if err != nil {
				logger.Errorf("[ProjectAuthMiddleware] %#v", err)
				encoder.Encode(respx.Fail(respx.ErrInternalServer))
				return
			}
		}

		projectMsg, err := m.projectClient.FindProjectByMemberId(r.Context(), &projectservice.FindProjectByMemberIdRequest{
			MemberId: r.Context().Value(KeyMemberId).(int64),
			ProjectId: projectId,
		})
		if err != nil {
			logger.Errorf("[ProjectAuthMiddleware] %#v", err)
			encoder.Encode(respx.Fail(respx.ErrInternalServer))
			return
		}
		if !projectMsg.GetIsMember() {
			encoder.Encode(respx.Fail(respx.ErrNotMember))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if projectMsg.Project.Private == PrivateProject {
			if projectMsg.GetIsOwner() {
				next(w, r)
				return
			} else {
				encoder.Encode(respx.Fail(respx.ErrNotMember))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	}
}
