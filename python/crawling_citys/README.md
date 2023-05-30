## python打包命令Pyinstaller
pyinstaller打包python程序为exe，参数：
-F：打包后只生成单个exe格式文件；

-D：默认选项，创建一个目录，包含exe文件以及大量依赖文件；

-c：默认选项，使用控制台(就是类似cmd的黑框)；只对 Windows 有效

-w：不使用控制台；这在GUI界面时非常有用。不过如果是命令行程序的话那就把这个选项删除吧！只对 Windows 有效

-p：添加搜索路径，让其找到对应的库；用路径分割符 (Windows 使用分号 ,Linux 使用冒号 ) 分割 , 指定多个目录 . 也可以使用多个 -p 参数来设置多个导入路径，让pyinstaller自己去找程序需要的资源，和使用 PYTHONPATH 效果相似

-i：改变生成程序的icon图标。只在windows下生效
-o DIR, –out=DIR	指定 spec 文件的生成目录 , 如果没有指定 , 而且当前目录是 PyInstaller 的根目录 , 会自动创建一个用于输出 (spec 和生成的可执行文件 ) 的目录 . 如果没有指定 , 而当前目录不是 PyInstaller 的根目录 , 则会输出到当前的目录下 .
–add-data <SRC;DEST or SRC:DEST> 添加资源文件（Windows使用;分号，大多数Unix使用:冒号，注意后面还有一个.点）可以添加dll ico等资源文件


-v FILE, –version=FILE	将 verfile 作为可执行文件的版本资源 ( 只对 Windows 系统有效 )
-n NAME, –name=NAME	可选的项目 ( 产生的 spec 的 ) 名字 . 如果省略 , 第一个脚本的主文件名将作为 spec 的名字