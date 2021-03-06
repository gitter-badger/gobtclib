package demos

import (
	"github.com/gobasis/log"
)

/*
Description:
a demo test of GetBlockCount, Get the current block count.
 * Author: architect.bian
 * Date: 2018/09/02 18:24
 */
func GetBlockCountTest() {
	blockCount, err := cli.GetBlockCount()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockCount", "count", blockCount)
}

/*
Description: 
a demo test of GetBestBlockHash, Get the current best block hash.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBestBlockHashTest() {
	hash, err := cli.GetBestBlockHash()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBestBlockHash", "hash", hash)
}

/*
Description:
a demo test of GetDifficulty, Get the current difficulty for mining difficulty target.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetDifficultyTest() {
	hash, err := cli.GetDifficulty()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetDifficulty", "difficulty", hash)
}

/*
Description:
a demo test of GetBlockHash, Get the height block hash of current best chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHashTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
}

/*
Description:
a demo test of GetBlockHeaderVerbose, Get the special blockheader data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHeaderVerboseTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	header, err := cli.GetBlockHeaderVerbose(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockHeaderVerbose", "header", header)
}

/*
Description:
a demo test of GetBlockHeader, Get the special blockheader data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHeaderTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	header, err := cli.GetBlockHeader(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockHeader", "header", header)
}

/*
Description:
a demo test of GetBlockChainInfo, Get the blockchain state info.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockChainInfoTest() {
	info, err := cli.GetBlockChainInfo()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockChainInfo", "info", info)
}

/*
Description:
a demo test of GetBlock, Get the special block data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlock(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlock", "block", block)
}

/*
Description:
a demo test of GetBlockVerbose, Get the special block data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockVerboseTest() {
	hash, err := cli.GetBlockHash(100)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlockVerbose(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockVerbose", "block", block)
}

/*
Description:
a demo test of GetBlockVerboseTx, Get the special block data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockVerboseTxTest() {
	hash, err := cli.GetBlockHash(100)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlockVerboseTx(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockVerboseTx", "block", block)
}

///*
//Description:
//a demo test of RescanBlocks.
// * Author: architect.bian
// * Date: 2018/09/14 13:13
// */
//func RescanBlocksTest() { //TODO unregistered
//	hash, err := cli.GetBlockHash(100)
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("GetBlockHash", "hash", hash)
//	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
//	block, err := cli.RescanBlocks([]client.Hash{*hash})
//	if err != nil {
//		log.Error("", "error", err)
//		panic(err)
//	}
//	log.Info("RescanBlocks", "block", block)
//}

///*
//Description:
//a demo test of InvalidateBlock, Permanently marks a block as invalid, as if
//it violated a consensus rule.
// * Author: architect.bian
// * Date: 2018/09/14 13:13
// */
//func InvalidateBlockTest() {
//	hash, err := cli.GetBlockHash(1)
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("GetBlockHash", "hash", hash)
//	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
//	e := cli.InvalidateBlock(hash)
//	if e != nil {
//		log.Error("", "error", e)
//		panic(e)
//	}
//	log.Info("InvalidateBlock", "result", "success")
//}

/*
Description:
a demo test of VerifyChain, Verifies blockchain database.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func VerifyChainTest() {
	verify, err := cli.VerifyChain()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("VerifyChain", "verify", verify)
}

/*
Description:
a demo test of VerifyChainLevel, Verifies blockchain database.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func VerifyChainLevelTest() {
	verify, err := cli.VerifyChainLevel(3)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("VerifyChain", "verify", verify)
}

/*
Description:
a demo test of VerifyChainBlocks, Verifies blockchain database.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func VerifyChainBlocksTest() {
	verify, err := cli.VerifyChainBlocks(3, 10)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("VerifyChain", "verify", verify)
}

/*
Description:
a demo test of GetRawMempool, Get all tx ids in memory pool as a string array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetRawMempoolTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		log.Info("GetRawMempool", "hashs", h)
	}
}

/*
Description:
a demo test of GetRawMempoolVerbose, Get all tx ids in memory pool as a tx array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetRawMempoolVerboseTest() {
	result, err := cli.GetRawMempoolVerbose()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, r := range result  {
		log.Info("GetRawMempoolVerbose", "hashs", r)
	}
}

/*
Description:
a demo test of GetRawMempool, Get all tx ids in memory pool as a string array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetMempoolEntryTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		log.Info("GetRawMempool", "hashs", h)
		result, _ := cli.GetMempoolEntry(h.String())
		log.Info("GetMempoolEntry", "entry", result)
		break
	}
}

/*
Description:
a demo test of GetTxOut, Get all tx ids in memory pool as a string array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetTxOutTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		log.Info("GetRawMempool", "hashs", h)
		result, _ := cli.GetTxOut(h, 0, true)
		log.Info("GetTxOut", "result", result)
		break
	}
}