package main
 
import (
    "os"
    "os/exec"
    "strconv"
    "gopkg.in/ini.v1"
)
 
func main() {
    phpver := os.Args[1]
    cfg, _ := ini.Load("tools/php-fpm.ini")
    spawnNum, _ := cfg.Section("php-fpm").Key("fpm").Int()
    port, _ := cfg.Section("php-fpm").Key("port").Int()
    
    for i := 0; i < spawnNum; i++ {
        th_port := port + i
        exec.Command("php/" + phpver + "/php-cgi.exe", "-b", "127.0.0.1:" + strconv.Itoa(th_port)).Start()
    }
}