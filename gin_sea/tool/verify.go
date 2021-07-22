package tool

//需要验证的字段
var (
	PageInfoVerify  = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	AboutTeamVerify = Rules{"NameCn": {NotEmpty()}, "NameEn": {NotEmpty()}, "PositionCn": {NotEmpty()}, "PositionEn": {NotEmpty()}}
)
