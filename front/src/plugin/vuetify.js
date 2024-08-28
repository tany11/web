import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi'
import { ja } from 'vuetify/locale'

const vuetify = createVuetify({
    components,
    directives,
    icons: {
        defaultSet: "mdi",
        aliases,
        sets: {
            mdi,
        },
    },
    theme: {
        // defaultTheme: 'light', // デフォルトテーマをlightに設定
        themes: {
            light: {
                dark: false,
                // ライトテーマの色を設定
            },
            dark: {
                dark: true,
                // ダークテーマの色を設定
            },
        },
    },
    locale: {
        locale: 'ja',
        fallback: 'en',
        messages: { ja }
    },
});

export default vuetify