package container

import (
	"fmt"
	"path/filepath"
)

func main(){
	mntURL := "/root/mnt/"
	containerURL := "/container"

	fmt.Println(filepath.Join(mntURL, containerURL))
	// nums := [...]int{1, 2,3 ,4, 5, 6}
	/* for index, item := range nums{
		
	}
 */
}