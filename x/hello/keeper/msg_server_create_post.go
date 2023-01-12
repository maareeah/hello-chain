package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"hello/x/hello/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// _ = ctx
	var post = types.Post{
		Creator: msg.Creator,
		Title: msg.Title,
		Body: msg.Body,
		Date: msg.Date,
	}

	// Add a post to the store and get back the ID
	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id:id}, nil
}
