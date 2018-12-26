package logger
import (
	"log"
	"os"
	"io/ioutil"
	"io"
	"time"
	"strings"
	"errors"
)
var (
	trace   *log.Logger // 记录所有日志
	info    *log.Logger // 重要的信息
	warning *log.Logger // 需要注意的信息
	errorlog   *log.Logger // 致命错误
	fileName    string
	LogDir   string
	hasInit bool   //是否已初始化
)

func init(){
	hasInit = false;
	fileName = time.Now().Format("2006-01-02-15")
	initLog();
}

func initLog() {
	if LogDir == "" {
		//info = log.New(os.Stdout, "Info: ", log.Ltime|log.Lshortfile)
		//warning = log.New(os.Stdout, "Warning: ", log.Ltime|log.Lshortfile)
		//errorlog = log.New(os.Stdout,  "Error", log.Ltime|log.Lshortfile)
		return ;
	}
	log.Println("init log setting")
	if !strings.HasSuffix(LogDir, "/"){
		LogDir =LogDir+"/"
	}
	file, err := os.OpenFile(LogDir+"logger-"+fileName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	trace = log.New(ioutil.Discard, "TRACE: ", log.Ltime|log.Lshortfile)
	//info = log.New(os.Stdout, "Info: ", log.Ltime|log.Lshortfile)
	//warning = log.New(os.Stdout, "Warning: ", log.Ltime|log.Lshortfile)
	info = log.New(io.MultiWriter(file, os.Stderr), "Info: ", log.Ltime|log.Lshortfile)
	warning = log.New(io.MultiWriter(file, os.Stderr), "Warning: ", log.Ltime|log.Lshortfile)
	errorlog = log.New(io.MultiWriter(file, os.Stderr),  "Error", log.Ltime|log.Lshortfile)
	hasInit = true;
	//file.Fd()
}

func main() {
	for {
		Trace("I have something standard to say")
		Info("Special Information")
		Warning("There is something you need to know about")
		Error("Something has failed")
		time.Sleep(time.Duration(10)*time.Second)
	}

}


func checkLog() error{
	//若当前文件名 与
	currentFileName := time.Now().Format("2006-01-02-15")//1月2日3时4分5秒6年)
	if(currentFileName != fileName){
		fileName = currentFileName
		initLog();
	}

	//如果无log目录，则报错
	if LogDir == "" {
		return errors.New("无日志目录")
	} else if hasInit == false {//如果 日志未初始化，并且 logDir有值，则初始化
		initLog();
	}
	return nil
}


func Trace(content string){
	trace.Println(content)
}

func Info(content string) error{
	err := checkLog()
	if err != nil {
		return err;
	}
	info.Println(content)
	return nil
}
func Warning(content string) error{
	err := checkLog()
	if err != nil {
		return err;
	}
	warning.Println(content)
	return nil
}

func Error(content string) error {
	err := checkLog()
	println(err,LogDir)
	if err != nil {
		return err;
	}
	errorlog.Println(content)
	return nil
}
