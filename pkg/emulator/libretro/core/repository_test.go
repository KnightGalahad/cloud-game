package core

import "testing"

func TestBuildbotRepo(t *testing.T) {
	testAddress := "http://test.me"
	tests := []struct {
		file        string
		compression string
		arch        ArchInfo
		resultUrl   string
	}{
		{
			file: "uber_core",
			arch: ArchInfo{
				os:   "linux",
				arch: "x86_64",
				Lib:  ".so",
			},
			resultUrl: testAddress + "/" + "linux/x86_64/latest/uber_core.so",
		},
		{
			file:        "uber_core",
			compression: "zip",
			arch: ArchInfo{
				os:   "linux",
				arch: "x86_64",
				Lib:  ".so",
			},
			resultUrl: testAddress + "/" + "linux/x86_64/latest/uber_core.so.zip",
		},
		{
			file: "uber_core",
			arch: ArchInfo{
				os:     "osx",
				arch:   "x86_64",
				vendor: "apple",
				Lib:    ".dylib",
			},
			resultUrl: testAddress + "/" + "apple/osx/x86_64/latest/uber_core.dylib",
		},
	}

	for _, test := range tests {
		repo := BuildbotRepo{address: testAddress, compression: test.compression}
		url := repo.GetLink(test.file, test.arch)
		if url != test.resultUrl {
			t.Errorf("seems that expected link address is incorrect for file %s %+v", test.file, test.arch)
		}
	}
}