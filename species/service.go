package species

import (
    "fmt"
    "github.com/creack/pty"
    "io"
    "os"
    "os/exec"
    "path"
)


/**
 * $(Call)
 * @Description:
 * @param filePath
 * @param script
 * @param calls
 */
func Call(filePath, script string, calls ...func(msg string) string)  {

    homeDir, _ := os.UserHomeDir()
    _ = os.MkdirAll(path.Join(homeDir, Workspace), os.ModePerm)

    filepath := path.Join(homeDir, Workspace, filePath)
    fp, _ := os.Create(filepath)
    _, _ = fp.WriteString(script)
    command := exec.Command("bash", filepath)
    f, err := pty.Start(command)
    if err != nil {
        fmt.Println(`error:`, err)
        return
    }
    _, _ = io.Copy(os.Stdout, f)
    for _, call := range calls {
        fmt.Println(call(`[SUCCESS] `))
    }

}
