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
        { text: "Getting Started", link: "/docs/getting-started/" },
        { text: "Documentation", link: "/docs/" }
      ],
      "/docs/": [
        { text: "Documentation", link: "/docs/" },
        { text: "Getting Started", link: "/docs/getting-started" },
        { text: "Installation", link: "/docs/install" },
        { text: "Configuration", link: "/docs/config/" },
        { text: "Contributing", link: "/docs/contributing" },
      ],
      "/docs/config/": [
        { text: "Configuration", link: "/docs/config/" },
        { text: "General", link: "/docs/config/general" },
        { text: "Keybinds", link: "/docs/config/keybinds" },
        { text: "Styles", link: "/docs/config/styles" },
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
