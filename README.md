# HR System Backend

人力資源管理系統後端服務，使用 Go 語言開發。

## 系統需求

- Go 1.21+
- MySQL 8.0+
- Make

## 快速開始

### 1. 環境設定
(設定我就寫死了比較簡單不用改)

### 2. 啟動服務

- 啟動開發環境：

  ```bash
  make dev
  ```

- 服務將在 `http://localhost:8080` 啟動

## API 端點

### 員工管理

- **POST** `/api/v1/employees` - 創建員工
- **GET** `/api/v1/employees` - 取得員工列表
- **GET** `/api/v1/employees/:id` - 取得特定員工

### 請假管理

- **POST** `/api/v1/leaves` - 申請請假
- **GET** `/api/v1/leaves` - 取得請假列表
- **GET** `/api/v1/leaves/:id` - 取得特定請假申請
- **PUT** `/api/v1/leaves/:id/approve` - 核准請假
- **PUT** `/api/v1/leaves/:id/reject` - 拒絕請假

## 開發指令

- **啟動開發環境**：

  ```bash
  make dev
  ```

- **資料庫遷移**：

  ```bash
  make migrate
  ```

- **執行測試**：
  ```bash
  make test
  ```

## 專案結構

- **測試檔案**：放在 `tests/` 資料夾中
- **資料庫遷移**：系統啟動時會自動執行資料庫遷移，初始化必要的表格和資料

## 開發說明

- 使用 `make dev` 啟動開發環境
- 修改程式碼後，服務會自動重新載入
- API 回應格式統一使用 JSON

## 測試

- 執行所有測試：

  ```bash
  make test
  ```

- 執行特定測試：
  ```bash
  go test ./tests/... -run TestSpecific
  ```

## License

MIT

## api 範例

生成請假
```
curl --location 'http://localhost:8080/api/v1/leaves' \
--header 'Content-Type: application/json' \
--data '{
    "employee_id": 1,
    "leave_type": "annual",
    "start_date": 1705708800,
    "end_date": 1705881600,
    "reason": "Family vacation"
  }'
```
請假列表
```
curl --location --request GET 'http://localhost:8080/api/v1/leaves' \
--header 'Content-Type: application/json' \
--data-raw '{
    "employeeCode": "EMP001",
    "name": "張小明",
    "email": "xming@example.com",
    "phoneNumber": "0912345678",
    "departmentId": 1,
    "positionId": 1,
    "status": "active",
    "joinDate": "2024-03-20T00:00:00Z"
  }'
```

### 請假類型

- annual: 年假
- sick: 病假
- personal: 事假
- marriage: 婚假
- funeral: 喪假
- maternity: 產假
- paternity: 陪產假

### 狀態碼說明

- 200: 請求成功
- 201: 創建成功
- 400: 請求參數錯誤
- 401: 未授權
- 403: 權限不足
- 404: 資源不存在
- 500: 伺服器錯誤
