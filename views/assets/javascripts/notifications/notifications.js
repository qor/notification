(function (factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as anonymous module.
    define(['jquery'], factory);
  } else if (typeof exports === 'object') {
    // Node / CommonJS
    factory(require('jquery'));
  } else {
    // Browser globals.
    factory(jQuery);
  }
})(function ($) {

  'use strict';

  var NAMESPACE = 'qor.notification';
  var EVENT_ENABLE = 'enable.' + NAMESPACE;
  var EVENT_DISABLE = 'disable.' + NAMESPACE;
  var EVENT_UNDO = 'undo.qor.action';
  var UNDO_TYPE = 'notification';

  function QorNotification(element, options) {
    this.$element = $(element);
    this.options = $.extend({}, QorNotification.DEFAULTS, $.isPlainObject(options) && options);
    this.init();
  }

  QorNotification.prototype = {
    constructor: QorNotification,

    init: function () {
      this.bind();
    },

    bind: function () {
      this.$element.on(EVENT_UNDO, $.proxy(this.undo, this));
    },

    undo: function (e, $actionButton, isUndo) {
      var data = $actionButton.data(),
          undoType = data.undoType,
          undoLabel = data.undoLabel,
          orignalLabel = data.label;

      if (undoType == UNDO_TYPE) {
        if (!isUndo) {
          $actionButton.html(undoLabel);
        } else {
          $actionButton.html(orignalLabel);
        }
      }
    },

    unbind: function () {
      this.$element.off(EVENT_UNDO, this.undo);
    },

    destroy: function () {
      this.unbind();
      this.$element.removeData(NAMESPACE);
    }
  };

  QorNotification.DEFAULTS = {};

  QorNotification.plugin = function (options) {
    return this.each(function () {
      var $this = $(this);
      var data = $this.data(NAMESPACE);
      var fn;

      if (!data) {
        if (/destroy/.test(options)) {
          return;
        }

        $this.data(NAMESPACE, (data = new QorNotification(this, options)));
      }

      if (typeof options === 'string' && $.isFunction(fn = data[options])) {
        fn.apply(data);
      }
    });
  };

  $(function () {
    var selector = '.qor-notifications';

    $(document).
      on(EVENT_DISABLE, function (e) {
        QorNotification.plugin.call($(selector, e.target), 'destroy');
      }).
      on(EVENT_ENABLE, function (e) {
        QorNotification.plugin.call($(selector, e.target));
      }).
      triggerHandler(EVENT_ENABLE);
  });

  return QorNotification;

});
