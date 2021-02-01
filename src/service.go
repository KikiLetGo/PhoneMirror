package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"./structs"
    "./utils"
    "./filesystem"
    "./constants"
    "os"
    "strings"


)
func main() {
	//http.HandleFunc("/",myResponse)
	router := gin.Default()
    router.MaxMultipartMemory = 1024*1024*1024//1G
	router.GET("/listDir",listDir)
	router.POST("/upload",uploadFile)
	router.GET("/download",downloadFile)
    router.GET("/checkMirrorDisk",checkMirrorDisk)
    router.POST("/createMirrorDisk",createMirrorDisk)

    router.GET("/readContacts",readContacts)

    router.POST("/uploadContacts",uploadContacts)

    router.GET("/listPhones",listPhones)



	router.Run(":8888")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/listDir",listDir)
	//mux.Handle("/", http.FileServer(http.Dir(dirPath+"/ui/")))
	//mux.Handle("/index", http.FileServer(http.Dir("/data/local/tmp/webservice/")))

	//http.Handle("/index", http.FileServer(http.Dir("/data/local/tmp/webservice/index.html")))
	
	//http.ListenAndServe(":8888",mux)
}

func listDir(c *gin.Context) {
	dir := c.Query("dir")
    device := c.Query("device")

	basePath := constants.MIRROR_DISK_BASE_PATH+device
    path := basePath+dir
    fmt.Println(path)

    //var files [] string
    fs,_:= ioutil.ReadDir(path)
    var fileInfoDatas structs.FileInfoDatas
    fileInfoDatas.FileInfos = []structs.FileInfo{}
    for _,file:=range fs{
        fileType := ""
        if file.IsDir(){
            fileType = "dir"
            
        }else{
            fileType = "document"
        }
        //files = append(files,path+"/"+file.Name() )
        fileInfo := structs.FileInfo{FileType:fileType,Name:file.Name(),Dir:dir}
        fileInfoDatas.FileInfos = append(fileInfoDatas.FileInfos,fileInfo)
    }
    fileInfoDatas.StatusCode = 200
    fileInfoDatas.Msg = "success"


    data,_ := json.Marshal(fileInfoDatas)
    filesJson := string(data)
    fmt.Println(filesJson)

	c.String(http.StatusOK, filesJson)
}
func checkMirrorDisk(c *gin.Context) {
    device := c.Query("device")
    exist,_:= utils.PathExists(constants.MIRROR_DISK_BASE_PATH+device)
    fmt.Println(exist)
    if(exist){
        c.String(http.StatusOK, "true")
    }else{
        c.String(http.StatusOK, "false")

    }
    //fs,_:= ioutil.ReadDir("/"+device)  
}
func createMirrorDisk(c *gin.Context) {
    fileDir := c.PostForm("fileDir")
    deviceId := c.PostForm("deviceId")
    externalStorageDirectory := c.PostForm("externalStorageDirectory")

    deviceBasePath := filesystem.CreateMirrorDisk(fileDir,externalStorageDirectory,deviceId)
    c.String(http.StatusOK,deviceBasePath)
      
}

func uploadFile(c *gin.Context) {
	file,err:=c.FormFile("file")
    path:=c.Query("path")

    if err!=nil{
        fmt.Println("can not fetch files:",err)
        c.String(http.StatusOK,"no")
        return
    }
    device := c.Query("device")
    fmt.Println("Filename:"+file.Filename)
    fmt.Println("path:"+path)
    fmt.Println("device:"+device)

    
    fmt.Println(file.Size)
    // 检查下文件的大小
    //var name=file.Filename
    savePath := constants.MIRROR_DISK_BASE_PATH+device+path
    fmt.Println("file savePath:"+savePath)


    ok:=c.SaveUploadedFile(file,savePath)
    if ok!=nil{
        fmt.Println("error when save",ok)
        c.String(http.StatusOK,"n2")
    }else{
        c.String(http.StatusOK,"save success")

    }
}
func downloadFile(c *gin.Context){
    fmt.Println("downloadFile")

	dir := c.Query("dir")
    device := c.Query("device")
    filename := c.Query("filename")


    path := constants.MIRROR_DISK_BASE_PATH+device+dir+filename
    fmt.Println("path:"+path)


	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
    c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(path)
}

func uploadContacts(c *gin.Context) {
    contacts := c.PostForm("contacts")
    device := c.PostForm("device")
    savePath := constants.MIRROR_DISK_BASE_PATH+device+"/contacts.json"
    fmt.Println("save constants json into:"+savePath)
    file,err := os.Create(savePath)
    if err != nil{
        fmt.Println(err.Error())
        return
    }
    defer file.Close()
    file.WriteString(contacts)
    c.String(http.StatusOK,"success")
      
}
func readContacts(c *gin.Context) {
    device := c.Query("device")
    savePath := constants.MIRROR_DISK_BASE_PATH+device+"/contacts.json"
    fmt.Println("read constants json from:"+savePath)
    file,err := os.Open(savePath)
    if err != nil{
        fmt.Println(err.Error())
        return
    }
    defer file.Close()
    contacts,err := ioutil.ReadAll(file)
    c.String(http.StatusOK,string(contacts))    
}

func listPhones(c *gin.Context) {
    phones := []string{}
    rd,_:= ioutil.ReadDir(constants.MIRROR_DISK_BASE_PATH)
    for _, fi := range rd{
        if fi.IsDir(){
            device := fi.Name()
            if !strings.HasPrefix(device,".") {
                fmt.Println("device:"+device)
                phones = append(phones,device)
            }
            
        }
    }
    phonesJson,_ := json.Marshal(phones)
    c.String(http.StatusOK,string(phonesJson))  

    
}
