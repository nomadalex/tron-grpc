package tron_grpc

import (
	"context"
	"github.com/fullstackwang/tron-grpc/api"
	"github.com/fullstackwang/tron-grpc/core"
	"google.golang.org/grpc"
)

func (g *GrpcClient) GetAccount(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*core.Account, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAccount(ctx, in, opts...)
}

func (g *GrpcClient) GetAccountById(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*core.Account, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAccountById(ctx, in, opts...)
}

func (g *GrpcClient) GetAccountBalance(ctx context.Context, in *core.AccountBalanceRequest, opts ...grpc.CallOption) (*core.AccountBalanceResponse, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAccountBalance(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockBalanceTrace(ctx context.Context, in *core.BlockBalanceTrace_BlockIdentifier, opts ...grpc.CallOption) (*core.BlockBalanceTrace, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockBalanceTrace(ctx, in, opts...)
}

func (g *GrpcClient) CreateTransaction(ctx context.Context, in *core.TransferContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateTransaction(ctx, in, opts...)
}

func (g *GrpcClient) CreateTransaction2(ctx context.Context, in *core.TransferContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateTransaction2(ctx, in, opts...)
}

func (g *GrpcClient) BroadcastTransaction(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.Return, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.BroadcastTransaction(ctx, in, opts...)
}

func (g *GrpcClient) UpdateAccount(ctx context.Context, in *core.AccountUpdateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateAccount(ctx, in, opts...)
}

func (g *GrpcClient) SetAccountId(ctx context.Context, in *core.SetAccountIdContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.SetAccountId(ctx, in, opts...)
}

func (g *GrpcClient) UpdateAccount2(ctx context.Context, in *core.AccountUpdateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateAccount2(ctx, in, opts...)
}

func (g *GrpcClient) VoteWitnessAccount(ctx context.Context, in *core.VoteWitnessContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.VoteWitnessAccount(ctx, in, opts...)
}

func (g *GrpcClient) UpdateSetting(ctx context.Context, in *core.UpdateSettingContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateSetting(ctx, in, opts...)
}

func (g *GrpcClient) UpdateEnergyLimit(ctx context.Context, in *core.UpdateEnergyLimitContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateEnergyLimit(ctx, in, opts...)
}

func (g *GrpcClient) VoteWitnessAccount2(ctx context.Context, in *core.VoteWitnessContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.VoteWitnessAccount2(ctx, in, opts...)
}

func (g *GrpcClient) CreateAssetIssue(ctx context.Context, in *core.AssetIssueContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateAssetIssue(ctx, in, opts...)
}

func (g *GrpcClient) CreateAssetIssue2(ctx context.Context, in *core.AssetIssueContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateAssetIssue2(ctx, in, opts...)
}

func (g *GrpcClient) UpdateWitness(ctx context.Context, in *core.WitnessUpdateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateWitness(ctx, in, opts...)
}

func (g *GrpcClient) UpdateWitness2(ctx context.Context, in *core.WitnessUpdateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateWitness2(ctx, in, opts...)
}

func (g *GrpcClient) CreateAccount(ctx context.Context, in *core.AccountCreateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateAccount(ctx, in, opts...)
}

func (g *GrpcClient) CreateAccount2(ctx context.Context, in *core.AccountCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateAccount2(ctx, in, opts...)
}

func (g *GrpcClient) CreateWitness(ctx context.Context, in *core.WitnessCreateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateWitness(ctx, in, opts...)
}

func (g *GrpcClient) CreateWitness2(ctx context.Context, in *core.WitnessCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateWitness2(ctx, in, opts...)
}

func (g *GrpcClient) TransferAsset(ctx context.Context, in *core.TransferAssetContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.TransferAsset(ctx, in, opts...)
}

func (g *GrpcClient) TransferAsset2(ctx context.Context, in *core.TransferAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.TransferAsset2(ctx, in, opts...)
}

func (g *GrpcClient) ParticipateAssetIssue(ctx context.Context, in *core.ParticipateAssetIssueContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ParticipateAssetIssue(ctx, in, opts...)
}

func (g *GrpcClient) ParticipateAssetIssue2(ctx context.Context, in *core.ParticipateAssetIssueContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ParticipateAssetIssue2(ctx, in, opts...)
}

func (g *GrpcClient) FreezeBalance(ctx context.Context, in *core.FreezeBalanceContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.FreezeBalance(ctx, in, opts...)
}

func (g *GrpcClient) FreezeBalance2(ctx context.Context, in *core.FreezeBalanceContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.FreezeBalance2(ctx, in, opts...)
}

func (g *GrpcClient) UnfreezeBalance(ctx context.Context, in *core.UnfreezeBalanceContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UnfreezeBalance(ctx, in, opts...)
}

func (g *GrpcClient) UnfreezeBalance2(ctx context.Context, in *core.UnfreezeBalanceContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UnfreezeBalance2(ctx, in, opts...)
}

func (g *GrpcClient) UnfreezeAsset(ctx context.Context, in *core.UnfreezeAssetContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UnfreezeAsset(ctx, in, opts...)
}

func (g *GrpcClient) UnfreezeAsset2(ctx context.Context, in *core.UnfreezeAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UnfreezeAsset2(ctx, in, opts...)
}

func (g *GrpcClient) WithdrawBalance(ctx context.Context, in *core.WithdrawBalanceContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.WithdrawBalance(ctx, in, opts...)
}

func (g *GrpcClient) WithdrawBalance2(ctx context.Context, in *core.WithdrawBalanceContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.WithdrawBalance2(ctx, in, opts...)
}

func (g *GrpcClient) UpdateAsset(ctx context.Context, in *core.UpdateAssetContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateAsset(ctx, in, opts...)
}

func (g *GrpcClient) UpdateAsset2(ctx context.Context, in *core.UpdateAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateAsset2(ctx, in, opts...)
}

func (g *GrpcClient) ProposalCreate(ctx context.Context, in *core.ProposalCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ProposalCreate(ctx, in, opts...)
}

func (g *GrpcClient) ProposalApprove(ctx context.Context, in *core.ProposalApproveContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ProposalApprove(ctx, in, opts...)
}

func (g *GrpcClient) ProposalDelete(ctx context.Context, in *core.ProposalDeleteContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ProposalDelete(ctx, in, opts...)
}

func (g *GrpcClient) BuyStorage(ctx context.Context, in *core.BuyStorageContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.BuyStorage(ctx, in, opts...)
}

func (g *GrpcClient) BuyStorageBytes(ctx context.Context, in *core.BuyStorageBytesContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.BuyStorageBytes(ctx, in, opts...)
}

func (g *GrpcClient) SellStorage(ctx context.Context, in *core.SellStorageContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.SellStorage(ctx, in, opts...)
}

func (g *GrpcClient) ExchangeCreate(ctx context.Context, in *core.ExchangeCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ExchangeCreate(ctx, in, opts...)
}

func (g *GrpcClient) ExchangeInject(ctx context.Context, in *core.ExchangeInjectContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ExchangeInject(ctx, in, opts...)
}

func (g *GrpcClient) ExchangeWithdraw(ctx context.Context, in *core.ExchangeWithdrawContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ExchangeWithdraw(ctx, in, opts...)
}

func (g *GrpcClient) ExchangeTransaction(ctx context.Context, in *core.ExchangeTransactionContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ExchangeTransaction(ctx, in, opts...)
}

func (g *GrpcClient) MarketSellAsset(ctx context.Context, in *core.MarketSellAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.MarketSellAsset(ctx, in, opts...)
}

func (g *GrpcClient) MarketCancelOrder(ctx context.Context, in *core.MarketCancelOrderContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.MarketCancelOrder(ctx, in, opts...)
}

func (g *GrpcClient) GetMarketOrderById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.MarketOrder, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetMarketOrderById(ctx, in, opts...)
}

func (g *GrpcClient) GetMarketOrderByAccount(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.MarketOrderList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetMarketOrderByAccount(ctx, in, opts...)
}

func (g *GrpcClient) GetMarketPriceByPair(ctx context.Context, in *core.MarketOrderPair, opts ...grpc.CallOption) (*core.MarketPriceList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetMarketPriceByPair(ctx, in, opts...)
}

func (g *GrpcClient) GetMarketOrderListByPair(ctx context.Context, in *core.MarketOrderPair, opts ...grpc.CallOption) (*core.MarketOrderList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetMarketOrderListByPair(ctx, in, opts...)
}

func (g *GrpcClient) GetMarketPairList(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.MarketOrderPairList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetMarketPairList(ctx, in, opts...)
}

func (g *GrpcClient) ListNodes(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NodeList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ListNodes(ctx, in, opts...)
}

func (g *GrpcClient) GetAssetIssueByAccount(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAssetIssueByAccount(ctx, in, opts...)
}

func (g *GrpcClient) GetAccountNet(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*api.AccountNetMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAccountNet(ctx, in, opts...)
}

func (g *GrpcClient) GetAccountResource(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*api.AccountResourceMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAccountResource(ctx, in, opts...)
}

func (g *GrpcClient) GetAssetIssueByName(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.AssetIssueContract, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAssetIssueByName(ctx, in, opts...)
}

func (g *GrpcClient) GetAssetIssueListByName(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAssetIssueListByName(ctx, in, opts...)
}

func (g *GrpcClient) GetAssetIssueById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.AssetIssueContract, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAssetIssueById(ctx, in, opts...)
}

