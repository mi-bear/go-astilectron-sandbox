let index = {
  about: function(html) {
    let c = document.createElement('div');
    c.innerHTML = html;
    asticode.modaler.setContent(c);
    asticode.modaler.show();
  },
  init: function() {
    // Init
    asticode.loader.init();
    asticode.modaler.init();
    asticode.notifier.init();

    // Wait for astilectron to be ready
    document.addEventListener('astilectron-ready', function() {
      // Listen
      index.listen();

      // Kuma
      index.kuma();
    })
  },
  kuma: function() {
    let message = {'name': 'kuma'};
    message.payload = "アメリカクロクマ"

    // Send message
    asticode.loader.show();
    astilectron.sendMessage(message, function(message) {
      // Init
      asticode.loader.hide();
      // Check error
      if (message.name === 'error') {
        asticode.notifier.error(message.payload);
        return
      }
      // View
      document.getElementById('name').innerHTML = message.payload.name;
    })
  },
  listen: function() {
    astilectron.onMessage(function(message) {
      // bootstrap.SendMessage の2つ目の引数で設定
      switch (message.name) {
        case 'about':
          index.about(message.payload);
          return {payload: payload};
          break;
        case 'check.out.menu':
          asticode.notifier.info(message.payload);
          break
        }
    });
  }
};
