let express = require('express')();
require('zazler')("sqlite3:///root/xssdata/data.db3", { read: "cross" } )
.then(sqlApi => {
  express.use('/api/', sqlApi.expressRequest);
  console.log('xssapi running on 8099')
  express.listen(8099);
});