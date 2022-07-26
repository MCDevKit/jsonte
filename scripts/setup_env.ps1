Write-Host "This script will download ANTLR4 to C:\antlr, create a runner script and modify CLASSPATH and PATH environment variables to enable antlr command system-wide"
$confirmation = Read-Host "Are you Sure You Want To Proceed:"
if ($confirmation -eq 'y') {
    # Create antlr folder
    New-Item -ItemType Directory -Force -Path "C:\antlr" | out-null
    # Download antlr
    Write-Host "Downloading ANTLR4"
    Invoke-WebRequest -Uri "https://www.antlr.org/download/antlr-4.10.1-complete.jar" -OutFile "C:\antlr\antlr-4.10.1-complete.jar"
    Write-Host "Setting environment variables"
    # Add to CLASSPATH
    [Environment]::SetEnvironmentVariable("CLASSPATH", ".;C:\antlr\antlr-4.10.1-complete.jar" + $env:CLASSPATH, [System.EnvironmentVariableTarget]::User)
    # Add to Path
    [Environment]::SetEnvironmentVariable("Path", "C:\antlr\;" + $env:Path, [System.EnvironmentVariableTarget]::User)
    # Create runner script
    Write-Host "Creating runner script"
    [IO.File]::WriteAllLines("C:\antlr\antlr.bat", "java org.antlr.v4.Tool %*")
    Read-Host "Done. Press any key to continue"
}
