// ---------------------------------------------------
// GetJSONDocList - get the list of jsondoc using
//                  fetch API and append items in
//                  table html
// ---------------------------------------------------
window.GetJSONDocList = function() {
    return fetch("/jsondocs/")
    .then(function(response) {
        return response.json();
    })
    .catch(function(err) {
        console.error(err);
    });
};

// ---------------------------------------------------
// SaveJSONDoc - create json doc using fetch API
// ---------------------------------------------------
window.SaveJSONDoc = function(jsonDoc, DocID) {
    return fetch("/jsondocs/" + DocID.toString() + "/", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(jsonDoc),
    })
    .then(function(response) {
        return response.json();
    })
    .catch(function(err) {
        console.error(err);
    });
};

// ---------------------------------------------------
// RemoveJSONDoc - remove json doc using fetch API
// ---------------------------------------------------
window.RemoveJSONDoc = function(DocID) {
    return fetch("/jsondocs/" + DocID.toString() + "/", {
        method: "DELETE",
    })
    .then(function(response) {
        return response.json();
    })
    .catch(function(err) {
        console.error(err);
    });
};
