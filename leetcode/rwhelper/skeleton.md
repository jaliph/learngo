

defer rwhelper.CloseWriter()
n := rwhelper.ReadInt()
for i := 0; i < n; i++ {
	rwhelper.Write("helll world", i)
}