package EntireInstanceModels

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/models/OrganizationModels"
	"time"
)

type EntireInstanceUseModel struct {
	models.BaseModel
	EntireInstanceIdentityCode                       string                                                      `gorm:"type:VARCHAR(20);COMMENT:所属器材;" json:"entire_instance_identity_code"`
	EntireInstance                                   EntireInstanceModel                                         `gorm:"foreignKey:EntireInstanceIdentityCode;references:IdentityCode;COMMENT:所属器材;" json:"entire_instance"`
	InAt                                             time.Time                                                   `gorm:"type:DATETIME;COMMENT:入所时间;" json:"in_at"`
	PrevInAt                                         time.Time                                                   `gorm:"type:DATETIME;COMMENT:上一次入所时间;" json:"prev_in_at"`
	OutAt                                            time.Time                                                   `gorm:"type:DATETIME;COMMENT:出所时间;" json:"out_at"`
	PrevOutAt                                        time.Time                                                   `gorm:"type:DATETIME;COMMENT:上一次出所时间;" json:"prev_out_at"`
	OrganizationLineUUID                             string                                                      `gorm:"type:CHAR(36);COMMENT:所属线别代码;" json:"organization_line_unique_code"`
	OrganizationLine                                 OrganizationModels.OrganizationLineModel                    `gorm:"foreignKey:OrganizationLineUUID;references:UUID;COMMENT:所属线别;" json:"organization_line"`
	PrevOrganizationLineUUID                         string                                                      `gorm:"type:CHAR(36);COMMENT:上一次线别代码;" json:"prev_organization_line_uuid"`
	PrevOrganizationLine                             OrganizationModels.OrganizationLineModel                    `gorm:"foreignKey:PrevOrganizationLineUUID;references:UUID;COMMENT:上一次所属线别;" json:"prev_organization_line"`
	OrganizationWorkshopUUID                         string                                                      `gorm:"type:CHAR(36);COMMENT:所属车间代码;" json:"organization_workshop_uuid"`
	OrganizationWorkshop                             OrganizationModels.OrganizationWorkshopModel                `gorm:"foreignKey:OrganizationWorkshopUUID;references:UUID;COMMENT:所属车间;" json:"organization_workshop"`
	PrevOrganizationWorkshopUUID                     string                                                 `gorm:"type:CHAR(36);COMMENT:上一次所属车间代码;" json:"prev_organization_workshop_uuid"`
	PrevOrganizationWorkshop                         OrganizationModels.OrganizationWorkshopModel                `gorm:"foreignKey:PrevOrganizationWorkshopUUID;references:UUID;COMMENT:上一次所属车间;" json:"prev_organization_workshop"`
	OrganizationStationUUID                          string                                                      `gorm:"type:CHAR(36);COMMENT:所属车站代码;" json:"organization_station_uuid"`
	OrganizationStation                              OrganizationModels.OrganizationStationModel                 `gorm:"foreignKey:OrganizationStationUUID;references:UUID;COMMENT:所属车间;" json:"organization_station"`
	PrevOrganizationStationUUID                      string                                                      `gorm:"type:CHAR(36);COMMENT:上一次所属车站代码;" json:"prev_organization_station_uuid"`
	PrevOrganizationStation                          OrganizationModels.OrganizationStationModel                 `gorm:"foreignKey:PrevOrganizationWorkshopUUID;references:UUID;COMMENT:上一次所属车间;" json:"prev_organization_station"`
	OrganizationWorkAreaUUID                         string                                                      `gorm:"type:CHAR(8);COMMENT:所属现场工区代码;" json:"organization_work_area_uuid"`
	OrganizationWorkArea                             OrganizationModels.OrganizationWorkAreaModel                `gorm:"foreignKey:OrganizationWorkAreaUUID;references:UUID;COMMENT:所属工区;" json:"organization_work_area"`
	PrevOrganizationWorkAreaUUID                     string                                                      `gorm:"type:CHAR(8);COMMENT:上一次所属现场工区代码;" json:"prev_organization_work_area_uuid"`
	PrevOrganizationWorkArea                         OrganizationModels.OrganizationWorkAreaModel                `gorm:"foreignKey:PrevOrganizationWorkAreaUUID;references:UUID;COMMENT:上一次所属工区;" json:"prev_organization_work_area"`
	LocationInstallPositionUUID                      string                                                      `gorm:"type:VARCHAR(64);COMMENT:所属室内上道位置代码;" json:"location_install_position_uuid"`
	LocationInstallPosition                          models.LocationIndoorCellModel                              `gorm:"foreignKey:LocationInstallPositionUUID;references:UUID;COMMENT:所属室内上道位置;" json:"location_install_position"`
	PrevLocationInstallPositionUUID                  string                                                      `gorm:"type:VARCHAR(64);COMMENT:上一次所属室内上道位置代码;" json:"prev_location_install_position_uuid"`
	PrevLocationInstallPosition                      models.LocationIndoorCellModel                              `gorm:"foreignKey:PrevLocationInstallPositionUUID;references:UUID;COMMENT:上一次所属室内上道位置;" json:"prev_location_install_position"`
	CrossroadName                                    string                                                      `gorm:"type:VARCHAR(64);COMMENT:所属道岔号名称;" json:"crossroad_name"`
	PrevCrossroadName                                string                                                      `gorm:"type:VARCHAR(64);COMMENT:上一次所属道岔号名称;" json:"prev_crossroad_name"`
	OpenDirection                                    string                                                      `gorm:"type:VARCHAR(64);COMMENT:开向;" json:"open_direction"`
	PrevOpenDirection                                string                                                      `gorm:"type:VARCHAR(64);COMMENT:上一次开向;" json:"prev_open_direction"`
	OrganizationSectionName                          string                                                      `gorm:"type:VARCHAR(64);COMMENT:所属区间名称;" json:"organization_section_name"`
	PrevOrganizationSectionName                      string                                                      `gorm:"type:VARCHAR(64);COMMENT:上一次所属区间名称;" json:"prev_organization_section_name"`
	OrganizationSection                              OrganizationModels.OrganizationSectionModel                 `gorm:"foreignKey:OrganizationSectionUniqueCode;references:UUID;COMMENT:所属区间;" json:"organization_section"`
	OrganizationSectionUniqueCode                    string                                                      `gorm:"type:CHAR(6);COMMENT:所属区间代码;" json:"organization_section_uuid"`
	PrevOrganizationSection                          OrganizationModels.OrganizationSectionModel                 `gorm:"foreignKey:PrevOrganizationSectionUniqueCode;references:UUID;COMMENT:上一次所属区间;" json:"prev_organization_section"`
	PrevOrganizationSectionUniqueCode                string                                                      `gorm:"type:CHAR(6);COMMENT:上一次所属区间代码;" json:"prev_organization_section_uuid"`
	SendOrReceive                                    string                                                      `gorm:"type:VARCHAR(64);COMMENT:送/受端;" json:"send_or_receive"`
	PrevSendOrReceive                                string                                                      `gorm:"type:VARCHAR(64);COMMENT:上一次送/受端;" json:"prev_send_or_receive"`
	LocationSignalPostMainOrIndicatorUUID            string                                                      `gorm:"type:CHAR(6);COMMENT:信号机主机或表示器;" json:"location_signal_post_main_or_indicator_uuid"`
	LocationSignalPostMainOrIndicator                models.LocationOutdoorSignalPostMainOrIndicatorModel        `gorm:"foreignKey:LocationSignalPostMainOrIndicatorUUID;references:UUID;COMMENT:所属信号灯主机或表示器;" json:"location_signal_post_main_or_indicator"`
	PrevLocationSignalPostMainOrIndicatorUUID        string                                                      `gorm:"type:CHAR(6);COMMENT:上一次信号机主机或表示器;" json:"prev_location_signal_post_main_or_indicator_uuid"`
	PrevLocationSignalPostMainOrIndicator            models.LocationOutdoorSignalPostMainOrIndicatorModel        `gorm:"foreignKey:PrevLocationSignalPostMainOrIndicatorUUID;references:UUID;COMMENT:上一次所属信号灯主机或表示器;" json:"prev_location_signal_post_main_or_indicator"`
	LocationSignalPostMainLightPositionUUID          string                                                      `gorm:"type:CHAR(2);COMMENT:信号机主机灯位代码;" json:"location_signal_post_main_light_position_uuid"`
	LocationSignalPostMainLightPosition              models.LocationOutdoorSignalPostMainLightPositionModel      `gorm:"foreignKey:LocationSignalPostMainLightPositionUUID;references:UUID;COMMENT:所属信号灯主机灯位代码;" json:"location_signal_post_main_light_position"`
	PrevLocationSignalPostMainLightPositionUUID      string                                                      `gorm:"type:CHAR(2);COMMENT:上一次信号机主机灯位代码;" json:"prev_location_signal_post_main_light_position_uuid"`
	PrevLocationSignalPostMainLightPosition          models.LocationOutdoorSignalPostMainLightPositionModel      `gorm:"foreignKey:PrevLocationSignalPostMainLightPositionUUID;references:UUID;COMMENT:上一次所属信号灯主机灯位代码;" json:"prev_location_signal_post_main_light_position"`
	LocationSignalPostIndicatorLightPositionUUID     string                                                      `gorm:"type:CHAR(2);COMMENT:信号机表示器灯位代码;" json:"location_signal_post_indicator_light_position_uuid"`
	LocationSignalPostIndicatorLightPosition         models.LocationOutdoorSignalPostIndicatorLightPositionModel `gorm:"foreignKey:LocationSignalPostIndicatorLightPositionUUID;references:UUID;COMMENT:所属信号灯表示器灯位代码;" json:"location_signal_post_indicator_light_position"`
	PrevLocationSignalPostIndicatorLightPositionUUID string                                                      `gorm:"type:CHAR(2);COMMENT:上一次信号机表示器灯位代码;" json:"prev_location_signal_post_indicator_light_position_uuid"`
	PrevLocationSignalPostIndicatorLightPosition     models.LocationOutdoorSignalPostIndicatorLightPositionModel `gorm:"foreignKey:PrevLocationSignalPostIndicatorLightPositionUUID;references:UUID;COMMENT:上一次所属信号灯表示器灯位代码;" json:"prev_location_signal_post_indicator_light_position"`
	OrganizationRailroadGradeCrossUUID               string                                                      `gorm:"type:CHAR(5);COMMENT:道口代码;" json:"organization_railroad_grade_cross_uuid"`
	OrganizationRailroadGradeCross                   OrganizationModels.OrganizationRailroadGradeCrossModel      `gorm:"foreignKey:OrganizationRailroadGradeCrossUUID;references:UUID;COMMENT:所属道口;" json:"organization_railroad_grade_cross"`
	PrevOrganizationRailroadGradeCrossUUID           string                                                      `gorm:"type:CHAR(5);COMMENT:上一次道口代码;" json:"prev_organization_railroad_grade_cross_uuid"`
	PrevOrganizationRailroadGradeCross               OrganizationModels.OrganizationRailroadGradeCrossModel      `gorm:"foreignKey:PrevOrganizationRailroadGradeCrossUUID;references:UUID;COMMENT:上一次所属道口;" json:"prev_organization_railroad_grade_cross"`
	OrganizationCenterUUID                           string                                                      `gorm:"type:CHAR(6);COMMENT:中心代码;" json:"organization_center_uuid"`
	OrganizationCenter                               OrganizationModels.OrganizationCenterModel                  `gorm:"foreignKey:OrganizationCenterUUID;references:UUID;COMMENT:所属中心;" json:"organization_center"`
	PrevOrganizationCenterUUID                       string                                                      `gorm:"type:CHAR(6);COMMENT:上一次中心代码;" json:"prev_organization_center_uuid"`
	PrevOrganizationCenter                           OrganizationModels.OrganizationCenterModel                  `gorm:"foreignKey:PrevOrganizationCenterUUID;references:UUID;COMMENT:上一次所属中心;" json:"prev_organization_center"`
}

// TableName 表名称
func (cls *EntireInstanceUseModel) TableName() string {
	return "entire_instance_uses"
}
