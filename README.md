# 曹操传开源重制

---
---
## 项目介绍

本项目采用GO语言作为主要开发语言

一方面可以练手GO语言

一方面可以参考OpenDiablo2的编码思路

ebiten虽然不怎么好用，但是暂时够用。渲染效率对于CCZ这样的回合制SLG游戏不是特别重要。

开发目标：
1. 可以快速适配不同的软硬件平台(Windows/Linux/MacOS on x64/Arm64)
2. 保持原版风格的基础上，调整画面分辨率。(640 \* 480 -\> 1280 \* 960)。
3. 将有一些游戏机制上的修改。
4. 方便喜欢MOD制作的朋友接入。

---
---

## 项目结构

c2common: 

c2core: 核心类库。

c2game: Game类。

resources: 各类素材图片/音频，暂时以文件存在。

main.go: 游戏入口

---
---
## 安装

系统需要GO语言环境[Recommend Version:1.17.0]

### Windows
install.bat (文件尚不存在)

### MacOS/Linux
bash install.sh (文件尚不存在)

---
---
## 开发记录

---
### 2022.01.28 建立项目

解决素材来源问题，大部分素材来自于以下开源项目
[uhziel/ccz_csdn](https://gitee.com/uhziel/ccz_csdn)

---
### 2022.01.29 使用腾讯Arc实验室开源的超分辨率模型重制素材

[Real-ESRGAN](https://github.com/xinntao/Real-ESRGAN "Real-ESTGAN")

文档要求python版本>=3.7，实测本机python3.6可以跑通。

1. 本机只能使用CPU，且只能用1核，处理速度较慢。不知道是否有参数可以指定CPU线程数。
2. SSL连接需要手动下载两个模型到指定位置。不知道是否和Python版本有关系。

使用4倍分辨率模型，目前看效果还可以。

单位动作动画效果有点糊，大型场景/背景、地图效果非常nice。

之后如果算法有更新，使用最新的模型重新生成素材即可。

---
### 2022.01.30 试验一下代码风格

```
#include<iostream>
int main()
{
    return 0;
}
````
