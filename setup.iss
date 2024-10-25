[Setup]
AppName=Lemon
AppVersion=1.0
DefaultDirName={pf}\Lemon
DefaultGroupName=Lemon
OutputBaseFilename=LemonSetup
Compression=lzma
SolidCompression=yes

[Files]
Source: "lemon.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\Lemon"; Filename: "{app}\lemon.exe"
Name: "{group}\Uninstall Lemon"; Filename: "{uninstallexe}"

[Code]
function AddToPath(Param: String): Boolean;
var
  Path: String;
  Key: String;
begin
  Key := 'Environment';
  if not RegQueryStringValue(HKEY_CURRENT_USER, Key, 'Path', Path) then begin
    Path := '';
  end;
  if Pos(Param, Path) = 0 then begin
    if (Path <> '') and (Path[Length(Path)] <> ';') then begin
      Path := Path + ';';
    end;
    Path := Path + Param;
    Result := RegWriteStringValue(HKEY_CURRENT_USER, Key, 'Path', Path);
  end else begin
    Result := True;
  end;
end;

procedure CurStepChanged(CurStep: TSetupStep);
begin
  if CurStep = ssPostInstall then begin
    if AddToPath(ExpandConstant('{app}')) then
      MsgBox('Lemon has been successfully installed and added to your PATH.', mbInformation, MB_OK)
    else
      MsgBox('Failed to add Lemon to your PATH.', mbError, MB_OK);
  end;
end;
