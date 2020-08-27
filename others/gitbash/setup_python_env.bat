@echo off

set "python_path="
for /f "tokens=2*" %%a in ('reg query "HKLM\SOFTWARE\Python\PythonCore" /s /f "InstallPath" ^| findstr /i "InstallPath"') do (
    set "temp_path=%%b"
    echo %temp_path% | findstr /i "python3" >nul
    if %errorlevel% equ 0 (
        set "python_path=%temp_path%"
        goto :FOUND
    )
)

echo 未找到 Python 3.X 安装路径，请手动安装 Python 并重新运行脚本。
pause
exit /b 1

:FOUND

for /f "skip=2 tokens=2*" %%A in ('reg query "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path') do (
    set "current_path=%%B"
)

echo %current_path% | find /i "%python_path%" >nul
if %errorlevel% equ 0 (
    echo Python 路径已经在环境变量中。
) else (
    rem 将 Python 路径添加到环境变量中
    set "new_path=%current_path%;%python_path%;%python_path%\Scripts"
    reg add "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path /t REG_EXPAND_SZ /d "%new_path%" /f

    echo Python 路径已成功添加到环境变量中。
)

setx PATH "%new_path%"

echo 请重新启动 Git Bash 以使环境变量生效。
pause
    