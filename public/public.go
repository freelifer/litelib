package public

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"math/big"
	"strconv"

	"fmt"
	"github.com/sony/sonyflake"
	"time"
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

// DefaultQueryForInt returns the keyed url query value if it exists
func DefaultQueryForInt(c *gin.Context, key string, defaultValue int) int {
	if value, ok := strconv.Atoi(c.Query(key)); ok == nil {
		return value
	}
	return defaultValue
}

// DefaultQueryForInt64 returns the keyed url query value if it exists
func DefaultQueryForInt64(c *gin.Context, key string, defaultValue int64) int64 {
	if value, ok := strconv.ParseInt(c.Query(key), 10, 64); ok == nil {
		return value
	}
	return defaultValue
}

// ParamFromId returns the keyed url param value if it exists
func ParamFromID(c *gin.Context, key string) (int64, error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}

func ParamFromUUID(c *gin.Context, key string) (uint64, error) {
	return strconv.ParseUint(c.Param(key), 0, 64)
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
