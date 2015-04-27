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

