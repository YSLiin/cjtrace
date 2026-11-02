
## Data Sources

The data source of the frequency

- comes from https://language.moe.gov.tw/001/Upload/files/SITE_CONTENT/M0001/PIN/biau1.htm,
  - which is from [國語辭典簡編本編輯資料字詞頻統計報告](https://language.moe.gov.tw/001/Upload/files/SITE_CONTENT/M0001/PIN/c11.htm),
- parsed by `internal/tools/genfreq/main.go`.

The data source of CJ codes

- comes from https://github.com/cangjie-system/rime-cangjie-integrated/blob/master/cangjie.5-base.dict.yaml,
- parsed by `internal/tools/gencode/main.go`.

```bash
go run internal/tools/genfreq/main.go -in ~/Downloads/freq.txt --minfreq 32 > freq.csv
go run internal/tools/gencode/main.go -in ~/Downloads/cangjie.5-base.dict.yaml > cjcode.csv
```
