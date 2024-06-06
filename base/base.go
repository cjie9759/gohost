package base

import (
	_ "embed"
	"flag"
	"sync"
	"time"

	"github.com/cjie9759/notify"
	"gorm.io/gorm"
)

var HostData *sync.Map

// var HostDataLock = new(sync.RWMutex)

// var HostData = make(syncmap[string][]HostInfo)

var (
	Is_server bool
	Is_user   bool
	Listen    Strs
	Name      string
	LosTime   time.Duration
	IsTest    bool

	Notifys *notify.NotifyGrop
	DB      *gorm.DB
	Uptime  time.Time
)

func Init() {
	if !IsTest {
		Is_server = false
		Is_user = false
		Listen = Strs{":80"}
		LosTime = 120

		flag.BoolVar(&Is_server, "s", false, "server")
		flag.BoolVar(&Is_user, "u", false, "getdata")
		flag.StringVar(&Name, "n", "", "node name")
		flag.Var(&Listen, "l", "listen addr")
		flag.DurationVar(&LosTime, "t", time.Minute, "Lost Time for alert")
		flag.Parse()
		Listen = Listen[1:]
	}
	if Is_server {
		Uptime = time.Now()

		// Notifys = notify.NewNotifyGrop([]notify.Notify{
		// 	wxrobot.NewNotify(wxrobot.Msgtype_text, webhook),
		// 	mail.NewMail(mail.Cfg{User: MAIL_USER, Pwd: MAIL_PWD, From: MAIL_FROM, To: []string{MAIL_TEST_TO}, Sub: "gohost"}),
		// 	cqrobot.NewNotify(CQ_GROUP_ID, CQ_URL),
		// })

		// var err error
		// DB, err = gorm.Open(sqlite.Open(dbdsn), &gorm.Config{
		// 	Logger: logger.New(
		// 		log.New(os.Stdout, "\r\n", log.LstdFlags),
		// 		logger.Config{
		// 			SlowThreshold:             time.Second / 5, // Slow SQL threshold
		// 			LogLevel:                  logger.Info,     // Log level
		// 			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
		// 			// ParameterizedQueries:      true,          // Don't include params in the SQL log
		// 			Colorful: true, // Disable color
		// 		})})
		// if err != nil {
		// 	log.Panic("db connect fail:", err)
		// }

		// err = DB.AutoMigrate(&HostInfo{})
		// if err != nil {
		// 	log.Panic("db connect fail:", err)
		// }
	}
}

// # 生成私钥
// openssl genrsa -out server.key 2048
// openssl genpkey -algorithm ED25519 -out server.key
// # 生成证书
// openssl req -new -x509 -key server.key -out server.crt -days 3650
// # 只读权限
// chmod 400 server.key
// openssl genrsa -out server.key 2048 &&openssl req -new -x509 -key server.key -out server.crt -days 3650
// openssl genrsa -out client.key 2048 &&openssl req -new -x509 -key client.key -out client.crt -days 3650

// //go:embed pem/fullchain.pem
// var Cert []byte

// //go:embed pem/privkey.pem
// var Key []byte

// / go:embed pem/client.crt
var CCert []byte

// /go:embed pem/client.key
var CKey []byte

// /go:embed pem/server.crt
var SCert []byte

// /go:embed pem/server.key
var SKey []byte
