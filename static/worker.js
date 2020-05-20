workerCode = function() {
  //  const workerread = getElementById('webworkersread');

    t = setInterval(function() {
        fetch('/api/workers')
            .then(response => response.body)
            .then(rs => {
                const reader = rs.getReader();

                var prms = reader.read();

                prms.then(function(value){
                    postMessage(value);
                }, function(e){
                    console.log(e);
                })
            })
            .catch(error => {
                console.log(error)
            })
    }, 2000)
}

workerCode();