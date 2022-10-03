package sync

import (
	"fix-workshop-ue/models"
	"fix-workshop-ue/tools"
	"fix-workshop-ue/wrongs"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
)

type PositionDepotService struct{}

// PositionDepotFromParagraphCenterForm 同步库房位置表单（段中心 → 检修车间）
type PositionDepotFromParagraphCenterForm struct {
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
	OrganizationWorkArea            models.OrganizationWorkAreaModel
	OrganizationParagraphUniqueCode string `json:"paragraph_unique_code"`
	OrganizationParagraph           models.OrganizationParagraphModel
}

// CheckBind 检查表单（仓库）
func (ins PositionDepotStorehouseForm) CheckBind() PositionDepotStorehouseForm {
	var ret *gorm.DB

	if ins.UniqueCode == "" {
		wrongs.PanicValidate("仓库表单验证错误：代码不能为空")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("仓库表单验证错误：名称不能为空")
	}
	if ins.OrganizationWorkshopUniqueCode == "" {
		wrongs.PanicValidate("仓库表单验证错误：所属车间代码不能为空")
	}
	ret = models.
		BootByModel(models.OrganizationWorkshopModel{}).
		SetWheres(tools.Map{"unique_code": ins.OrganizationWorkshopUniqueCode}).
		PrepareByDefaultDbDriver().
		First(&ins.OrganizationWorkshop)
	wrongs.PanicWhenIsEmpty(ret, "仓库表单验证错误：所属车间")
	if ins.OrganizationWorkAreaUniqueCode != "" {
		ret = models.
			BootByModel(models.OrganizationWorkAreaModel{}).
			SetWheres(tools.Map{"unique_code": ins.OrganizationWorkAreaUniqueCode}).
			PrepareByDefaultDbDriver().
			First(&ins.OrganizationWorkArea)
		wrongs.PanicWhenIsEmpty(ret, "仓库表单验证错误：所属工区")
	}
	if ins.OrganizationParagraphUniqueCode == "" {
		wrongs.PanicValidate("仓库表单验证错误：段代码不能为空")
	}
	ret = models.
		BootByModel(models.OrganizationParagraphModel{}).
		SetWheres(tools.Map{"unique_code": ins.OrganizationParagraphUniqueCode}).
		PrepareByDefaultDbDriver().
		First(&ins.OrganizationParagraph)
	wrongs.PanicWhenIsEmpty(ret, "仓库表单验证错误：所属段")

	return ins
}

// PositionDepotSectionForm 仓库区域表单
type PositionDepotSectionForm struct {
	CreatedAt                         string `json:"created_at"`
	UpdatedAt                         string `json:"updated_at"`
	UniqueCode                        string `json:"unique_code"`
	Name                              string `json:"name"`
	PositionDepotStorehouseUniqueCode string `json:"storehouse_unique_code"`
	PositionDepotStorehouse           models.PositionDepotStorehouseModel
	//OrganizationParagraphUniqueCode   string `json:"paragraph_unique_code"`
	//OrganizationParagraph             models.OrganizationParagraphModel
}

// CheckBind 检查绑定（仓库区域）
func (ins PositionDepotSectionForm) CheckBind() PositionDepotSectionForm {
	var ret *gorm.DB

	if ins.UniqueCode == "" {
		wrongs.PanicValidate("仓库区域表单验证错误：代码不能为空")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("仓库区域表单验证错误：名称不能为空")
	}
	if ins.PositionDepotStorehouseUniqueCode == "" {
		wrongs.PanicValidate("仓库区域表单验证错误：所属仓库代码不能为空")
	}
	ret = models.
		BootByModel(models.PositionDepotStorehouseModel{}).
		SetWheres(tools.Map{"unique_code": ins.PositionDepotStorehouseUniqueCode}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionDepotStorehouse)
	wrongs.PanicWhenIsEmpty(ret, "仓库区域表单验证错误：所属仓库")

	return ins
}

// PositionDepotRowForm 仓库排表单
type PositionDepotRowForm struct {
	CreatedAt                      string `json:"created_at"`
	UpdatedAt                      string `json:"updated_at"`
	UniqueCode                     string `json:"unique_code"`
	Name                           string `json:"name"`
	PositionDepotSectionUniqueCode string `json:"area_unique_code"`
	PositionDepotSection           models.PositionDepotSectionModel
	PositionDepotRowTypeUniqueCode string `json:"row_type_unique_code"`
	PositionDepotRowType           models.PositionDepotRowTypeModel
}

