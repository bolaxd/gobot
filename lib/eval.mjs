(async function helo() {
    try {
	const evals = await eval(`;(async () => {${process.argv[2]}})()`)
	console.log(evals);
    } catch (e) {
	console.log(e);
    }
    process.exit()
})();


setInterval(() => {
    if (process.uptime() > 300) {
	process.exit();
	console.log("waktu sessi ngeval habis");
    }}, 1000);