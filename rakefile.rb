task :build => ["get"]do 
	sh "go build"
end

task :get do 
	sh "go get"
end

namespace :run do
	
	task :local => ["build", "config:local"] do
		sh  "./simplyactiveapi"
	end

	task :dev => ["build", "config:dev"] do
		sh  "./simplyactiveapi"
	end
end

namespace :config do

	task :local do
		ENV['PORT']="4200"
		ENV['MONGOURI']="127.0.0.1:27017"
		ENV['DBNAME']="simplyactivedb"
		ENV['APIKEY']="1234"
		ENV['CNAME']="meals"
	end

	task :dev do
		ENV['PORT']="4300"
		ENV['MONGOURI']="mongodb://admin:admin@ds029793.mongolab.com:29793/simplyactivedb"
		ENV['DBNAME']="simplyactivedb"
		ENV['APIKEY']="12358"
		ENV['CNAME']="meals"
	end

	task :mimic do
		ENV['PORT']="3800"
		ENV['MONGOURI']="127.0.0.1:27017"
		ENV['DBNAME']="mongoid"
		ENV['APIKEY']="1234"
		ENV['CNAME']="meals"
	end
end