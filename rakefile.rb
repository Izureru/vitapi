task :build do
	
		sh  "go get"
		sh  "go build"
end

namespace :run do

	task :dev do
		ENV['PORT']="8080"
		sh  "go get"
		sh  "go build"
		sh  "./simplyactiveapi"
	end

	task :prod do
		sh  "go get"
		sh  "go build"
		sh  "./simplyactiveapi"
	end
end