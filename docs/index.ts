import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Cruise",
  description: "A minimal Docker TUI client",
  appearance: 'force-dark',
  base: '/cruise/',
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'About', link: '/about/' },
      { text: 'Docs', link: '/docs/' }
    ],

    sidebar: {
      "/": [
        { text: "About", link: "/about/" },
        { text: "Getting Started", link: "/getting-started/" },
        { text: "Documentation", link: "/docs/" }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/NucleoFusion/cruise' }
    ],
    lastUpdated: {
      text: "Last Updated",
    },
    search: {
      provider: 'local',
    }
  }
})
