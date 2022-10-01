const baseURL = 'https://1cca-1-46-130-241.ap.ngrok.io/course';
fetch(baseURL)
  .then(resp => resp.json())
  .then(data => {
    displayData(data)
    appendData(data)
  }).catch(err => console.log('error : ' + err));

function displayData(data) {
  document.querySelector("pre").innerHTML = JSON.stringify(data, null, 2);
}

function appendData(data) {
  var mainContainer = document.getElementById("myData")
  for (let i = 0; i < data.length; i++) {
    var div = document.createElement("div")
    div.innerHTML = 'CourseID: ' + data[i].id + ' ' + data[i].name + ' ' + data[i].price + ' ' + data[i].instructor
    mainContainer.appendChild(div)
  }
}