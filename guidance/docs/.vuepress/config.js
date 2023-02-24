module.exports = {
    title: 'Go-To-Byte 抖声项目文档',
    description: '极简版抖音相关文档',
    base: '/Docs-DouSheng/',
    theme: 'reco',
    locales: {
        '/': {
            lang: 'zh-CN'
        }
    },
    themeConfig: {
        nav: [
            { text: '首页', link: '/' },
            {text: 'Github', link: 'https://github.com/Go-To-Byte/DouSheng' ,}
        ],
        subSidebar: 'auto',
        sidebar: [
            {
                title: '欢迎查看',
                path: '/',
                collapsable: false, // 不折叠
                children: [
                    { title: "关于", path: "/" },
                    { title: "Go-To-Byte", path: "/index/team" },
                ]
            },
            {
                title: "初识Dousheng",
                path: '/handbook/why',
                collapsable: false, // 折叠
                children: [
                    { title: "为什么选择 Go", path: "/handbook/why" },
                    { title: "开始", path: "/handbook/start" },
                    { title: "功能特点", path: "/handbook/feature" },
                ],
            },
            {
                title: "Dousheng的相关架构",
                path: '/framework/evolve',
                collapsable: false, // 折叠
                children: [
                    { title: "Dousheng的架构演变之路", path: "/framework/evolve" },
                    { title: "如何管理DouSheng应用的生命周期", path: "/framework/lifecycle" },
                    { title: "如何将IoC使用到架构中", path: "/framework/ioc" },
                ],
            },
            {
                title: "Dousheng的公共库",
                path: '/kit/base',
                collapsable: true, // 折叠
                children: [
                    { title: "dou_kit基础", path: "/kit/base" },
                    { title: "统一error处理", path: "/kit/err" },
                    { title: "如何注入项目配置", path: "/kit/conf" },
                    { title: "Makefile工程化管理项目", path: "/kit/make" },
                    { title: "如何提供HTTP服务", path: "/kit/httpserver" },
                    { title: "如何提供GRPC服务", path: "/kit/grpcserver" },
                    { title: "怎么还算优雅的引入其他服务的protobuf", path: "/kit/proto" },
                ],
            },
            {
                title: "Dousheng的业务实现",
                path: '/service/user',
                collapsable: true, // 折叠
                children: [
                    { title: "用户中心", path: "/service/user" },
                    { title: "视频服务", path: "/service/video" },
                    { title: "评论服务", path: "/service/comment" },
                    { title: "关系服务", path: "/service/relation" },
                ],
            },
            {
                title: "其他",
                path: '/other/gowork',
                collapsable: true, // 折叠
                children: [
                    { title: "一份简单的测试报告", path: "/other/test" },
                    { title: "使用go.work的好处", path: "/other/gowork" },
                ],
            },
        ]
    }
}