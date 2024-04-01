  function getCookie(name) {
      name += "=";
      var cookies = decodeURIComponent(document.cookie).split(";");
      for (var i = 0, len = cookies.length; i < len; i++) {
        var cookie = $.trim(cookies[i]);
        if (cookie.indexOf(name) === 0) {
          return cookie.substr(name.length);
        }
      }

      return "";
    }
