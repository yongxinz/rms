// Code generated by goctl. DO NOT EDIT.
package types

type PageReq struct {
	PageIndex int64 `form:"pageIndex,default=1"`
	PageSize  int64 `form:"pageSize,default=20"`
}

type Pagination struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	Count     int64 `json:"count"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
	Code     string `json:"code"`
}

type LoginResp struct {
	CurrentAuthority string `json:"currentAuthority"`
	Expire           int64  `json:"expire"`
	Token            string `json:"token"`
}

type CaptchaResp struct {
	Data string `json:"data"`
	Id   string `json:"id"`
}

type UserInfoResp struct {
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	UserName     string   `json:"userName"`
	UserId       int64    `json:"userId"`
	DeptId       int64    `json:"deptId"`
	Name         string   `json:"name"`
	Buttons      []string `json:"buttons"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
}

type UserListReq struct {
	PageReq
}

type UserListData struct {
	UserId    int64        `json:"userId"`
	UserName  string       `json:"username"`
	NickName  string       `json:"nickName"`
	Phone     string       `json:"phone"`
	Status    string       `json:"status"`
	Email     string       `json:"email"`
	Salt      string       `json:"salt"`
	Avatar    string       `json:"avatar"`
	Sex       string       `json:"sex"`
	Remark    string       `json:"remark"`
	CreatedAt string       `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
	CreateBy  int64        `json:"createBy"`
	UpdateBy  int64        `json:"updateBy"`
	RoleId    int64        `json:"roleId"`
	DeptId    int64        `json:"deptId"`
	PostId    int64        `json:"postId"`
	Dept      DeptListData `json:"dept"`
}

type UserListResp struct {
	List []*UserListData `json:"list"`
	Pagination
}

type RoleListReq struct {
	PageReq
}

type RoleListData struct {
	RoleId    int64  `json:"roleId"`
	RoleName  string `json:"roleName"`
	RoleKey   string `json:"roleKey"`
	RoleSort  int64  `json:"roleSort"`
	Flag      string `json:"flag"`
	DataScope string `json:"dataScope"`
	Admin     int64  `json:"admin"`
	Status    int64  `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type RoleListResp struct {
	List []RoleListData `json:"list"`
	Pagination
}

type DeptListReq struct {
	PageReq
}

type DeptListData struct {
	DeptId    int64  `json:"deptId"`
	DeptPath  string `json:"deptPath"`
	DeptName  string `json:"deptName"`
	Phone     string `json:"phone"`
	Status    int64  `json:"status"`
	Email     string `json:"email"`
	Leader    string `json:"leader"`
	Sort      int64  `json:"sort"`
	ParentId  int64  `json:"parentId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type DeptTreeResp struct {
	Id       int64          `json:"id"`
	Label    string         `json:"label"`
	Children []DeptTreeResp `json:"children"`
}

type SysConfigResp struct {
	SysAppName string `json:"sys_app_name"`
	SysAppLogo string `json:"sys_app_logo"`
}

type ConfigPwReq struct {
	ConfigKey string `path:"configKey"`
}

type ConfigPwResp struct {
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
}

type MenuRoleResp struct {
	MenuId     int64          `json:"menuId"`
	MenuName   string         `json:"menuName"`
	MenuType   string         `json:"menuType"`
	Title      string         `json:"title"`
	Permission string         `json:"permission"`
	Params     string         `json:"params"`
	Path       string         `json:"path"`
	Paths      string         `json:"paths"`
	Action     string         `json:"action"`
	Apis       string         `json:"apis"`
	SysApi     string         `json:"sysApi"`
	Breadcrumb string         `json:"breadcrumb"`
	Component  string         `json:"component"`
	ParentId   int64          `json:"parentId"`
	Sort       int64          `json:"sort"`
	DataScope  string         `json:"dataScope"`
	Icon       string         `json:"icon"`
	IsFrame    string         `json:"isFrame"`
	Visible    string         `json:"visible"`
	Is_select  bool           `json:"is_select"`
	NoCache    int64          `json:"noCache"`
	CreateBy   int64          `json:"createBy"`
	CreatedAt  string         `json:"createdAt"`
	UpdateBy   int64          `json:"updateBy"`
	UpdatedAt  string         `json:"updatedAt"`
	Children   []MenuRoleResp `json:"children"`
}

type DictDataOpReq struct {
	DictType string `form:"dictType"`
}

type DictDataOpResp struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PostListReq struct {
	PageReq
}

type PostListData struct {
	PostId    int64  `json:"postId"`
	PostName  string `json:"postName"`
	PostCode  string `json:"postCode"`
	Sort      int64  `json:"sort"`
	Status    int64  `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type PostListResp struct {
	List []PostListData `json:"list"`
	Pagination
}
