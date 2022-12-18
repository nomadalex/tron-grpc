package client

import (
	"context"
	"github.com/fullstackwang/tron-grpc/api"
	"github.com/fullstackwang/tron-grpc/core"
	"google.golang.org/grpc"
)

func (c *Client) GetAccount(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*core.Account, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAccount(ctx, in, opts...)
}

func (c *Client) GetAccountById(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*core.Account, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAccountById(ctx, in, opts...)
}

func (c *Client) GetAccountBalance(ctx context.Context, in *core.AccountBalanceRequest, opts ...grpc.CallOption) (*core.AccountBalanceResponse, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAccountBalance(ctx, in, opts...)
}

func (c *Client) GetBlockBalanceTrace(ctx context.Context, in *core.BlockBalanceTrace_BlockIdentifier, opts ...grpc.CallOption) (*core.BlockBalanceTrace, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockBalanceTrace(ctx, in, opts...)
}

func (c *Client) CreateTransaction(ctx context.Context, in *core.TransferContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateTransaction(ctx, in, opts...)
}

func (c *Client) CreateTransaction2(ctx context.Context, in *core.TransferContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateTransaction2(ctx, in, opts...)
}

func (c *Client) BroadcastTransaction(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.Return, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.BroadcastTransaction(ctx, in, opts...)
}

func (c *Client) UpdateAccount(ctx context.Context, in *core.AccountUpdateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateAccount(ctx, in, opts...)
}

func (c *Client) SetAccountId(ctx context.Context, in *core.SetAccountIdContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.SetAccountId(ctx, in, opts...)
}

func (c *Client) UpdateAccount2(ctx context.Context, in *core.AccountUpdateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateAccount2(ctx, in, opts...)
}

func (c *Client) VoteWitnessAccount(ctx context.Context, in *core.VoteWitnessContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.VoteWitnessAccount(ctx, in, opts...)
}

func (c *Client) UpdateSetting(ctx context.Context, in *core.UpdateSettingContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateSetting(ctx, in, opts...)
}

func (c *Client) UpdateEnergyLimit(ctx context.Context, in *core.UpdateEnergyLimitContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateEnergyLimit(ctx, in, opts...)
}

func (c *Client) VoteWitnessAccount2(ctx context.Context, in *core.VoteWitnessContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.VoteWitnessAccount2(ctx, in, opts...)
}

func (c *Client) CreateAssetIssue(ctx context.Context, in *core.AssetIssueContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateAssetIssue(ctx, in, opts...)
}

func (c *Client) CreateAssetIssue2(ctx context.Context, in *core.AssetIssueContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateAssetIssue2(ctx, in, opts...)
}

func (c *Client) UpdateWitness(ctx context.Context, in *core.WitnessUpdateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateWitness(ctx, in, opts...)
}

func (c *Client) UpdateWitness2(ctx context.Context, in *core.WitnessUpdateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateWitness2(ctx, in, opts...)
}

func (c *Client) CreateAccount(ctx context.Context, in *core.AccountCreateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateAccount(ctx, in, opts...)
}

func (c *Client) CreateAccount2(ctx context.Context, in *core.AccountCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateAccount2(ctx, in, opts...)
}

func (c *Client) CreateWitness(ctx context.Context, in *core.WitnessCreateContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateWitness(ctx, in, opts...)
}

func (c *Client) CreateWitness2(ctx context.Context, in *core.WitnessCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateWitness2(ctx, in, opts...)
}

func (c *Client) TransferAsset(ctx context.Context, in *core.TransferAssetContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.TransferAsset(ctx, in, opts...)
}

func (c *Client) TransferAsset2(ctx context.Context, in *core.TransferAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.TransferAsset2(ctx, in, opts...)
}

func (c *Client) ParticipateAssetIssue(ctx context.Context, in *core.ParticipateAssetIssueContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ParticipateAssetIssue(ctx, in, opts...)
}

func (c *Client) ParticipateAssetIssue2(ctx context.Context, in *core.ParticipateAssetIssueContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ParticipateAssetIssue2(ctx, in, opts...)
}

func (c *Client) FreezeBalance(ctx context.Context, in *core.FreezeBalanceContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.FreezeBalance(ctx, in, opts...)
}

func (c *Client) FreezeBalance2(ctx context.Context, in *core.FreezeBalanceContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.FreezeBalance2(ctx, in, opts...)
}

