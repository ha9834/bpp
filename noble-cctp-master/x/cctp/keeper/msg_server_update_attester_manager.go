// Copyright 2024 Circle Internet Group, Inc.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package keeper

import (
	"context"

	"github.com/circlefin/noble-cctp/x/cctp/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAttesterManager(goCtx context.Context, msg *types.MsgUpdateAttesterManager) (*types.MsgUpdateAttesterManagerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentOwner := k.GetOwner(ctx)
	if currentOwner != msg.From {
		return nil, errors.Wrapf(types.ErrUnauthorized, "this message sender cannot update the attester manager")
	}

	_, err := sdk.AccAddressFromBech32(msg.NewAttesterManager)
	if err != nil {
		return nil, errors.Wrapf(types.ErrInvalidAddress, "invalid attester manager address (%s)", err)
	}

	currentAttesterManager := k.GetAttesterManager(ctx)
	k.SetAttesterManager(ctx, msg.NewAttesterManager)

	event := types.AttesterManagerUpdated{
		PreviousAttesterManager: currentAttesterManager,
		NewAttesterManager:      msg.NewAttesterManager,
	}
	err = ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgUpdateAttesterManagerResponse{}, err
}