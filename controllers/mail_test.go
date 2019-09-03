package controllers

import (
	"reflect"
	"testing"

	"github.com/teed7334-restore/counter/beans"
)

func Test_doSendMail(t *testing.T) {
	type args struct {
		params *beans.SendMail
	}
	tests := []struct {
		name string
		args args
		want *beans.Response
	}{
		{
			name: "一般測試",
			args: args{
				params: &beans.SendMail{
					To:      "teed7334@gmail.com",
					Cc:      "teed7334@163.com",
					Subject: "這是一封測試信",
					Content: "這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信<br />這是一封測試信",
				},
			},
			want: &beans.Response{
				Status:  true,
				Channel: "Mail",
				Message: "Mail/SendMail</UseService>{\"To\":\"teed7334@gmail.com\",\"Cc\":\"teed7334@163.com\",\"Subject\":\"\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\",\"Content\":\"\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\\u003cbr /\\u003e\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\\u003cbr /\\u003e\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\\u003cbr /\\u003e\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\\u003cbr /\\u003e\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\\u003cbr /\\u003e\351\200\231\346\230\257\344\270\200\345\260\201\346\270\254\350\251\246\344\277\241\"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doSendMail(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doSendMail() = %v, want %v", got, tt.want)
			}
		})
	}
}
