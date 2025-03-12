let summaryChart;

const createBarGraph = (id, data, title, labels, options) => {
    const barCtx = document.getElementById(id).getContext("2d");

    const barGraphData = {
        labels: labels,
        datasets: [
            {
                label: title,
                data: data,
                backgroundColor: "rgb(20, 200, 100)",
                borderColor: "rgb(0, 0, 0)",
                borderWidth: 2,
            },
        ],
    };


    if (summaryChart) {
        summaryChart.destroy()
    };

    summaryChart = new Chart(barCtx, {
        type: "bar",
        data: barGraphData,
        options: options
    });
};

function createLineGraph(id, data, title, labels, options) {
    const lineCtx = document.getElementById(id).getContext("2d");

    const lineGraphData = {
        labels: labels,
        datasets: [
            {
                label: title,
                data: data,
                backgroundColor: "rgb(200, 20, 100)",
                borderColor: "rgb(0, 0, 0)",
                borderWidth: 2,
            },
        ],
    };

    const withdrawalsChart = new Chart(lineCtx, {
        type: "line",
        data: lineGraphData,
        options: options,
    });
}
