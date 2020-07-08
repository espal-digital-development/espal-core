let offlineSince;
const timeout = (promise, ms) => Promise.race([new Promise((_, reject) => setTimeout(reject, ms)), promise]);

window.addEventListener('load', () => {
    const toolbar = document.createElement('div');
    toolbar.style.position = 'absolute';
    toolbar.style.right = '0px';
    toolbar.style.bottom = '0px';
    toolbar.style.width = '150px';
    toolbar.style.height = '40px';
    toolbar.style.backgroundColor = '#333';
    toolbar.style.color = '#fff';
    toolbar.style.boxShadow = '-1px -1px 0.5 #282828';

    const liveRefresh = document.createElement('input');
    liveRefresh.type = 'checkbox';
    liveRefresh.checked = localStorage.getItem('liveRefresh') == 'true';
    liveRefresh.style.marginLeft = '25px';
    liveRefresh.addEventListener('change', (e) => localStorage.setItem('liveRefresh', e.target.checked));

    toolbar.appendChild(liveRefresh);
    document.body.appendChild(toolbar);

    const intervalID = setInterval(() => {
        if (!liveRefresh.checked) return;

        if (offlineSince && new Date().getTime() - offlineSince > 2500) {
            console.warn('Stopped the liveRefresh watcher. Server seems offline');
            clearInterval(intervalID);
            localStorage.setItem('liveRefresh', false);
            return;
        }

        timeout(fetch(location.origin + '/health', { method: 'head' }), 100)
            .then(({ ok }) => (ok && offlineSince ? location.reload() : ''))
            .catch(() => (offlineSince = offlineSince || new Date().getTime()));
    }, 500);
});
