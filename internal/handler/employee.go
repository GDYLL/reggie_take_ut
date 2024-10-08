package handler

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"strconv"
	"time"

	"github.com/prynnekey/go-reggie/global"
	"github.com/prynnekey/go-reggie/utils"
)

type EmployeeController struct {
}

func (e EmployeeController) Login(w http.ResponseWriter, r *http.Request) {

	var empInput model.Employee
	if err := json.NewDecoder(r.Body).Decode(&empInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	username := empInput.Username
	password := utils.MD5(empInput.Password)

	var empStored model.Employee
	if err := global.DB.Table("employee").Where("username = ?", username).First(&empStored).Error; err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	if empStored.Status == 0 {
		http.Error(w, "账号已禁用", http.StatusUnauthorized)
		return
	}
	if empStored.Password == password {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(common.Success("login success"))
		if err != nil {
			http.Error(w, "系统错误", http.StatusInternalServerError)
			return
		}
		return
	} else {
		http.Error(w, "用户名或密码不正确", http.StatusBadRequest)
		return
	}

}

func (e EmployeeController) Save(w http.ResponseWriter, r *http.Request) {
	var empInput model.Employee
	if err := json.NewDecoder(r.Body).Decode(&empInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	var empStored model.Employee
	if err := global.DB.Table("employee").Where("username = ?", empInput.Username).First(&empStored).Error; err != nil {
		// 如果根据username找到记录，则说明已有同名用户，此时应该报错
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 设置默认密码
			password := utils.MD5("123456")
			// 设置创建时间和更新时间
			now := time.Now()
			// 填充Employee结构体的字段
			empInput.Password = password
			empInput.CreateTime = now
			empInput.UpdateTime = now

			if err := global.DB.Table("employee").Create(&empInput).Error; err != nil {
				// 处理错误
				http.Error(w, "Error inserting new employee record", http.StatusInternalServerError)
				return
			}
			// 返回成功的响应
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(common.Success("保存成功"))
			if err != nil {
				http.Error(w, "系统错误", http.StatusInternalServerError)
				return
			}
		} else {
			// 查询时出现错误，但不是因为记录未找到
			http.Error(w, "Database query failed", http.StatusInternalServerError)
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(common.Success("用户名已存在"))
		if err != nil {
			http.Error(w, "系统错误", http.StatusInternalServerError)
			return
		}
		return
	}
	return

}

func (e EmployeeController) Page(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")
	emName := r.URL.Query().Get("name")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum <= 0 {
		http.Error(w, "无效的页码", http.StatusBadRequest)
		return
	}

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeNum <= 0 {
		http.Error(w, "无效的页大小", http.StatusBadRequest)
		return
	}

	offset := (pageNum - 1) * pageSizeNum
	var employees []model.Employee
	if emName != "" {
		// 添加模糊查询条件
		err = global.DB.Table("employee").Where("name LIKE ?", "%"+emName+"%").
			Offset(offset).Limit(pageSizeNum).Find(&employees).Error
	} else {
		err = global.DB.Table("employee").Offset(offset).Limit(pageSizeNum).Find(&employees).Error
	}
	if err != nil {
		// 如果查询失败，返回错误信息
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	var total int64
	err = global.DB.Table("employee").Count(&total).Error
	if err != nil {
		// 如果查询失败，返回错误信息
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	responseData := model.ResponseData{
		Records: employees,
		Total:   total,
	}
	w.Header().Set("Content-Type", "application/json")

	if json.NewEncoder(w).Encode(common.Success(responseData)) != nil {
		http.Error(w, "JSON 编码失败", http.StatusInternalServerError)
	}
	return

}

func (e EmployeeController) Get(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "无效的ID", http.StatusBadRequest)
		return
	}
	var employee model.Employee
	if err := global.DB.Table("employee").Where("id = ?", id).First(&employee).Error; err != nil {
		// 如果查询失败，返回错误信息
		http.Error(w, "查询失败", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(common.Success(employee))
	if err != nil {
		http.Error(w, "系统错误", http.StatusInternalServerError)
		return
	}
	return

}

func (e EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	var empInput model.Employee
	if err := json.NewDecoder(r.Body).Decode(&empInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	id := empInput.ID
	if id == "" {
		http.Error(w, "无效的ID", http.StatusBadRequest)
		return
	}
	if err := global.DB.Table("employee").Where("id = ?", id).Updates(&empInput).Error; err != nil {
		// 如果更新失败，返回错误信息
		http.Error(w, "更新失败", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(common.Success("更新成功"))
	if err != nil {
		http.Error(w, "系统错误", http.StatusInternalServerError)
		return
	}
	return

}
