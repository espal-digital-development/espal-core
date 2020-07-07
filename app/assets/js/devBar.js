var wentOffline = false;
var offlineSince;

window.onload = () => {
    let toolbar = document.createElement('div');
    toolbar.style.position = 'absolute';
    toolbar.style.right = '0px';
    toolbar.style.bottom = '0px';
    toolbar.style.width = '150px';
    toolbar.style.height = '40px';
    toolbar.style.backgroundColor = '#333333';
    toolbar.style.color = '#ffffff';
    toolbar.style.boxShadow = '-1px -1px 0.5 #282828';

    let liveRefresh = document.createElement('input');
    liveRefresh.type = 'checkbox';
    liveRefresh.checked = localStorage.getItem('liveRefresh') == 'true';
    liveRefresh.style.marginLeft = '25px';
    liveRefresh.addEventListener('change', (e) => {
        localStorage.setItem('liveRefresh', e.target.checked);
    });

    toolbar.appendChild(liveRefresh);
    document.body.appendChild(toolbar);

    setInterval(() => {
        if (!liveRefresh.checked) {
            return;
        }
        if (offlineSince && new Date().getTime() - offlineSince > 2500) {
            console.log('Stopped the liveRefresh watcher. Server seems offline');
            liveRefresh.checked = false;
            wentOffline = false;
            offlineSince = null;
            localStorage.setItem('liveRefresh', false);
            return;
        }
        fetch(location.origin + '/health', { method: 'head' })
            .then((status) => {
                if (status.ok) {
                    if (wentOffline) {
                        // Refresh when the server is back online again
                        window.location.reload(true);
                        wentOffline = false;
                    }
                    offlineSince = null;
                }
            })
            .catch(() => {
                wentOffline = true;
                if (!offlineSince) {
                    offlineSince = new Date().getTime();
                }
            });
    }, 500);
};