func (g *GrpcClient) GetNowBlock(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.Block, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetNowBlock(ctx, in, opts...)
}

func (g *GrpcClient) GetNowBlock2(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.BlockExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetNowBlock2(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockByNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*core.Block, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockByNum(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockByNum2(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.BlockExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockByNum2(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionCountByBlockNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionCountByBlockNum(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Block, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockById(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockByLimitNext(ctx context.Context, in *api.BlockLimit, opts ...grpc.CallOption) (*api.BlockList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockByLimitNext(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockByLimitNext2(ctx context.Context, in *api.BlockLimit, opts ...grpc.CallOption) (*api.BlockListExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockByLimitNext2(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockByLatestNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.BlockList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockByLatestNum(ctx, in, opts...)
}

func (g *GrpcClient) GetBlockByLatestNum2(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.BlockListExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlockByLatestNum2(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionById(ctx, in, opts...)
}

func (g *GrpcClient) DeployContract(ctx context.Context, in *core.CreateSmartContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.DeployContract(ctx, in, opts...)
}

func (g *GrpcClient) GetContract(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.SmartContract, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetContract(ctx, in, opts...)
}

func (g *GrpcClient) GetContractInfo(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.SmartContractDataWrapper, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetContractInfo(ctx, in, opts...)
}

func (g *GrpcClient) TriggerContract(ctx context.Context, in *core.TriggerSmartContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.TriggerContract(ctx, in, opts...)
}

func (g *GrpcClient) TriggerConstantContract(ctx context.Context, in *core.TriggerSmartContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.TriggerConstantContract(ctx, in, opts...)
}

func (g *GrpcClient) ClearContractABI(ctx context.Context, in *core.ClearABIContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ClearContractABI(ctx, in, opts...)
}

func (g *GrpcClient) ListWitnesses(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.WitnessList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ListWitnesses(ctx, in, opts...)
}

func (g *GrpcClient) GetDelegatedResource(ctx context.Context, in *api.DelegatedResourceMessage, opts ...grpc.CallOption) (*api.DelegatedResourceList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetDelegatedResource(ctx, in, opts...)
}

func (g *GrpcClient) GetDelegatedResourceAccountIndex(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.DelegatedResourceAccountIndex, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetDelegatedResourceAccountIndex(ctx, in, opts...)
}

func (g *GrpcClient) ListProposals(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.ProposalList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ListProposals(ctx, in, opts...)
}

func (g *GrpcClient) GetPaginatedProposalList(ctx context.Context, in *api.PaginatedMessage, opts ...grpc.CallOption) (*api.ProposalList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetPaginatedProposalList(ctx, in, opts...)
}

func (g *GrpcClient) GetProposalById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Proposal, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetProposalById(ctx, in, opts...)
}

func (g *GrpcClient) ListExchanges(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.ExchangeList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ListExchanges(ctx, in, opts...)
}

func (g *GrpcClient) GetPaginatedExchangeList(ctx context.Context, in *api.PaginatedMessage, opts ...grpc.CallOption) (*api.ExchangeList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetPaginatedExchangeList(ctx, in, opts...)
}

func (g *GrpcClient) GetExchangeById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Exchange, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetExchangeById(ctx, in, opts...)
}

func (g *GrpcClient) GetChainParameters(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.ChainParameters, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetChainParameters(ctx, in, opts...)
}

func (g *GrpcClient) GetAssetIssueList(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAssetIssueList(ctx, in, opts...)
}

func (g *GrpcClient) GetPaginatedAssetIssueList(ctx context.Context, in *api.PaginatedMessage, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetPaginatedAssetIssueList(ctx, in, opts...)
}

func (g *GrpcClient) TotalTransaction(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.TotalTransaction(ctx, in, opts...)
}

func (g *GrpcClient) GetNextMaintenanceTime(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetNextMaintenanceTime(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionSign(ctx context.Context, in *core.TransactionSign, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionSign(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionSign2(ctx context.Context, in *core.TransactionSign, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionSign2(ctx, in, opts...)
}

func (g *GrpcClient) CreateAddress(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateAddress(ctx, in, opts...)
}

func (g *GrpcClient) EasyTransferAsset(ctx context.Context, in *api.EasyTransferAssetMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.EasyTransferAsset(ctx, in, opts...)
}

func (g *GrpcClient) EasyTransferAssetByPrivate(ctx context.Context, in *api.EasyTransferAssetByPrivateMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.EasyTransferAssetByPrivate(ctx, in, opts...)
}

func (g *GrpcClient) EasyTransfer(ctx context.Context, in *api.EasyTransferMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.EasyTransfer(ctx, in, opts...)
}

func (g *GrpcClient) EasyTransferByPrivate(ctx context.Context, in *api.EasyTransferByPrivateMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.EasyTransferByPrivate(ctx, in, opts...)
}

func (g *GrpcClient) GenerateAddress(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.AddressPrKeyPairMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GenerateAddress(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionInfoById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.TransactionInfo, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionInfoById(ctx, in, opts...)
}

func (g *GrpcClient) AccountPermissionUpdate(ctx context.Context, in *core.AccountPermissionUpdateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.AccountPermissionUpdate(ctx, in, opts...)
}

func (g *GrpcClient) AddSign(ctx context.Context, in *core.TransactionSign, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.AddSign(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionSignWeight(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.TransactionSignWeight, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionSignWeight(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionApprovedList(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.TransactionApprovedList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionApprovedList(ctx, in, opts...)
}

func (g *GrpcClient) GetNodeInfo(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.NodeInfo, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetNodeInfo(ctx, in, opts...)
}

func (g *GrpcClient) GetRewardInfo(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetRewardInfo(ctx, in, opts...)
}

func (g *GrpcClient) GetBrokerageInfo(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBrokerageInfo(ctx, in, opts...)
}

func (g *GrpcClient) UpdateBrokerage(ctx context.Context, in *core.UpdateBrokerageContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.UpdateBrokerage(ctx, in, opts...)
}

func (g *GrpcClient) CreateShieldedTransaction(ctx context.Context, in *api.PrivateParameters, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateShieldedTransaction(ctx, in, opts...)
}

func (g *GrpcClient) GetMerkleTreeVoucherInfo(ctx context.Context, in *core.OutputPointInfo, opts ...grpc.CallOption) (*core.IncrementalMerkleVoucherInfo, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetMerkleTreeVoucherInfo(ctx, in, opts...)
}

func (g *GrpcClient) ScanNoteByIvk(ctx context.Context, in *api.IvkDecryptParameters, opts ...grpc.CallOption) (*api.DecryptNotes, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ScanNoteByIvk(ctx, in, opts...)
}

func (g *GrpcClient) ScanAndMarkNoteByIvk(ctx context.Context, in *api.IvkDecryptAndMarkParameters, opts ...grpc.CallOption) (*api.DecryptNotesMarked, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ScanAndMarkNoteByIvk(ctx, in, opts...)
}

func (g *GrpcClient) ScanNoteByOvk(ctx context.Context, in *api.OvkDecryptParameters, opts ...grpc.CallOption) (*api.DecryptNotes, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ScanNoteByOvk(ctx, in, opts...)
}

func (g *GrpcClient) GetSpendingKey(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetSpendingKey(ctx, in, opts...)
}

func (g *GrpcClient) GetExpandedSpendingKey(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.ExpandedSpendingKeyMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetExpandedSpendingKey(ctx, in, opts...)
}

func (g *GrpcClient) GetAkFromAsk(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetAkFromAsk(ctx, in, opts...)
}

func (g *GrpcClient) GetNkFromNsk(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetNkFromNsk(ctx, in, opts...)
}

func (g *GrpcClient) GetIncomingViewingKey(ctx context.Context, in *api.ViewingKeyMessage, opts ...grpc.CallOption) (*api.IncomingViewingKeyMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetIncomingViewingKey(ctx, in, opts...)
}

func (g *GrpcClient) GetDiversifier(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.DiversifierMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetDiversifier(ctx, in, opts...)
}

func (g *GrpcClient) GetNewShieldedAddress(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.ShieldedAddressInfo, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetNewShieldedAddress(ctx, in, opts...)
}

func (g *GrpcClient) GetZenPaymentAddress(ctx context.Context, in *api.IncomingViewingKeyDiversifierMessage, opts ...grpc.CallOption) (*api.PaymentAddressMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetZenPaymentAddress(ctx, in, opts...)
}

func (g *GrpcClient) GetRcm(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetRcm(ctx, in, opts...)
}

func (g *GrpcClient) IsSpend(ctx context.Context, in *api.NoteParameters, opts ...grpc.CallOption) (*api.SpendResult, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.IsSpend(ctx, in, opts...)
}

func (g *GrpcClient) CreateShieldedTransactionWithoutSpendAuthSig(ctx context.Context, in *api.PrivateParametersWithoutAsk, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateShieldedTransactionWithoutSpendAuthSig(ctx, in, opts...)
}

func (g *GrpcClient) GetShieldTransactionHash(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetShieldTransactionHash(ctx, in, opts...)
}

func (g *GrpcClient) CreateSpendAuthSig(ctx context.Context, in *api.SpendAuthSigParameters, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateSpendAuthSig(ctx, in, opts...)
}

func (g *GrpcClient) CreateShieldNullifier(ctx context.Context, in *api.NfParameters, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateShieldNullifier(ctx, in, opts...)
}

func (g *GrpcClient) CreateShieldedContractParameters(ctx context.Context, in *api.PrivateShieldedTRC20Parameters, opts ...grpc.CallOption) (*api.ShieldedTRC20Parameters, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateShieldedContractParameters(ctx, in, opts...)
}

func (g *GrpcClient) CreateShieldedContractParametersWithoutAsk(ctx context.Context, in *api.PrivateShieldedTRC20ParametersWithoutAsk, opts ...grpc.CallOption) (*api.ShieldedTRC20Parameters, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateShieldedContractParametersWithoutAsk(ctx, in, opts...)
}

func (g *GrpcClient) ScanShieldedTRC20NotesByIvk(ctx context.Context, in *api.IvkDecryptTRC20Parameters, opts ...grpc.CallOption) (*api.DecryptNotesTRC20, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ScanShieldedTRC20NotesByIvk(ctx, in, opts...)
}

func (g *GrpcClient) ScanShieldedTRC20NotesByOvk(ctx context.Context, in *api.OvkDecryptTRC20Parameters, opts ...grpc.CallOption) (*api.DecryptNotesTRC20, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.ScanShieldedTRC20NotesByOvk(ctx, in, opts...)
}

func (g *GrpcClient) IsShieldedTRC20ContractNoteSpent(ctx context.Context, in *api.NfTRC20Parameters, opts ...grpc.CallOption) (*api.NullifierResult, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.IsShieldedTRC20ContractNoteSpent(ctx, in, opts...)
}

func (g *GrpcClient) GetTriggerInputForShieldedTRC20Contract(ctx context.Context, in *api.ShieldedTRC20TriggerContractParameters, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTriggerInputForShieldedTRC20Contract(ctx, in, opts...)
}

func (g *GrpcClient) CreateCommonTransaction(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.CreateCommonTransaction(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionInfoByBlockNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.TransactionInfoList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionInfoByBlockNum(ctx, in, opts...)
}

func (g *GrpcClient) GetBurnTrx(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBurnTrx(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionFromPending(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionFromPending(ctx, in, opts...)
}

func (g *GrpcClient) GetTransactionListFromPending(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.TransactionIdList, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetTransactionListFromPending(ctx, in, opts...)
}

func (g *GrpcClient) GetPendingSize(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetPendingSize(ctx, in, opts...)
}

func (g *GrpcClient) GetBlock(ctx context.Context, in *api.BlockReq, opts ...grpc.CallOption) (*api.BlockExtention, error) {
	ctx, cancel := g.makeContext(ctx)
	defer cancel()
	return g.client.GetBlock(ctx, in, opts...)
}
