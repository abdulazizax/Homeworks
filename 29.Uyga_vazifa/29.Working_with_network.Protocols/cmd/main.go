package main

import (
	fk "db/29.Working_with_network.Protocols/functions"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	db := fk.DBConnect()

	man := fk.Manager{
		DB: db,
	}

	man.DropUsersTableIfExists()
	// 2:
	man.CreatePeopleTable()
	// 3:
	man.InserIntoPeopleTable()

	fmt.Println("Users who firstname is 'Adrian':")
	persons := man.GetAdrian()
	for _, v := range persons {
		pp.Println(v)
	}

	// 4:
	fmt.Printf("Get query plan for 'Adrian' before creating index: \n")
	new := man.GetQueryPlanForAdrian()
	fmt.Printf("%v\n\n", new)

	// 5:
	man.CreateIndexOnFirstName()

	// 6:
	fmt.Printf("Get query plan for 'Adrian' after creating index: \n")
	new = man.GetQueryPlanForAdrian()
	fmt.Printf("%v\n\n", new)

	// 7:
	man.DropIndexOnFirstName()

	fmt.Println("Users who firstname is 'Adrian' and lastname is 'Gross':")
	persons = man.GetAdrianGross()
	for _, v := range persons {
		pp.Println(v)
	}

	fmt.Printf("Get query plan for 'Adrian' and 'Gross' before creating index: \n")
	new = man.GetQueryPlanForAdrianGross()
	fmt.Printf("%v\n\n", new)

	// 8:
	man.CreateIndexOnFirstNameAndLastName()

	// 9:
	fmt.Printf("Get query plan for 'Adrian' and 'Gross' after creating index: \n")
	new = man.GetQueryPlanForAdrianGross()
	fmt.Printf("%v\n\n", new)

	fmt.Println("After we created the index, the search time was reduced.")
}
