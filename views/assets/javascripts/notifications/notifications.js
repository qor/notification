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
  var UNDO_CONTAINER = '.qor-notifications__item';
  var BUTTON_UNDO = '.qor-notifications__item-undo';

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

    undo: function (e, $actionButton, isUndo, data) {
      var actionData = $actionButton.data(),
          undoType = actionData.undoType,
          $item = $actionButton.closest(UNDO_CONTAINER).length ? $actionButton.closest(UNDO_CONTAINER) : $actionButton.closest(BUTTON_UNDO).prev(UNDO_CONTAINER),
          $template,
          template = QorNotification.UNDO_HTML;

      data.undoLabel = actionData.undoLabel,
      $template = $(window.Mustache.render(template, data));
      !isUndo && $template.find('button').data(actionData);

      if (undoType == UNDO_TYPE) {
        $item.before(data.notification);
        if (!isUndo) {
          $item.after($template);
        } else {
          $item.next(BUTTON_UNDO).remove();
        }
        $item.remove();
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
  QorNotification.UNDO_HTML = '<div class="qor-notifications__item-undo"><span>[[message]]</span><button class="mdl-button mdl-button--colored qor-action-button is_undo" type="button" data-undo-type="notification">[[undoLabel]]</button></div>';

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
