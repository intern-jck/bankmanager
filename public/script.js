let currentStatement = {};

window.onload = () => {
  currentStatement = {};
};

let data = {};
function createGraphs(data) {
  // Create summary graph
  const summaryBeginning = data.Summary.Beginning;
  const summaryEnding = data.Summary.Ending;
  const summaryDeposits = data.Summary.Deposits;
  const summaryChecks = data.Summary.Checks;
  const summaryDebit = data.Summary.Debit;
  const summaryElectronic = data.Summary.Electronic;

  const summaryData = [
    parseFloat(summaryBeginning),
    parseFloat(summaryEnding),
    parseFloat(summaryDeposits),
    parseFloat(summaryChecks),
    parseFloat(summaryDebit),
    parseFloat(summaryElectronic),
  ];

  const summaryLabels = [
    "Beginning",
    "Ending",
    "Deposits",
    "Checks",
    "Debit",
    "Electronic",
  ];

  const summaryCtx = document.getElementById("summaryChart").getContext("2d");

  const sumGraphData = {
    labels: summaryLabels,
    datasets: [
      {
        label: "Summary 2018",
        data: summaryData,
        backgroundColor: "rgb(20, 200, 100)",
        borderColor: "rgb(0, 0, 0)",
        borderWidth: 2,
      },
    ],
  };

  const summaryChart = new Chart(summaryCtx, {
    type: "bar",
    data: sumGraphData,
    options: {},
  });

  // Create Withdrawals Graph
  const withdrawalAmounts = [];
  const withdrawalLabels = [];

  for (let w of data.Withdrawals) {
    withdrawalAmounts.push(parseFloat(w.Amount));
    withdrawalLabels.push(w.Date);
  }

  const withdrawalsCtx = document
    .getElementById("withdrawalsChart")
    .getContext("2d");

  const withGraphData = {
    labels: withdrawalLabels,
    datasets: [
      {
        label: "Withdrawals",
        data: withdrawalAmounts,
        backgroundColor: "rgb(200, 20, 100)",
        borderColor: "rgb(0, 0, 0)",
        borderWidth: 2,
      },
    ],
  };

  const withdrawalsChart = new Chart(withdrawalsCtx, {
    type: "line",
    data: withGraphData,
    options: {},
  });

  // Create Deposits Graph
  const depositAmounts = [];
  const depositLabels = [];

  for (let d of data.Deposits) {
    depositAmounts.push(parseFloat(d.Amount));
    depositLabels.push(d.Date);
  }

  const depositCtx = document.getElementById("depositsChart").getContext("2d");

  const depositsGraphData = {
    labels: depositLabels,
    datasets: [
      {
        label: "Deposits",
        data: depositAmounts,
        backgroundColor: "rgb(20, 200, 100)",
        borderColor: "rgb(0, 0, 0)",
        borderWidth: 2,
      },
    ],
  };

  const depositsChart = new Chart(depositCtx, {
    type: "line",
    data: depositsGraphData,
    options: {},
  });
}