func (c *Client) UnfreezeBalance(ctx context.Context, in *core.UnfreezeBalanceContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UnfreezeBalance(ctx, in, opts...)
}

func (c *Client) UnfreezeBalance2(ctx context.Context, in *core.UnfreezeBalanceContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UnfreezeBalance2(ctx, in, opts...)
}

func (c *Client) UnfreezeAsset(ctx context.Context, in *core.UnfreezeAssetContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UnfreezeAsset(ctx, in, opts...)
}

func (c *Client) UnfreezeAsset2(ctx context.Context, in *core.UnfreezeAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UnfreezeAsset2(ctx, in, opts...)
}

func (c *Client) WithdrawBalance(ctx context.Context, in *core.WithdrawBalanceContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.WithdrawBalance(ctx, in, opts...)
}

func (c *Client) WithdrawBalance2(ctx context.Context, in *core.WithdrawBalanceContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.WithdrawBalance2(ctx, in, opts...)
}

func (c *Client) UpdateAsset(ctx context.Context, in *core.UpdateAssetContract, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateAsset(ctx, in, opts...)
}

func (c *Client) UpdateAsset2(ctx context.Context, in *core.UpdateAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateAsset2(ctx, in, opts...)
}

func (c *Client) ProposalCreate(ctx context.Context, in *core.ProposalCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ProposalCreate(ctx, in, opts...)
}

func (c *Client) ProposalApprove(ctx context.Context, in *core.ProposalApproveContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ProposalApprove(ctx, in, opts...)
}

func (c *Client) ProposalDelete(ctx context.Context, in *core.ProposalDeleteContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ProposalDelete(ctx, in, opts...)
}

func (c *Client) BuyStorage(ctx context.Context, in *core.BuyStorageContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.BuyStorage(ctx, in, opts...)
}

func (c *Client) BuyStorageBytes(ctx context.Context, in *core.BuyStorageBytesContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.BuyStorageBytes(ctx, in, opts...)
}

func (c *Client) SellStorage(ctx context.Context, in *core.SellStorageContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.SellStorage(ctx, in, opts...)
}

func (c *Client) ExchangeCreate(ctx context.Context, in *core.ExchangeCreateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ExchangeCreate(ctx, in, opts...)
}

func (c *Client) ExchangeInject(ctx context.Context, in *core.ExchangeInjectContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ExchangeInject(ctx, in, opts...)
}

func (c *Client) ExchangeWithdraw(ctx context.Context, in *core.ExchangeWithdrawContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ExchangeWithdraw(ctx, in, opts...)
}

func (c *Client) ExchangeTransaction(ctx context.Context, in *core.ExchangeTransactionContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ExchangeTransaction(ctx, in, opts...)
}

func (c *Client) MarketSellAsset(ctx context.Context, in *core.MarketSellAssetContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.MarketSellAsset(ctx, in, opts...)
}

func (c *Client) MarketCancelOrder(ctx context.Context, in *core.MarketCancelOrderContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.MarketCancelOrder(ctx, in, opts...)
}

func (c *Client) GetMarketOrderById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.MarketOrder, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetMarketOrderById(ctx, in, opts...)
}

func (c *Client) GetMarketOrderByAccount(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.MarketOrderList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetMarketOrderByAccount(ctx, in, opts...)
}

func (c *Client) GetMarketPriceByPair(ctx context.Context, in *core.MarketOrderPair, opts ...grpc.CallOption) (*core.MarketPriceList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetMarketPriceByPair(ctx, in, opts...)
}

func (c *Client) GetMarketOrderListByPair(ctx context.Context, in *core.MarketOrderPair, opts ...grpc.CallOption) (*core.MarketOrderList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetMarketOrderListByPair(ctx, in, opts...)
}

func (c *Client) GetMarketPairList(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.MarketOrderPairList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetMarketPairList(ctx, in, opts...)
}

func (c *Client) ListNodes(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NodeList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ListNodes(ctx, in, opts...)
}

func (c *Client) GetAssetIssueByAccount(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAssetIssueByAccount(ctx, in, opts...)
}

func (c *Client) GetAccountNet(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*api.AccountNetMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAccountNet(ctx, in, opts...)
}

func (c *Client) GetAccountResource(ctx context.Context, in *core.Account, opts ...grpc.CallOption) (*api.AccountResourceMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAccountResource(ctx, in, opts...)
}

