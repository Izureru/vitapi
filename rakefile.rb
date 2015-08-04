task :run do
	
		# sh  "pkill simplea
		sh  "go get"
		sh  "go build"
		sh  "./simplyactiveapi"
end

task :build do
	
		sh  "go get"
		sh  "go build"
end