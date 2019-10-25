package public

import (
	"crypto/rand"
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"

	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	DB struct {
		Engine *xorm.Engine
		Tables []interface{}
	}

	Gin struct {
		g *gin.Engine
	}

	SF *sonyflake.Sonyflake
)

func init() {
	fmt.Println("-------models init")
	var st sonyflake.Settings
	st.StartTime = time.Now()

	SF = sonyflake.NewSonyflake(st)
	if SF == nil {
		panic("sonyflake not created")
	}

	// ip, _ := lower16BitPrivateIP()
	// machineID = uint64(ip)
}

func Insert(i interface{}) (err error) {
	sess := DB.Engine.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(i); err != nil {
		return err
	}

	return sess.Commit()
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandomString returns generated random string in given length of characters.
// It also returns possible error during generation.
func RandomString(n int) (string, error) {
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(alphanum)))

	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buffer[i] = alphanum[index]
	}

	return string(buffer), nil
}

func randomInt(max *big.Int) (int, error) {
	rand, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(rand.Int64()), nil
}
