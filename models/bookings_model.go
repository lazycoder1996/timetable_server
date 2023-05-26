package models

type Bookings struct {
	ID     int    `gorm:"primaryKey" binding:"required"`
	UserID int    `gorm:"foreignKey:user_fkey" binding:"required"`
	Date   string `binding:"required"`
	Start  string `binding:"required"`
	End    string `binding:"required"`
	Room   string `gorm:"foreignKey:room_fkey" binding:"required"`
}
