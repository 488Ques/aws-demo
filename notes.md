# Tài khoản nhân viên
- Có quyền chỉnh sửa (thêm, sửa, xóa, tăng giảm số lượng) các sản phẩm trong kho của xe tải mà nhân viên đó phụ trách và chỉ có quyền xem các sản phẩm trong kho khác
- Có chức năng tạo đơn hàng xuất kho
- Khi số lượng đạt ngưỡng tối thiểu (sắp hết hàng) thì dùng Twillio để gửi tin nhắn cho quản lý

# Tài khoản quản lý
- Có quyền chỉnh sửa các sản phẩm ở tất cả các kho của xe tải
- Có quyền thêm xóa nhân viên

# Trong mỗi bảng đều có 1 thuộc tính là company_id là làm theo kiểu multi tenancy, cái nào mà là khóa ngoại thì là thuộc tính cần có trong bảng còn không phải thì chỉ để vào để xác định xem dòng này thuộc company nào

# Bảng xe tải: Bảng này lưu trữ thông tin về từng xe tải dịch vụ. Nó có các cột sau:
- id: Mã định danh duy nhất cho mỗi xe tải.
- name: Tên hoặc số của xe tải.
- vị trí: Vị trí hiện tại của xe tải.
- trạng thái: Trạng thái của xe tải, chẳng hạn như "có sẵn", "đang sử dụng" hoặc "đang bảo trì".

# Bảng tồn kho: Bảng này lưu trữ thông tin chi tiết về các mặt hàng tồn kho được vận chuyển bởi mỗi xe tải.
- Inventory_id: Mã định danh duy nhất cho từng mặt hàng trong kho.
- truck_id: Khóa ngoại liên kết từng mặt hàng trong kho với xe tải chứa mặt hàng đó.
- Product_name: Tên của mục, chẳng hạn như "cáp", "bóng đèn" hoặc "công tắc".
- Product_quantity: Số lượng hiện tại của mặt hàng trong kho trong xe tải.
- minimum_quantity: Số lượng tối thiểu của mặt hàng cần được chở trong xe tải.
- Company_id
- Inventory_status

# Bảng Company lưu trữ thông tin về công ty bảo trì điện:
- ID:
- name:
- Address:
- Email:
- Company_status:

# Bảng Staff lưu trữ thông tin về nhân viên của công ty bảo trì điện:
- Staff_id
- Truck_id: ID của xe tải được chỉ định cho họ; đối với nhân viên quản lý thì giá trị này null
- Company_id: ID của công ty họ làm việc.
- Staff_name:
- Phone_number:
- Is_manager: nếu là nhân viên quản lý thì True ngược lại thì false
- Staff_status:

# Bảng Truck được lưu trữ thông tin về từng xe tải:
- Truck_id :
- Company_id: ID của công ty sở hữu xe tải này
- Truck_Name:
- Location:
- Truck_status