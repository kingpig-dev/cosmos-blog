package blog

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "blog/api/blog/blog"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PostAll",
					Use:       "list-post",
					Short:     "List all post",
				},
				{
					RpcMethod:      "Post",
					Use:            "show-post [id]",
					Short:          "Shows a post by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "LoanAll",
					Use:       "list-loan",
					Short:     "List all loan",
				},
				{
					RpcMethod:      "Loan",
					Use:            "show-loan [id]",
					Short:          "Shows a loan by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePost",
					Use:            "create-post [title] [body]",
					Short:          "Create post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "body"}},
				},
				{
					RpcMethod:      "UpdatePost",
					Use:            "update-post [id] [title] [body]",
					Short:          "Update post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "title"}, {ProtoField: "body"}},
				},
				{
					RpcMethod:      "DeletePost",
					Use:            "delete-post [id]",
					Short:          "Delete post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "RequestLoan",
					Use:            "request-loan [amount] [fee] [collateral] [deadline]",
					Short:          "Send a request-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "fee"}, {ProtoField: "collateral"}, {ProtoField: "deadline"}},
				},
				{
					RpcMethod:      "ApproveLoan",
					Use:            "approve-loan [id]",
					Short:          "Send a approve-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "RepayLoan",
					Use:            "repay-loan [id]",
					Short:          "Send a repay-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "LiquidateLoan",
					Use:            "liquidate-loan [id]",
					Short:          "Send a liquidate-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CancelLoan",
					Use:            "cancel-loan [id]",
					Short:          "Send a cancel-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
