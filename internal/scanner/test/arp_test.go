package scanner

import "testing"


func TestARPTable(t *testing.T) {


	table := GetARPTable()


	for ip, mac := range table {

		t.Log(
			ip,
			mac,
		)
	}


}