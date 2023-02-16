module.exports = {
    title: 'Go-To-Byte 抖声项目文档',
    description: '极简版抖音相关文档',
    base: '/test_docs/',
    theme: 'reco',
    locales: {
        '/': {
            lang: 'zh-CN'
        }
    },
    themeConfig: {
        nav: [
            { text: '首页', link: '/' },
            {
                text: 'dousheng',
                items: [
                    { text: 'Github', link: 'https://github.com/Go-To-Byte/DouSheng' },
                    { text: '掘金', link: 'https://juejin.cn/user/4332537967820215' }
                ]
            }
        ],
        subSidebar: 'auto',
        sidebar: [
            {
                title: '欢迎查看',
                path: '/',
                collapsable: false, // 不折叠
                children: [
                    { title: "看前必读", path: "/" }
                ]
            },
            {
                title: "IOC相关",
                path: '/handbook/ioc/ioc',
                collapsable: false, // 不折叠
                children: [
                    { title: "IOC的设计", path: "/handbook/ioc/ioc" },
                    { title: "IOC的测试", path: "/handbook/ioc/ioctest" }
                ],
            }
        ]
    }
}