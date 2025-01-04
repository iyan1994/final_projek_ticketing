package models

import "time"

type User struct {
	IdUser       int       `gorm:"type:int;primaryKey;autoIncrement;column:id_user"` // Kolom id_user sebagai primary key dan auto increment
	Username     string    `gorm:"type:varchar(200);column:username;not null;unique"`
	Password     string    `gorm:"type:varchar(150);column:password;not null"`
	Name         string    `gorm:"type:varchar(150);column:name"`
	Email        string    `gorm:"type:varchar(150);column:email;not null;unique"`
	Token        string    `gorm:"type:varchar(220);column:token"`
	ExpiredToken time.Time `gorm:"type:datetime;column:expired_token"`
	IdRole       int       `gorm:"type:int;index;column:id_role;not null"` // Foreign Key yang mengarah ke role.id_role
	Image        string    `gorm:"type:varchar(220);column:image"`
	Title        string    `gorm:"type:varchar(220);column:title;not null"`
	NoTelepon    int       `gorm:"type:int;column:no_telepon;not null"`
	Address      string    `gorm:"column:address;type:text;not null"`
	CreatedAt    time.Time `gorm:"type:datetime;"`
	UpdatedAt    time.Time `gorm:"type:datetime;"`
	Role         Role      `gorm:"foreignKey:IdRole;constraint:OnDelete:RESTRICT;"` // Relasi dengan Role
}

type Role struct {
	IdRole    int       `gorm:"type:int;primaryKey;autoIncrement;column:id_role"` // Kolom id_role sebagai primary key dan auto increment
	NameRole  string    `gorm:"type:varchar(200);column:name_role;not null;unique"`
	CreatedAt time.Time `gorm:"type:datetime;"`
	UpdatedAt time.Time `gorm:"type:datetime;"`
}

type Ticket struct {
	IdTicket   int       `gorm:"type:int;primaryKey;autoIncrement;column:id_ticket"` // Kolom id_ticket sebagai primary key dan auto increment
	IdUser     int       `gorm:"type:int;index;column:id_user;not null"`             // Foreign Key yang mengarah ke User.id_user
	IdCategory int       `gorm:"type:int;index;column:id_category;not null"`         // Foreign Key yang mengarah ke Kategoru.id_kategori
	Status     string    `gorm:"type:varchar(220);column:status;not null"`
	NoTiket    string    `gorm:"type:varchar(220);column:no_tiket;not null"`
	Subjek     string    `gorm:"type:varchar(220);column:subjek;not null"`
	Deksripsi  string    `gorm:"type:text;column:deksripsi;not null"`
	CreatedAt  time.Time `gorm:"type:datetime;"`
	UpdatedAt  time.Time `gorm:"type:datetime;"`
	User       User      `gorm:"foreignKey:IdUser;constraint:OnDelete:RESTRICT;"`     // Relasi ke User
	Category   Category  `gorm:"foreignKey:IdCategory;constraint:OnDelete:RESTRICT;"` // Relasi ke Kategori

}

type Category struct {
	IdCategory   int       `gorm:"type:int;primaryKey;autoIncrement;column:id_category"` // Kolom id_kategori sebagai primary key dan auto increment
	NameCategory string    `gorm:"type:varchar(200);column:name_category;not null"`
	CreatedAt    time.Time `gorm:"type:datetime;"`
	UpdatedAt    time.Time `gorm:"type:datetime;"`
}

type Feedback struct {
	IdFeedback   int       `gorm:"type:int;primaryKey;autoIncrement;column:id_feedback"` // Kolom id_feedback sebagai primary key dan auto increment
	IdTicket     int       `gorm:"type:int;index;column:id_ticket;not null"`             // Foreign Key yang mengarah ke Ticket.id_ticket
	Satisfaction string    `gorm:"type:varchar(100);column:satisfaction;not null"`
	Deksripsi    string    `gorm:"type:text;column:deksripsi;not null"`
	CreatedAt    time.Time `gorm:"type:datetime;"`
	UpdatedAt    time.Time `gorm:"type:datetime;"`
	Ticket       Ticket    `gorm:"foreignKey:IdTicket;constraint:OnDelete:RESTRICT;"` // Relasi ke ticket

}

