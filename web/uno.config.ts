// uno.config.ts
import { defineConfig, presetUno } from 'unocss'
import presetIcons from '@unocss/preset-icons'

export default defineConfig({
    // ...UnoCSS options
    presets: [
        presetIcons({
            extraProperties: {
                "display": "inline-block",
                'vertical-align': 'middle',
            }
        }),
        presetUno(),
    ]
})
