task :run do
	
		# sh  "pkill simplea
		sh  "go get"
		sh  "go build"
		sh  "./simplyactive-api"
end

task :build do
	
		sh  "go get"
		sh  "go build"
end