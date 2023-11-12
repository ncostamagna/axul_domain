package user

var roles map[string]uint64

var READ_ROLE uint64 = 1
var WRITE_ROLE uint64 = 1 << 1
var UPDATE_ROLE uint64 = 1 << 2
var DELETE_ROLE uint64 = 1 << 3
var ADMIN_R_ROLE uint64 = 1 << 4
var ADMIN_RW_ROLE uint64 = 1 << 5
var OWNER_ROLE uint64 = 1 << 6
	

func init() {

	roles = map[string]uint64{
		"read":     READ_ROLE,
		"write":    WRITE_ROLE,
		"update":   UPDATE_ROLE,
		"delete":   DELETE_ROLE,
		"admin_r":  ADMIN_R_ROLE,
		"admin_rw": ADMIN_RW_ROLE,
		"owner":    OWNER_ROLE,
	}

}


type Role struct {
	ID           string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID     	 string         `json:"user_id" gorm:"type:char(36);not null;"`
	App			 string 		`json:"app" gorm:"type:char(36);not null;"`
	Role         uint64         `json:"role"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}


func (r *Role) AddRole(role string) error {
	rValue := roles[role]
	if !rValue {
		return errors.New("invalid role")
	} 

	if !(rValue & r.Role) {
		r.Role = rValue + r.Role
	} 
	
	return nil
}


func (r *Role) RemoveRole(role string) error {
	rValue := roles[role]
	if !rValue {
		return errors.New("invalid role")
	} 

	if rValue & r.Role {
		r.Role =  r.Role - rValue
	} 

	return nil
}