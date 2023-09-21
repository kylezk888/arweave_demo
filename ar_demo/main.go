package main

import (
	"encoding/json"
	"fmt"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"go_ether/ar_demo/kits"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var c = kits.HttpClient{}

func CreateWallet(path string, endpoint string) (*goar.Wallet, error) {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var jsonList []map[string]string
	err = json.Unmarshal(byteValue, &jsonList)
	if err != nil {
		log.Fatal(err)
	}

	byteSlices := make([][]byte, len(jsonList))

	for i, jsonObj := range jsonList {
		byteSlice, err := json.Marshal(jsonObj)
		if err != nil {
			log.Fatal(err)
		}
		byteSlices[i] = byteSlice
	}
	wallet, err := goar.NewWallet(byteSlices[0], endpoint)
	//wallet, err := goar.NewWallet(byteSlices[1], "http://localhost:1984")
	//wallet, err := goar.NewWalletFromPath("./ar_demo/test-keyfile.json", "https://arweave.net")
	if err != nil {
		panic(err)
	}
	fmt.Println("Address:", wallet.Signer.Address)
	return wallet, err
}

func LocalMint(wallet *goar.Wallet) {

	resp, err := c.Get(fmt.Sprintf("http://localhost:1984/mint/%s/100000000000000000", wallet.Signer.Address))
	if err != nil {
		panic(err)
	}
	fmt.Println("mint resp:", resp)
	resp, err = c.Get("http://localhost:1984/mine")
	if err != nil {
		panic(err)
	}
	fmt.Println("mine block resp:", resp)
	amount, err := wallet.Client.GetWalletBalance(wallet.Signer.Address)
	if err != nil {
		panic(err)
	}
	fmt.Println("amount", amount)
}

func TestDownloadData(wallet *goar.Wallet, txId string) {
	var duration []int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("start download")
			now := time.Now()
			_, err := wallet.Client.GetTransactionDataByGateway(txId)
			//_, err := wallet.Client.GetTransactionData(txId)
			if err != nil {
				fmt.Printf("GetTransactionDataByGateway error,txId:%s,err:%+v", txId, err)
			}
			duration = append(duration, int(time.Now().Unix()-now.Unix()))
			fmt.Println("end download")
		}()
	}
	wg.Wait()
	fmt.Println("duration:", duration)
	kits.GetStats(duration)

}

func TestDownloadFromPath(wallet *goar.Wallet, path string) {
	var duration []int
	var wg sync.WaitGroup
	var downloadData kits.ParametersIterator

	loadCSV, err := kits.LoadCSV(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range loadCSV {
		split := strings.Split(r, ",")
		downloadData.Add(split[0])
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			txId := downloadData.Next().(string)
			fmt.Println("start download")
			now := time.Now()
			_, err := wallet.Client.GetTransactionDataByGateway(txId)
			//_, err := wallet.Client.GetTransactionData(txId)
			if err != nil {
				fmt.Printf("GetTransactionDataByGateway error,txId:%s,err:%+v", txId, err)
			}
			duration = append(duration, int(time.Now().Unix()-now.Unix()))
			fmt.Println("end download")
		}()
		time.Sleep(time.Second * 1)
	}
	wg.Wait()
	//fmt.Println("duration:", duration)
	kits.GetStats(duration)

}

func TestDownloadDataOne(wallet *goar.Wallet, txId string) {
	var duration []int
	fmt.Println("start download")
	now := time.Now()
	//data, err := wallet.Client.GetTransactionDataByGateway(txId)
	data, err := wallet.Client.GetTransactionData(txId)
	if err != nil {
		fmt.Printf("GetTransactionDataByGateway error,txId:%s,err:%+v", txId, err)
	}
	duration = append(duration, int(time.Now().Unix()-now.Unix()))
	sizeInBytes := len(data)
	sizeInKB := float64(sizeInBytes) / 1024
	fmt.Println("end download")
	fmt.Printf("size:%0.00f \n", sizeInKB)
	kits.GetStats(duration)
}

func TestUpload(wallet *goar.Wallet, dataSize int) {

	var wg sync.WaitGroup

	buffer := kits.GetFileByBuffer(dataSize)
	//buffer :=GetFileByBuffer(dataSize)

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//bBalance, err := wallet.Client.GetWalletBalance(wallet.Signer.Address)
			//if err != nil {
			//	fmt.Println("GetWalletBalance err,err:", err)
			//}
			startTime := time.Now()
			fmt.Println("startTime:", startTime.Unix())
			tx, err := wallet.SendData(
				buffer.Bytes(),
				[]types.Tag{
					types.Tag{
						Name:  "testSendDataSpeedUp",
						Value: fmt.Sprintf("nrtest_%s", dataSize),
					},
				})
			if err != nil {
				fmt.Println("SendData err,err:", err)
			}
			duration := time.Now().Unix() - startTime.Unix()
			//aBalance, err := wallet.Client.GetWalletBalance(wallet.Signer.Address)
			//if err != nil {
			//	fmt.Println("GetWalletBalance err,err:", err)
			//}
			//fmt.Printf("bBalance:%s, aBalance:%s\n", bBalance.String(), aBalance.String())
			//sub := bBalance.Sub(bBalance, aBalance)
			//subWinston := big.NewInt(0).Sub(utils.ARToWinston(bBalance), utils.ARToWinston(aBalance))
			//txID,startTime,endTime-startTime,ar消耗花费，winston消耗花费
			fmt.Printf("%s,%d,%d \n", tx.ID, startTime.Unix(), duration)
		}()
		time.Sleep(time.Second * 1)
	}
	wg.Wait()
}

func main() {
	//s := GetSizeByString("1m")
	//fmt.Println("size", s)
	//buffer := GetFileByBuffer(int(s))
	//fmt.Println(len(buffer.Bytes()))
	//endpoint := "http://localhost:1984"
	endpoint := "https://arweave.net"
	wallet, err := CreateWallet("./ar_demo/test-keyfile.json", endpoint)
	if err != nil {
		panic(err)
	}
	fmt.Println(wallet.Signer.Address)
	////36WJaNnsFvi_zaVAq7KPAcI1fRRytYIAO7X3LoQiEU8
	//otherAmount, err := wallet.Client.GetWalletBalance("nQKiFZE11MiXjY18qib_M4vz_AHyO3cf6gxxbemtIJY")
	//amount, err := wallet.Client.GetWalletBalance(wallet.Signer.Address)
	////arToWinston := utils.ARToWinston(amount)
	//fmt.Println(amount.String(), otherAmount.String())
	//sub := amount.Sub(amount, otherAmount)
	//fmt.Println(sub.String())

	//LocalMint(wallet)

	//TestUpload(wallet, 1024*1024)

	//TestDownloadData(wallet, "76CWjuI6itsb-rDiI5Y0bfM5hRRCdKBUm9QBxFPCXvQ")
	//8DEd9GE8YGkfbSA-GUAkCjP5iJ58pG_d_VDmMOvf82E
	//TestDownloadDataOne(wallet, "fsmZ_IEg_prSot51d65LbJAQy4rNqj2HBj-AwCCiEi0")
	TestDownloadFromPath(wallet, "/Users/kyle/Desktop/code/workPlace/golang/demo.csv")

}
