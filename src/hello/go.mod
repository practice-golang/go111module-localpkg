module example.com/me/hello

require (
	hello/spkg v0.0.0
	mypkg v0.0.0
	vpkg v0.0.0
)

replace (
	hello/spkg => ./spkg
	mypkg => ../mypkg
	vpkg => ./vendor/vpkg
)
