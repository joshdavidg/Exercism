param (
    [string] $Stub,
    [switch] $Easy,
    [switch] $Medium,
    [switch] $Hard
)

cmd.exe /c "exercism download --exercise=${Stub} --track=go"

$json = Get-Content -Raw $env:APPDATA\exercism\user.json | ConvertFrom-Json
$dir = "{0}\go\" -f $json.workspace

if ($Easy) {
    Move-Item "$dir\$Stub\" "$dir\Easy\"
} elseif ($Medium) {
    Move-Item "$dir\$Stub\" "$dir\Medium\"
} elseif ($Hard) {
    Move-Item "$dir\$Stub\" "$dir\Hard\"
}