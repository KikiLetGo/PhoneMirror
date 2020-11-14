package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
)
func main() {
	//http.HandleFunc("/",myResponse)
	router := gin.Default()
	router.GET("/listDir",listDir)
	router.POST("/upload",uploadFile)

	router.Run(":8888")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/listDir",listDir)
	//mux.Handle("/", http.FileServer(http.Dir(dirPath+"/ui/")))
	//mux.Handle("/index", http.FileServer(http.Dir("/data/local/tmp/webservice/")))

	//http.Handle("/index", http.FileServer(http.Dir("/data/local/tmp/webservice/index.html")))
	
	//http.ListenAndServe(":8888",mux)
}
func listRootDir(path string)(string) {
	var files [] string
    fs,_:= ioutil.ReadDir(path)
    for _,file:=range fs{
        // if file.IsDir(){
        //     fmt.Println(path+file.Name())
        //     //getFileList(path+file.Name()+"/")
        // }else{
        //     fmt.Println(path+file.Name())
        // }
        //files.PushBack(path+file.Name())
        files = append(files,path+"/"+file.Name() )
    }

	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    data,_ := json.Marshal(files)
    filesJson := string(data)
    fmt.Println(filesJson)
    return filesJson
}
func listDir(c *gin.Context) {
	dir := c.Query("dir")
	//query := r.URL.Query()
	//dir := query["dir"][0]
	fmt.Println(dir)
	dirPath, _ := os.Getwd()
	_  = dirPath
	//files := listRootDir("/storage/emulated/0")
	files := listRootDir(dir)
	c.String(http.StatusOK, files)
	//w.Write([]byte(files))
}
func uploadFile(c *gin.Context) {
	file,err:=c.FormFile("file")
   

    if err!=nil{
        fmt.Println("can not fetch files",err)
        c.String(http.StatusOK,"no")
        return
    }
    fmt.Println(file.Size)
    // 检查下文件的大小
    var path=file.Filename
    ok:=c.SaveUploadedFile(file,`./upload/`+path)
    if ok!=nil{
        fmt.Println("error when save",ok)
        c.String(http.StatusOK,"n2")
    }
    c.String(http.StatusOK,"success")
}
