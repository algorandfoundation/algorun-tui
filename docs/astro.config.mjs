// @ts-check
import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";
import tailwind from "@astrojs/tailwind";
import mdx from '@astrojs/mdx';

// https://astro.build/config
export default defineConfig({
  integrations: [
    starlight({
      title: "NodeKit",
      logo: {
        light: "./public/nodekit-light.png",
        dark: "./public/nodekit-dark.png",
        alt: "NodeKit for Algorand",
        replacesTitle: true,
      },
      social: {
        github: "https://github.com/algorandfoundation/nodekit",
      },
      components: {
        ThemeProvider: "./src/components/CustomThemeProvider.astro",
      },
      customCss: ["./src/tailwind.css"],
    }),
    mdx(),
    tailwind({ applyBaseStyles: true }),
  ],
});
