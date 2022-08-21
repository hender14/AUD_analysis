package gateway

import (
	"testing"
	"time"

	"github.com/hender14/app/domain"

	"github.com/golang/mock/gomock"
)

func TestQueryEmail(t *testing.T) {
	var r gomock.TestReporter
	input := &domain.InUser{FirstName: "test", LastName: "test", Email: "test@gmail.com", Password: "test", Password_confirm: "test"}
	output := []domain.SignUser{}
	output_tmp := domain.SignUser{ID: "test", FirstName: "test", LastName: "test", Email: "test@gmail.com", Password: []byte("test"), Year: time.Now()}
	output = append(output, output_tmp)
	// mockのコントローラを作成します
	ctrl := gomock.NewController(r)
	// ctrl := gomock.NewController(NewConcurrentTestReporter(t))
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成します
	mockApiClinet := NewMockCRUD(ctrl)
	mockApiClinet_f := NewMockCRUD(ctrl)
	// mockApiClinet := mock_port.NewMockUserRepository(ctrl)
	// 作成したmockに対して期待する呼び出しと返り値を定義します
	// EXPECT()では呼び出されたかどうか
	// Request()ではそのメソッド名が指定した引数で呼び出されたかどうか
	// Return()では返り値を指定します
	mockApiClinet.EXPECT().Fsquery(gomock.Any()).Return(nil, nil).Times(1)
	mockApiClinet_f.EXPECT().Fsquery(gomock.Any()).Return(output, nil).Times(1)
	// mockApiClinet.EXPECT().RegisterAccoount(gomock.Any()).Do(func(s *domain.SignUser) {
	// 	// Do を使ってモック関数への引数を得ることができる。
	// 	// Do に渡す引数は`actUser`を持つクロージャ関数となる。
	// 	check = s
	// }).Return(nil).Times(1)
	// mockApiClinet.EXPECT().RegisterAccoount(gomock.Any()).Return(nil).Times(1)

	d := &UserRepository{}
	d_f := &UserRepository{}
	// d.User = &ApiClientMock{} // mockを登録
	d.context = mockApiClinet     // mockを登録
	d_f.context = mockApiClinet_f // mockを登録
	expected := input

	// res, err := d.Sign(input)
	_, err := d.QueryEmail_none(expected.Email)
	if err != nil {
		t.Fatal("Register error!", err)
	}

	_, err = d_f.QueryEmail_none(expected.Email)
	if err == nil {
		t.Fatal("Register error!", err)
	}
}
