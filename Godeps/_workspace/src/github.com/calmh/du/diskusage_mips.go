// +build mipso32

package du

/*
#include <sys/statfs.h>
#include <stdlib.h>

typedef struct statfs_wrapper
{
    long f_bsize;
    long f_blocks;
    long f_bfree;
    long f_bavail;
} statfs_wrapper;

void get_statfs(char* path, statfs_wrapper* stat_w){
    struct statfs stat;
    statfs(path,&stat);
    stat_w->f_bsize = stat.f_bsize;
    stat_w->f_blocks = stat.f_blocks;
    stat_w->f_bfree = stat.f_bfree;
    stat_w->f_bavail = stat.f_bavail;
}
*/
import "C"
import (
	"unsafe"
)

//type StatfsWrapperGo C.statfs_wrapper

// Get returns the Usage of a given path, or an error if usage data is
// unavailable.
func Get(path string) (Usage, error) {
	stat := new(StatfsWrapperGo)
	stat_p := unsafe.Pointer(stat)
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.get_statfs(cpath, (*C.statfs_wrapper)(stat_p))
	u := Usage{
		FreeBytes:  int64(stat.Bfree) * int64(stat.Bsize),
		TotalBytes: int64(stat.Blocks) * int64(stat.Bsize),
		AvailBytes: int64(stat.Bavail) * int64(stat.Bsize),
	}
	return u, nil
}
