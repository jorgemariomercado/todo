import * as esbuild from 'esbuild'

let result = await esbuild.build({
    entryPoints: ["web/index.js"],
    outdir: "public/assets",
    allowOverwrite: true,
    minify: true,
    loader: {'.js': 'jsx'},
    sourcemap: false,
    jsx: 'automatic',
    bundle: true,
})

console.log(result)
