package leetcodealgorithms

import "strconv"

/*
Marie invented a Time Machine and wants to test it by time-traveling to visit Russia on the Day of the Programmer (the 256th day of the year) during a year in the inclusive range from 1700 to 2700.

From 1700 to 1917, Russia's official calendar was the Julian calendar;
since 1919 they used the Gregorian calendar system. The transition from the Julian to Gregorian calendar system occurred in 1918,
when the next day after January 31st was February 14th.
This means that in 1918, February 14th was the 32nd day of the year in Russia.
*/

func dayOfProgrammer(year int32) string {

	str := strconv.Itoa(int(year))

	if year < 1918 {
		if year%4 == 0 {
			return "12.09." + str
		} else {
			return "13.09." + str
		}
	} else if year > 1918 {
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			return "12.09." + str
		} else {
			return "13.09." + str
		}
	} else {
		return "26.09.1918"
	}

}
