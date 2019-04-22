package member

import (
	"log"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/utils"
)

func RetrieveMember() ([]bukakoneksi.Member, error) {
	var members []bukakoneksi.Member

	_, err := bukakoneksi.Dbmap.Select(&members, "select * from members")
	if err != nil {
		log.Println(err)
		return members, err
	}

	var resMembers []bukakoneksi.Member
	for _, member := range members {
		member.PictureURL = utils.GetPic(member.PictureImage)
		resMembers = append(resMembers, member)
	}

	return resMembers, nil
}

func RetrieveMemberByID(ID int) (bukakoneksi.Member, error) {
	var member bukakoneksi.Member

	err := bukakoneksi.Dbmap.SelectOne(&member, "select * from members where id = ?", ID)
	if err != nil {
		log.Println(err)
		return member, err
	}
	member.PictureURL = utils.GetPic(member.PictureImage)

	return member, nil
}

func RetrieveMemberByName(name string) (bukakoneksi.Member, error) {
	var member bukakoneksi.Member

	err := bukakoneksi.Dbmap.SelectOne(&member, "select * from members where name = ?", name)
	if err != nil {
		log.Println(err)
		return member, err
	}
	member.PictureURL = utils.GetPic(member.PictureImage)

	return member, nil
}