func (c *Client) GetAssetIssueByName(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.AssetIssueContract, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAssetIssueByName(ctx, in, opts...)
}

func (c *Client) GetAssetIssueListByName(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAssetIssueListByName(ctx, in, opts...)
}

func (c *Client) GetAssetIssueById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.AssetIssueContract, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAssetIssueById(ctx, in, opts...)
}

func (c *Client) GetNowBlock(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.Block, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetNowBlock(ctx, in, opts...)
}

func (c *Client) GetNowBlock2(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.BlockExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetNowBlock2(ctx, in, opts...)
}

func (c *Client) GetBlockByNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*core.Block, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockByNum(ctx, in, opts...)
}

func (c *Client) GetBlockByNum2(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.BlockExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockByNum2(ctx, in, opts...)
}

func (c *Client) GetTransactionCountByBlockNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionCountByBlockNum(ctx, in, opts...)
}

func (c *Client) GetBlockById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Block, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockById(ctx, in, opts...)
}

func (c *Client) GetBlockByLimitNext(ctx context.Context, in *api.BlockLimit, opts ...grpc.CallOption) (*api.BlockList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockByLimitNext(ctx, in, opts...)
}

func (c *Client) GetBlockByLimitNext2(ctx context.Context, in *api.BlockLimit, opts ...grpc.CallOption) (*api.BlockListExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockByLimitNext2(ctx, in, opts...)
}

func (c *Client) GetBlockByLatestNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.BlockList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockByLatestNum(ctx, in, opts...)
}

func (c *Client) GetBlockByLatestNum2(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.BlockListExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlockByLatestNum2(ctx, in, opts...)
}

func (c *Client) GetTransactionById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionById(ctx, in, opts...)
}

func (c *Client) DeployContract(ctx context.Context, in *core.CreateSmartContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.DeployContract(ctx, in, opts...)
}

func (c *Client) GetContract(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.SmartContract, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetContract(ctx, in, opts...)
}

func (c *Client) GetContractInfo(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.SmartContractDataWrapper, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetContractInfo(ctx, in, opts...)
}

func (c *Client) TriggerContract(ctx context.Context, in *core.TriggerSmartContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.TriggerContract(ctx, in, opts...)
}

func (c *Client) TriggerConstantContract(ctx context.Context, in *core.TriggerSmartContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.TriggerConstantContract(ctx, in, opts...)
}

func (c *Client) ClearContractABI(ctx context.Context, in *core.ClearABIContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ClearContractABI(ctx, in, opts...)
}

func (c *Client) ListWitnesses(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.WitnessList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ListWitnesses(ctx, in, opts...)
}

func (c *Client) GetDelegatedResource(ctx context.Context, in *api.DelegatedResourceMessage, opts ...grpc.CallOption) (*api.DelegatedResourceList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetDelegatedResource(ctx, in, opts...)
}

func (c *Client) GetDelegatedResourceAccountIndex(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.DelegatedResourceAccountIndex, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetDelegatedResourceAccountIndex(ctx, in, opts...)
}

func (c *Client) ListProposals(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.ProposalList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ListProposals(ctx, in, opts...)
}

func (c *Client) GetPaginatedProposalList(ctx context.Context, in *api.PaginatedMessage, opts ...grpc.CallOption) (*api.ProposalList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetPaginatedProposalList(ctx, in, opts...)
}

func (c *Client) GetProposalById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Proposal, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetProposalById(ctx, in, opts...)
}

func (c *Client) ListExchanges(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.ExchangeList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ListExchanges(ctx, in, opts...)
}

func (c *Client) GetPaginatedExchangeList(ctx context.Context, in *api.PaginatedMessage, opts ...grpc.CallOption) (*api.ExchangeList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetPaginatedExchangeList(ctx, in, opts...)
}

func (c *Client) GetExchangeById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Exchange, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetExchangeById(ctx, in, opts...)
}

func (c *Client) GetChainParameters(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.ChainParameters, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetChainParameters(ctx, in, opts...)
}

func (c *Client) GetAssetIssueList(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAssetIssueList(ctx, in, opts...)
}

func (c *Client) GetPaginatedAssetIssueList(ctx context.Context, in *api.PaginatedMessage, opts ...grpc.CallOption) (*api.AssetIssueList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetPaginatedAssetIssueList(ctx, in, opts...)
}

func (c *Client) TotalTransaction(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.TotalTransaction(ctx, in, opts...)
}

