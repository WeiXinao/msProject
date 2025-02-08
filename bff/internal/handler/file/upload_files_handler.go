package file

import (
	"net/http"

	filelogic "github.com/WeiXinao/msProject/bff/internal/logic/file"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := filelogic.NewUploadFilesLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadFiles(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respx.Success(resp))
		}
	}
}
