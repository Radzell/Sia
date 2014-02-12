package disk

import (
	"os"
	//"bytes"
)

type SwarmStorage struct{
	SwarmId string
	amountused uint64
}

func  CreateSwarmSystem(swarmid string) (*SwarmStorage,error){
	var i SwarmStorage
	i.SwarmId = swarmid
	i.amountused=0
	err:=os.Mkdir(swarmid,os.ModeDir|os.ModePerm)
	if err!=nil&&!os.IsExist(err){
		return nil,err
	}
	return &i,nil
}

func (r SwarmStorage) CreateFile(filehash string, length uint64) error{
	file,err:=os.Create(r.SwarmId+string(os.PathSeparator)+filehash)
	defer file.Close()
	if os.IsExist(err){

	}
	if err==nil{
		b:=[]byte{0}
		for i :=uint64(0);i<length;i++{
			_,_=file.Write(b)
		}
		return nil
	}
	return err
}

func (r SwarmStorage) DeleteFile(filehash string) error{
	err:=os.Remove(r.SwarmId+string(os.PathSeparator)+filehash)
	return err
}

func (r SwarmStorage) WriteFile(filehash string, start uint64, data []byte) error{
	path:=r.SwarmId+string(os.PathSeparator)+filehash
	file,err:=os.Open(path)
	if err!=nil{
		return err
	}
	file.WriteAt(data,int64(start))
	file.Close()
	return nil

}

















