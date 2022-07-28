module.exports = { 
    main: function (event, context) {
      const utcDate1 = new Date();
      var n = 0
      while (n < 100) {
         console.log(utcDate1.toUTCString() + "test logging to console:", n);
        n++;
      }

      return "Hello World!";
    }
  } 