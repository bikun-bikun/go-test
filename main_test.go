package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct{
		name string
		a, b, want int
	}{
		{"Simple", 0, 1, 1},
		{"Minus", -1 , -1, -2},
		{"Both", -3, 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if got := Sum(tt.a, tt.b); got != tt.want{
				t.Errorf("Want = %d, got = %d", tt.want, got)
			}
		})
	}

}

func TestGetFizzBuzz(t *testing.T) {
	tests := []struct{
		name string
		p int
		want string
	}{
		{"test case 1",1,"1"},
		{"test case 2",2,"2"},
		{"test case 3",3,"Fizz"},
		{"test case 4",4,"4"},
		{"test case 5",5,"Buzz"},
		{"test case 6",6,"Fizz"},
		{"test case 7",7,"7"},
		{"test case 8",8,"8"},
		{"test case 9",9,"Fizz"},
		{"test case 10",10,"Buzz"},
		{"test case 11",11,"11"},
		{"test case 12",12,"Fizz"},
		{"test case 13",13,"13"},
		{"test case 14",14,"14"},
		{"test case 15",15,"FizzBuzz"},
		{"test case 16",16,"16"},
		{"test case 17",17,"17"},
		{"test case 18",18,"Fizz"},
		{"test case 19",19,"19"},
		{"test case 20",20,"Buzz"},
		{"test case 21",21,"Fizz"},
		{"test case 22",22,"22"},
		{"test case 23",23,"23"},
		{"test case 24",24,"Fizz"},
		{"test case 25",25,"Buzz"},
		{"test case 26",26,"26"},
		{"test case 27",27,"Fizz"},
		{"test case 28",28,"28"},
		{"test case 29",29,"29"},
		{"test case 30",30,"FizzBuzz"},
		{"test case 31",31,"31"},
		{"test case 32",32,"32"},
		{"test case 33",33,"Fizz"},
		{"test case 34",34,"34"},
		{"test case 35",35,"Buzz"},
		{"test case 36",36,"Fizz"},
		{"test case 37",37,"37"},
		{"test case 38",38,"38"},
		{"test case 39",39,"Fizz"},
		{"test case 40",40,"Buzz"},
		{"test case 41",41,"41"},
		{"test case 42",42,"Fizz"},
		{"test case 43",43,"43"},
		{"test case 44",44,"44"},
		{"test case 45",45,"FizzBuzz"},
		{"test case 46",46,"46"},
		{"test case 47",47,"47"},
		{"test case 48",48,"Fizz"},
		{"test case 49",49,"49"},
		{"test case 50",50,"Buzz"},
		{"test case 51",51,"Fizz"},
		{"test case 52",52,"52"},
		{"test case 53",53,"53"},
		{"test case 54",54,"Fizz"},
		{"test case 55",55,"Buzz"},
		{"test case 56",56,"56"},
		{"test case 57",57,"Fizz"},
		{"test case 58",58,"58"},
		{"test case 59",59,"59"},
		{"test case 60",60,"FizzBuzz"},
		{"test case 61",61,"61"},
		{"test case 62",62,"62"},
		{"test case 63",63,"Fizz"},
		{"test case 64",64,"64"},
		{"test case 65",65,"Buzz"},
		{"test case 66",66,"Fizz"},
		{"test case 67",67,"67"},
		{"test case 68",68,"68"},
		{"test case 69",69,"Fizz"},
		{"test case 70",70,"Buzz"},
		{"test case 71",71,"71"},
		{"test case 72",72,"72"}, //Todo: Fizz
		{"test case 73",73,"73"},
		{"test case 74",74,"74"},
		{"test case 75",75,"FizzBuzz"},
		{"test case 76",76,"76"},
		{"test case 77",77,"77"},
		{"test case 78",78,"Fizz"},
		{"test case 79",79,"79"},
		{"test case 80",80,"Buzz"},
		{"test case 81",81,"Fizz"},
		{"test case 82",82,"82"},
		{"test case 83",83,"83"},
		{"test case 84",84,"Fizz"},
		{"test case 85",85,"Buzz"},
		{"test case 86",86,"Fizz"}, //Todo: 86
		{"test case 87",87,"Fizz"},
		{"test case 88",88,"88"},
		{"test case 89",89,"89"},
		{"test case 90",90,"FizzBuzz"},
		{"test case 91",91,"91"},
		{"test case 92",92,"92"},
		{"test case 93",93,"Fizz"},
		{"test case 94",94,"94"},
		{"test case 95",95,"Buzz"},
		{"test case 96",96,"Fizz"},
		{"test case 97",97,"97"},
		{"test case 98",98,"98"},
		{"test case 99",99,"Fizz"},
		{"test case 100",100,"Buzz"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T){
			t.Parallel()
			if got := getFizzBuzz(tt.p); got != tt.want{
				t.Errorf("Want = %v, got = %v", tt.want, got)
			}
		})
	}

/*
	for _, tt := range tests {
		if got := getFizzBuzz(tt.p); got != tt.want{
			t.Errorf("Want = %v, got = %v", tt.want, got)
		}
	}
*/

}