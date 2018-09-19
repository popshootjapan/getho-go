package getho_test

import (
	"testing"

	getho "github.com/popshootjapan/getho-go"
)

func TestApi(t *testing.T) {
	cli, err := getho.New("candid-lion-93555", &getho.Options{Scheme: "https://", Host: "getho.io"})
	if err != nil {
		t.Fatalf(err.Error())
	}

	input := &getho.EthGetBalanceInput{
		Address:        "0xb1f407dcc37cdc0d5193c09f499d3766aa4c5743",
		BlockParameter: getho.NewBlockParameter("latest"),
	}

	req, resp := cli.EthGetBalanceRequest(input)
	err = req.Call()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Error != nil {
		t.Fatalf(resp.Error.Error())
	}

	t.Logf("result: %v", resp.Result)
}