// CheckBind 表单检查（仓库排）
func (ins PositionDepotRowForm) CheckBind() PositionDepotRowForm {
	var ret *gorm.DB

	if ins.UniqueCode == "" {
		wrongs.PanicValidate("仓库排表单验证错误：代码不能为空")
	}
	if ins.Name == "" {
		wrongs.PanicValidate("仓库排表单验证错误：名称不能为空")
	}
	if ins.PositionDepotSectionUniqueCode == "" {
		wrongs.PanicValidate("仓库排表单验证错误：所属仓库区域代码不能为空")
	}
	ret = models.
		BootByModel(models.PositionDepotSectionModel{}).
		SetWheres(tools.Map{"unique_code": ins.PositionDepotSectionUniqueCode}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionDepotSection)
	wrongs.PanicWhenIsEmpty(ret, "仓库拍表单验证错误：所属仓库区域")
	if ins.PositionDepotRowTypeUniqueCode == "" {
		wrongs.PanicValidate("仓库排表单验证错误：所属仓库排类型代码不能为空")
	}
	ret = models.
		BootByModel(models.PositionDepotRowTypeModel{}).
		SetWheres(tools.Map{"unique_code": ins.PositionDepotRowTypeUniqueCode}).
		PrepareByDefaultDbDriver().
		First(&ins.PositionDepotRowType)

	return ins
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
func (ins PositionDepotService) FromParagraphCenter(ctx *gin.Context) {
	positionDepotFromParagraphCenterForm := new(PositionDepotFromParagraphCenterForm)
	if err := ctx.ShouldBindJSON(&positionDepotFromParagraphCenterForm); err != nil {
		wrongs.PanicForbidden("数据格式不正确（只接受JSON格式）")
	}

	var (
		createdCountStorehouse uint64
		updatedCountStorehouse uint64
		createdCountSection    uint64
		updatedCountSection    uint64
		createdCountRow        uint64
		updatedCountRow        uint64
		responseStrings        []string
	)

	// 处理仓库部分
	if len(positionDepotFromParagraphCenterForm.PositionDepotStorehousesForm) > 0 {
		createdCountStorehouse, updatedCountStorehouse = ins.processPositionDepotStorehouses(positionDepotFromParagraphCenterForm)
	}

	// 处理仓库区域部分
	if len(positionDepotFromParagraphCenterForm.PositionDepotSectionsForm) > 0 {
		createdCountStorehouse, updatedCountSection = ins.processPositionDepotSections(positionDepotFromParagraphCenterForm)
	}

	// 处理仓库排部分
	if len(positionDepotFromParagraphCenterForm.PositionDepotRowsForm) > 0 {
		createdCountRow, updatedCountRow = ins.processPositionDepotRows(positionDepotFromParagraphCenterForm)
	}

	responseStrings = []string{
		fmt.Sprintf("成功新建仓库：%d", createdCountStorehouse),
		fmt.Sprintf("成功编辑仓库：%d", updatedCountStorehouse),
		fmt.Sprintf("成功新建仓库区域：%d", createdCountSection),
		fmt.Sprintf("成功编辑仓库区域：%d", updatedCountSection),
		fmt.Sprintf("成功新建仓库排：%d", createdCountRow),
		fmt.Sprintf("成功编辑仓库排：%d", updatedCountRow),
	}

	ctx.JSON(tools.CorrectBoot(strings.Join(responseStrings, ctx.DefaultQuery("sep", "\r\n"))).Ok(tools.Map{}))
}

// processPositionDepotStorehouses 处理仓库部分
func (PositionDepotService) processPositionDepotStorehouses(positionDepotFromParagraphCenterForm *PositionDepotFromParagraphCenterForm) (createdCount, updatedCount uint64) {
	createdCount = 0
	updatedCount = 0

	// 处理仓库部分
	if len(positionDepotFromParagraphCenterForm.PositionDepotStorehousesForm) > 0 {
		for _, positionDepotStorehouseForm := range positionDepotFromParagraphCenterForm.PositionDepotStorehousesForm {
			positionDepotStorehouseForm = positionDepotStorehouseForm.CheckBind() // 检查绑定

			// 声明变量
			var (
				ret                     *gorm.DB
				positionDepotStorehouse models.PositionDepotStorehouseModel
			)

			// 检查仓库是否存在
			ret = models.BootByModel(models.PositionDepotStorehouseModel{}).
				SetWheres(tools.Map{"unique_code": positionDepotStorehouseForm.UniqueCode}).
				PrepareByDefaultDbDriver().
				First(&positionDepotStorehouse)
			if !wrongs.PanicWhenIsEmpty(ret, "") {
				// 新建仓库
				positionDepotStorehouse = models.PositionDepotStorehouseModel{
					BaseModel:                models.BaseModel{Uuid: uuid.NewV4().String(), Sort: 0},
					UniqueCode:               positionDepotStorehouseForm.UniqueCode,
					Name:                     positionDepotStorehouseForm.Name,
					OrganizationWorkshopUuid: positionDepotStorehouseForm.OrganizationWorkshop.Uuid,
					OrganizationWorkAreaUuid: positionDepotStorehouseForm.OrganizationWorkArea.Uuid,
				}
				models.
					BootByModel(models.PositionDepotStorehouseModel{}).
					PrepareByDefaultDbDriver().
					Create(&positionDepotStorehouse)
				createdCount++
			} else {
				// 编辑仓库数据
				positionDepotStorehouse.Name = positionDepotStorehouseForm.Name
				positionDepotStorehouse.OrganizationWorkshopUuid = positionDepotStorehouseForm.OrganizationWorkshop.Uuid
				positionDepotStorehouse.OrganizationWorkAreaUuid = positionDepotStorehouseForm.OrganizationWorkArea.Uuid
				models.BootByModel(models.PositionDepotStorehouseModel{}).SetWheres(tools.Map{"uuid": positionDepotStorehouse.Uuid}).PrepareByDefaultDbDriver().Save(&positionDepotStorehouse)
				updatedCount++
			}
		}
	}

	return
}

// processPositionDepotSections 处理仓库区域部分
func (PositionDepotService) processPositionDepotSections(positionDepotFromParagraphCenterForm *PositionDepotFromParagraphCenterForm) (createdCount, updatedCount uint64) {
	createdCount = 0
	updatedCount = 0

	// 处理仓库部分
	if len(positionDepotFromParagraphCenterForm.PositionDepotSectionsForm) > 0 {
		for _, positionDepotSectionForm := range positionDepotFromParagraphCenterForm.PositionDepotSectionsForm {
			positionDepotSectionForm = positionDepotSectionForm.CheckBind() // 检查绑定

			// 声明变量
			var (
				ret                  *gorm.DB
				positionDepotSection models.PositionDepotSectionModel
			)

			// 检查仓库区域是否存在
			ret = models.BootByModel(models.PositionDepotSectionModel{}).
				SetWheres(tools.Map{"unique_code": positionDepotSectionForm.UniqueCode}).
				PrepareByDefaultDbDriver().
				First(&positionDepotSection)
			if !wrongs.PanicWhenIsEmpty(ret, "") {
				// 新建仓库区域
				positionDepotSection = models.PositionDepotSectionModel{
					BaseModel:                   models.BaseModel{Uuid: uuid.NewV4().String(), Sort: 0},
					UniqueCode:                  positionDepotSectionForm.UniqueCode,
					Name:                        positionDepotSectionForm.Name,
					PositionDepotStorehouseUuid: positionDepotSectionForm.PositionDepotStorehouse.Uuid,
				}
				models.
					BootByModel(models.PositionDepotSectionModel{}).
					PrepareByDefaultDbDriver().
					Create(&positionDepotSection)
				createdCount++
			} else {
				// 编辑仓库区域数据
				positionDepotSection.Name = positionDepotSectionForm.Name
				positionDepotSection.PositionDepotStorehouseUuid = positionDepotSectionForm.PositionDepotStorehouse.Uuid
				models.BootByModel(models.PositionDepotSectionModel{}).SetWheres(tools.Map{"uuid": positionDepotSection.Uuid}).PrepareByDefaultDbDriver().Save(&positionDepotSection)
				updatedCount++
			}
		}
	}

	return
}

// processPositionRowSections 处理仓库排部分
func (PositionDepotService) processPositionDepotRows(positionDepotFromParagraphCenterForm *PositionDepotFromParagraphCenterForm) (createdCount, updatedCount uint64) {
	createdCount = 0
	updatedCount = 0

	// 处理仓库部分
	if len(positionDepotFromParagraphCenterForm.PositionDepotRowsForm) > 0 {
		for _, positionDepotRowForm := range positionDepotFromParagraphCenterForm.PositionDepotRowsForm {
			positionDepotRowForm = positionDepotRowForm.CheckBind() // 检查绑定

			// 声明变量
			var (
				ret              *gorm.DB
				positionDepotRow models.PositionDepotRowModel
			)

			// 检查仓库排是否存在
			ret = models.BootByModel(models.PositionDepotRowModel{}).
				SetWheres(tools.Map{"unique_code": positionDepotRowForm.UniqueCode}).
				PrepareByDefaultDbDriver().
				Debug().
				First(&positionDepotRow)
			if !wrongs.PanicWhenIsEmpty(ret, "") {
				// 新建仓库排
				positionDepotRow = models.PositionDepotRowModel{
					BaseModel:                models.BaseModel{Uuid: uuid.NewV4().String(), Sort: 0},
					UniqueCode:               positionDepotRowForm.UniqueCode,
					Name:                     positionDepotRowForm.Name,
					PositionDepotSectionUuid: positionDepotRowForm.PositionDepotSection.Uuid,
					PositionDepotRowTypeUuid: positionDepotRowForm.PositionDepotRowType.Uuid,
				}
				models.
					BootByModel(models.PositionDepotRowModel{}).
					PrepareByDefaultDbDriver().
					Create(&positionDepotRow)
				createdCount++
			} else {
				// 编辑仓库排数据
				positionDepotRow.Name = positionDepotRowForm.Name
				positionDepotRow.PositionDepotSectionUuid = positionDepotRowForm.PositionDepotSection.Uuid
				models.BootByModel(models.PositionDepotRowModel{}).SetWheres(tools.Map{"uuid": positionDepotRow.Uuid}).PrepareByDefaultDbDriver().Save(&positionDepotRow)
				updatedCount++
			}
		}
	}

	return
}
