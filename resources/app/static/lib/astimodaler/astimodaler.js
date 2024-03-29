if (typeof asticode === 'undefined') {
  var asticode = {};
}
asticode.modaler = {
  close: function() {
    if (typeof asticode.modaler.onclose !== 'undefined' && asticode.modaler.onclose !== null) {
      asticode.modaler.onclose();
    }
    asticode.modaler.hide();
  },
  hide: function() {
    document.getElementById('astimodaler').style.display = 'none';
  },
  init: function() {
    document.body.innerHTML = `<div class="astimodaler" id="astimodaler">
      <div class="astimodaler-background"></div>
      <div class="astimodaler-table">
        <div class="astimodaler-wrapper">
          <div id="astimodaler-body">
            <i class="<i class="far fa-window-close astimodaler-close" onclick="asticode.modaler.close()"></i>
            <div id="astimodaler-content"></div>
          </div>
        </div>
      </div>
    </div>` + document.body.innerHTML;
  },
  setContent: function(content) {
    document.getElementById('astimodaler-content').innerHTML = '';
    document.getElementById('astimodaler-content').appendChild(content);
  },
  setWidth: function(width) {
    document.getElementById('astimodaler-body').style.width = width;
  },
  show: function() {
    document.getElementById('astimodaler').style.display = 'block';
  }
};