function sunday_first(){
	/*Start at January 1st 1901*/
	var currDate = new Date(1901,0,1);
	/*Keep track of the Sundays*/
	var sundayCount = 0;
	var untilDate = new Date(2000,11,31);
	do{
		/*If it is a Sunday and it's the first of the month*/
		if (currDate.getDay() === 0 && currDate.getDate() === 1)
			sundayCount += 1;
		currDate.setDate(currDate.getDate()+1);
	}while(currDate.valueOf() !== untilDate.valueOf());
	return sundayCount;
};

print(sunday_first());