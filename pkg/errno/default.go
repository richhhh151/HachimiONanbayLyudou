package errno

var (
	Success = NewErrNo(SuccessCode, "ok")

	AuthError          = NewErrNo(AuthErrorCode, "鉴权失败")            // 鉴权失败，通常是内部错误，如解析失败
	AuthInvalid        = NewErrNo(AuthInvalidCode, "鉴权无效")          // 鉴权无效，如令牌颁发者不是 west2-online
	AuthAccessExpired  = NewErrNo(AuthAccessExpiredCode, "访问令牌过期")  // 访问令牌过期
	AuthRefreshExpired = NewErrNo(AuthRefreshExpiredCode, "刷新令牌过期") // 刷新令牌过期
	AuthMissing        = NewErrNo(AuthInvalidCode, "缺失合法鉴权数据")      // 鉴权缺失，如访问令牌缺失

	ParamError = NewErrNo(ParamErrorCode, "参数错误") // 参数校验失败，可能是参数为空、参数类型错误等

	InternalServiceError = NewErrNo(InternalServiceErrorCode, "内部服务错误")
)
