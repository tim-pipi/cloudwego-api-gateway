import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
	vus: 1,
	duration: '5s',
};

export default function () {
	const url = 'http://127.0.0.1:8080/HelloService/echo';
	const message = 'Hello, World!';
	const payload = JSON.stringify({
		message: message,
	});

	const params = {
		headers: {
			'Content-Type': 'application/json',
		},
	};

	const res = http.request('GET', url, payload, params);
	check(res, {
		'Echo status is 200': (r) => r.status === 200,
		'Echo Content-Type includes application/json': (r) =>
			res.headers['Content-Type'].includes('application/json'),
		'Echo message is correct': (r) =>
			r.status === 200 && r.body.includes(message),
	});

	const test = JSON.parse(res.body)
	console.log(test)
	console.log(typeof(test))

	sleep(1);
}
