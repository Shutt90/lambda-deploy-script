# Lamba Deploy Script

## How to use this script
To start using the deploy script you can simply build it, providing you have Go installed just open and run go build.
This will create the binary file.

For unix users you will need to make this executable by running the following command
(`sudo chmod +x main`)

You should then be able to run the deploy script as long as the position is within the root directory of your lambdas. <br />
-lambdas <br/>
 deploy <br/>
--some lamda <br/>
 lambda files <br/>
--another lambda <br/>
 lambda files <br/>

This will run over all of the children folders which have the names you typed, until you type 'exit'