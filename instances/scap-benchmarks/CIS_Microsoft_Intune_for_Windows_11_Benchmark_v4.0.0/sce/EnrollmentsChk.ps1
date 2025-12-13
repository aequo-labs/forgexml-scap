<#
.SYNOPSIS
    SCE Check for two Enrollment settings, Intune Benchmark
.DESCRIPTION
    Takes a string from a WorkBench environment variable, splits it and uses it in
    an assessment for keys with unique ProviderID GUIDS.
.NOTES
    Format WB artifact string using below syntax
    <Key>::<ValueName>::<ValueType>::<Operator>::<Value>
.EXAMPLE
    ConfigRefresh::Enabled::DWORD::EQ::1
    ConfigRefresh::Cadence::DWORD::LE::90
#>

$WBString = $env:XCCDF_VALUE_REGEX

$EnrollmentsPath = "HKLM:\SOFTWARE\Microsoft\Enrollments\"

if (Test-Path -Path HKLM:\SOFTWARE\Microsoft\Enrollments\ -ErrorAction SilentlyContinue) {
    $Providers = Get-ChildItem -Path $EnrollmentsPath | Where-Object { $_.Property -contains 'ProviderID' }    
}
else {
    Write-Error "$EnrollmentsPath does not exist or is inaccessible."
    exit $env:XCCDF_RESULT_ERROR
}

foreach ($Provider in $Providers) {
    $ProviderID = $Provider.GetValue('ProviderID')
    if ($ProviderID -eq 'MS DM Server') {
        $ProviderGUID = $Provider.PSChildName
        break
    } 
}

if ([string]::IsNullOrEmpty($ProviderGUID)) {
    Write-Error "A matching provider GUID for Intune was not found."
    exit $env:XCCDF_RESULT_ERROR
}

# Split WorkBench string and construct variables
$RegSplit = $WBString -split '::'
$RegPath = "HKLM:\SOFTWARE\Microsoft\Enrollments\$ProviderGUID\{0}" -f $RegSplit[0]
$RefValueName = $RegSplit[1]
$RefValueType = $RegSplit[2]
$RefOperator = $RegSplit[3]
$RefValue = switch ($RefValueType) {
    'DWORD' { [int]$RegSplit[4] }
    'STRING' { [string]$RegSplit[4] }
    Default { [string]$RegSplit[4] }
}
if (Test-Path -Path $RegPath) {
    $RegItem = Get-Item -Path $RegPath
} else {
    # Failed because Key does not exist
    Write-Host "The key $($RegSplit[0]) does not exist."
    exit $env:XCCDF_RESULT_FAIL
}

$RegValue = $RegItem.GetValue($RefValueName)

# Construct Evidence
if ($RegItem.Property -contains $RefValueName) {
    $RegValueType = $RegItem.GetValueKind($RefValueName)
    $RegValueTypeExistence = "[Exists]"
    $RegValueName = $RefValueName
    $RegValueNameExistence = "[Exists]"
} else {
    $RegValueType = $RefValueType
    $RegValueName = $RefValueName
    $RegValueTypeExistence = "[Does not exist]"
    $RegValueNameExistence = "[Does not exist]"
}

$RegValueType = switch ($RegValueType) {
    'Dword' { 'REG_DWORD' }
    'String' { 'REG_SZ ' }
    Default { $RegValueType }
}
$Evidence = [PSCustomObject]@{
    Hive  = $RegItem.PSDrive.Root
    Key   = $RegPath
    Name  = "$RegValueName $RegValueNameExistence"
    Type  = "$RegValueType $RegValueTypeExistence"
    Value = if ($null -eq $RegValue) { "<value> [Does not exist]" } else { "$RegValue [Exists]" }
}
$Operator = switch ($RefOperator) {
    'EQ' { '-eq' } # Equal to
    'LT' { '-lt' } # Less than
    'GT' { '-gt' } # Greater Than
    'LE' { '-le' } # Less Than or Equal to
    'GE' { '-ge' } # Greater Than or Equal to
    Default { '-eq' }
}

if ($RegValueNameExistence -match 'Exists') {
    $AuditPass = Invoke-Expression "$RegValue $Operator $RefValue"
} else {
    $AuditPass = $false
    if ($RefValueName -eq 'Cadence') {
        # Cadence defaults to 90 when Enabled is 1, but Intune does not write the value 90 to the registry when the default value is 
        # specified via settings catalog. We need a pass for Cadence if it _does not exist_
        $AuditPass = $true
        Add-Member -InputObject $Evidence -MemberType NoteProperty -Name "Note" -Value "Passed due to existence check."
    } elseif ($RefValueName -eq 'Enabled') {
        Add-Member -InputObject $Evidence -MemberType NoteProperty -Name "Note" -Value "Failed due to existence check."
    }
}

# Results and Evidence
$AuditResult = if ($AuditPass) { 'PASS' } else { 'FAIL' }
$Result = "Audit Result: ** $AuditResult **`n"
$Result += "Assessment evidence:`n"
$Result += "----------------------------------------`n"
$Result += $Evidence | Format-List | Out-String
Write-Host $Result

if ($AuditPass) { exit $env:XCCDF_RESULT_PASS } else { exit $env:XCCDF_RESULT_FAIL }