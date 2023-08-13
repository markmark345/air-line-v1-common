package responses

import "github.com/markmark345/air-line-v1-common/api/model"

var lookups = []model.Lookup{
	{
		Key:      "success",
		HTTPCode: 200,
		Code:     1000,
		DescEN:   "Success",
		DescTH:   "สำเร็จ",
	},
	{
		Key:      "created",
		HTTPCode: 201,
		Code:     1000,
		DescEN:   "Creation successful",
		DescTH:   "สร้างสำเร็จ",
	},
	{
		Key:      "updated",
		HTTPCode: 200,
		Code:     1000,
		DescEN:   "Data successfully updated",
		DescTH:   "อัปเดตข้อมูลสำเร็จ",
	},
	{
		Key:      "user_not_found",
		HTTPCode: 201,
		Code:     5001,
		DescEN:   "Current user not found",
		DescTH:   "ไม่พบบัญชีผู้ใช้ปัจจุบัน",
	},
	{
		Key:      "user_not_found",
		HTTPCode: 201,
		Code:     5001,
		DescEN:   "Current user not found",
		DescTH:   "ไม่พบบัญชีผู้ใช้ปัจจุบัน",
	},
	{
		Key:      "invaild_parameters",
		HTTPCode: 400,
		Code:     1101,
		DescEN:   "Invaild Parameters",
		DescTH:   "พารามิเตอร์ไม่ถูกต้อง",
	},
	{
		Key:      "missing_required_parameters",
		HTTPCode: 400,
		Code:     1102,
		DescEN:   "Missing required parameters",
		DescTH:   "ไม่มีพารามิเตอร์ที่ต้องการ",
	},
	{
		Key:      "missing_authroization_header",
		HTTPCode: 401,
		Code:     1103,
		DescEN:   "Missing Authorization Header",
		DescTH:   "ไม่มีส่วนหัวการอนุญาต",
	},
	{
		Key:      "invalid_access_token",
		HTTPCode: 401,
		Code:     1104,
		DescEN:   "Invalid Access Token",
		DescTH:   "โทเค็นการเข้าถึงไม่ถูกต้อง",
	},
	{
		Key:      "internal_server_error",
		HTTPCode: 500,
		Code:     9999,
		DescEN:   "Internal Server Error",
		DescTH:   "ข้อผิดพลาดภายในเซิร์ฟเวอร์",
	},
}

func GetLookup(msgKey string) (*model.Lookup, *model.Response) {
	data := make(map[string]model.Lookup)
	for _, lookup := range lookups {
		data[lookup.Key] = lookup
	}

	getKey, ok := data[msgKey]
	if !ok {
		res := &model.Response{
			Status: model.Status{
				Code:        500,
				Description: "Internal Server Error",
				Details: model.Details{
					Source: "common.GetLookup",
					Reason: "lookup json not found key",
				},
			},
		}
		return nil, res
	}

	result := &model.Lookup{
		Key:      getKey.Key,
		HTTPCode: getKey.HTTPCode,
		Code:     getKey.Code,
		DescEN:   getKey.DescEN,
		DescTH:   getKey.DescTH,
	}

	return result, nil
}
