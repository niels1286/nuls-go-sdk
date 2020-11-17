// @Title
// @Description
// @Author  Niels  2020/11/17
package txdata

import (
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"math/big"
)

//调用合约的交易扩展字段的具体协议
type CallContract struct {
	Sender          []byte
	ContractAddress []byte
	Value           *big.Int
	GasLimit        uint64
	Price           uint64
	MethodName      string
	MethodDesc      string
	ArgsCount       uint8
	Args            [][]string
}

//反序列化
func (a *CallContract) Parse(reader *seria.ByteBufReader) error {
	var err error
	a.Sender, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	a.ContractAddress, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	a.Value, err = reader.ReadBigInt()
	if err != nil {
		return err
	}
	a.GasLimit, err = reader.ReadUint64()
	if err != nil {
		return err
	}
	a.Price, err = reader.ReadUint64()
	if err != nil {
		return err
	}
	a.MethodName, err = reader.ReadStringWithLen()
	if err != nil {
		return err
	}
	a.MethodDesc, err = reader.ReadStringWithLen()
	if err != nil {
		return err
	}
	a.ArgsCount, err = reader.ReadByte()
	if err != nil {
		return err
	}
	a.Args = [][]string{}
	for i := uint8(0); i < a.ArgsCount; i++ {
		count, _ := reader.ReadByte()
		if 0 == count {
			a.Args = append(a.Args, []string{})
		} else {
			val := []string{}
			for x := uint8(0); x < count; x++ {
				str, _ := reader.ReadStringWithLen()
				val = append(val, str)
				a.Args = append(a.Args, val)
			}
		}
	}
	return nil
}

//序列化方法
func (a *CallContract) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBytes(a.Sender)
	writer.WriteBytes(a.ContractAddress)
	writer.WriteBigint(a.Value)
	writer.WriteUInt64(a.GasLimit)
	writer.WriteUInt64(a.Price)
	writer.WriteString(a.MethodName)
	writer.WriteString(a.MethodDesc)
	writer.WriteByte(a.ArgsCount)
	if nil != a.Args && len(a.Args) > 0 {
		for _, arg := range a.Args {
			if nil == arg || len(arg) == 0 {
				writer.WriteByte(0)
			} else {
				writer.WriteByte(uint8(len(arg)))
				for _, val := range arg {
					writer.WriteString(val)
				}
			}
		}
	}

	return writer.Serialize(), nil
}
