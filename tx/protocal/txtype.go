/*
 *  MIT License
 *  Copyright (c) 2019-2020 niels.wang
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 *
 */
//@Description 这里列举了所有NULS网络支持的交易类型
package txprotocal

const (
	//coinbase 共识奖励交易类型
	TX_TYPE_COIN_BASE = 1

	//the type of the transfer transaction
	TX_TYPE_TRANSFER = 2

	//设置账户别名
	//*Set the transaction type of account alias.
	TX_TYPE_ACCOUNT_ALIAS = 3

	//新建共识节点
	TX_TYPE_REGISTER_AGENT = 4

	//委托参与共识
	TX_TYPE_DEPOSIT = 5

	//取消委托
	TX_TYPE_CANCEL_DEPOSIT = 6

	//黄牌惩罚交易
	TX_TYPE_YELLOW_PUNISH = 7

	//红牌惩罚交易
	TX_TYPE_RED_PUNISH = 8

	//注销共识节点交易类型
	TX_TYPE_STOP_AGENT = 9

	//跨链转账交易类型
	TX_TYPE_CROSS_CHAIN = 10

	//注册平行链交易类型
	TX_TYPE_REGISTER_CHAIN_AND_ASSET = 11

	//从NULS网络中注销一条平行链的交易的类型
	TX_TYPE_DESTROY_CHAIN_AND_ASSET = 12

	//为平行链登记一种资产的交易类型
	TX_TYPE_ADD_ASSET_TO_CHAIN = 13

	//删除链上资产交易类型
	TX_TYPE_REMOVE_ASSET_FROM_CHAIN = 14

	//创建智能合约交易的类型
	TX_TYPE_CREATE_CONTRACT = 15

	//调用智能合约的交易的类型
	TX_TYPE_CALL_CONTRACT = 16

	//删除智能合约交易的类型
	TX_TYPE_DELETE_CONTRACT = 17

	//合约内部转账交易的类型
	//contract transfer tx type
	TX_TYPE_CONTRACT_TRANSFER = 18

	//合约执行手续费返还交易的类型
	//合约在调用时，会模拟执行一次，用来估算需要花费的手续费金额，实际在调用过程中，给出了更多的手续费（出于一定要执行成功的目的），
	//当合约实际执行时，可能会出现手续费富余，这部分gas将以“合约手续费返回交易”的方式返回给调用者
	TX_TYPE_CONTRACT_RETURN_GAS = 19

	//合约创建共识节点交易
	TX_TYPE_CONTRACT_CREATE_AGENT = 20

	//合约委托交易
	TX_TYPE_CONTRACT_DEPOSIT = 21

	//合约撤销委托交易
	TX_TYPE_CONTRACT_CANCEL_DEPOSIT = 22

	//合约停止节点交易
	TX_TYPE_CONTRACT_STOP_AGENT = 23

	//跨链验证人变更交易
	TX_TYPE_VERIFIER_CHANGE = 24

	//跨链验证人初始化交易
	TX_TYPE_VERIFIER_INIT = 25
)
