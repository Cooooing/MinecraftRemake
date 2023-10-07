<h1 align="center">Minecraft Remake</h1>
<div align="center">
    <strong>使用 go 对 minecraft 的重制</strong>
</div>

## 简介

出于学习 go 语言和 OpenGL 的目的，对 minecraft 进行重制。

## 运行

暂无

## 拉取代码

~~~bash
git clone https://github.com/Cooooing/MinecraftRemake.git
~~~

## 构建可执行程序

Windows：
~~~bash
go build -ldflags -H=windowsgui main.go 
~~~

Linux：
暂无

## 关于

本来是决定使用 Java 和 lwjgl 对 minecraft 进行重制的。
但是考虑到 Java 的性能问题，和“我不能死在 Java 上”。
以及，Java 打包后的程序需要 Java 环境，而 Go 则可以根据平台打包成不同平台的可执行文件。
我相信不少玩家刚接触 minecraft 的时候，都或多或少遇到过 Java 的问题。
所以选择了 Go 语言。
另外，希望 Go 的多线程是助力，而不是阻碍。（得熟练使用 Go 啊...

希望这个项目能帮助到大家。

## 鸣谢

[我的世界Minecraft 彩虹像素 RainbowPixel~☆材质包 - 爱发电](https://afdian.net/a/Nan2uu)

## 开源许可

本项目使用 MIT 许可。
