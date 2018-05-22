## 资源地址

https://flaviocopes.com/go-git-contributions/

## 小结

Go 的标准库比 Python 要好用很多，和 linux 本身的设计紧密相连，可以说从 linux 中看到的正好就是是 Go 所提供的，比如 linux 的 文件系统中，
文件系统本身的各种功能在 Go 中就是很好体现。

比如判断一个文件本身是文件还是文件夹，直接使用 FileInfo.IsDir() 判断即可。FileInfo 在 linux 系统中正好就是对应 ls 所提供的信息
```Go
± % go doc os.FileInfo
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
    A FileInfo describes a file and is returned by Stat and Lstat.


func Lstat(name string) (FileInfo, error)
func Stat(name string) (FileInfo, error)

```

FileInfo 是一个 Interface，意味着任何满足此 Interface 的都可以使用，也即是任何平台的文件系统的共性都提供了。

反观 Python 的实现是通过 os.stat 传入 path 来实现，不直观，而 stat 本身就是文件的属性信息，其实就应该是 FileInfo