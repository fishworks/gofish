$version = "v0.5.0"
if (-Not $env:VERSION -eq $null) {
  $version = "$env:VERSION"
}
$url = "https://gofi.sh/releases/gofish-$version-windows-amd64.zip"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$tempDir = Join-Path $env:TEMP 'Fish'
if (![System.IO.Directory]::Exists($tempDir)) {[void][System.IO.Directory]::CreateDirectory($tempDir)}
$file = Join-Path $env:TEMP "gofish-$version-windows-amd64.zip"

# Download fish
Write-Output "Downloading '$url'"
(new-object System.Net.WebClient).DownloadFile($url, $file)

$installPath = "$env:SYSTEMDRIVE\ProgramData\bin"
if (![System.IO.Directory]::Exists($installPath)) {[void][System.IO.Directory]::CreateDirectory($installPath)}
Write-Output "Preparing to install into '$installPath'"

$binaryPath = "$installPath\gofish.exe"
Expand-Archive -Path "$file" -DestinationPath "$tempDir" -Force
if ([System.IO.File]::Exists("$binaryPath")) {[void][System.IO.File]::Delete("$binaryPath")}
Move-Item -Path "$tempDir\windows-amd64\gofish.exe" -Destination "$binaryPath"

# Add gofish to the path
if ($($env:Path).ToLower().Contains($($installPath).ToLower()) -eq $false) {
  Write-Output "Adding '$installPath' to system PATH"
  $newPath = [Environment]::GetEnvironmentVariable('Path',[System.EnvironmentVariableTarget]::Machine) + ";$installPath";
  [Environment]::SetEnvironmentVariable('Path',$newPath,[System.EnvironmentVariableTarget]::Machine);
}

Write-Output "gofish installed to '$binaryPath'"
Write-Output "Restart your terminal, then run 'gofish init' to get started!"


# -----BEGIN PGP PUBLIC KEY BLOCK-----

