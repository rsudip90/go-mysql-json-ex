// call api when dom is ready
$(function() {
    InitUnApprovedApplicantsTable();
});

// ---------------------------------------------------
// "Create Item" BUTTON CLICK HANDLER
// ---------------------------------------------------
$(document).on("click", "#sidebar a#unapproved", function(event) {
    InitUnApprovedApplicantsTable();
});

// ---------------------------------------------------
// Initalize table entries with json docs
// ---------------------------------------------------
window.InitUnApprovedApplicantsTable = function() {
    // get list and render in DOM
    GetJSONDocList()
    .then(function(jsonData) {
        // loop over json docs
        jsonData.map(function(jsonDoc) {
            // append json doc in html
            RenderJSONDocInDOM(jsonDoc);
        });
    })
};

// ---------------------------------------------------
// "Create Item" BUTTON CLICK HANDLER
// ---------------------------------------------------
$(document).on("click", "button#create_json", function(event) {
    // fill the inputs
    $("#modal_form form input").val("");

    // new DocID -- value 0
    $("#modal_form form input#DocID").val(0);

    // show the modal form
    $("#modal_form").modal("show");
});

// ---------------------------------------------------
// EDIT BUTTON CLICK HANDLER
// ---------------------------------------------------
$(document).on("click", "#applicants_queue tbody tr td button.edit", function(event) {
    var DocID = parseInt($(this).attr("data-id")),
        tr    = $("tr[data-id=" + DocID + "]");

    // get the values to fill the inputs
    var name        = tr.find("td[class=name]").text(),
        email       = tr.find("td[class=email]").text(),
        cellphone   = tr.find("td[class=cellphone]").text(),
        address     = tr.find("td[class=address]").text();

    // fill the inputs
    $("#modal_form form").find("input#DocID").val(DocID);
    $("#modal_form form").find("input#Name").val(name);
    $("#modal_form form").find("input#Email").val(email);
    $("#modal_form form").find("input#CellPhone").val(cellphone);
    $("#modal_form form").find("input#Address").val(address);

    // show the modal form
    $("#modal_form").modal("show");
});

// ---------------------------------------------------
// SAVE BUTTON CLICK HANDLER
// ---------------------------------------------------
$(document).on("click", "#form_save", function(event) {
    event.preventDefault();

    // button
    var saveBtn = event.target;

    // disable the button
    saveBtn.disabled = true;

    // get inputs from the form
    var name        = $("#modal_form form").find("input#Name").val(),
        email       = $("#modal_form form").find("input#Email").val(),
        cellphone   = $("#modal_form form").find("input#CellPhone").val(),
        address     = $("#modal_form form").find("input#Address").val();

    // DocID is number
    var DocID = parseInt($("#modal_form form").find("input#DocID").val());

    // data
    var jsonDoc = {
        "Name": name,
        "Email": email,
        "CellPhone": cellphone,
        "Address": address,
    };

    // create data
    SaveJSONDoc(jsonDoc, DocID)
    .then(function(newJSONDoc) {

        // append json doc
        RenderJSONDocInDOM(newJSONDoc);

        // enable the button
        saveBtn.disabled = false;

        // clear inputs in form
        $("#modal_form form input").val("");

        // hide the modal form
        $("#modal_form").modal("hide");
    });
});

// ---------------------------------------------------
// Remove Item HANDLER
// ---------------------------------------------------
$(document).on("click", "#applicants_queue tbody tr td button.remove", function(event) {
    event.preventDefault();

    // disable button
    $(this)[0].disabled = true;

    // get DocID
    var DocID = parseInt($(this).attr("data-id"));

    RemoveJSONDoc(DocID)
    .then(function(data) {
        $("#applicants_queue tbody").find("tr[data-id=" + data.DocID + "]").remove();
    });
});

// ---------------------------------------------------
// RenderJSONDocInDOM - append the jsondoc data
//                        in table of jsondocs
// ---------------------------------------------------
window.RenderJSONDocInDOM = function(jsonDoc) {
    // if jsonDoc available in dom then just update the content
    var targetTR = $("tr[data-id=" + jsonDoc.DocID + "]");
    if (targetTR.length > 0) { // update DOM with data
        targetTR.find("td[class=name]").text(jsonDoc.Data.Name);
        targetTR.find("td[class=email]").text(jsonDoc.Data.Email);
        targetTR.find("td[class=cellphone]").text(jsonDoc.Data.CellPhone);
        targetTR.find("td[class=address]").text(jsonDoc.Data.Address);
    } else {
        var tr = `<tr data-id="` + jsonDoc.DocID  + `">`;
        tr += `<th scope="row">`+jsonDoc.DocID+`</th>`;
        tr += `<td class="name">`+jsonDoc.Data.Name+`</td>`;
        tr += `<td class="email">`+jsonDoc.Data.Email+`</td>`;
        tr += `<td class="cellphone">`+jsonDoc.Data.CellPhone+`</td>`;
        tr += `<td class="address">`+jsonDoc.Data.Address+`</td>`;
        tr += `<td><button class="btn btn-info edit" data-id="` + jsonDoc.DocID + `">Edit</button></td>`;
        tr += `<td><button class="btn btn-danger remove" data-id="` + jsonDoc.DocID + `">X</button></td>`;
        tr += "</tr>";
        $("#applicants_queue tbody").append(tr);
    }
};
