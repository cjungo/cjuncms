import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";
import Icons from "unplugin-icons/vite";
import IconsResolver from "unplugin-icons/resolver";
import { resolve } from "path";

// https://vitejs.dev/config/
export default ({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "VITE_");
  const target = env["VITE_SERVER_URL"];
  
  return defineConfig({
    base: "/",
    build: {
      emptyOutDir: true,
      outDir: resolve(__dirname, "..", "server", "public"),
    },
    server: {
      proxy: {
        "/proxy/": {
          target: target,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/proxy/, ""),
        },
      },
    },
    plugins: [
      vue(),
      AutoImport({
        resolvers: [
          ElementPlusResolver(),
          // 自动导入图标组件
          IconsResolver({
            prefix: "Icon",
          }),
        ],
      }),
      Components({
        dirs: ["src/dialogs", "src/components", "src/layouts"],
        resolvers: [
          // 自动注册图标组件
          IconsResolver({
            // 都要加 IEp 开头才能自动加载。
            enabledCollections: ["ep"],
          }),
          ElementPlusResolver(),
        ],
      }),
      Icons({
        autoInstall: true,
      }),
    ],
  });
};
