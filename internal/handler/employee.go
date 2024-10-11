package handler

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"reggie_take_ut/internal/model"
	"reggie_take_ut/pkg/common"
	"reggie_take_ut/pkg/session"
	"strconv"
	"strings"
	"time"

	"github.com/prynnekey/go-reggie/global"
	"github.com/prynnekey/go-reggie/utils"
)

type EmployeeController struct {
}

func (e EmployeeController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. 解析请求体
	var empInput model.Employee
	if err := json.NewDecoder(r.Body).Decode(&empInput); err != nil {
		json.NewEncoder(w).Encode(common.Result{}.Error("Invalid request"))
		return
	}

	// 2. 获取并验证用户信息
	username := empInput.Username
	password := utils.MD5(empInput.Password)

	var empStored model.Employee
	result := global.DB.Table("employee").Where("username = ?", username).First(&empStored)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			json.NewEncoder(w).Encode(common.Result{}.Error("用户名或密码错误"))
			return
		}
		json.NewEncoder(w).Encode(common.Result{}.Error("系统错误"))
		return
	}

	// 3. 验证账号状态和密码
	if empStored.Status == 0 {
		json.NewEncoder(w).Encode(common.Result{}.Error("账号已禁用"))
		return
	}

	if empStored.Password != password {
		json.NewEncoder(w).Encode(common.Result{}.Error("用户名或密码错误"))
		return
	}

	// 4. 登录成功，设置session
	session, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		log.Printf("获取session失败: %v", err)
		json.NewEncoder(w).Encode(common.Result{}.Error("系统错误"))
		return
	}

	// 设置session值，存储更多有用的信息
	session.Values["employee_id"] = empStored.ID // 存储员工ID
	session.Values["username"] = username
	session.Values["authenticated"] = true

	// 保存session
	if err := session.Save(r, w); err != nil {
		log.Printf("保存session失败: %v", err)
		json.NewEncoder(w).Encode(common.Result{}.Error("系统错误"))
		return
	}

	// 5. 返回成功响应
	json.NewEncoder(w).Encode(common.Result{}.Success(map[string]interface{}{
		"id":       empStored.ID,
		"username": empStored.Username,
		"name":     empStored.Name,
	}))
}

func (e EmployeeController) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session, err := session.Store.Get(r, session.SessionName)
	if err != nil {
		json.NewEncoder(w).Encode(common.Result{}.Error("系统错误"))
		return
	}

	// 清除所有 session 值
	session.Values = make(map[interface{}]interface{})

	// 设置 session 过期
	session.Options.MaxAge = -1

	// 保存 session
	if err := session.Save(r, w); err != nil {
		json.NewEncoder(w).Encode(common.Result{}.Error("系统错误"))
		return
	}

	json.NewEncoder(w).Encode(common.Result{}.Success("退出成功"))
}

func (e EmployeeController) Save(w http.ResponseWriter, r *http.Request) {
	var empInput model.Employee
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&empInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	var empStored model.Employee
	if err := global.DB.Table("employee").Where("username = ?", empInput.Username).First(&empStored).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			password := utils.MD5(strconv.Itoa(123456))
			now := time.Now()
			empInput.Password = password
			empInput.CreateTime = now
			empInput.UpdateTime = now

			if err := global.DB.Table("employee").Create(&empInput).Error; err != nil {
				http.Error(w, "Error inserting new employee record", http.StatusInternalServerError)
				return
			}
			err := json.NewEncoder(w).Encode(common.Result{}.Success("保存成功"))
			if err != nil {
				http.Error(w, "系统错误", http.StatusInternalServerError)
				return
			}
			return
		} else {
			err := json.NewEncoder(w).Encode(common.Result{}.Success("Database query failed"))
			if err != nil {
				http.Error(w, "系统错误", http.StatusInternalServerError)
				return
			}
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(common.Result{}.Error("用户名已存在"))
		if err != nil {
			http.Error(w, "系统错误", http.StatusInternalServerError)
			return
		}
		return
	}
}

func (e EmployeeController) Page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	if json.NewEncoder(w).Encode(common.Result{}.Success(responseData)) != nil {
		http.Error(w, "JSON 编码失败", http.StatusInternalServerError)
		return
	}
	return
}

func (e EmployeeController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/employee/")
	if id == "" {
		http.Error(w, "员工ID不能为空", http.StatusBadRequest)
		return
	}

	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的员工ID格式", http.StatusBadRequest)
		return
	}

	var employee model.Employee
	result := global.DB.Table("employee").Where("id = ?", id).First(&employee)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "未找到该员工", http.StatusNotFound)
			return
		}

		log.Printf("查询员工信息失败: %v", result.Error)
		http.Error(w, "查询员工信息失败", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(common.Result{}.Success(employee)); err != nil {
		log.Printf("JSON编码失败: %v", err)
		http.Error(w, "系统错误", http.StatusInternalServerError)
		return
	}
	return
}

func (e EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	var empInput model.Employee
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&empInput); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	id := empInput.ID
	status := empInput.Status
	if id == "" {
		http.Error(w, "无效的ID", http.StatusBadRequest)
		return
	}
	if err := global.DB.Table("employee").Where("id = ?", id).Update("status", status).Error; err != nil {
		// 如果更新失败，返回错误信息
		http.Error(w, "更新失败", http.StatusInternalServerError)
		return
	}
	err := json.NewEncoder(w).Encode(common.Result{}.Success("更新成功"))
	if err != nil {
		http.Error(w, "系统错误", http.StatusInternalServerError)
		return
	}
	return
}
