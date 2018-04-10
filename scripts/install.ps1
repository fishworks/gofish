$version = "0.1.0"
$url = "https://gofi.sh/releases/gofish-v$version-windows-amd64.zip"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$tempDir = Join-Path $env:TEMP 'Fish'
if (![System.IO.Directory]::Exists($tempDir)) {[void][System.IO.Directory]::CreateDirectory($tempDir)}
$file = Join-Path $env:TEMP "gofish-v$version-windows-amd64.zip"

# Download fish
Write-Output "Downloading $url"
(new-object System.Net.WebClient).DownloadFile($url, $file)

$installPath = "$env:SYSTEMDRIVE\ProgramData\bin"
if (![System.IO.Directory]::Exists($installPath)) {[void][System.IO.Directory]::CreateDirectory($installPath)}
Write-Output "Preparing to install into $installPath"

Expand-Archive -Path "$file" -DestinationPath "$tempDir" -Force
Move-Item -Path "$tempDir\windows-amd64\gofish.exe" -Destination "$installPath\gofish.exe"

Write-Output "gofish installed into $installPath\gofish.exe"
Write-Output "Restart your terminal, then run 'gofish init' to get started!"

# Add gofish to the path
if ($($env:Path).ToLower().Contains($($installPath).ToLower()) -eq $false) {
  $newPath = [Environment]::GetEnvironmentVariable('Path',[System.EnvironmentVariableTarget]::Machine) + ";$installPath";
  [Environment]::SetEnvironmentVariable('Path',$newPath,[System.EnvironmentVariableTarget]::Machine);
}
