//@Author: Tuan Bach Van
//@Github: https://github.com/bachvtuan/shortmongoid
//@Website: https://dethoima.info
package shortmongoid

import (
	"strconv"
	"time"
	"strings"
	"math"
  "errors"
  //"fmt"
  "encoding/hex"
)
var symbols []string

func init() {
	symbols = strings.Split("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-","")
}

func toBase( num int ) (string, error ){
  decimal :=  num ;
  base := 64
  var temp int;
  var conversion string;

  if (base > len(symbols) || base <= 1) {
    return "", errors.New("Radix must be less than "+ strconv.Itoa(len(symbols))+" and greater than 1");
  }

  for decimal > 0 {
    temp = int(math.Floor( float64( decimal / base ) ));
    conversion = symbols[(decimal - (base * temp))] + conversion;
    decimal = temp;
  }

  return conversion, nil;
}

func ShortId( mongoId string ) (string, error) {

  if len(mongoId) != 24{
    return "", errors.New("Expect mongoId with length is 24")
  }


  //try validate hexa string
  _, err := hex.DecodeString( mongoId )

  if err != nil{

    return "", err
  }

  //Get 8 first characters
  first8 := mongoId[:8]
  // Convert to decimal
  timeDecimal, err := strconv.ParseInt(first8, 16, 64) 

  if err != nil{
    return "", err
  }

  //Convert to date, read here http://stackoverflow.com/questions/7327296/how-do-i-extract-the-created-date-out-of-a-mongo-objectid
  timeStamp:= time.Unix(0, timeDecimal*int64(time.Second))
  //Get last 6 chracters from mongoId
  last6 := mongoId[len(mongoId)-6:]

  lastDecimal, err := strconv.ParseInt(last6, 16, 64) 

  if err != nil{
    return "", err 
  }

  //Continue get last 6 number from decimal number
  last6String := strconv.Itoa( int(lastDecimal) )
  last6Decimal, err := strconv.Atoi(last6String[ len(last6String)-6: ])

  if err != nil{
    return "", err
  }

  nano := timeStamp.UnixNano()
//  fmt.Println( nano )
  umillisec := nano / 1000
  //combine to  time to make strong unique
  //fmt.Println(int(umillisec)  + last6Decimal)
  return toBase( int(umillisec)  + last6Decimal )

}