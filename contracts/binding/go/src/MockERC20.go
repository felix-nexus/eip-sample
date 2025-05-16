// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// MockERC20MetaData contains all meta data concerning the MockERC20 contract.
var MockERC20MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
	ID:  "MockERC20",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611de8380380611de883398181016040528101906100319190610296565b828282825f9081610042919061052e565b508160019081610052919061052e565b508060ff1660808160ff16815250504660a0818152505061007761008960201b60201c565b60c08181525050505050505050610766565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516100b99190610699565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646306040516020016100f8959493929190610715565b60405160208183030381529060405280519060200120905090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101728261012c565b810181811067ffffffffffffffff821117156101915761019061013c565b5b80604052505050565b5f6101a3610113565b90506101af8282610169565b919050565b5f67ffffffffffffffff8211156101ce576101cd61013c565b5b6101d78261012c565b9050602081019050919050565b8281835e5f83830152505050565b5f6102046101ff846101b4565b61019a565b9050828152602081018484840111156102205761021f610128565b5b61022b8482856101e4565b509392505050565b5f82601f83011261024757610246610124565b5b81516102578482602086016101f2565b91505092915050565b5f60ff82169050919050565b61027581610260565b811461027f575f5ffd5b50565b5f815190506102908161026c565b92915050565b5f5f5f606084860312156102ad576102ac61011c565b5b5f84015167ffffffffffffffff8111156102ca576102c9610120565b5b6102d686828701610233565b935050602084015167ffffffffffffffff8111156102f7576102f6610120565b5b61030386828701610233565b925050604061031486828701610282565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061036c57607f821691505b60208210810361037f5761037e610328565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103e17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103a6565b6103eb86836103a6565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61042f61042a61042584610403565b61040c565b610403565b9050919050565b5f819050919050565b61044883610415565b61045c61045482610436565b8484546103b2565b825550505050565b5f5f905090565b610473610464565b61047e81848461043f565b505050565b5b818110156104a1576104965f8261046b565b600181019050610484565b5050565b601f8211156104e6576104b781610385565b6104c084610397565b810160208510156104cf578190505b6104e36104db85610397565b830182610483565b50505b505050565b5f82821c905092915050565b5f6105065f19846008026104eb565b1980831691505092915050565b5f61051e83836104f7565b9150826002028217905092915050565b6105378261031e565b67ffffffffffffffff8111156105505761054f61013c565b5b61055a8254610355565b6105658282856104a5565b5f60209050601f831160018114610596575f8415610584578287015190505b61058e8582610513565b8655506105f5565b601f1984166105a486610385565b5f5b828110156105cb578489015182556001820191506020850194506020810190506105a6565b868310156105e857848901516105e4601f8916826104f7565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f815461062581610355565b61062f81866105fd565b9450600182165f8114610649576001811461065e57610690565b60ff1983168652811515820286019350610690565b61066785610607565b5f5b8381101561068857815481890152600182019150602081019050610669565b838801955050505b50505092915050565b5f6106a48284610619565b915081905092915050565b5f819050919050565b6106c1816106af565b82525050565b6106d081610403565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106ff826106d6565b9050919050565b61070f816106f5565b82525050565b5f60a0820190506107285f8301886106b8565b61073560208301876106b8565b61074260408301866106b8565b61074f60608301856106c7565b61075c6080830184610706565b9695505050505050565b60805160a05160c0516116586107905f395f61070a01525f6106d601525f6106b101526116585ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c806370a082311161008a5780639dc29fac116100645780639dc29fac1461025e578063a9059cbb1461027a578063d505accf146102aa578063dd62ed3e146102c6576100e8565b806370a08231146101e05780637ecebe001461021057806395d89b4114610240576100e8565b806323b872dd116100c657806323b872dd14610158578063313ce567146101885780633644e515146101a657806340c10f19146101c4576100e8565b806306fdde03146100ec578063095ea7b31461010a57806318160ddd1461013a575b5f5ffd5b6100f46102f6565b6040516101019190610eab565b60405180910390f35b610124600480360381019061011f9190610f5c565b610381565b6040516101319190610fb4565b60405180910390f35b61014261046e565b60405161014f9190610fdc565b60405180910390f35b610172600480360381019061016d9190610ff5565b610474565b60405161017f9190610fb4565b60405180910390f35b6101906106af565b60405161019d9190611060565b60405180910390f35b6101ae6106d3565b6040516101bb9190611091565b60405180910390f35b6101de60048036038101906101d99190610f5c565b61072f565b005b6101fa60048036038101906101f591906110aa565b61073d565b6040516102079190610fdc565b60405180910390f35b61022a600480360381019061022591906110aa565b610752565b6040516102379190610fdc565b60405180910390f35b610248610767565b6040516102559190610eab565b60405180910390f35b61027860048036038101906102739190610f5c565b6107f3565b005b610294600480360381019061028f9190610f5c565b610801565b6040516102a19190610fb4565b60405180910390f35b6102c460048036038101906102bf9190611129565b61090e565b005b6102e060048036038101906102db91906111c6565b610bfb565b6040516102ed9190610fdc565b60405180910390f35b5f805461030290611231565b80601f016020809104026020016040519081016040528092919081815260200182805461032e90611231565b80156103795780601f1061035057610100808354040283529160200191610379565b820191905f5260205f20905b81548152906001019060200180831161035c57829003601f168201915b505050505081565b5f8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161045c9190610fdc565b60405180910390a36001905092915050565b60025481565b5f5f60045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146105a1578281610524919061128e565b60045f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b8260035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105ed919061128e565b925050819055508260035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161069b9190610fdc565b60405180910390a360019150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f0000000000000000000000000000000000000000000000000000000000000000461461070857610703610c1b565b61072a565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b6107398282610ca5565b5050565b6003602052805f5260405f205f915090505481565b6005602052805f5260405f205f915090505481565b6001805461077490611231565b80601f01602080910402602001604051908101604052809291908181526020018280546107a090611231565b80156107eb5780601f106107c2576101008083540402835291602001916107eb565b820191905f5260205f20905b8154815290600101906020018083116107ce57829003601f168201915b505050505081565b6107fd8282610d70565b5050565b5f8160035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461084e919061128e565b925050819055508160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108fc9190610fdc565b60405180910390a36001905092915050565b42841015610951576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109489061130b565b60405180910390fd5b5f600161095c6106d3565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a60055f8f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190600101919050558b6040516020016109e196959493929190611338565b60405160208183030381529060405280519060200120604051602001610a0892919061140b565b604051602081830303815290604052805190602001208585856040515f8152602001604052604051610a3d9493929190611441565b6020604051602081039080840390855afa158015610a5d573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610ad057508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b06906114ce565b60405180910390fd5b8560045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610bea9190610fdc565b60405180910390a350505050505050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f604051610c4b9190611588565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001610c8a95949392919061159e565b60405160208183030381529060405280519060200120905090565b8060025f828254610cb691906115ef565b925050819055508060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d649190610fdc565b60405180910390a35050565b8060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610dbc919061128e565b925050819055508060025f82825403925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e2f9190610fdc565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e7d82610e3b565b610e878185610e45565b9350610e97818560208601610e55565b610ea081610e63565b840191505092915050565b5f6020820190508181035f830152610ec38184610e73565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ef882610ecf565b9050919050565b610f0881610eee565b8114610f12575f5ffd5b50565b5f81359050610f2381610eff565b92915050565b5f819050919050565b610f3b81610f29565b8114610f45575f5ffd5b50565b5f81359050610f5681610f32565b92915050565b5f5f60408385031215610f7257610f71610ecb565b5b5f610f7f85828601610f15565b9250506020610f9085828601610f48565b9150509250929050565b5f8115159050919050565b610fae81610f9a565b82525050565b5f602082019050610fc75f830184610fa5565b92915050565b610fd681610f29565b82525050565b5f602082019050610fef5f830184610fcd565b92915050565b5f5f5f6060848603121561100c5761100b610ecb565b5b5f61101986828701610f15565b935050602061102a86828701610f15565b925050604061103b86828701610f48565b9150509250925092565b5f60ff82169050919050565b61105a81611045565b82525050565b5f6020820190506110735f830184611051565b92915050565b5f819050919050565b61108b81611079565b82525050565b5f6020820190506110a45f830184611082565b92915050565b5f602082840312156110bf576110be610ecb565b5b5f6110cc84828501610f15565b91505092915050565b6110de81611045565b81146110e8575f5ffd5b50565b5f813590506110f9816110d5565b92915050565b61110881611079565b8114611112575f5ffd5b50565b5f81359050611123816110ff565b92915050565b5f5f5f5f5f5f5f60e0888a03121561114457611143610ecb565b5b5f6111518a828b01610f15565b97505060206111628a828b01610f15565b96505060406111738a828b01610f48565b95505060606111848a828b01610f48565b94505060806111958a828b016110eb565b93505060a06111a68a828b01611115565b92505060c06111b78a828b01611115565b91505092959891949750929550565b5f5f604083850312156111dc576111db610ecb565b5b5f6111e985828601610f15565b92505060206111fa85828601610f15565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061124857607f821691505b60208210810361125b5761125a611204565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61129882610f29565b91506112a383610f29565b92508282039050818111156112bb576112ba611261565b5b92915050565b7f5045524d49545f444541444c494e455f455850495245440000000000000000005f82015250565b5f6112f5601783610e45565b9150611300826112c1565b602082019050919050565b5f6020820190508181035f830152611322816112e9565b9050919050565b61133281610eee565b82525050565b5f60c08201905061134b5f830189611082565b6113586020830188611329565b6113656040830187611329565b6113726060830186610fcd565b61137f6080830185610fcd565b61138c60a0830184610fcd565b979650505050505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f6113d5600283611397565b91506113e0826113a1565b600282019050919050565b5f819050919050565b61140561140082611079565b6113eb565b82525050565b5f611415826113c9565b915061142182856113f4565b60208201915061143182846113f4565b6020820191508190509392505050565b5f6080820190506114545f830187611082565b6114616020830186611051565b61146e6040830185611082565b61147b6060830184611082565b95945050505050565b7f494e56414c49445f5349474e45520000000000000000000000000000000000005f82015250565b5f6114b8600e83610e45565b91506114c382611484565b602082019050919050565b5f6020820190508181035f8301526114e5816114ac565b9050919050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f815461151481611231565b61151e81866114ec565b9450600182165f8114611538576001811461154d5761157f565b60ff198316865281151582028601935061157f565b611556856114f6565b5f5b8381101561157757815481890152600182019150602081019050611558565b838801955050505b50505092915050565b5f6115938284611508565b915081905092915050565b5f60a0820190506115b15f830188611082565b6115be6020830187611082565b6115cb6040830186611082565b6115d86060830185610fcd565b6115e56080830184611329565b9695505050505050565b5f6115f982610f29565b915061160483610f29565b925082820190508082111561161c5761161b611261565b5b9291505056fea2646970667358221220d4867e3ca63815f6694b838eab2f1b6ecd213fba6e1de455102bd7c889926fd664736f6c634300081c0033",
}

