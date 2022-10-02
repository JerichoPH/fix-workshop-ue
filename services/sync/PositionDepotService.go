package sync

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionDepotService struct{}

// LocationForm 库房位置表单
type LocationForm struct {
	PositionDepotStorehousesForm []PositionDepotStorehouseForm `json:"storehouses"`
	PositionDepotSectionsForm    []PositionDepotSectionForm    `json:"areas"`
	PositionDepotRowsForm        []PositionDepotRowForm        `json:"platoons"`
	PositionDepotCabinetsForm    []PositionDepotCabinetForm    `json:"shelves"`
	PositionDepotTiersForm       []PositionDepotTierForm       `json:"tiers"`
	PositionDepotCellsForm       []PositionDepotCellForm       `json:"positions"`
}

// PositionDepotStorehouseForm 仓库表单
type PositionDepotStorehouseForm struct {
	CreatedAt                       string `json:"created_at"`
	UpdatedAt                       string `json:"updated_at"`
	UniqueCode                      string `json:"unique_code"`
	Name                            string `json:"name"`
	OrganizationWorkshopUniqueCode  string `json:"workshop_unique_code"`
	OrganizationWorkshop            models.OrganizationWorkshopModel
	OrganizationWorkAreaUniqueCode  string `json:"work_area_unique_code"`
	organizationWorkAreaUuid        string
	OrganizationWorkArea            models.OrganizationWorkAreaModel
	OrganizationParagraphUniqueCode string `json:"paragraph_unique_code"`
	OrganizationParagraph           models.OrganizationParagraphModel
}

// CheckBind 检查表单
func (cls PositionDepotStorehouseForm) CheckBind() PositionDepotStorehouseForm {
	var ret *gorm.DB

	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库表单验证错误：代码不能为空")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库表单验证错误：名称不能为空")
	}
	if cls.OrganizationWorkshopUniqueCode == "" {
		wrongs.PanicValidate("仓库表单验证错误：所属车间代码不能为空")
	}
	ret = models.BootByModel(models.OrganizationWorkshopModel{}).SetWheres(tools.Map{"unique_code": cls.OrganizationWorkshopUniqueCode}).PrepareByDefaultDbDriver().First(&cls.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "仓库表单验证错误：所属车间")
	if cls.OrganizationWorkAreaUniqueCode != "" {
		ret = models.BootByModel(models.OrganizationWorkAreaModel{}).SetWheres(tools.Map{"unique_code": cls.OrganizationWorkAreaUniqueCode}).PrepareByDefaultDbDriver().First(&cls.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "仓库表单验证错误：所属工区")
		cls.organizationWorkAreaUuid = cls.OrganizationWorkArea.Uuid
	}
	if cls.OrganizationParagraphUniqueCode == "" {
		wrongs.PanicValidate("仓库表单验证错误：段代码不能为空")
	}
	ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"unique_code": cls.OrganizationParagraphUniqueCode}).PrepareByDefaultDbDriver().First(&cls.OrganizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "仓库表单验证错误：所属段")

	return cls
}

// PositionDepotSectionForm 仓库区域表单
type PositionDepotSectionForm struct {
	CreatedAt                         string `json:"created_at"`
	UpdatedAt                         string `json:"updated_at"`
	UniqueCode                        string `json:"unique_code"`
	Name                              string `json:"name"`
	PositionDepotStorehouseUniqueCode string `json:"storehouse_unique_code"`
	PositionDepotStorehouse           models.PositionDepotStorehouseModel
	OrganizationParagraphUniqueCode   string `json:"paragraph_unique_code"`
	OrganizationParagraph             models.OrganizationParagraphModel
}

// CheckBind 检查绑定
func (cls PositionDepotSectionForm) CheckBind() PositionDepotSectionForm {
	var ret *gorm.DB

	if cls.UniqueCode == "" {
		wrongs.PanicValidate("仓库区域表单验证错误：代码不能为空")
	}
	if cls.Name == "" {
		wrongs.PanicValidate("仓库区域表单验证错误：名称不能为空")
	}
	if cls.PositionDepotStorehouseUniqueCode == "" {
		wrongs.PanicValidate("仓库区域表单验证错误：所属仓库代码不能为空")
	}
	ret = models.BootByModel(models.PositionDepotSectionModel{}).SetWheres(tools.Map{"unique_code": cls.PositionDepotStorehouseUniqueCode}).PrepareByDefaultDbDriver().First(&cls.PositionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "仓库区域表单验证错误：所属仓库")
	if cls.OrganizationParagraphUniqueCode != "" {
		wrongs.PanicValidate("仓库区域表单验证错误：所属段代码不能为空")
	}
	ret = models.BootByModel(models.OrganizationParagraphModel{}).SetWheres(tools.Map{"unique_code": cls.OrganizationParagraphUniqueCode}).PrepareByDefaultDbDriver().First(&cls.OrganizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "仓库区域表单验证错误：所属段")

	return cls
}

type PositionDepotRowForm struct {
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	UniqueCode          string `json:"unique_code"`
	Name                string `json:"name"`
	AreaUniqueCode      string `json:"area_unique_code"`
	ParagraphUniqueCode string `json:"paragraph_unique_code"`
}

