package tutorials

import "fmt"

// declare a User struct
type User struct {
    id int
    username string
    userAge int
    activeDays int
    email string
    likes int
    followers int
}


// declare a custom type for an anonymous function that accepts User as parameter and returns bool
// this By essentially refers to as a function placeholder or type; which we'll use for methods (functions as type for certain methods)
type By func(user User) bool

// We declare a Filter method that is applicable to By types; i.e. functions that accept a User and return bool;
func (by By) Filter(usrSlice []User) []User {
    
    // creates an empty sice/array of type User; which will store a list of filtered users
    filtered := make([]User, 0)

    // for each element in the user slice
    for _, elem := range usrSlice {
        // if by; which is of type By; essentially an anonymous function that takes in User and return bool; returns true
        if by(elem) {
            // append this element to the empty slice of users created earlier
            filtered = append(filtered, elem)
        }
    }

    // return the updated slice of users
    return filtered

}

func MethodsWithFnAsReceiversDemo() {
    // initialize a list of Users
    users := []User {
        User {1, "thesajan", 32, 11, "thisisnotmyemail@getalife.com", 33, 12 },
        User {2, "dummyboy", 9, 34, "dummyemail@notyourconcern.com", 829, 91 },
        User {3, "killerbee", 31, 88, "killerbee@naruto.com", 4354, 98 },
        User {4, "volibear", 16, 98, "volibear-admin@lol.com", 324, 176 },
        User {5, "senna", 25, 45, "senna-admin@lol.com", 345, 556 },
    }
    
    // print all users for testing
    fmt.Println("All users are: ", users)


    // The anonymous functions declared earlier using By type can now be signified with certain conditions
    // Users with many followers
    isUserWithManyFollowers := func(user User) bool {
        return user.followers > 100
    }

    // Users who are active;
    isUserActive := func(user User) bool {
        return user.activeDays > 30
    }

    // Mature users;
    isUserMature := func(user User) bool {
        return user.userAge >= 18
    }
    
    // Now we finally use the Filter method that is applicable to such anonymous function (signified by type By)
    // and get the desired list of Users under different conditions
    userWithManyFollowersList := By(isUserWithManyFollowers).Filter(users)
    fmt.Println("Many following users are: ", userWithManyFollowersList)

    activeUsersList := By(isUserActive).Filter(users)
    fmt.Println("Active users list: ", activeUsersList)

    matureUsersLists := By(isUserMature).Filter(users)
    fmt.Println("Mature users list: ", matureUsersLists)
}

