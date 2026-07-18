import { createRequire } from 'node:module'
import { dirname, resolve } from 'node:path'

import { defineConfig, mergeConfig } from 'vite'

import baseConfig from './vite.config'

const require = createRequire(import.meta.url)

function packageFile(packageName: string, relativePath: string): string {
  return resolve(dirname(require.resolve(`${packageName}/package.json`)), relativePath)
}

export default defineConfig((env) => {
  const resolvedBaseConfig = typeof baseConfig === 'function' ? baseConfig(env) : baseConfig

  return mergeConfig(resolvedBaseConfig, {
    define: {
      'process.env.NODE_ENV': JSON.stringify('production'),
    },
    ssr: {
      noExternal: true,
    },
    resolve: {
      alias: [
        {
          find: '@vue/server-renderer',
          replacement: packageFile('@vue/server-renderer', 'dist/server-renderer.esm-bundler.js'),
        },
        {
          find: /^vue$/,
          replacement: packageFile('vue', 'dist/vue.runtime.esm-bundler.js'),
        },
      ],
    },
    build: {
      ssr: true,
      target: 'es2020',
      minify: 'oxc',
      rollupOptions: {
        input: {
          server: 'src/entry-server.ts',
        },
        output: {
          format: 'cjs',
          entryFileNames: '[name].js',
          codeSplitting: false,
        },
      },
    },
  })
})
