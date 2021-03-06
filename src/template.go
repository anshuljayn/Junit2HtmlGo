package src

const ReportTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1" name="viewport">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.3/Chart.min.js"></script>
    <style>
        * {
            box-sizing: border-box;
        }

        body {

            font-family: 'Source Sans Pro';
            margin: 0;
        }

        table {
            border: 1px solid lightgray;
            margin-top: 1em;
        }

        div {
            display: block;
        }

        /*Header classes*/
        .header {
            overflow: hidden;
            height: 45px;
            padding: 5px;
            /*text-align: center;*/
            vertical-align: middle;
            background: #9fe2bf;
            color: white;
        }

        .header span {
            vertical-align: middle;
            /*display: block;*/
        }

        .header span.right {
            float: right;
            vertical-align: bottom;
        }

        /* Increase the font size of the h1 element */
        .header h1 {
            font-size: 40px;
        }

        .brand-logo {
            left: 35px;
            height: 45px;
            /*background-color: yellow;*/
        }

        .brand-logo > img {
            width: 35px;
            /*margin-top: 10px;*/
            /*margin-left: 10px;*/
        }

        img {
            border-radius: 5px 5px 0 0;
        }

        .right {
            float: right !important
        }

        .report-name {
            margin-left: 90px;
            padding-top: 1px;
            padding-bottom: 1px;
            font-size: 20px;
            text-align: center;
            vertical-align: middle;
            color: #333333;
        }

        /*Navigation bar*/
        .navbar {
            overflow: hidden;
            background-color: #333;
        }

        .navbar a {
            float: left;
            display: block;
            color: white;
            text-align: center;
            padding: 14px 20px;
            text-decoration: none;
        }

        /* Right-aligned link */
        .navbar a.right {
            float: right;
        }

        /* Change color on hover */
        .navbar a:hover {
            background-color: #ddd;
            color: black;
        }

        .material-icons {
            font-family: 'Material Icons';
            font-weight: normal;
            font-style: normal;
            font-size: 24px;
            line-height: 1;
            letter-spacing: normal;
            text-transform: none;
            display: inline-block;
            white-space: nowrap;
            word-wrap: normal;
            direction: ltr;
            -webkit-font-feature-settings: 'liga';
            -webkit-font-smoothing: antialiased;
        }

        a.activeTab{
            float: right;
            padding-right: 1em;
        }

        a span {
            padding: 3px 6px;
            border-radius: 4px;
            background-color: #1565C0;
            color: white;
            font-size: 20px;
        }
        /*Report Content*/
        .report {
            padding: 5px;
        }

        .test-name{
            margin-top: 0;
            color: white;
            background-color: #333333;
            font-size: 24px;
        }

        /*Detail view*/
        .main {
            display: flex;
            flex-wrap: wrap;
        }

        .split {
            height: 100%;
            position: fixed;
            z-index: 1;
            /*top: 0;*/
            overflow-x: hidden;
            padding-top: 20px;
        }

        .leftSplit {
            left: 0;
            width: 25%;
            /*background-color: #111;*/
        }

        .rightSplit {
            right: 0;
            width: 75%;
            /*background-color: red;*/
        }

        .suite_collection {
            margin-right: 10px;
            margin-left: 10px;
            /*padding: 0px 10px 0px 10px;*/
        }

        .view-summary{
            margin-right: 10px;
            margin-left: 10px;
        }

        .test-heading {
            font-weight: 600;
            border: 1px solid lightgray;
        }

        .test-heading.active {
            background-color: lightgray;
        }

        .heading {
            /*word-wrap: break-word;*/
            word-break: break-all;
            padding: 5px 5px 15px 5px;
            font-size: 18px;

        }

        .passed {
            color: #32cd32;
            font-size: 15px;
        }

        .skipped {
            color: grey;
            font-size: 15px;
        }

        .failed {
            color: tomato;
            font-size: 15px;
        }

        .sub-heading {
            padding: 5px 5px 5px 5px;
            font-size: 12px;
        }

        .duration {

        }

        .test-content {
            padding-bottom: 100px;
        }

        .hide {
            display: none !important
        }

        .testcase {
            padding-bottom: 15px;
            width:100%
        }

        td.fitwidth {
            width: 1px;
            white-space: nowrap;
        }

        .headColumn {
            background-color: black;
            color: whitesmoke;
        }

        .testcaseRow {
            padding: 3px 0px 3px 0px;
        }

        /*summary view*/
        .tabs {
            padding: 0px;
            margin: 10px
        }

        .suite-table {
            width: 100%;
            /*margin-bottom: 1em;*/
        }

        .suite {
            background-color: #000000;
            color: whitesmoke;
            size: 30px;
        }

        /*error-view*/

        td.fitwidth, th.fitwidth {
            width: 1px;
            white-space: nowrap;
            text-align: left;
        }

        .error {
            font-family: "Courier New",serif;
            color: red;
        }

        /*stats-view*/
        .row {
            margin-left: 0.75rem;
            margin-right: 0.75rem;
            margin-top: 0.75rem;
            padding-top: 10px
        }

        .col {
            float: left;
            box-sizing: border-box;
            padding: 0 0.75rem;
            min-height: 1px
        }

        .col.split5 {
            width: 20%;
            margin-left: auto;
            left: auto;
            right: auto
        }

        .card {
            box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
            transition: 0.3s;
            border-radius: 5px;
        }

        .card:hover {
            box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
        }

        .container {
            padding: 2px 16px;
        }

        /* Responsive layout - when the screen is less than 700px wide, make the two columns stack on top of each other instead of next to each other */
        @media screen and (max-width: 700px) {
            .main {
                flex-direction: column;
            }
        }

        /* Responsive layout - when the screen is less than 400px wide, make the navigation links stack on top of each other instead of next to each other */
        @media screen and (max-width: 400px) {
            .navbar a {
                float: none;
                width: 100%;
            }
        }

        /* Responsive layout - when the screen is less than 700px wide, make the two columns stack on top of each other instead of next to each other */
        @media screen and (max-width: 700px) {
            .main {
                flex-direction: column;
            }
        }

        /* Responsive layout - when the screen is less than 400px wide, make the navigation links stack on top of each other instead of next to each other */
        @media screen and (max-width: 400px) {
            .navbar a {
                float: none;
                width: 100%;
            }
        }

        @font-face {
            font-family: 'Material Icons';
            font-style: normal;
            font-weight: 300;
            src: local('Material Icons'), local('MaterialIcons-Regular'), url(https://fonts.gstatic.com/s/materialicons/v18/2fcrYFNaTjcS6g4U3t-Y5ZjZjT5FdEJ140U2DJYC3mY.woff2) format('woff2');
        }
    </style>
    <title>Test Execution Report</title>
</head>
<body>
    <div class="header">
        <a class="brand-logo" href="#!"><img
                src="{{.BrandLogo}}"></a>
        <span class="report-name">{{.ReportName}}</span>
        <span class='right'>Jul 19, 2021 09:01:19 AM</span>
    </div>
    <div class="navbar">
        <a href='#!' onclick="displayView('detail-view')"><i class='material-icons'
                                                             style="color: yellow">loupe</i></a>
        <a href='#!' onclick="displayView('summary-view')"><i class='material-icons'
                                                              style="color:white;">view_list</i></a>
        <a href='#!' onclick="displayView('errors-view')"><i class='material-icons'
                                                             style="color:red;">error_outline</i></a>
        <a href='#!' onclick="displayView('stats-view')"><i class='material-icons'
                                                            style="color:deepskyblue">dashboard</i></a>
        <a href='#!' class="activeTab"><span id="blink">detail-view</span></a>
    </div>
    <div class="report">
        <div class="main" id='detail-view'>
            <div class="split leftSplit" id='suiteCollection'>
                {{.SuiteCollections}}
            </div>
            <div class="split rightSplit">
                <div class='view-summary' id='testBlock'>
                    <h3 class='test-name'>{{.SuiteName}}</h3>
                    <div class='test-content'>{{.SuiteContent}}</div>
                </div>
            </div>
        </div>
        <div class='tabs hide' id='summary-view'>
            <table width="100%">
                <tr>
                    <td width="50%">
                        <canvas id="myChart1" style="width:100%;max-width:600px"></canvas>
                    </td>
                    <td width="50%">
                        <canvas id="myChart2" style="width:100%;max-width:600px"></canvas>
                    </td>
                </tr>
            </table>
            {{.SummaryViewSuiteTables}}
        </div>
        <div class='tabs hide' id='errors-view'>
            {{.ErrorViewSuiteTables}}
        </div>
        <div class='hide' id='stats-view'>
            <div class="row">
                <div class="col split5">
                    <div class="card">
                        <div card-id=1 class="container">
                            <h4><b>Test Suite</b></h4>
                            <p>{{.NoOfTestSuites}}</p>
                        </div>
                    </div>
                </div>
                <div class="col split5">
                    <div class="card">
                        <div card-id=2 class="container">
                            <h4><b>Test Cases</b></h4>
                            <p>{{.NoOfTestCases}}</p>
                        </div>
                    </div>
                </div>
                <div class="col split5">
                    <div class="card">
                        <div card-id=3 class="container">
                            <h4><b>Start</b></h4>
                            <p>{{.StartTime}}</p>
                        </div>
                    </div>
                </div>
                <div class="col split5">
                    <div class="card">
                        <div card-id=4 class="container">
                            <h4><b>End</b></h4>
                            <p>{{.EndTime}}</p>
                        </div>
                    </div>
                </div>
                <div class="col split5">
                    <div class="card">
                        <div card-id=5 class="container">
                            <h4><b>Time Taken</b></h4>
                            <p>{{.Duration}}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <script>
            $('body').on('click', 'div.suite_collection', function () {
                $('#suiteCollection').children('div').each(function (index) {
                    $(this).children('div.test-heading').removeClass('active')
                })
                $(this).children('div.test-heading').addClass('active');
                $('#testBlock').children("h3").text($(this).children('div.test-heading').children('div.heading').text());
                $('#testBlock').find('div.test-content').html($(this).find('div.test-content').html());
            });
    </script>
    <script>
        function displayView(v){
            document.getElementById('blink').innerText = v
            $('#detail-view').addClass('hide');
            $('#summary-view').addClass('hide');
            $('#errors-view').addClass('hide');
            $('#stats-view').addClass('hide');
            $('#'+v).removeClass('hide');
            window.scrollTo(0,0);
        }
    </script>
    <script type="text/javascript">
        var blink = document.getElementById('blink');
        setInterval(function() {
            blink.style.opacity = (blink.style.opacity == .5 ? 1 : .5);
        }, 1500);
    </script>
    <script>
            var xValues = ["Pass", "Fail", "Skip"];
            var ytsValues = [{{.NoOfTSPass}}, {{.NoOfTSFail}}, {{.NoOfTSSkip}}]
            var ytcValues = [{{.NoOfTestCasesPass}}, {{.NoOfTestCasesFail}}, {{.NoOfTestCasesSkip}}]
            var barColors = [
                "green",
                "tomato",
                "#97898d"
            ];

            new Chart("myChart1", {
                type: "doughnut",
                data: {
                    labels: xValues,
                    datasets: [{
                        backgroundColor: barColors,
                        data: ytsValues
                    }]
                },
                options: {
                    title: {
                        display: true,
                        text: "Test Suite"
                    }
                }
            });
            new Chart("myChart2", {
                type: "doughnut",
                data: {
                    labels: xValues,
                    datasets: [{
                        backgroundColor: barColors,
                        data: ytcValues
                    }]
                },
                options: {
                    title: {
                        display: true,
                        text: "Test Cases"
                    }
                }
            });
        </script>
</body>
</html>
`
