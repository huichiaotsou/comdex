package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *LockedVault) Validate() error {
	if m.LockedVaultID == 0 {
		return fmt.Errorf("id cannot be empty")
	}
	if m.VaultID == 0 {
		return fmt.Errorf("id cannot be empty")
	}
	if m.PairID == 0 {
		return fmt.Errorf("pair_id cannot be empty")
	}
	if m.Admin == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return errors.Wrapf(err, "invalid owner %s", m.Admin)
	}
	if m.VaultOwner == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.VaultOwner); err != nil {
		return errors.Wrapf(err, "invalid owner %s", m.VaultOwner)
	}
	if m.AmountIn.IsNil() {
		return fmt.Errorf("amount_in cannot be nil")
	}
	if m.AmountIn.IsNegative() {
		return fmt.Errorf("amount_in cannot be negative")
	}
	if m.AmountOut.IsNil() {
		return fmt.Errorf("amount_out cannot be nil")
	}
	if m.AmountOut.IsNegative() {
		return fmt.Errorf("amount_out cannot be negative")
	}

	return nil
}

