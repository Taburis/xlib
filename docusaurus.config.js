const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');
const math = require('remark-math');
const katex = require('rehype-katex');

/** @type {import('@docusaurus/types').DocusaurusConfig} */
module.exports = {
  title: 'XLibrary',
  tagline:  '',

  url: 'https://taburis.github.io',
  baseUrl: '/xlib/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/Ton618_icon.png',
  //favicon: 'img/favicon.ico',
  organizationName: 'Taburis', // Usually your GitHub org/user name.
  projectName: 'xlib', // Usually your repo name.
  trailingSlash: true,
  stylesheets: [
    {
        href: "https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/katex.min.css",
        integrity: "sha384-Um5gpz1odJg5Z4HAmzPtgZKdTBHZdw8S29IecapCSB31ligYPhHQZMIlWLYQGVoc",
        crossorigin: "anonymous",
    },
  ],
/*
  stylesheets: [
    {
      href: 'https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.css',
      type: 'text/css',
      integrity:
        'sha384-AfEj0r4/OFrOo5t7NnNe46zW/tFgW6x/bCJG8FqQCEo3+Aro6EYUG4+cU+KJWu/X',
      crossorigin: 'anonymous',
    },
  ],
  */
  themeConfig: {
    navbar: {
      title: 'XLibrary',
      logo: {
        alt: 'My Site Logo',
        src: 'img/Ton618.png',
        //src: 'img/logo.svg',
      },
      items: [
        {
          type: 'doc',
          docId: 'intro',
          position: 'left',
          label: 'Notes',
          //label: 'Tutorial',
        },
        {
          type: 'doc',
          position: 'left',
          docId: 'Projects/index',
          label: 'Projects',
          //label: 'Tutorial',
        },
        {to: '/blog', label: 'Blog', position: 'left'},
        {
          //href: 'https://github.com/facebook/docusaurus',
          //label: 'GitHub',
          //position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
		  /*
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Notes',
              to: '/docs/intro',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Stack Overflow',
              href: 'https://stackoverflow.com/questions/tagged/docusaurus',
            },
            {
              label: 'Discord',
              href: 'https://discordapp.com/invite/docusaurus',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/docusaurus',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Blog',
              to: '/blog',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/facebook/docusaurus',
            },
          ],
        },
      ],
	*/
      copyright: `Copyright © ${new Date().getFullYear()} Xiao W. Built with Docusaurus.`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
		  routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl:
            'https://github.com/Taburis/xlib/edit/documentation/',
            //'https://github.com/facebook/docusaurus/edit/master/website/',
          remarkPlugins: [math],
		  rehypePlugins: [katex],
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github.com/Taburis/xlib/edit/documentation/blog/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  
};
