# pdf-designer

a designer for PDF, useful for writing resumes.

## origin of the project

> When I started writing my resume, I saw many websites that provided templates. These websites provide online resume design. So I came up with the idea, "As a front-end worker, why don't I create an exclusive design tool for my resume?"

## Installation and how to use it

You can download it on the [release page](https://github.com/ddosakura/pdf-designer/releases), or

```bash
go get -v -u https://github.com/ddosakura/pdf-designer
```

See how to use it:

```bash
pdf-designer help
```

Find the templates:

+ [Built-in](./assets/template)
+ [Official](https://github.com/ddosakura/pdf-designer-template)

**Welcome to share your template**

### How to Share Your Templates

1. fork the prokect and PR (add in Built-in or change the list in the README)
2. send email to [ddosakura@qq.com](mailto:ddosakura@qq.com), it will be possible to join the [official­-template](https://github.com/ddosakura/pdf-designer-template).
3. fork [official­-template](https://github.com/ddosakura/pdf-designer-template) and PR

### How to add images

+ [iconfont](https://www.iconfont.cn/manage/index?spm=a313x.7781069.1998910419.11&manage_type=myprojects)
+ online images
+ Built-in url: `/wp/*/*.{jpg|png}`

## TODO

+ ~~[ ] Replace statik with [sblock](https://github.com/ddosakura/sblock), after sblock had its first stable version~~ **Because statik has a bug that is difficult to handle when reading subdirectories, it uses sblock directly.**
+ [ ] replace [sblock](https://github.com/ddosakura/sblock) with release version and [aferox](https://github.com/ddosakura/aferox)
+ [x] auto load iconfont
+ [ ] Create a dedicated iconfont project
+ [ ] relative path support in templates
+ [ ] ignore *.txt in saving
+ [ ] change the size of page (not only A4)
+ [ ] pdf-designer-ui project
