$version = "0.1.0"
$url = "https://bacongobbler.blob.core.windows.net/gofish/fish-v$version-windows-amd64.zip"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$tempDir = Join-Path $env:TEMP 'Fish'
if (![System.IO.Directory]::Exists($tempDir)) {[void][System.IO.Directory]::CreateDirectory($tempDir)}
$file = Join-Path $env:TEMP "fish-v$version-windows-amd64.zip"

# Download fish
Write-Output "Downloading $url"
(new-object System.Net.WebClient).DownloadFile($url, $file)

$installPath = "$env:SYSTEMDRIVE\ProgramData\bin"
if (![System.IO.Directory]::Exists($installPath)) {[void][System.IO.Directory]::CreateDirectory($installPath)}
Write-Output "Preparing to install into $installPath"

Expand-Archive -Path "$file" -DestinationPath "$tempDir" -Force
Move-Item -Path "$tempDir\windows-amd64\fish.exe" -Destination "$installPath\fish.exe"

Write-Output "fish installed into $installPath\fish.exe"
Write-Output "Restart your terminal, then run 'fish init' to get started!"

# Add fish to the path
if ($($env:Path).ToLower().Contains($($installPath).ToLower()) -eq $false) {
  $newPath = [Environment]::GetEnvironmentVariable('Path',[System.EnvironmentVariableTarget]::Machine) + ";$installPath";
  [Environment]::SetEnvironmentVariable('Path',$newPath,[System.EnvironmentVariableTarget]::Machine);
}