type ImageTicket struct {
	IdImageTicket int       `gorm:"type:int;primaryKey;autoIncrement;column:id_image_ticket"` // Kolom id_image_ticket sebagai primary key dan auto increment
	IdTicket      int       `gorm:"type:int;index;column:id_ticket;not null"`                 // Foreign Key yang mengarah ke Ticket.id_ticket
	Image         string    `gorm:"type:varchar(200);column:image;not null"`
	Deksripsi     string    `gorm:"type:text;column:deksripsi;not null"`
	CreatedAt     time.Time `gorm:"type:datetime;"`
	UpdatedAt     time.Time `gorm:"type:datetime;"`
	Ticket        Ticket    `gorm:"foreignKey:IdTicket;constraint:OnDelete:RESTRICT;"` // Relasi ke ticket

}

type AssignTicket struct {
	IdAssignTicket int       `gorm:"type:int;primaryKey;autoIncrement;column:id_assign_ticket"` // Kolom id_assign_ticket sebagai primary key dan auto increment
	IdTicket       int       `gorm:"type:int;index;column:id_ticket;not null"`                  // Foreign Key yang mengarah ke User.id_user
	IdPriority     int       `gorm:"type:int;index;column:id_priority;not null"`                // Foreign Key yang mengarah ke Kategoru.id_prioritas
	IdTeknisi      int       `gorm:"type:int;column:id_teknisi;not null"`                       // Foreign Key yang mengarah ke Kategoru.id_prioritas
	IdAdmin        int       `gorm:"type:int;column:id_admin;not null"`                         // Foreign Key yang mengarah ke Kategoru.id_prioritas
	StartTicket    time.Time `gorm:"type:datetime;column:start_ticket"`
	CloseTicket    time.Time `gorm:"type:datetime;column:close_ticket"`
	FinishTime     time.Time `gorm:"type:varchar(200);column:finish_time"`
	CreatedAt      time.Time `gorm:"type:datetime;"`
	UpdatedAt      time.Time `gorm:"type:datetime;"`
	Priority       Priority  `gorm:"foreignKey:IdPriority;constraint:OnDelete:RESTRICT;"` // Relasi ke Kategori
	Ticket         Ticket    `gorm:"foreignKey:IdTicket;constraint:OnDelete:RESTRICT;"`   // Relasi ke ticket

}

type Solution struct {
	IdSolution     int          `gorm:"type:int;primaryKey;autoIncrement;column:id_solution"` // Kolom id_solution sebagai primary key dan auto increment
	IdAssignTicket int          `gorm:"type:int;index;column:id_assign_ticket;not null"`      // Foreign Key yang mengarah ke assign_ticket.id_assign_ticket
	Deksripsi      string       `gorm:"type:text;column:deksripsi;not null"`
	Image          string       `gorm:"type:varchar(200);column:image;not null"`
	CreatedAt      time.Time    `gorm:"type:datetime;"`
	UpdatedAt      time.Time    `gorm:"type:datetime;"`
	AssignTicket   AssignTicket `gorm:"foreignKey:IdAssignTicket;constraint:OnDelete:RESTRICT;"` // Relasi ke assign_ticket

}

type Priority struct {
	IdPriority   int       `gorm:"type:int;primaryKey;autoIncrement;column:id_priority"` // Kolom id_prioritas sebagai primary key dan auto increment
	NamePriority string    `gorm:"type:varchar(100);column:name_priority;not null"`
	Deksripsi    string    `gorm:"type:varchar(100);column:deksripsi;not null"`
	CreatedAt    time.Time `gorm:"type:datetime;"`
	UpdatedAt    time.Time `gorm:"type:datetime;"`
}
