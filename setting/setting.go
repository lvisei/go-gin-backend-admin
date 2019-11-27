package setting

import (
    "log"
    "time"

    "gopkg.in/ini.v1"
)

var (
    Cfg *ini.File

    AppMode string

    HttpPort int
    ReadTimeout time.Duration
    WriteTimeout time.Duration

    PageSize int
    JwtSecret string
)

func init() {
    var err error
    Cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    LoadBase()
    LoadServer()
    LoadApp()
}

func LoadBase() {
		AppMode = Cfg.Section("").Key("AppMode").MustString("debug")
		fmt.Println("App Mode:", AppMode)
}

func LoadServer() {
    sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("Fail to get section 'server': %v", err)
    }

		HttpPort = sec.Key("HttpPort").MustInt(8000)
		fmt.Println("HttpPort:", HttpPort)
    ReadTimeout = time.Duration(sec.Key("ReadTimeout").MustInt(60)) * time.Second
    WriteTimeout =  time.Duration(sec.Key("WriteTimeout").MustInt(60)) * time.Second    
}

func LoadApp() {
    sec, err := Cfg.GetSection("app")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
    }

    JwtSecret = sec.Key("JwtSecret").MustString("@!$3#&^*~`")
    PageSize = sec.Key("PageSize").MustInt(10)
}