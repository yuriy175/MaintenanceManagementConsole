package Interfaces

import (
	"time"

	"../Models"
)

type IDalService interface {
	Start()
	////
	GetStudiesInWork(equipName string, startDate time.Time, endDate time.Time) []Models.StudyInWorkModel
	GetSystemInfo(equipName string, startDate time.Time, endDate time.Time) *Models.FullSystemInfoModel
	GetOrganAutoInfo(equipName string, startDate time.Time, endDate time.Time) []Models.OrganAutoInfoModel
	GetGeneratorInfo(equipName string, startDate time.Time, endDate time.Time) []Models.GeneratorInfoModel
	GetSoftwareInfo(equipName string, startDate time.Time, endDate time.Time) *Models.FullSoftwareInfoModel
	GetDicomInfo(equipName string, startDate time.Time, endDate time.Time) []Models.DicomsInfoModel
	GetStandInfo(equipName string, startDate time.Time, endDate time.Time) []Models.StandInfoModel

	GetPermanentSystemInfo(equipName string) *Models.SystemInfoModel
	GetPermanentSoftwareInfo(equipName string) *Models.SoftwareInfoModel

	//user repository
	UpdateUser(user *Models.UserViewModel) *Models.UserModel
	GetUsers() []Models.UserModel
	GetUserByName(surname string, email string, password string) *Models.UserModel
}