// MockERC20 is an auto generated Go binding around an Ethereum contract.
type MockERC20 struct {
	abi abi.ABI
}

// NewMockERC20 creates a new instance of MockERC20.
func NewMockERC20() *MockERC20 {
	parsed, err := MockERC20MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MockERC20{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MockERC20) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _name, string _symbol, uint8 _decimals) returns()
func (mockERC20 *MockERC20) PackConstructor(_name string, _symbol string, _decimals uint8) []byte {
	enc, err := mockERC20.abi.Pack("", _name, _symbol, _decimals)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (mockERC20 *MockERC20) PackDOMAINSEPARATOR() []byte {
	enc, err := mockERC20.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (mockERC20 *MockERC20) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := mockERC20.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (mockERC20 *MockERC20) PackAllowance(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := mockERC20.abi.Pack("allowance", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (mockERC20 *MockERC20) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := mockERC20.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (mockERC20 *MockERC20) PackApprove(spender common.Address, amount *big.Int) []byte {
	enc, err := mockERC20.abi.Pack("approve", spender, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (mockERC20 *MockERC20) UnpackApprove(data []byte) (bool, error) {
	out, err := mockERC20.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (mockERC20 *MockERC20) PackBalanceOf(arg0 common.Address) []byte {
	enc, err := mockERC20.abi.Pack("balanceOf", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (mockERC20 *MockERC20) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := mockERC20.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 value) returns()
func (mockERC20 *MockERC20) PackBurn(from common.Address, value *big.Int) []byte {
	enc, err := mockERC20.abi.Pack("burn", from, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (mockERC20 *MockERC20) PackDecimals() []byte {
	enc, err := mockERC20.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (mockERC20 *MockERC20) UnpackDecimals(data []byte) (uint8, error) {
	out, err := mockERC20.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (mockERC20 *MockERC20) PackMint(to common.Address, value *big.Int) []byte {
	enc, err := mockERC20.abi.Pack("mint", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (mockERC20 *MockERC20) PackName() []byte {
	enc, err := mockERC20.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (mockERC20 *MockERC20) UnpackName(data []byte) (string, error) {
	out, err := mockERC20.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (mockERC20 *MockERC20) PackNonces(arg0 common.Address) []byte {
	enc, err := mockERC20.abi.Pack("nonces", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (mockERC20 *MockERC20) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := mockERC20.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (mockERC20 *MockERC20) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := mockERC20.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (mockERC20 *MockERC20) PackSymbol() []byte {
	enc, err := mockERC20.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (mockERC20 *MockERC20) UnpackSymbol(data []byte) (string, error) {
	out, err := mockERC20.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (mockERC20 *MockERC20) PackTotalSupply() []byte {
	enc, err := mockERC20.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (mockERC20 *MockERC20) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := mockERC20.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (mockERC20 *MockERC20) PackTransfer(to common.Address, amount *big.Int) []byte {
	enc, err := mockERC20.abi.Pack("transfer", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (mockERC20 *MockERC20) UnpackTransfer(data []byte) (bool, error) {
	out, err := mockERC20.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (mockERC20 *MockERC20) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := mockERC20.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (mockERC20 *MockERC20) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := mockERC20.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// MockERC20Approval represents a Approval event raised by the MockERC20 contract.
type MockERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const MockERC20ApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (MockERC20Approval) ContractEventName() string {
	return MockERC20ApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (mockERC20 *MockERC20) UnpackApprovalEvent(log *types.Log) (*MockERC20Approval, error) {
	event := "Approval"
	if log.Topics[0] != mockERC20.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MockERC20Approval)
	if len(log.Data) > 0 {
		if err := mockERC20.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mockERC20.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MockERC20Transfer represents a Transfer event raised by the MockERC20 contract.
type MockERC20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const MockERC20TransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (MockERC20Transfer) ContractEventName() string {
	return MockERC20TransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (mockERC20 *MockERC20) UnpackTransferEvent(log *types.Log) (*MockERC20Transfer, error) {
	event := "Transfer"
	if log.Topics[0] != mockERC20.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MockERC20Transfer)
	if len(log.Data) > 0 {
		if err := mockERC20.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mockERC20.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
