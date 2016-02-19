# shortmongoid
Generate short id( unique id ) from _id from mongodb for Golang.

I'm inspired from https://github.com/treygriffith/short-mongo-id which is written in Nodejs, I modify to work for Golang and make it more unique than original project but it will generate longer string than original one..

## Example

```
package main
import (
	"fmt"
	"github.com/bachvtuan/shortmongoid"
)

func main() {
	a := "56397963d531b73d1fb6cba98"
	key, error := shortmongoid.ShortId(a)
	if error != nil{
		fmt.Printf("Error: %s\n",error)
	}else{
		fmt.Printf("Key :%s\n",key)
		//Key ::58WVYxt6F
	}

}

```


## License 
MIT