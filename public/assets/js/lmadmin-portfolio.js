$(function () {
    $("#img-uploads").fileinput({
        uploadUrl: "/lmadmin/upload_images",
        minImageWidth: 250,
        minImageHeight: 250,
        maxImageWidth: 5000,
        maxImageHeight: 5000,
        initialPreview: [
            //"<img src='/web/images/2/2/thumbnail.png' class='file-preview-image' >"
        ]
    })
    .on('fileuploaded', function(event, data, previewId, index) {
        $('.form-horizontal').append('<input type="hidden" name="images[]" value="' + data.response.image_id + '" />');
    });

});