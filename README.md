# table
A command line table output tool

## example
```
  //
  type Student struct {
		Name   string   `head:"NAME"`
		Age    int      `head:"AGE"`
		Gender string   `head:"GENDER"`
		Hobby  []string `head:"HOBBY"`
	}

  var stuSlice = []Student{
		{
			Name:   "Zhang san",
			Age:    23,
			Gender: "Man",
			Hobby:  []string{"Basketball", "Ping pang ball"},
		},
		{
			Name:   "Li si",
			Age:    22,
			Gender: "WoMan",
			Hobby:  []string{"Swimming"},
		},
		{
			Name:   "Wang wu",
			Age:    21,
			Gender: "Man",
		},
	}
  
``` 
#### PrintTableSimple
```
 table, err := CreateTableCell(stuSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		PrintTableSimple(table)
	}
```
#### Output result
```
NAME        AGE   GENDER   HOBBY         
Zhang san   23    Man      Basketball    
                           Ping pang ball
Li si       22    WoMan    Swimming      
Wang wu     21    Man   
```

#### PrintTableBox
```
 table, err := CreateTableCell(stuSlice)
	if err != nil {
		fmt.Println(err)
	} else {
		PrintTableBox(table)
	}
```
#### Output result
``` 
┌───────────┬─────┬────────┬─────────────────┐
│ NAME      │ AGE │ GENDER │ HOBBY           │
├───────────┼─────┼────────┼─────────────────┤
│ Zhang san │ 23  │ Man    │ Basketball      │
│           │     │        │ Ping pang ball  │
├───────────┼─────┼────────┼─────────────────┤
│ Li si     │ 22  │ WoMan  │ Swimming        │
├───────────┼─────┼────────┼─────────────────┤
│ Wang wu   │ 21  │ Man    │                 │
└───────────┴─────┴────────┴─────────────────┘
``` 


  
