package manager 

import (
    "fmt"
    "io/ioutil"
    "io"
    "os"
    "syscall"
    //"errors"
    "path/filepath"
)

func (p *Manager) CopyDirectory(scrDir, dest string) error {
    entries, err := ioutil.ReadDir(scrDir)
    if err != nil {
        return err
    }
    if !p.Exists(dest){os.MkdirAll(dest, 0755)}
    for _, entry := range entries {
        sourcePath := filepath.Join(scrDir, entry.Name())
        destPath := filepath.Join(dest, entry.Name())

        fileInfo, err := os.Stat(sourcePath)
        if err != nil {
            return err
        }

        stat, ok := fileInfo.Sys().(*syscall.Stat_t)
        if !ok {
            return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
        }

        switch fileInfo.Mode() & os.ModeType{
        case os.ModeDir:
            if err := p.CreateIfNotExists(destPath, 0755); err != nil {
                fmt.Println("-------------------------3")
                return err
            }
            if err := p.CopyDirectory(sourcePath, destPath); err != nil {
                fmt.Println("----------------------------4")
                return err
            }
        case os.ModeSymlink:
            if err := p.CopySymLink(sourcePath, destPath); err != nil {
                fmt.Println("--------------------------5")
                return err
            }
        default:
            if err := p.Copy(sourcePath, destPath); err != nil {
                fmt.Println("-------------------------------6")
                return err
            }
        }

        if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
            return err
        }

        isSymlink := entry.Mode()&os.ModeSymlink != 0
        if !isSymlink {
            if err := os.Chmod(destPath, entry.Mode()); err != nil {
                return err
            }
        }
    }
    return nil
}

func (p *Manager) Copy(srcFile, dstFile string) error {
    out, err := os.Create(dstFile)
    if err != nil {
        return err
    }

    defer out.Close()

    in, err := os.Open(srcFile)
    defer in.Close()
    if err != nil {
        return err
    }

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    return nil
}

func (p *Manager) Exists(filePath string) bool {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return false
    }

    return true
}

func (p *Manager) CreateIfNotExists(dir string, perm os.FileMode) error {
    if p.Exists(dir) {
        return nil
    }

    if err := os.MkdirAll(dir, perm); err != nil {
        return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
    }

    return nil
}

func (p *Manager) CopySymLink(source, dest string) error {
    link, err := os.Readlink(source)
    if err != nil {
        return err
    }
    return os.Symlink(link, dest)
}