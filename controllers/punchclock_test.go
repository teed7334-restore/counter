package controllers

import (
	"reflect"
	"testing"

	"github.com/teed7334-restore/counter/beans"
	homekeeper "github.com/teed7334-restore/homekeeper/beans"
)

func Test_doUploadDailyPunchclockData(t *testing.T) {
	type args struct {
		params *homekeeper.DailyPunchclockData
	}
	tests := []struct {
		name string
		args args
		want *beans.Response
	}{
		{
			name: "一般測試",
			args: args{
				params: &homekeeper.DailyPunchclockData{
					Employee: &homekeeper.EmployeeOnChain{
						Identify:  "00190",
						FirstName: "Peter",
						LastName:  "Cheng",
					},
					Punchclock: &homekeeper.Punchclock{
						Begin: &homekeeper.TimeStruct{
							Year:   "2019",
							Month:  "08",
							Day:    "28",
							Hour:   "09",
							Minute: "24",
							Second: "00",
						},
						End: &homekeeper.TimeStruct{
							Year:   "2019",
							Month:  "08",
							Day:    "28",
							Hour:   "20",
							Minute: "05",
							Second: "00",
						},
					},
				},
			},
			want: &beans.Response{
				Status:  true,
				Channel: "PunchClock",
				Message: "PunchClock/UploadDailyPunchclockData</UseService>{\"employee\":{\"$class\":\"\",\"identify\":\"00190\",\"firstName\":\"Peter\",\"lastName\":\"Cheng\"},\"punchclock\":{\"begin\":{\"year\":\"2019\",\"month\":\"08\",\"day\":\"28\",\"hour\":\"09\",\"minute\":\"24\",\"second\":\"00\"},\"end\":{\"year\":\"2019\",\"month\":\"08\",\"day\":\"28\",\"hour\":\"20\",\"minute\":\"05\",\"second\":\"00\"}}}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doUploadDailyPunchclockData(tt.args.params)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doUploadDailyPunchclockData() = %v, want %v", got, tt.want)
			}
		})
	}
}
