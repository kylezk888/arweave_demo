package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sort"
	"strings"
)

//func foo() {
//	start := time.Now()
//	time.Sleep(100 * time.Millisecond)
//	elapsed := time.Since(start)
//
//	/*
//	   Report your test result as a success
//	*/
//	boomer.RecordSuccess("http", "foo", elapsed.Nanoseconds()/int64(time.Millisecond), int64(10))
//}
//
//func bar() {
//	start := time.Now()
//	time.Sleep(100 * time.Millisecond)
//	elapsed := time.Since(start)
//
//	/*
//	   Report your test result as a failure
//	*/
//	boomer.RecordFailure("udp", "bar", elapsed.Nanoseconds()/int64(time.Millisecond), "udp error")
//}

const (
	TxTypeEmpty            = iota //0
	TxTypeRegisterZns             //1
	TxTypeCreatePair              //2
	TxTypeUpdatePairRate          //3
	TxTypeDeposit                 //4
	TxTypeDepositNft              //5
	TxTypeTransfer                //6
	TxTypeSwap                    //7
	TxTypeAddLiquidity            //8
	TxTypeRemoveLiquidity         //9
	TxTypeWithdraw                //10
	TxTypeCreateCollection        //11
	TxTypeMintNft                 //12
	TxTypeTransferNft             //13
	TxTypeAtomicMatch             //14
	TxTypeCancelOffer             //15
	TxTypeWithdrawNft             //16
	TxTypeFullExit                //17
	TxTypeFullExitNft             //18
	TxTypeOffer                   //19
)

func Foo() {
	fmt.Printf("我是 %s, %s 在调用我!\n", printMyName(), printCallerName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, %s 又在调用我!\n", printMyName(), printCallerName())
}
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

type Test struct {
	name string
	age  int64
}

func demo99() (res []string) {
	i := 1
	fmt.Println("i", i)
	return res
}

type Test1 struct {
	n string
}

var name = (*Test1)(nil)

//SlashContractAddr
//rpcClient.SetTimeOut(60 * time.Second)

//	for i := 0; i < 310; i++ {
//		proposal, err := s.BCNormalAccount[0].RpcClient.GetSideChainProposal(id, s.Config.BSCChainid)
//		if err == nil && proposal.GetStatus().String() == "Executed" {
//			log.Infof("proposalInfo: %s, %v", proposal.GetStatus().String(), proposal)
//			return nil
//		}
//		time.Sleep(time.Second * 1)
//	}
//
// return fmt.Errorf("proposal timeout,proposal id:%d", id)

func ToIntByPrecise(str string, precise uint64) *big.Int {
	result := new(big.Int)
	splits := strings.Split(str, ".")
	if len(splits) == 1 { // doesn't contain "."
		var i uint64 = 0
		for ; i < precise; i++ {
			str += "0"
		}
		intValue, ok := new(big.Int).SetString(str, 10)
		if ok {
			result.Set(intValue)
		}
	} else if len(splits) == 2 {
		value := new(big.Int)
		ok := false
		floatLen := uint64(len(splits[1]))
		if floatLen <= precise { // add "0" at last of str
			parseString := strings.Replace(str, ".", "", 1)
			var i uint64 = 0
			for ; i < precise-floatLen; i++ {
				parseString += "0"
			}
			value, ok = value.SetString(parseString, 10)
		} else { // remove redundant digits after "."
			splits[1] = splits[1][:precise]
			parseString := splits[0] + splits[1]
			value, ok = value.SetString(parseString, 10)
		}
		if ok {
			result.Set(value)
		}
	}
	return result
}

func getStats(timestamps []int) {
	sort.Ints(timestamps)

	n := len(timestamps)
	p90Index := int(float64(n) * 0.9)
	p90 := timestamps[p90Index]

	var sum int
	var max, min int

	for _, value := range timestamps {
		sum += value
		if value > max {
			max = value
		}
		if value < min || min == 0 {
			min = value
		}
	}

	avg := sum / n

	fmt.Println("p90:", p90)
	fmt.Println("Avg:", avg)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}

func main() {
	var timestamps []int
	for i := 100; i > 1; i-- {
		timestamps = append(timestamps, i)
	}
	getStats(timestamps)
	//	url := "https://op-bnb-testnet-opensearch.nodereal.io/_dashboards/api/console/proxy?path=transfer-721-opbnb/_search&method=GET"
	//
	//	// 创建请求体
	//	requestBody := `{
	//  "_source": {
	//    "includes": ["721TokenId"]
	//  },
	//  "track_total_hits": true,
	//  "query":{
	//    "range": {
	//      "blockNumber": {
	//        "gte": 0,
	//        "lte": 5611093
	//      }
	//    }
	//  },
	//  "collapse": {
	//    "field": "contractAddress.keyword"
	//  },
	//  "size": 10000
	//}`
	//
	//	req, err := http.NewRequest("POST", url, bytes.NewBufferString(requestBody))
	//	if err != nil {
	//		fmt.Println("Error creating request:", err)
	//		return
	//	}
	//
	//	// 设置请求头
	//	req.Header.Set("Content-Type", "application/json")
	//	req.Header.Set("Osd-Xsrf", "opensearchDashboards")
	//	req.Header.Set("Authority", "op-bnb-testnet-opensearch.nodereal.io")
	//
	//	// 设置Cookie
	//	cookie := &http.Cookie{
	//		Name:  "security_authentication",
	//		Value: "Fe26.2**7085b9309a8dfb339493b3952950998e75660d1d94c53b28a91db319d5c843ad*xYzW-JReYB3vmpi26lRkVw*AOgGUEMTTM9qFhstGBTWwvt0XprEpls9Hqrgu1mm3D96k_-5jzrAYitwEcicZM5wWT5sZLl7En5R0dOZskFx_RXk33lrv4dB7Mny72W2D8dwShOpHSo2J8Yp9UDtqpdsEdZ_DT5AeGig-bikt1pQKyi1J5FkZjwXnqkLbV0btohj2ubv_0tRki8284bDWzfY0w2jmOy6bvpq28HrS5HHeD9F3x71ssC9T3m-UG5QKKriccPm6PiLHkiHUA4Nc6VDYkGqrEx5mSZs_b7Mx7F0VV5xsnak-CrJU6At6sLHWQM**1c0d2a94c4251c3ab6a065675fa0cc1f8450aa4126ec22ff2f768c78ec1fe345*Sy1XjmEfSnSHY_G0poUW861XlIDfhVriuu95s31XcOE",
	//	}
	//	req.AddCookie(cookie)
	//
	//	// 发送请求
	//	client := &http.Client{}
	//	resp, err := client.Do(req)
	//	if err != nil {
	//		fmt.Println("Error sending request:", err)
	//		return
	//	}
	//	defer resp.Body.Close()
	//
	//	// 读取响应主体
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		fmt.Println("Error reading response body:", err)
	//		return
	//	}
	//
	//	// 将响应主体转换为字符串
	//	bodyString := string(body)
	//
	//	fmt.Println("resp:", bodyString)
}
