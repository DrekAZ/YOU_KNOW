module.exports = {
  title: 'YOU KNOW(仮)',
  description: 'vue',
  themeConfig: {
    //description: 'YOU KNOW',
    //nav: [
      //{ text: 'About', link: '/about/'}
    //],
    /*sidebar: [
      '/'
    ],*/
    footer: {
      copyright: [
        {
          text: 'Privacy Policy',
          link: '',
        },
        {
          text: 'Contact',
          link: '',
        },
      ],
    }
  },
  base: '/',
  //hostname: '',
  head: [
    ['link', { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons' }],
    ['link', { rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/npm/@mdi/font@4.x/css/materialdesignicons.min.css'}],
  ],
  markdown: {
    lineNumbers: true,
    anchor: {
    },
  },
  plugins: [
    ['@vuepress/plugin-blog', {
      directories: [
        {
          id: 'posts',
          dirname: '_posts',
          path: '/posts/',
          layout: 'PageList',
          itemLayout: 'Article',
          itemPermalink: '/posts/:year/:month/:day/:slug',
          pagenation: {
            lengthPerPage: 20,
          }
        },
        {
          id: 'official',
          dirname: '_official',
          path: '/official/',
          layout: 'PageList',
          itemLayout: 'Article',
          itemPermalink: '/official/:slug',
          pagenation: {
            lengthPerPage: 20,
          }
        }
      ],
    }],
  ],
  markdown: {
    extendMarkdown: md => {
      md.set({ injected: true })
    },
  },
}