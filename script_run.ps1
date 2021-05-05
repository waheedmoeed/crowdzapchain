#Write-Host "removing previous state"
<<<<<<< HEAD
#rm  ~/.nsd
#rm  ~/.nscli

=======
#Remove-Item c:/.appd
#Remove-Item c:/.appli
Set-Location  .\GolandProjects\crowdzapchain
>>>>>>> b16dcb6f9d59c8304e305186a71fa966f3e51e09
Write-Host "creating default configuration for chain"

#change dir where built app located
Set-Location  .\cmd\appd
.\appd.exe init test --chain-id='relchain'

Write-Host "basic configuration for cli"
#change dir where built app located
Set-Location ..\appcli
.\appcli.exe config output json
.\appcli.exe config indent true
.\appcli.exe config trust-node true
.\appcli.exe config chain-id relchain
.\appcli.exe config keyring-backend test

Write-Host "adding key for abdul"
.\appcli.exe keys add abdulvalidator
Write-Host "adding key for haroon"
#.\appcli.exe keys add haroon
Write-Host "adding key for nadeem"
#.\appcli.exe keys add nadeemtest

#change dir where built app located
Set-Location  ..\appd

Write-Host "adding abdulvalidator account to genesis state"

#appd add-genesis-account cosmos12hhfqm2qqm23n6zkdl46vlegrsnlna8ld349v9 1000relcoin,100000000stake


.\appd.exe add-genesis-account $(..\appcli\appcli.exe keys show abdulvalidator -a) 1000rel,100000000stake

Write-Host "adding contract account to genesis state"
.\appd.exe add-genesis-account cosmos1ss4h432u3ltnf624kaj9k4ackukhzazntm2ryr 1000rel
.\appd.exe add-genesis-account cosmos167n2fd2qwzazthe8gpccv3ewj3f03kz3hw3ncr 1000rel
.\appd.exe add-genesis-account cosmos18x9my7vst6svm65sug8cjt2h060j5nd9mkdj3p 1000rel

Write-Host "creating keyring store..."
.\appd.exe gentx --name abdulvalidator --keyring-backend test

Write-Host "Collecting genesis txs..."
.\appd.exe collect-gentxs

Write-Host "Validating genesis file..."
.\appd.exe validate-genesis




##############
#############
Write-Host "Modifing Genesis State"
$pathToJsonGenesis = "C:\.appd\config\genesis.json"
$a = Get-Content $pathToJsonGenesis -raw | ConvertFrom-Json

#getting addresses from genesis state
#$abdulAddress = $a.app_state.auth.accounts.Get(0).value.address
#$haroonAddress = $a.app_state.auth.accounts.Get(1).value.address
#$waheedTestAddress = $a.app_state.auth.accounts.Get(1).value.address
$abdulAddress = 'cosmos1ss4h432u3ltnf624kaj9k4ackukhzazntm2ryr'
$haroonAddress = 'cosmos167n2fd2qwzazthe8gpccv3ewj3f03kz3hw3ncr'
$nadeemAddress = 'cosmos18x9my7vst6svm65sug8cjt2h060j5nd9mkdj3p'
#Write-Host "$du"

#####
#####

$pathToJsonContract = "C:\Users\Afaq Malik\GolandProjects\crowdzapchain\contract_temp.json"
$contract = Get-Content $pathToJsonContract -raw | ConvertFrom-Json

$contract.rel_contract.rel_contractors[0].contractor_address = $abdulAddress
$contract.rel_contract.rel_contractors[1].contractor_address = $haroonAddress
$contract.rel_contract.rel_contractors[2].contractor_address = $nadeemAddress
#$contract.rel_contract.rel_contractors[2].contractor_address = $waheedTestAddress


$contract.rel_contract.distributed_coins_logs[0].contractor_address = $abdulAddress
$contract.rel_contract.distributed_coins_logs[1].contractor_address = $haroonAddress
$contract.rel_contract.distributed_coins_logs[2].contractor_address = $nadeemAddress
#$contract.rel_contract.distributed_coins_logs[2].contractor_address = $waheedTestAddress

#####
#####

$a.app_state.relcontractors.rel_contract = $contract.rel_contract

$a | ConvertTo-Json -Depth 32 | Set-Content $pathToJsonGenesis
Write-Host "Validating genesis file..."
.\appd.exe validate-genesis
Write-Host "Start chain server and rest server"
#appd start
#appcli rest-server --chain-id crowdzap --trust-node