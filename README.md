
## Data Sources

The data source of the frequency

- comes from https://language.moe.gov.tw/001/Upload/files/SITE_CONTENT/M0001/PIN/biau1.htm,
  - which is from [國語辭典簡編本編輯資料字詞頻統計報告](https://language.moe.gov.tw/001/Upload/files/SITE_CONTENT/M0001/PIN/c11.htm),
- saved as `BIAU1.TXT` (UTF-8),
- parsed by `internal/tools/genfreq/main.go`.

The data source of CJ codes

- comes from https://github.com/openvanilla/McTabimWeb/blob/main/src/data/cin/cj5.json,
- saved as `cj5.json`,
- parsed by `internal/tools/gencode/main.go`.

The parsed `cjdict/cjcode.csv` and `cjdict/freq.csv` are committed and are the golden data for
the tests; the source files are not tracked. To regenerate them, point the environment variables
at the downloaded files (use absolute paths: `go generate` does not expand `~`):

```bash
export CJ5_JSON="$HOME/Downloads/cj5.json" BIAU1_TXT="$HOME/Downloads/BIAU1.TXT"
go generate ./cjdict/
```

Update the golden counts in `cjdict/dict_test.go` if the output changes.