func (c *Client) GetNextMaintenanceTime(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetNextMaintenanceTime(ctx, in, opts...)
}

func (c *Client) GetTransactionSign(ctx context.Context, in *core.TransactionSign, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionSign(ctx, in, opts...)
}

func (c *Client) GetTransactionSign2(ctx context.Context, in *core.TransactionSign, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionSign2(ctx, in, opts...)
}

func (c *Client) CreateAddress(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateAddress(ctx, in, opts...)
}

func (c *Client) EasyTransferAsset(ctx context.Context, in *api.EasyTransferAssetMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.EasyTransferAsset(ctx, in, opts...)
}

func (c *Client) EasyTransferAssetByPrivate(ctx context.Context, in *api.EasyTransferAssetByPrivateMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.EasyTransferAssetByPrivate(ctx, in, opts...)
}

func (c *Client) EasyTransfer(ctx context.Context, in *api.EasyTransferMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.EasyTransfer(ctx, in, opts...)
}

func (c *Client) EasyTransferByPrivate(ctx context.Context, in *api.EasyTransferByPrivateMessage, opts ...grpc.CallOption) (*api.EasyTransferResponse, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.EasyTransferByPrivate(ctx, in, opts...)
}

func (c *Client) GenerateAddress(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.AddressPrKeyPairMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GenerateAddress(ctx, in, opts...)
}

func (c *Client) GetTransactionInfoById(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.TransactionInfo, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionInfoById(ctx, in, opts...)
}

func (c *Client) AccountPermissionUpdate(ctx context.Context, in *core.AccountPermissionUpdateContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.AccountPermissionUpdate(ctx, in, opts...)
}

func (c *Client) AddSign(ctx context.Context, in *core.TransactionSign, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.AddSign(ctx, in, opts...)
}

func (c *Client) GetTransactionSignWeight(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.TransactionSignWeight, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionSignWeight(ctx, in, opts...)
}

func (c *Client) GetTransactionApprovedList(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.TransactionApprovedList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionApprovedList(ctx, in, opts...)
}

func (c *Client) GetNodeInfo(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*core.NodeInfo, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetNodeInfo(ctx, in, opts...)
}

func (c *Client) GetRewardInfo(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetRewardInfo(ctx, in, opts...)
}

func (c *Client) GetBrokerageInfo(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBrokerageInfo(ctx, in, opts...)
}

func (c *Client) UpdateBrokerage(ctx context.Context, in *core.UpdateBrokerageContract, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.UpdateBrokerage(ctx, in, opts...)
}

func (c *Client) CreateShieldedTransaction(ctx context.Context, in *api.PrivateParameters, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateShieldedTransaction(ctx, in, opts...)
}

func (c *Client) GetMerkleTreeVoucherInfo(ctx context.Context, in *core.OutputPointInfo, opts ...grpc.CallOption) (*core.IncrementalMerkleVoucherInfo, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetMerkleTreeVoucherInfo(ctx, in, opts...)
}

func (c *Client) ScanNoteByIvk(ctx context.Context, in *api.IvkDecryptParameters, opts ...grpc.CallOption) (*api.DecryptNotes, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ScanNoteByIvk(ctx, in, opts...)
}

func (c *Client) ScanAndMarkNoteByIvk(ctx context.Context, in *api.IvkDecryptAndMarkParameters, opts ...grpc.CallOption) (*api.DecryptNotesMarked, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ScanAndMarkNoteByIvk(ctx, in, opts...)
}

func (c *Client) ScanNoteByOvk(ctx context.Context, in *api.OvkDecryptParameters, opts ...grpc.CallOption) (*api.DecryptNotes, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ScanNoteByOvk(ctx, in, opts...)
}

func (c *Client) GetSpendingKey(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetSpendingKey(ctx, in, opts...)
}

func (c *Client) GetExpandedSpendingKey(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.ExpandedSpendingKeyMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetExpandedSpendingKey(ctx, in, opts...)
}

func (c *Client) GetAkFromAsk(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetAkFromAsk(ctx, in, opts...)
}

func (c *Client) GetNkFromNsk(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetNkFromNsk(ctx, in, opts...)
}

func (c *Client) GetIncomingViewingKey(ctx context.Context, in *api.ViewingKeyMessage, opts ...grpc.CallOption) (*api.IncomingViewingKeyMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetIncomingViewingKey(ctx, in, opts...)
}

func (c *Client) GetDiversifier(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.DiversifierMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetDiversifier(ctx, in, opts...)
}

