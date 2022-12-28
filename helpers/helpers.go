package helpers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

func GenerateUniqueID() (string, string) {

	node, err := snowflake.NewNode(1)
	snowflake.Epoch = time.Now().UnixMilli()

	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	unique_id := node.Generate()
	short_url := unique_id.Base32()
	id := unique_id.String()
	
	return id, short_url

}
