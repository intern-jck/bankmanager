window.onload = () => {
  console.log("JavaScript loaded!");
};

// let socket = null;

// document.getElementById("cmd-send-btn").onclick = function (event) {
//   if (!socket) {
//     return false;
//   }

//   const cmdInput = document.getElementById("cmd-input").value;
//   socket.send(cmdInput);
// };

// document.getElementById("ws-close-btn").onclick = function (evt) {
//   if (!socket) {
//     console.log("NO SOCKET");
//     return false;
//   }

//   console.log("CLOSING SOCKET");
//   socket.close();
//   socket = null;
// };

// document.getElementById("ws-open-btn").onclick = function (evt) {
//   if (socket) {
//     console.log("OPEN SOCKET");
//     return false;
//   }

//   console.log("CREATING SOCKET");
//   // if in a script file, only need endpoint
//   socket = new WebSocket("/esp");
//   return true;
//   // socket.onopen = function (evt) {
//   //   console.log("OPEN");
//   // };

//   // ws.onclose = function (evt) {
//   //   console.log("CLOSE");
//   //   ws = null;
//   // };

//   // ws.onmessage = function (evt) {
//   //   console.log("RESPONSE: " + evt.data);
//   // };

//   // ws.onerror = function (evt) {
//   //   console.log("ERROR: " + evt.data);
//   // };

//   // return false;
// };

// let row_count = 0;

// function addDataToTable(data) {
//   const table = document.getElementById("esp-data-body");
//   const row = table.insertRow(-1);

//   const cell_1 = row.insertCell(0);
//   const cell_2 = row.insertCell(1);
//   const cell_3 = row.insertCell(2);

//   cell_1.textContent = data.Status;
//   cell_2.textContent = data.Cmd ? data.Cmd : "None";
//   cell_3.textContent = data.Val.toString();

//   row_count += 1;
// }

// window.onload = (event) => {
//   console.log("Page Load");
//   socket = new WebSocket("/esp");

//   socket.addEventListener("open", (event) => {
//     socket.send("socket open");
//   });

//   socket.addEventListener("message", (event) => {
//     console.log("ESPSRV: ", event.data);
//     const espData = JSON.parse(event.data);
//     addDataToTable(espData);
//   });
// };
