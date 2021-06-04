# 网易直播/互动直播接口

```
import (
	"fmt"

	"github.com/sjmark/wyLive"
)

var live wyLive.WYLv
func init() {
	live = wyLive.InitWYLvChannel("appKey", "appSecret")
}

func createLive() {
	wyRes, err := live.CreateChannel("thisLive")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wyRes)
}

func main(){
	createLive()
}


```