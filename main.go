package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	// "database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string `db:"name"`
	Password string `db:"password"`
	Mission  int    `db:"mission_lvl"`
	Coin     int    `db:"coin"`
	Extra    int    `db:"extra_words"`
	Chance   int    `db:"chance"`
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func MainMenu() {
	fmt.Println("Welcome to this game:")
	fmt.Println("******* Main Menu *******")
	fmt.Println("1.Sign up")
	fmt.Println("2.Login")
	fmt.Println("3.Rank")
	fmt.Println("4.challenge")
	fmt.Println("5.Exit")
	fmt.Println("Please enter your choice:")
}
func InternalMenu(login_user User) {
	fmt.Println("******* Internal Menu *******")
	fmt.Printf("username : %s\n", login_user.Name)
	fmt.Printf("mission level : %d\n", login_user.Mission)
	fmt.Printf("Coin : %d\n", login_user.Coin)
	fmt.Printf("Chance : %d\n", login_user.Chance)
	fmt.Printf("\n")
	fmt.Println("1.Continue game")
	fmt.Println("2.Choose level")
	fmt.Println("3.Wheel of Luck")
	fmt.Println("4.Edit profile")
	fmt.Println("5.Logout")
	fmt.Printf("\n\n")
	fmt.Println("Please enter your choice:")
}
func Get_UserInfo(player *User) {
	db, err := sqlx.Connect("postgres", "user=a dbname=amirza sslmode=disable password=123 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	var insert string = fmt.Sprintf("SELECT name,password, mission_lvl,coin,extra_words,chance FROM user_info where name='%s';", player.Name)
	rows, _ := db.Queryx(insert)
	// 		errr := db.Select(&player, insert)
	// 		fmt.Println(errr)
	// //////////////////##############need error handling fir empty query ############################
	// 		if errr != nil{
	// 			os.Exit(2)
	// 		}
	for rows.Next() {
		err := rows.StructScan(&player)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%#v\n", player)
	}

}
func editProfile(player *User) {
	db, err := sqlx.Connect("postgres", "user=a dbname=amirza sslmode=disable password=123 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	var tmp_player User
	for i := 0; i < 3; i++ {
		fmt.Println("Please enter your current password: ")
		var old_password string
		fmt.Scan(&old_password)
		get_pass := fmt.Sprintf("select password from user_info where name='%s';", player.Name)
		rows, _ := db.Queryx(get_pass)
		for rows.Next() {
			err := rows.StructScan(&tmp_player)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("%#v\n", tmp_player)
		}
		if old_password == tmp_player.Password {

			var new_password string
			var new_password_check string
			fmt.Println("enter your new password")
			fmt.Scan(&new_password)
			fmt.Println("enter your new password again")
			fmt.Scan(&new_password_check)
			if new_password == new_password_check {
				var insert string = fmt.Sprintf("UPDATE user_info SET password ='%s' where name='%s';", new_password, player.Name)
				db.Exec(insert)
				log.Printf("password successfully changed")

			} else {
				fmt.Println("password is not match")
			}

		}
	}
	// db.QueryRow(get_pass).Scan(&output.Country)

}

type Word struct {
	Chars string
	Main  int
	Sub   int
	Words string
}

func main() {
	player := User{}
	MainMenu()
	var choose int

	fmt.Scanln(&choose)

	CallClear()

	switch choose {
	case 1:
		//##########################################  make user  #########################################################

		fmt.Println("1")
	case 2:
		fmt.Println("insert the name of player")
		fmt.Scanln(&player.Name)
		Get_UserInfo(&player)
		// CallClear()
		InternalMenu(player)

		choose = 0
		fmt.Scanln(&choose)
		CallClear()
		switch choose {
		case 1:
			//############### continue ###################

		case 2:
			//############### choose lvl ###################

		case 3:
			CallClear()
			//############### wheel of luck ###################

		case 4:
			CallClear()
			//############### editProfile ###################

			editProfile(&player)
			//     goto sub;
			//         break;
		case 5:
			CallClear()
		}
	case 3:
		fmt.Println("3")
		//############### showRank ###################

		// showRank();

	case 4:
		fmt.Println("4")
		//############### challenge ###################

	case 5:
		os.Exit(1)
	}
}
