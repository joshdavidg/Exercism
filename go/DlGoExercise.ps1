param (
    [string] $Stub,
    [switch] $Easy,
    [switch] $Medium,
    [switch] $Hard
)

cmd.exe /c "exercism download --exercise=${Stub} --track=go"

if ($Easy) {
    Move-Item ".\${Stub}\" ".\Easy\"
} elseif ($Medium) {
    Move-Item ".\${Stub}\" ".\Medium\"
} elseif ($Hard) {
    Move-Item ".\${Stub}\" ".\Hard\"
}