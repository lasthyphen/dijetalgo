// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/lasthyphen/dijetalgo/utils/units"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1,
		"allocations": [
			{
				"ethAddr": "0x1B00f59fff05F6591c13e32740377eAE72661061",
				"djtxAddr": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"initialAmount": 22500000000000000
			},
			{
				"ethAddr": "0xd86b355443158939c2f1b2A00961F8453b33E74E",
				"djtxAddr": "X-dijets1v8wat5z4cxh7wh873d7n6d9m6mpnynr8sgl059",
				"initialAmount": 22500000000000000
			},
			{
				"ethAddr": "0x80231567cD6E270c8360B80b97034Ec26dad83b8",
				"djtxAddr": "X-dijets16yd4ams4xdp9c6ht9zfnp90225ukwmqnj964sw",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 175000000000000,
						"locktime": 1674257383
					},
					{
						"amount": 175000000000000,
						"locktime": 1676935783
					},
					{
						"amount": 175000000000000,
						"locktime": 1679354983
					},
					{
						"amount": 175000000000000,
						"locktime": 1682029783
					},
					{
						"amount": 175000000000000,
						"locktime": 1684621783
					},
					{
						"amount": 175000000000000,
						"locktime": 1687300183
					},
					{
						"amount": 175000000000000,
						"locktime": 1689892183
					},
					{
						"amount": 175000000000000,
						"locktime": 1692570583
					},
					{
						"amount": 175000000000000,
						"locktime": 1695248983
					},
					{
						"amount": 175000000000000,
						"locktime": 1697840983
					},
					{
						"amount": 175000000000000,
						"locktime": 1700522983
					},
					{
						"amount": 175000000000000,
						"locktime": 1702682983
					}
				]
			},
			{
				"ethAddr": "0x9010F91C986dA7018c0918E1562abD039DB63Bdb",
				"djtxAddr": "X-dijets174rcm4tva8f7t4m7u2nxe7a236nv9pcwhdmt5v",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 1500000000000000,
						"locktime": 1674257383
					},
					{
						"amount": 1500000000000000,
						"locktime": 1676935783
					},
					{
						"amount": 1500000000000000,
						"locktime": 1679354983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1682029783
					},
					{
						"amount": 1500000000000000,
						"locktime": 1684621783
					},
					{
						"amount": 1500000000000000,
						"locktime": 1687300183
					},
					{
						"amount": 1500000000000000,
						"locktime": 1689892183
					},
					{
						"amount": 1500000000000000,
						"locktime": 1692570583
					},
					{
						"amount": 1500000000000000,
						"locktime": 1695248983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1697840983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1700522983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1702682983
					}
				]
			},
			{
				"ethAddr": "0x30bD7DcdE7848A360B7226EBa7A6427C488d5832",
				"djtxAddr": "X-dijets19qeq9fg568ztpace8mdecq4jd9rx5dj64fc4sf",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 1500000000000000,
						"locktime": 1674257383
					},
					{
						"amount": 1500000000000000,
						"locktime": 1676935783
					},
					{
						"amount": 1500000000000000,
						"locktime": 1679354983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1682029783
					},
					{
						"amount": 1500000000000000,
						"locktime": 1684621783
					},
					{
						"amount": 1500000000000000,
						"locktime": 1687300183
					},
					{
						"amount": 1500000000000000,
						"locktime": 1689892183
					},
					{
						"amount": 1500000000000000,
						"locktime": 1692570583
					},
					{
						"amount": 1500000000000000,
						"locktime": 1695248983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1697840983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1700522983
					},
					{
						"amount": 1500000000000000,
						"locktime": 1702682983
					}
				]
			},
			{
				"ethAddr": "0x6E04B23a9Fa424289A10eef126b04Dc13B38069c",
				"djtxAddr": "X-dijets1eg0gnn7q2uvk66kht7yqn8xtugz7uvutfk0txf",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 20833333300000,
						"locktime": 1674257383
					},
					{
						"amount": 20833333300000,
						"locktime": 1676935783
					},
					{
						"amount": 20833333300000,
						"locktime": 1679354983
					},
					{
						"amount": 20833333300000,
						"locktime": 1682029783
					},
					{
						"amount": 20833333300000,
						"locktime": 1684621783
					},
					{
						"amount": 20833333300000,
						"locktime": 1687300183
					},
					{
						"amount": 20833333300000,
						"locktime": 1689892183
					},
					{
						"amount": 20833333300000,
						"locktime": 1692570583
					},
					{
						"amount": 20833333300000,
						"locktime": 1695248983
					},
					{
						"amount": 20833333300000,
						"locktime": 1697840983
					},
					{
						"amount": 20833333300000,
						"locktime": 1700522983
					},
					{
						"amount": 20833333300000,
						"locktime": 1702682983
					}
				]
			},
			{
				"ethAddr": "0x3507A1131aba9D07714b76Eb8a245F434198690B",
				"djtxAddr": "X-dijets1ggwu285vknzzjs2rqyvzyrkjd0guc5lfgam9qp",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 83333333333333,
						"locktime": 1674257383
					},
					{
						"amount": 83333333333333,
						"locktime": 1676935783
					},
					{
						"amount": 83333333333333,
						"locktime": 1679354983
					},
					{
						"amount": 83333333333333,
						"locktime": 1682029783
					},
					{
						"amount": 83333333333333,
						"locktime": 1684621783
					},
					{
						"amount": 83333333333333,
						"locktime": 1687300183
					},
					{
						"amount": 83333333333333,
						"locktime": 1689892183
					},
					{
						"amount": 83333333333333,
						"locktime": 1692570583
					},
					{
						"amount": 83333333333333,
						"locktime": 1695248983
					},
					{
						"amount": 83333333333333,
						"locktime": 1697840983
					},
					{
						"amount": 83333333333333,
						"locktime": 1700522983
					},
					{
						"amount": 83333333333333,
						"locktime": 1702682983
					}
				]
			},
			{
				"ethAddr": "0x2070D7F1b4052c14f31a0C65AD8E011A6cEc33A7",
				"djtxAddr": "X-dijets1rrwvthn2cxpmhfx376hsrnx04me0s24ajsqgsq",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0x83101459D424C8f538F37F293D9151702859fDD6",
				"djtxAddr": "X-dijets1gq7fqklm93veh3gsxahu6a4uuu04d4shuqxnc8",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xCB4f9d421e2164D85a77d1ae0eff5d3a1e56b454",
				"djtxAddr": "X-dijets12p06h3dktfnt8jlrkvr93u7h4vjhx8wx39esng",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xa42451b72EC0210528e77D6febAD848253896250",
				"djtxAddr": "X-dijets1pag5qx6v6udm0myfpa6mg2jd8kf033wp5jzfzg",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xa85c8831526c36bCa2821e3491e319441D94f928",
				"djtxAddr": "X-dijets1fmq580x966pmzkt9h4uhe5djjdn42vlx3cwl34",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xb09085d85C416d353d7a469263e52688ab2dfE4A",
				"djtxAddr": "X-dijets1chlc4au6w90etdmvl6uhdcle025jcrvxzthc5e",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xa28a4B8c614aCF3F3448daD91cA80A2B3034A8D9",
				"djtxAddr": "X-dijets1mvqs89w6e3qn758cdmutv759tsyac97hf6m7vx",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xb4124cEB3451635DAcedd11767f004d8a28c6eE7",
				"djtxAddr": "X-dijets163yr65mhglrvu5wngu4qkw6lj4pqh2zggut3fk",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0x0FE71278D4F6FabcBfF80A3C58947724e40a77fd",
				"djtxAddr": "X-dijets1umw42smlwq6em4krng53zch0gpzjt8vlkqu3e9",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xea0df620254F5D577dDBB7d5F63a68d1d961C979",
				"djtxAddr": "X-dijets169chkgna7crmt3y7ujcqp6xg2qx3fzhg4nj7xn",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			},
			{
				"ethAddr": "0xb09085d85C416d353d7a469263e52688ab2dfE4A",
				"djtxAddr": "X-dijets1chlc4au6w90etdmvl6uhdcle025jcrvxzthc5e",
				"initialAmount": 0,
				"unlockSchedule": [
					{
						"amount": 800000000000000
					}
				]
			}
		],
		"startTime": 1671499328,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 54000,
		"initialStakedFunds": [
			"X-dijets16yd4ams4xdp9c6ht9zfnp90225ukwmqnj964sw",
			"X-dijets1v8wat5z4cxh7wh873d7n6d9m6mpnynr8sgl059",
			"X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
			"X-dijets1t7693hu760mdtce2e6jecmhd32rurh8ey0jl8j",
			"X-dijets1m5z9rmsrtftqwq2cdu4ddyax27unqrqwwrhy3p"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg",
				"rewardAddress": "X-dijets16yd4ams4xdp9c6ht9zfnp90225ukwmqnj964sw",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ",
				"rewardAddress": "X-dijets1v8wat5z4cxh7wh873d7n6d9m6mpnynr8sgl059",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu",
				"rewardAddress": "X-dijets1t7693hu760mdtce2e6jecmhd32rurh8ey0jl8j",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5",
				"rewardAddress": "X-dijets1m5z9rmsrtftqwq2cdu4ddyax27unqrqwwrhy3p",
				"delegationFee": 200000
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":98200,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0,\"apricotPhase1BlockTimestamp\":0,\"apricotPhase2BlockTimestamp\":0,\"apricotPhase3BlockTimestamp\":0,\"apricotPhase4BlockTimestamp\":0,\"apricotPhase5BlockTimestamp\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "Always act in good faith with full transparency and accountability"
	}`

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliDjtx,
			CreateAssetTxFee:      10 * units.MilliDjtx,
			CreateSubnetTxFee:     1 * units.Djtx,
			CreateBlockchainTxFee: 1 * units.Djtx,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement:  .8, // 80%
			MinValidatorStake:  2 * units.KiloDjtx,
			MaxValidatorStake:  3 * units.MegaDjtx,
			MinDelegatorStake:  25 * units.Djtx,
			MinDelegationFee:   20000, // 2%
			MinStakeDuration:   2 * 7 * 24 * time.Hour,
			MaxStakeDuration:   365 * 24 * time.Hour,
			StakeMintingPeriod: 365 * 24 * time.Hour,
		},
		EpochConfig: EpochConfig{
			EpochFirstTransition: time.Unix(1607626800, 0),
			EpochDuration:        6 * time.Hour,
		},
	}
)