func (c *Client) GetNewShieldedAddress(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.ShieldedAddressInfo, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetNewShieldedAddress(ctx, in, opts...)
}

func (c *Client) GetZenPaymentAddress(ctx context.Context, in *api.IncomingViewingKeyDiversifierMessage, opts ...grpc.CallOption) (*api.PaymentAddressMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetZenPaymentAddress(ctx, in, opts...)
}

func (c *Client) GetRcm(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetRcm(ctx, in, opts...)
}

func (c *Client) IsSpend(ctx context.Context, in *api.NoteParameters, opts ...grpc.CallOption) (*api.SpendResult, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.IsSpend(ctx, in, opts...)
}

func (c *Client) CreateShieldedTransactionWithoutSpendAuthSig(ctx context.Context, in *api.PrivateParametersWithoutAsk, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateShieldedTransactionWithoutSpendAuthSig(ctx, in, opts...)
}

func (c *Client) GetShieldTransactionHash(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetShieldTransactionHash(ctx, in, opts...)
}

func (c *Client) CreateSpendAuthSig(ctx context.Context, in *api.SpendAuthSigParameters, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateSpendAuthSig(ctx, in, opts...)
}

func (c *Client) CreateShieldNullifier(ctx context.Context, in *api.NfParameters, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateShieldNullifier(ctx, in, opts...)
}

func (c *Client) CreateShieldedContractParameters(ctx context.Context, in *api.PrivateShieldedTRC20Parameters, opts ...grpc.CallOption) (*api.ShieldedTRC20Parameters, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateShieldedContractParameters(ctx, in, opts...)
}

func (c *Client) CreateShieldedContractParametersWithoutAsk(ctx context.Context, in *api.PrivateShieldedTRC20ParametersWithoutAsk, opts ...grpc.CallOption) (*api.ShieldedTRC20Parameters, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateShieldedContractParametersWithoutAsk(ctx, in, opts...)
}

func (c *Client) ScanShieldedTRC20NotesByIvk(ctx context.Context, in *api.IvkDecryptTRC20Parameters, opts ...grpc.CallOption) (*api.DecryptNotesTRC20, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ScanShieldedTRC20NotesByIvk(ctx, in, opts...)
}

func (c *Client) ScanShieldedTRC20NotesByOvk(ctx context.Context, in *api.OvkDecryptTRC20Parameters, opts ...grpc.CallOption) (*api.DecryptNotesTRC20, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.ScanShieldedTRC20NotesByOvk(ctx, in, opts...)
}

func (c *Client) IsShieldedTRC20ContractNoteSpent(ctx context.Context, in *api.NfTRC20Parameters, opts ...grpc.CallOption) (*api.NullifierResult, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.IsShieldedTRC20ContractNoteSpent(ctx, in, opts...)
}

func (c *Client) GetTriggerInputForShieldedTRC20Contract(ctx context.Context, in *api.ShieldedTRC20TriggerContractParameters, opts ...grpc.CallOption) (*api.BytesMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTriggerInputForShieldedTRC20Contract(ctx, in, opts...)
}

func (c *Client) CreateCommonTransaction(ctx context.Context, in *core.Transaction, opts ...grpc.CallOption) (*api.TransactionExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.CreateCommonTransaction(ctx, in, opts...)
}

func (c *Client) GetTransactionInfoByBlockNum(ctx context.Context, in *api.NumberMessage, opts ...grpc.CallOption) (*api.TransactionInfoList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionInfoByBlockNum(ctx, in, opts...)
}

func (c *Client) GetBurnTrx(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBurnTrx(ctx, in, opts...)
}

func (c *Client) GetTransactionFromPending(ctx context.Context, in *api.BytesMessage, opts ...grpc.CallOption) (*core.Transaction, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionFromPending(ctx, in, opts...)
}

func (c *Client) GetTransactionListFromPending(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.TransactionIdList, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetTransactionListFromPending(ctx, in, opts...)
}

func (c *Client) GetPendingSize(ctx context.Context, in *api.EmptyMessage, opts ...grpc.CallOption) (*api.NumberMessage, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetPendingSize(ctx, in, opts...)
}

func (c *Client) GetBlock(ctx context.Context, in *api.BlockReq, opts ...grpc.CallOption) (*api.BlockExtention, error) {
	ctx, cancel := c.makeContext(ctx)
	defer cancel()
	return c.client.GetBlock(ctx, in, opts...)
}
