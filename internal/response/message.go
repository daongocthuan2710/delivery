package response

import "fmt"

func messageRequired(field string) string {
	return fmt.Sprintf("%s không được trống", field)
}

func messageInvalid(field string) string {
	return fmt.Sprintf("%s không hợp lệ", field)
}

func messageNotFound(field string) string {
	return fmt.Sprintf("không tìm thấy %s", field)
}

func messageExisted(field string) string {
	return fmt.Sprintf(" %s đã tồn tại", field)
}
