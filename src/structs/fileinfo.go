package structs
type FileInfo struct {
	FileType string `json:"fileType"`
	Name string `json:"name"`
	Dir string `json:"dir"`
}
type FileInfoDatas struct{
	StatusCode int `json:"statusCode,int"`
	Msg string `json:"msg"`
	FileInfos []FileInfo `json:"datas,FileInfos"`

}