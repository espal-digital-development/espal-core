import { h, ref } from 'vue';
import { authService } from './services';

export default {
    setup() {
        const jwtValidString = ref('');

        const horizontalRule = h('hr');

        return () => [
            h('img', { src: require('../images/logo.png').default }),
            h('h1', ['Hello!']),
            horizontalRule,
            h('p', [authService.jwt]),
            h(
                'button',
                {
                    onclick: () =>
                        authService.login({
                            email: 'no@one.com',
                            password: 'haha',
                            rememberMe: false
                        })
                },
                ['Get JWT']
            ),
            horizontalRule,
            h('p', [jwtValidString.value]),
            h(
                'button',
                {
                    onclick: async () => (jwtValidString.value = await authService.checkJWT())
                },
                ['Check JWT']
            )
        ];
    }
};