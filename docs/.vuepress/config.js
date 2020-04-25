module.exports = {
  title: 'You know',
  description: 'VUE',
  locales: {
    '/': {
      lang: 'ja',
    }
  },
  meta: [
    { charset: 'utf-8' }, 
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
  ],
  head: [],
  base: '/',

  plugins: [
    [
    '@vuepress/blog',
      {
        directories: [{
          id: 'post',
          dirname: '_posts',
          path: '/post',
        }],
        sitemap: {},
      }
    ],
  ],

  themeConfig: {
    sidebar: 'auto',
    sidebarDepth: 3
  }
}