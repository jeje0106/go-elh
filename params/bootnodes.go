// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

var MainnetBootnodes = []string{
"enode://d0ceac9b87639b17aed3d7736e0b4ac55d2b07f6ff5c2e0b24b7a28600b6df5fafee7798b1af4612ca4233185d7cd6c13769f8501606cf0a5190e0987b76b4cc@107.174.121.125:39390",
"enode://a2423412e020792767cee53d4b2dc3f5ffcd0cb4aa6af676485f3c071ad6984ec6f8da4e74f6a98301b64c52dcb698ec2243a28ecf698fabf1cf1b6a33658840@107.174.115.176:39390",
"enode://ba18c0581ff96d58c053e350efd5ef63daf346cf38f0112d073a65102b2a4ad81e9d1808d0f5c5092756ab35a15ad52163df50fddf2fb9ef5c50709e89486d2d@107.174.156.71:39390",
"enode://0f160286381f1b9dd543b8a7b69890689b01dbc01513b1ae163e0ac41ac6dde385161f393cf2b9642d464768d9a2bcda9430b5eb3f632072440c3b562db0a6f8@141.148.224.171:39390",
"enode://55e7f3666f6c697f47cb7be968d954edc26565ba59b73f4d79a918c3c218f9b23aff728f6b956e0f43c9f632a7568488248fd17138cdbd4b75b57b1342fe2b53@193.108.113.23:39390",
"enode://7fd9cc4b1455ba0c7faaa2ee603421b52e0abc78c814b897ccea0fba5aaa41da47c32fcd1f406a577130c29ce31f000689ffd809400863ff5412e427495d6a42@107.174.115.73:39390",
"enode://ab1dd1cc6129d72e28df63df4cc31d1805ee2428d09fe6dfeb5c28ef33875388be04427cd10d745bd27b2de92aef23617f7aed3faab9355877536d4ab10dc1b3@107.174.115.178:39390",
"enode://809a4913bd23dd5e0f8d1ba2c991654666cfe96370c63c16312cacf2edaa3fd49b667d0107045f48c8e507858659486e35b29d42df1e62a63271defe21603a69@89.33.193.104:39390",
"enode://a30b4753bb20bb9045520f8b0911267ceaa7e32ad63411f09d241e89c4542edc055270fe8d7d66f7d245352587e598687e3934f457ad79ee1ac6fcefe2c761d7@107.174.115.67:39390",
"enode://61b87a31fc6e3f6860bbe9342f3b663898ff55e2f3a4b8a142e01a7a54f579ccc53ecd6d4692940f40fd8b8234f5f1048d102c9214e20898fbcbfaf888e7226f@60.247.148.7:39390",
"enode://b5b058e0a89479dc9d59363eca3803217f237d8102aaaee07b46731dec4ddd3d55dce50ed3cd0899b3961444c8425a1d6ff5977b2db9f66110bcf013c0788c1a@107.174.115.63:39390",
"enode://3d50111783a5baec12fb161e481e4ee02e2e286c28ae217bec436ce718df6f1540ef3fbeb2991eb95b1fdda7a5f2b68fa50a8d678087271d1600bd43dfd56bea@85.14.247.139:39390",
"enode://bcd38043a206fcf026ce096a003c5cb0405f4b7fedc4fe63f51f45e3958e933e88c868f24cdf6dad9ec2933ce607056836947bba12834262748ff966d96388dc@173.82.248.81:39390",
"enode://b8e7af80803f7f0d890f6c8f22870139b0c2dd542bb381340524efbbfcf4ddb425cef1bb0e149ab6cea3156c76b55844c4d4d2ae0341dbd35de0325fbf0119e0@109.108.71.107:33034",
"enode://7c258e7e97ec035534275d1bd05b8793f188b4cc5b7cfd843d01146bd50ef8c36e7bf9ad0b942bbb3bd56add9ebb219b70bbdd143956b101e607c25efb79c3bf@198.211.9.237:39390",
"enode://0c31bf5f0e786ab00075c44ba30c132efa4637523c0fe0ce286026b3b8ce6b4328f6e64cb55720a6159364d23c7db0aae5734ba4e09d57a252cc90b7b697b47f@161.97.155.137:32323",
"enode://d8704c176eb1abee3356add5b2206cd29af30353146f285b3b67688d803e3072f38773331668375659b193336123099c8c13237d1be8fb861d170100fd97c89a@62.141.39.38:30403",
"enode://b5693ff162ca2dedf021b7a14e987abc5b77c162e581c069565ed43b6fdf95b38e8a2982bb7041816bcb14adcf2c828cb30e00124fe72172b930534f3ec8cdb8@142.132.227.211:30303",
"enode://04d8da86a9757e52299ae7fec7f8357da47a2a9012740fdd257b0f7b8cd9fb4665466f0d936238d5a2c45c947147e7fb7009ba5da4a8c963595bca046009943c@5.161.78.180:30303",
}
var RopstenBootnodes = []string{
}
var SepoliaBootnodes = []string{
}
var RinkebyBootnodes = []string{
}
var GoerliBootnodes = []string{
}
var KilnBootnodes = []string{
}
var V5Bootnodes = []string{
"enr:-KO4QM6oq6g3YstgN9jZtEv6Qy9XihDfixRw7t4slnsSuOHYRs53D3KGcTe0uEfkPjqD73n7WrGwjFriWuu_fcx_vUGGAYf6dFh3g2V0aMfGhMHpFjOAgmlkgnY0gmlwhCW7TCWJc2VjcDI1NmsxoQNK26z_KuxybXxBnj4X4S6VIOUjZDjvcMcpORq-66e1kYRzbmFwwIN0Y3CCmd6DdWRwgpne",
"enr:-KO4QPxTB0KISMOZ9pxuJQksd-wuz0br7GBPmGwAx5e3U76fb3uWIUVA4WLJgoN59xbMMXFWSJlQTKP71fIJupAwkeCGAYf6dt2Pg2V0aMfGhMHpFjOAgmlkgnY0gmlwhLyl1UWJc2VjcDI1NmsxoQJZm7S5-_ViFpXPKeRHlZZYegFruzMpzqJeqHr_BEfvIIRzbmFwwIN0Y3CCmd6DdWRwgpne",
"enr:-KO4QC9Nn1vgKhUw9daUV1AkQom5Ey2CnEkPxQvMAyL4t14-fyxYRIi4DtNpDWsJlbckkFocmEOxLl2Fom_JbGfj5MKGAYf6dqQWg2V0aMfGhMHpFjOAgmlkgnY0gmlwhCW7g2mJc2VjcDI1NmsxoQJufI7hIwxrEJQD92XI4AZsZ7wjPPWw-p4xCSLI1IpPrYRzbmFwwIN0Y3CCmd6DdWRwgpne",
"enr:-J24QHYZUEzTXdfHdxRwrd3UWqG7x0DvZ8gTHbbyEQklQIuDIzKQjGxOxGE5e3SKaYoh3JqmkI2Pdf6_1zWyRY7uAA6GAYgEDb3jg2V0aMfGhMHpFjOAgmlkgnY0gmlwhDPDKMmJc2VjcDI1NmsxoQLfSjOW_UdgMfGOtXERRN_khuac5Ai-w_OqW1aJSzlTh4N0Y3CCmd6DdWRwgpne",
}
var AgaMainnetBootnodes = []string{
"enode://4adbacff2aec726d7c419e3e17e12e9520e5236438ef70c729391abeeba7b591c9edaf6e380cbbf0857a44b7b9a328402b7719ceb2c1ec0aaf569fd39ec9026d@37.187.76.37:39390",
"enode://599bb4b9fbf5621695cf29e4479596587a016bbb3329cea25ea87aff0447ef20a3ca20a8c60c2325d8a67a08ccfa5b1d50f35cb46d955093991fd08e57677f6e@188.165.213.69:39390",
"enode://6e7c8ee1230c6b109403f765c8e0066c67bc233cf5b0fa9e310922c8d48a4fad59275af292272a6429fedface52eba99db2990cdc7f8e2e0009e51e2ff07694e@37.187.131.105:39390",
"enode://df4a3396fd476031f18eb5711144dfe486e69ce408bec3f3aa5b56894b395387730f6323effcdddb8ab34a9360238271ad004da4c954d167c81af8831477e88a@51.195.40.201:39390",
}
var AgaTestnetBootnodes = []string{
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
