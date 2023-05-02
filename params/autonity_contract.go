package params

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/autonity/autonity/crypto"

	"github.com/autonity/autonity/common"
	"github.com/autonity/autonity/common/acdefault"
	"github.com/autonity/autonity/log"
	"github.com/autonity/autonity/p2p/enode"
)

// Autonity contract config. It'is used for deployment.
type AutonityContractGenesis struct {
	// Bytecode of validators contract
	// would like this type to be []byte but the unmarshalling is not working
	Bytecode string `json:"bytecode,omitempty" toml:",omitempty"`
	// Json ABI of the contract
	ABI              string         `json:"abi,omitempty" toml:",omitempty"`
	MinBaseFee       uint64         `json:"minBaseFee"`
	EpochPeriod      uint64         `json:"epochPeriod"`
	UnbondingPeriod  uint64         `json:"unbondingPeriod"`
	BlockPeriod      uint64         `json:"blockPeriod"`
	MaxCommitteeSize uint64         `json:"maxCommitteeSize"`
	Operator         common.Address `json:"operator"`
	Treasury         common.Address `json:"treasury"`
	TreasuryFee      uint64         `json:"treasuryFee"`
	DelegationRate   uint64         `json:"delegationRate"`
	Validators       []*Validator   `json:"validators"`
}

// Prepare prepares the AutonityContractGenesis by filling in missing fields.
// It returns an error if the configuration is invalid.
func (ac *AutonityContractGenesis) Prepare() error {

	if len(ac.Bytecode) == 0 && len(ac.ABI) > 0 ||
		len(ac.Bytecode) > 0 && len(ac.ABI) == 0 {
		return errors.New("it is an error to set only of autonity contract abi or bytecode")
	}

	if len(ac.Bytecode) == 0 && len(ac.ABI) == 0 {
		log.Info("Setting up Autonity default dPoS contract")
		ac.ABI = acdefault.ABI()
		ac.Bytecode = acdefault.Bytecode()
	} else {
		log.Info("Setting up custom Autonity protocol contract")
	}
	if reflect.DeepEqual(ac.Operator, common.Address{}) {
		ac.Operator = acdefault.Governance()
	}

	if ac.MaxCommitteeSize == 0 {
		return errors.New("invalid max committee size")
	}

	if len(ac.GetValidators()) == 0 {
		return errors.New("no initial validators")
	}

	for i, v := range ac.Validators {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("error parsing validator %d, err: %v", i+1, err)
		}
	}
	return nil
}

//User - is used to put predefined accounts to genesis
type Validator struct {
	Treasury          common.Address  `abi:"treasury"`
	Address           *common.Address `abi:"nodeAddress"`
	OracleAddress     common.Address  `abi:"oracleAddress"`
	Enode             string          `abi:"enode"`
	CommissionRate    *big.Int        `abi:"commissionRate"`
	BondedStake       *big.Int        `abi:"bondedStake"`
	TotalSlashed      *big.Int        `abi:"totalSlashed"`
	LiquidContract    *common.Address `abi:"liquidContract"`
	LiquidSupply      *big.Int        `abi:"liquidSupply"`
	RegistrationBlock *big.Int        `abi:"registrationBlock"`
	State             *uint8          `abi:"state"`
}

// getAddressFromEnode gets the account address from the user enode.
func (u *Validator) getAddressFromEnode() (common.Address, error) {
	n, err := enode.ParseV4NoResolve(u.Enode)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse enode %q, error:%v", u.Enode, err)
	}
	return crypto.PubkeyToAddress(*n.Pubkey()), nil
}

func (u *Validator) Validate() error {
	if len(u.Enode) == 0 {
		return errors.New("enode must be specified")
	}
	if u.BondedStake == nil || u.BondedStake.Cmp(new(big.Int)) == 0 {
		return errors.New("bonded stake must be specified")
	}
	a, err := u.getAddressFromEnode()
	if err != nil {
		return err
	}
	// If address is set check it matches the address from the enode
	if u.Address != nil && *u.Address != a {
		return fmt.Errorf("mismatching address %q and address from enode %q", u.Address.String(), a.String())
	}
	u.Address = &a
	if u.TotalSlashed == nil {
		u.TotalSlashed = new(big.Int)
	}
	if u.LiquidSupply == nil {
		u.LiquidSupply = new(big.Int)
	}
	if u.CommissionRate == nil {
		u.CommissionRate = new(big.Int)
	}
	if u.RegistrationBlock == nil {
		u.RegistrationBlock = new(big.Int)
	}
	if u.LiquidContract == nil {
		u.LiquidContract = new(common.Address)
	}
	if u.State == nil {
		u.State = new(uint8)
	}
	if u.CommissionRate != nil && u.CommissionRate.Cmp(big.NewInt(0)) != 0 {
		return fmt.Errorf("commission rate for enode %q not allowed", a.String())
	}
	return nil
}

func (ac *AutonityContractGenesis) GetValidators() []*Validator {
	return ac.Validators
}

func (ac *AutonityContractGenesis) GetValidatorsCopy() []Validator {
	validators := make([]Validator, len(ac.Validators))
	for i := range ac.Validators {
		validators[i] = *ac.Validators[i]
	}
	return validators
}
