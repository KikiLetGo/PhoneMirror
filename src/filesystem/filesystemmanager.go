package filesystem
import (
	"../structs"
	"encoding/json"
	"fmt"
	"os"
	"../constants"

)
func CreateMirrorDisk(fileStruct string,externalStorageDirectory string,device string)(string) {
	fmt.Println(fileStruct)
	fmt.Println("CreateMirrorDisk starting")
	

	var rootDir structs.FileDir
	json.Unmarshal([]byte(fileStruct),&rootDir)
	showDiskStructs(&rootDir,"|_")
	deviceBasePath := constants.MIRROR_DISK_BASE_PATH+device+externalStorageDirectory

	err:=os.MkdirAll(deviceBasePath, os.ModePerm) 
	fmt.Println("deviceBasePath created:"+deviceBasePath)

	if err!=nil{
		fmt.Println(err) 
		return deviceBasePath
	}

	recurrenceCreateDirs(deviceBasePath,&rootDir)
	return deviceBasePath
	
	// ndata,_:=json.Marshal(obj)
	// fmt.Println(string(ndata))
}
func showDiskStructs(fileDir *structs.FileDir,deepFlag string) {
	documents := fileDir.Documents
	documentsLen := len(documents)

	var i int
	for i = 0; i < documentsLen; i++ {
		document := documents[i]
		fmt.Println(deepFlag+document)
   	}

	dirs := fileDir.Dirs

	dirsLen := len(dirs)

   	for i = 0; i < dirsLen; i++ {
		dir := dirs[i]
		fmt.Println(deepFlag+dir.Name)
		showDiskStructs(&dir," "+deepFlag)

   	}
   	
}
func recurrenceCreateDirs(parentPath string,rootDir *structs.FileDir) {
	fmt.Println("recurrenceCreateDirs parentPath:"+parentPath)


	dirs := rootDir.Dirs
	dirsLen := len(dirs)

	var i int
   	for i = 0; i < dirsLen; i++ {
		dir := dirs[i]
		path := parentPath+"/"+dir.Name
		fmt.Println("crate dir:"+path)
		err:=os.Mkdir(path, os.ModePerm) 
		if err!=nil{
			fmt.Println(err) 
		}
		recurrenceCreateDirs(path,&dir)



   	}
	
}

