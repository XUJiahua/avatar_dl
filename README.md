
### 微信头像批量下载

文件命名规则：uri 的 md5 值。

```
$ ./avatar_dl -h

A brief description of your application

Usage:
  avatar_dl [flags]

Flags:
      --config string               config file (default is $HOME/.avatar_dl.yaml)
  -d, --downloadFolderName string    (default "./download")
  -f, --failureFilename string       (default "failure_uris.txt")
  -h, --help                        help for avatar_dl
  -i, --inputFilename string         (default "sample100.csv")
  -l, --logFilename string           (default "logrus.log")
  -t, --toggle                      Help message for toggle
```

