$version = "0.1.0"
$url = "https://bacongobbler.blob.core.windows.net/gofish/fish-v$version-windows-amd64.zip"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$file = Join-Path $env:TEMP "fish-v$version-windows-amd64.zip"

function Download-String {
param (
  [string]$url
 )
  $downloader = Get-Downloader $url

  return $downloader.DownloadString($url)
}

function Download-File {
param (
  [string]$url,
  [string]$file
 )

  $downloader = new-object System.Net.WebClient
  $downloader.DownloadFile($url, $file)
}

# Download fish
Write-Output "Downloading $url"
Download-File $url $file

Write-Output "Preparing to install into $installPath"
$installPath = "$env:SYSTEMDRIVE\ProgramData\bin"
Expand-Archive -Path "$file" -DestinationPath "$installPath" -Force

Write-Output "fish installed into $installPath\fish"
Write-Output "Run 'fish init' to get started!"

# Add fish to the path
if ($($env:Path).ToLower().Contains($($installPath).ToLower()) -eq $false) {
  $env:Path = [Environment]::GetEnvironmentVariable('Path',[System.EnvironmentVariableTarget]::Machine);
}
