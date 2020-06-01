module.exports = {
  title: "Commitsar",
  tagline: "Conventional commit compliance made easy",
  url: "https://commitsar.tech",
  baseUrl: "/",
  favicon: "img/favicon.ico",
  organizationName: "aevea", // Usually your GitHub org/user name.
  projectName: "commitsar", // Usually your repo name.
  themeConfig: {
    navbar: {
      title: "Commitsar",
      logo: {
        alt: "My Site Logo",
        src: "img/logo.svg",
      },
      links: [
        {
          to: "docs/",
          activeBasePath: "docs",
          label: "Docs",
          position: "left",
        },
        { to: "blog", label: "Blog", position: "left" },
        {
          href: "https://github.com/facebook/docusaurus",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        {
          title: "More",
          items: [
            {
              label: "Blog",
              to: "blog",
            },
            {
              label: "GitHub",
              href: "https://github.com/aevea/commitsar",
            },
          ],
        },
      ],
      copyright: `Copyright © ${
        new Date().getFullYear()
      } My Project, Inc. Built with Docusaurus.`,
    },
  },
  presets: [
    [
      "@docusaurus/preset-classic",
      {
        docs: {
          // It is recommended to set document id as docs home page (`docs/` path).
          homePageId: "intro",
          sidebarPath: require.resolve("./sidebars.js"),
          // Please change this to your repo.
          editUrl: "https://github.com/aevea/commitsar/edit/master/www/docs",
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl: "https://github.com/aevea/commitsar/edit/master/www/blog/",
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      },
    ],
  ],
};
