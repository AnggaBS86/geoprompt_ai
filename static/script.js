let map;
let markers = [];

function initMap() {
    map = new google.maps.Map(document.getElementById('map'), {
        zoom: 12,
        center: { lat: -7.47, lng: 110.22 }, // Default to Magelang
    });
}

$(document).ready(function () {
    $('#send').click(function () {
        const query = $('#prompt').val();
        if (query.trim() == "") {
            alert("Input cannot be empty!");
            $('#prompt').focus();
            return;
        }

        $('#loading').show();
        $('#result').text('');
        $('#map-links').empty();
        $("#prompt").attr('disabled', true);
        $('#send').attr('disabled', true);
        clearMarkers();

        $.ajax({
            url: '/ask',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ query }),
            success: function (data) {
                $('#send').attr('disabled', false);
                $("#prompt").attr('disabled', false);
                $('#loading').hide();
                $('#result').text(JSON.stringify(data, null, 2));
                updateMap(data);
                updateLinks(data);
            },
            error: function (xhr) {
                $('#loading').hide();
                $('#result').text("❌ Error: " + xhr.responseText);
            }
        });
    });
});

function updateMap(coords) {
    if (!coords.length) return;

    const bounds = new google.maps.LatLngBounds();

    coords.forEach(coord => {
        const position = { lat: coord.latitude, lng: coord.longitude };
        const marker = new google.maps.Marker({ map, position });
        markers.push(marker);
        bounds.extend(position);
    });

    map.fitBounds(bounds);
}

function clearMarkers() {
    markers.forEach(marker => marker.setMap(null));
    markers = [];
}

function updateLinks(coords) {
    const $list = $('#map-links');
    coords.forEach((coord, index) => {
        const url = `https://www.google.com/maps/search/?api=1&query=${coord.latitude},${coord.longitude}`;
        const li = `<li><a href="${url}" target="_blank">View Location ${index + 1} on Google Maps</a></li>`;
        $list.append(li);
    });
}
