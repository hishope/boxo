package data

const (
	Data_Raw       int64 = 0
	Data_Directory int64 = 1
	Data_File      int64 = 2
	Data_Metadata  int64 = 3
	Data_Symlink   int64 = 4
	Data_HAMTShard int64 = 5
)

var DataTypeNames = map[int64]string{
	Data_Raw:       "Raw",
	Data_Directory: "Directory",
	Data_File:      "File",
	Data_Metadata:  "Metadata",
	Data_Symlink:   "Symlink",
	Data_HAMTShard: "HAMTShard",
}

var DataTypeValues = map[string]int64{
	"Raw":       Data_Raw,
	"Directory": Data_Directory,
	"File":      Data_File,
	"Metadata":  Data_Metadata,
	"Symlink":   Data_Symlink,
	"HAMTShard": Data_HAMTShard,
}
