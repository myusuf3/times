# times
simplified and consistent times and timezones library in golang.

Golang has an awesome `time` package compared to most languages but the
return values aren't always the most sane or consistent.

This library is fairly opinionated on how you should deal with times
internally in your system and how you should deal with timezones.

The general idea being that internally to your app simple user UTC as
your timezone/local and only change out to localized/normalized version
of your time for representation view in your application. e.g. when the
user in a specific locale is interacting with your application.

This library aims to provide a consisent api for such interactions as a
wrapper around the `time` package in golang proper.

For example the Parse function in the `time` package is full of caveats
which can be less than convient trying to remember in the heat of the moment!

>In the absence of a time zone indicator, Parse returns a time in UTC.

>When parsing a time with a zone offset like -0700, if the offset corresponds to a time zone used by the current location (Local), then Parse uses that location and zone in the returned time. Otherwise it records the time as being in a fabricated location with time fixed at the given zone offset.

Check out the double name reuse of Unix in the package as well. Granted one operates on
a `Time` instance but still, you had all the colors of the wind.

# func (Time) Unix

> Unix returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.


#func Unix

> Unix returns the local Time corresponding to the given Unix time, sec seconds and nsec nanoseconds since January 1, 1970 UTC. It is valid to pass nsec outside the range [0, 999999999].