# Authcord golang module
The Authcord library is a golang module that allows you to easily communicate with the Authcord API.

(many updates to come)

# Documentation 
see the <a href="https://docs.authcord.xyz">Authcord docs</a>

# Installation
simply run this command:
```
go get github.com/AuthCord/authcord-golang
```

# Usage

Bellow will teach you how to use all of the user level functions

Authcord requires an APIKEY so you can set that by using the NewAuthcordClient function and passing in your APIKEY in example of this:
```go
client := authcord.NewAuthcordClient("YOUR_USER_LEVEL_KEY")
```
```After setting your user level key you can use the check hwid function:```
```go
package main

import (
	"authcord"
	"fmt"
)

func main() {
	// get users HWID
	hwid := authcord.GetHWID()
	// Initialize a new Authcord client with your API key
	client := authcord.NewAuthcordClient("YOUR_USER_LEVEL_KEY")

	// Check the HWID of a user
	response, err := client.CheckHWID(hwid)
	if err != nil {
		fmt.Println("Error checking HWID:", err)
		return
	}
	fmt.Println("Response:", response)

}
```

Bellow will teach you how to use all of the admin level functions

To use admin level functions you must set your admin level key you can do so by using the NewAuthcordAdmin function and passing in your Adminkey Example:

```go
admin := authcord.NewAuthcordAdmin("YOUR_ADMIN_LEVEL_KEY")
```
After setting your user level key you can use the following functions:

```Add HWID```:
```go 
package main

import (
	"authcord"
	"fmt"
)

func main() {
	// get users HWID
	hwid := authcord.GetHWID()
	
	// Initialize a new Authcord admin with your API key
	admin := authcord.NewAuthcordAdmin("YOUR_ADMIN_LEVEL_KEY")
	
	// add HWID
	res, err := admin.AddHWID(hwid)
	if err != nil {
		fmt.Println("Error adding HWID:", err)
		return
	}
	fmt.Println("Response:", res)

}
```

```Delete HWID```:
```go
package main

import (
	"authcord"
	"fmt"
)

func main() {
	// get users HWID
	hwid := authcord.GetHWID()

	// delete HWID
	admin := authcord.NewAuthcordAdmin("YOUR_ADMIN_LEVEL_KEY")
	res, err := admin.DeleteHWID(hwid)
	if err != nil {
		fmt.Println("Error deleting HWID:", err)
		return
	}
	fmt.Println("Response:", res)

}
```

# Authentication system examples:

```go
package main

import (
	"authcord"
	"fmt"
)

func main() {
	// get users HWID
	hwid := authcord.GetHWID()
	// Initialize a new Authcord client with your API key
	client := authcord.NewAuthcordClient("YOUR_USER_LEVEL_KEY")

	// Check the HWID of a user
	response, err := client.CheckHWID(hwid)
	if err != nil {
		fmt.Println("Error checking HWID:", err)
		return
	}

	if response == "Valid" {
		fmt.Println("Logged in...")
		// your code
	} else {
		fmt.Println("Failed to login...")
		// your code
	}

}
```

 



