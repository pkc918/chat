// uno.config.ts
import { defineConfig, presetUno } from 'unocss'
import presetIcons from '@unocss/preset-icons'

export default defineConfig({
    // ...UnoCSS options
    presets: [
        presetIcons({
            extraProperties: {
                "display": "inline-block",
                'vertical-align': 'middle'
            }
        }),
        presetUno()

    ],
    rules: [
        ["login-bg", {"background-image": 'url("https://w.wallhaven.cc/full/jx/wallhaven-jxl31y.png")', "background-repeat": "no-repeat"}]
    ]
})