# mQINBFp8op4BEADDk9QQYbTaq+nXTvxxTyjBqmTS3CsN94y8TfxslVLbQym3wuT5
# 9zWUv+JRlIZoqZiXJvrXFuJnUgTuzniYkrLtxvUWRKY8cISqcuA26d40YuBeQXWl
# TvHAeDiJ3vyLRbS38/tzcEYThojuc0CMIklqDwrwI2J3DAegkfc4jAB60Q9oYo/M
# IlsRxS1jaLMoe3xUFJV8Lq9BQOehqNSpL+L0lCHsXmLJjczuAE+pYReIuAlyeke5
# o5/t4fdEqc5hziTN7XJGF5qAHK4+iEuzYx1M1UHuHqCOzZX8P/KOXT90Iru8HZ11
# r/NueLUHIK/PALFBd8tC1zu03pyEKBgSd/qdsyFvIlJHQIEUOe66RaBOLhhBm9o7
# 97dWceYrktl0xIJxLRlcwQAb8/+xivDsQAdERU3xwVbWk9UQJx8Jy9fjI8bvJF/n
# 3XKvGIX/FZVzgicZExJVbN/braP9obHswvUe5mKDwsllQPz7FuTjoaxAj+Px9iPt
# l5jLzk++pjEXSmbljeW6lDBh4T997szZTs8Vp1LdJprpZZcyOmA5ct/uilNsXJaY
# 8K1agfHNTb5n+wDwrVnepwPo7bpwONgktzgnvxa6Vhi4UEvE5JXx6R9YFM8cpS61
# TDVAg82i0PGZa98RKxRs/NH2ynUsplhkzIDbbDWZZPYCwp5EMAk0i/NyuwARAQAB
# tCpNYXR0aGV3IEZpc2hlciA8bWF0dC5maXNoZXJAbWljcm9zb2Z0LmNvbT6JAjgE
# EwECACIFAlp8op4CGwMGCwkIBwMCBhUIAgkKCwQWAgMBAh4BAheAAAoJEM3sZ2h+
# +qNOo50P/040uUwZ4FE7iHkwn7Da2Nqbe1hvadcvWUfVkXxCHTKXnAd3RLeEmPBE
# a+vkhNPA3vd57EGqDfBw2zdMVx+jbKgrUxwUrn506TgwnHFq+gaqm+JVQQrEz0/Z
# ZeO7xqlbSrwHP5jwMKFPCUF6QdBuyCqH4GnUiCMa69Vp3hEx1bGqEPeSAMWoBq1k
# i5H8A/oB/9498J8SSNGOhCN0igG50IgQw1bHoHvWMmsbNslFHnQ8s+M6q0LZzC6w
# 6gUnntvcNouWZqggMyBDh9QT8APH2c1iRuwPZU0t821O7RX9kzjDfDPpC/9QLW9c
# 2nEF5QU0EGF7HCd+7feqQHiDZT7B/VEXoWt89j7Fz8lDLs+jXl+G8MG+0F+tNcq6
# YkdNjTDw7HG4TQnreUnb+pKbRSb7By3gQA59ovwtRhpLCwIXLjO5YyYvyzjlKrQt
# 0jYtkwBCrulKq2EnF6lsg1eowTNQ2FaGBZZCCT46IDT1z39bSfKxHA0lZyE5FApS
# QKFqTkNnspIMLBpblYAyYeF4wLwyHmKXUZioTq09fd7QXWKUwIj5q5xJHBDtwlly
# HcGTjb6mJACsK46hBL5LIlYqCiJmaLilWQ02K+drVBAqFnQrGI7stH/8S/yP6+PQ
# EUjjmbDAxuplLhlo094qqmpPDJxBoUnEOdbERppKXrlUZf6HeOkluQINBFp8op4B
# EADdlBwTZ8XKGJUf0oINAQ7/4GZJCinBGWKHRmM+HrWsB+O13FEeWvGHNJQ5hswa
# aKaIylq4+uF8nnpf3OV4zEfZEguECIzqHo78JUAGRbVREnqzBN7eLQHeAnyvt/yd
# veEza1Uu/G/Gyi/QfL+nGfInjPcsqOY6kWfhlSYGLWMmi86cBwAHSxxCgRjLGZtB
# lID1qLXPeCZyqyvPLMuENIt01I6qnFQL+RLaXFpCjeXCF2XY5lT6VZ9DPJxmb/dE
# ZHFCZeplCPBIJd7ihMPO4j8gen7nuh1n6egU6h+aWJ4APtwkAQi8ESOM8yGTgPnW
# AFqEAIkVdR9erf894LZ6WKfuPKhtQULoVIBMM0idmmccqlejG6DlFsCua1i1Mce7
# xgfqRd1GxW6U0NzY4BPdyU91ggc+hhkUmpfB4F/ggZOs8N4ckiDFnXkN6ShVpXNt
# eMcaltsv5XpHqsBBy6vixn9vMdWavViBGLDapiloba97CbiHLajGuUDeyOHZSdCK
# 38ecMOfVHhWGI5QOb0ayDwI4pSsBmxdGc/mNsMLNS5ScTJBIsZKMKMnPIUzAFgZp
# /c4dCk/+cqAssejoX24UElkmoucrFJtxIt0blDlODybm8kTSHf0gVOYuK41OvbkJ
# cz5tLIUfZMRdlMfz7wSGVXamITjoEwSQ9uzD2fXA8ui1pQARAQABiQIfBBgBAgAJ
# BQJafKKeAhsMAAoJEM3sZ2h++qNOj/gQAI1MN167Gz9Zb4uzep+GzYSB2Bp6AGnw
# VHAIVFTTuUfz9yUa2Cv+nVuUdh4W0CHfwWUbaUqlP9NiYiMi/tcwgSiCdeCK1G/x
# PmJyyp17bfOiFhw042kbpX2yp2N2e3IUfNjOOyxh+X44rVcLLAuwIqVpDG8RCn6b
# /wBfXc2YCq4wWX+QYqkRrb3Qd0cOVfGDyyaSiLrj20oi9KO0PjwZ1qKRlk7xBC6R
# /0IvZwwFRVJfFbLlmCHmGkaSwpwvMLumN/wlV3T8Px/XGIYuK7RxhslZhB2b75Sj
# xnF8TfQPyTRpYVscujad0HdPd9uNrsQxrOsXwzMIyyJ+Nsn5IEt8jyMFS7mw3orW
# V9nmxBrdBUlUzxiCuwwzHrxdUXRwA6gJVNS5rk8m1tpIl+JNao66i3fGbYSSvGV4
# wqOwwpl0gLhnx5Xy3lwfrdLIDSgRDZ0i815Pcr4Q+j4O/PJXWt6J83aH8FiCTEuM
# BAjo8CuHR37jaO7ZliGsJW/kWVntM8uW7C1r+46sqG7nBlHmNW/E79/C9XqtT0Rp
# NOkfRaUyYKkjSgckeM7VzNCT4iLIUPUHGm+YmwHuh83IMRNuoELlMpoBr7DgqUtY
# N6lsOkvlvsLIL+TzuWb0sUKbtEUh1jiQUDtzGig1LGsimRBevH2IvHgBWxzPVpw6
# Qz0L9+I2ai68
# =dBiU
# -----END PGP PUBLIC KEY BLOCK-----
