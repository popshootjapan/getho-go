package getho_test

import (
	"testing"

	getho "github.com/cryptokyo/getho-go"
)

func TestApi(t *testing.T) {
	cli, err := getho.New("forward-tick-cf5499a16187")
	if err != nil {
		t.Fatalf(err.Error())
	}

	input := &getho.EthGetBalanceInput{
		Address:        "0x5c66b0d82df26e8FE165Be6628F5f5e1f1bccD5C",
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
