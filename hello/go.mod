module hello

require (
	mypkg v0.0.0
	spkg v0.0.0
)

replace (
	mypkg => ../mypkg
	spkg => ./spkg
)
