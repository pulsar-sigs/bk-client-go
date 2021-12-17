package entry

import "log"

func addEntry(ledgerId, entryId int64, masterKey, body []byte) {

	flag := proto.AddRequest_RECOVERY_ADD

	addRequest := &proto.AddRequest{
		LedgerId:  &ledgerId,
		EntryId:   &entryId,
		MasterKey: masterKey,
		Body:      body,
		Flag:      &flag,
	}
	log.Println(addRequest)
}
