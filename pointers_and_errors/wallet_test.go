package mypointer

import "testing"

// func TestWallet(t *testing.T) {
// 	wallet := Wallet{}

// 	wallet.Deposit(Bitcoin(10))
// 	got := wallet.Balance()
// 	want := Bitcoin(10)

// 	if got != want {
// 		t.Errorf("got %s, want %s", got, want)
// 	}
// }

// refactor to
// func TestWallet(t *testing.T) {

// 	t.Run("deposit", func(t *testing.T) {
// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))
// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})

// 	t.Run("withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}

// 		wallet.Withdraw(Bitcoin(10))
// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})
// }

// refactor to below

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	// tyring to withdraw insufficient funds
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get aan error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

// Helper function to check the balance

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
