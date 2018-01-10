// onoff sends a get request to /on if type=on, to /off otherwise
function onoff(type) {
  console.log(type)
  var xhttp = new XMLHttpRequest();
  xhttp.open("GET", type);
  xhttp.send()
}
