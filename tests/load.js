import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  vus: 100,
  duration: '30s',
  thresholds: {
    http_req_duration: ['p(95)<200'], // 95% of requests must complete below 200ms
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
  },
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

  sleep(1);
}
