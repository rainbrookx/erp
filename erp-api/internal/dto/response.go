package dto

// Response 统一返回结构体
type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    T      `json:"data"`
}

func Success[T any](data T) Response[T] {
	return Response[T]{
		Code:    200,
		Message: "成功",
		Data:    data,
	}
}

func Unauthorized[T any]() Response[T] {
	return Response[T]{
		Code:    401,
		Message: "没有权限",
	}
}

func Forbidden[T any]() Response[T] {
	return Response[T]{
		Code:    403,
		Message: "禁止访问",
	}
}

func NotFound[T any]() Response[T] {
	return Response[T]{
		Code:    404,
		Message: "记录不存在",
	}
}