type PositionDepotCabinetForm struct {
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	UniqueCode          string `json:"unique_code"`
	Name                string `json:"name"`
	PlatoonUniqueCode   string `json:"platoon_unique_code"`
	ParagraphUniqueCode string `json:"paragraph_unique_code"`
}

type PositionDepotTierForm struct {
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	UniqueCode          string `json:"unique_code"`
	Name                string `json:"name"`
	ShelfUniqueCode     string `json:"shelf_unique_code"`
	ParagraphUniqueCode string `json:"paragraph_unique_code"`
}

type PositionDepotCellForm struct {
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	UniqueCode          string `json:"unique_code"`
	Name                string `json:"name"`
	TierUniqueCode      string `json:"tier_unique_code"`
	ParagraphUniqueCode string `json:"paragraph_unique_code"`
}

// FromParagraphCenter 段中心 → 同步仓库数据
func (PositionDepotService) FromParagraphCenter(ctx *gin.Context) {
	locationForm := new(LocationForm)
	if err := ctx.ShouldBindJSON(&locationForm); err != nil {
		wrongs.PanicForbidden("数据格式不正确")
	}

	// 处理仓库部分
	if len(locationForm.PositionDepotStorehousesForm) > 0 {
		for _, positionDepotStorehouseForm := range locationForm.PositionDepotStorehousesForm {
			positionDepotStorehouseForm.CheckBind() // 检查绑定

			// 声明变量
			var (
				ret                     *gorm.DB
				positionDepotStorehouse models.PositionDepotStorehouseModel
			)

			// 检查仓库是否存在
			ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": positionDepotStorehouse.UniqueCode}).
				PrepareByDefaultDbDriver().
				First(&positionDepotStorehouse)
			if !wrongs.PanicWhenIsEmpty(ret, "") {
				// 新建仓库
				models.BootByModel(models.PositionDepotStorehouseModel{}).PrepareByDefaultDbDriver().Create(&models.PositionDepotStorehouseModel{
					BaseModel:                models.BaseModel{Uuid: uuid.NewV4().String(), Sort: 0},
					UniqueCode:               positionDepotStorehouseForm.UniqueCode,
					Name:                     positionDepotStorehouseForm.Name,
					OrganizationWorkshopUuid: positionDepotStorehouseForm.OrganizationWorkshop.Uuid,
					OrganizationWorkAreaUuid: positionDepotStorehouseForm.organizationWorkAreaUuid,
				})
			} else {
				// 编辑仓库数据s
				positionDepotStorehouse.Name = positionDepotStorehouseForm.Name
				positionDepotStorehouse.OrganizationWorkshopUuid = positionDepotStorehouseForm.OrganizationWorkshop.Uuid
				positionDepotStorehouse.OrganizationWorkAreaUuid = positionDepotStorehouseForm.organizationWorkAreaUuid
				models.BootByModel(models.PositionDepotStorehouseModel{}).SetWheres(tools.Map{"uuid": positionDepotStorehouse.Uuid}).PrepareByDefaultDbDriver().Save(&positionDepotStorehouse)
			}
		}
	}

	// 处理仓库区域部分
	if len(locationForm.PositionDepotSectionsForm) > 0 {
		for _, positionDepotSectionForm := range locationForm.PositionDepotSectionsForm {
			positionDepotSectionForm.CheckBind() // 检查表单

			// 声明变量
			var (
				ret                  *gorm.DB
				positionDepotSection models.PositionDepotSectionModel
			)

			ret = models.BootByModel(models.PositionDepotSectionModel{}).SetWheres(tools.Map{"unique_code": positionDepotSectionForm.UniqueCode}).PrepareByDefaultDbDriver().First(&positionDepotSection)
			if wrongs.PanicWhenIsEmpty(ret, "") {
				// 新建仓库区域
				models.BootByModel(models.PositionDepotSectionModel{}).PrepareByDefaultDbDriver().Create(&models.PositionDepotSectionModel{
					BaseModel:                   models.BaseModel{Uuid: uuid.NewV4().String(), Sort: 0},
					UniqueCode:                  positionDepotSectionForm.UniqueCode,
					Name:                        positionDepotSectionForm.Name,
					PositionDepotStorehouseUuid: positionDepotSectionForm.PositionDepotStorehouse.Uuid,
				})
			} else {
				// 编辑仓库区域
				positionDepotSection.Name = positionDepotSectionForm.Name
				positionDepotSection.PositionDepotStorehouseUuid = positionDepotSectionForm.PositionDepotStorehouse.Uuid
				models.BootByModel(models.PositionDepotSectionModel{}).SetWheres(tools.Map{"uuid": positionDepotSection.Uuid}).PrepareByDefaultDbDriver().Save(&positionDepotSection)
			}
		}
	}

	// 处理仓库排部分
	//if len(locationForm.PositionDepotRowsForm) > 0 {
	//	for _, positionDepotRowForm := range locationForm.PositionDepotRowsForm {
	//		positionDepotRowForm.CheckBind()
	//	}
	//}
}
