package interactor

import (
	"testing"

	"github.com/hender14/app/domain"
	"github.com/hender14/app/usecase/port"

	"github.com/golang/mock/gomock"
)

func TestSign(t *testing.T) {
	var r gomock.TestReporter
	var check *domain.SignUser
	// password, _ := domain.EncodePassword("test")
	input := &domain.InUser{FirstName: "test", LastName: "test", Email: "test@gmail.com", Password: "test", Password_confirm: "test"}
	// sinput := &domain.SignUser{FirstName: "test", LastName: "test", Email: "test@gmail.com"}
	// mockのコントローラを作成します
	ctrl := gomock.NewController(r)
	// ctrl := gomock.NewController(NewConcurrentTestReporter(t))
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成します
	mockApiClinet_repo := port.NewMockUserRepository(ctrl)
	mockApiClinet_out := port.NewMockUserOutputPort(ctrl)
	// mockApiClinet := mock_port.NewMockUserRepository(ctrl)
	// 作成したmockに対して期待する呼び出しと返り値を定義します
	// EXPECT()では呼び出されたかどうか
	// Request()ではそのメソッド名が指定した引数で呼び出されたかどうか
	// Return()では返り値を指定します
	mockApiClinet_repo.EXPECT().QueryEmail(input).Return(nil).Times(1)
	mockApiClinet_repo.EXPECT().RegisterAccoount(gomock.Any()).Do(func(s *domain.SignUser) {
		// Do を使ってモック関数への引数を得ることができる。
		// Do に渡す引数は`actUser`を持つクロージャ関数となる。
		check = s
	}).Return(nil).Times(1)
	mockApiClinet_out.EXPECT().Render(gomock.Any(), gomock.Any()).Times(1)
	mockApiClinet_out.EXPECT().RenderError(gomock.Any(), gomock.Any()).Times(0)
	// mockApiClinet.EXPECT().RegisterAccoount(gomock.Any()).Return(nil).Times(1)

	d := &UserInteractor{}
	// d.User = &ApiClientMock{} // mockを登録
	d.User = mockApiClinet_repo      // mockを登録
	d.OutputPort = mockApiClinet_out // mockを登録
	expected := input

	// res, err := d.Sign(input)
	// res, err := d.Sign(expected)
	d.Sign(expected)
	// if err != nil {
	// 	t.Fatal("Register error!", err)
	// }
	// if res.FirstName != expected.FirstName || res.LastName != expected.LastName || res.Email != expected.Email {
	// 	t.Fatal("Value does not match.")
	// }
	if len(check.Password) != 60 {
		t.Fatal("Value does not match.")
	}
}
