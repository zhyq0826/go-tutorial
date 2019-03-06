package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type statusObj struct {
	Status    int    `json:"status"`
	Timestamp int64  `json:"timestamp"`
	Token     string `json:"token"`
	RequestId string `json:"requestId"`
	Appid     string `json:"appid"`
	BigTag    string `json:"bigTag"`
}

type statusList struct {
	Statuses []statusObj `json:"statuses"`
}

func main() {
	in, err := os.Open("/Users/zhaoyongqiang/workspace/v12/huawei_push_2019-03-04.log") // For read access.
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.OpenFile("02.sql", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer in.Close()
	defer out.Close()

	scanner := bufio.NewScanner(in)
	count := 0
	writer := bufio.NewWriter(out)

	defer writer.Flush()

	for scanner.Scan() {
		count += 1
		list := statusList{}
		line := scanner.Text()
		line = strings.Replace(line, "u'", "'", -1)
		line = strings.Replace(line, "'", "\"", -1)
		err = json.Unmarshal([]byte(line), &list)
		if count%100000 == 0 {
			fmt.Println(time.Now(), count)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, obj := range list.Statuses {
			sql := fmt.Sprintf("INSERT INTO `huaweipush` (`id`, `status`, `timestamp`, `token`, `requestId`, `appid`, `biTag`) VALUES (0, %d, '%s', '%s', '%s', %s, '%s');\n",
				obj.Status, time.Unix(obj.Timestamp/1000, 0).Format("2006-01-02 15:04:05"), obj.Token, obj.RequestId, obj.Appid, obj.BigTag)
			writer.Write([]byte(sql))
		}

	}
}
