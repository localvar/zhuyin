This README is written in Chinese because this library is for Chinese phonetic notation, the user need to know some Chinese to use it.

If you're looking for the PHP version of this library, please checkout:　https://github.com/DictPedia/ZhuyinPinyin (此库的PHP版本)

# 注音

注音是中文处理的基本需求之一，在编辑低幼读物或古籍时更是常用。但现有的输入法和文字处理软件（如Word）处理起注音来却异常麻烦，要靠软键盘不停的切来切去才行。所以我就用Go语言实现了这个库，用来简化注音的处理。

中文注音有多种形式，这个库仅支持 **注音** 和 **汉语拼音** 这两种。而且，由于我个人不熟悉 **注音**，所以对它的处理完全基于查询到的资料，而没有任何实践经验，如有错误，敬请指出。

这个库的主要功能是把 **zhang1** 这样的输入转换成 **zhāng**（汉语拼音）或 **ㄓㄤ**（注音），也可以把 **zhāng** 或 **ㄓㄤ** 转换回 **zhang1**，同时也支持 **汉语拼音** 和 **注音** 之间的互转。

## 用法简介

这个库对外暴露了六个函数，每个函数都只有一个字符串型的输入参数和一个字符串型的返回值，当输入正确时，返回转换结果；如果输入有错误，则返回空字符串。下面是这六个函数的简介。

* EncodePinyin：把 **zhang1** 转换成 **zhāng**
* DecodePinyin：把 **zhāng** 转换成 **zhang1**
* EncodeZhuyin：把 **zhang1** 转换成 **ㄓㄤ**
* DecodeZhuyin：把 **ㄓㄤ** 转换成 **zhang1**
* PinyinToZhuyin：把 **zhāng** 转换成 **ㄓㄤ**
* ZhuyinToPinyin：把 **ㄓㄤ** 转换成 **zhāng**

## 授权许可

这个库使用 **MIT** 许可证发布。
