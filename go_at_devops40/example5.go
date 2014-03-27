func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close() // Метод Close будет гарантированно вызван перед выходом из функции CopyFile.

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close() // Аналогично, dst будет закрыт сразу после вызова io.Copy(dst, src).

    return io.Copy(dst, src)
}