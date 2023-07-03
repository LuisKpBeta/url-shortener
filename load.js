import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 1000,
  duration: '1s',
};
// export const options = {
//   stages: [
//     { duration: '2s', target: 1 },
//   ],
// };

export function insert() {
  const url = 'http://localhost:8080/'
  const payload = JSON.stringify({
    url: 'www.google.com',
  });
  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };


  const res = http.post(url, payload, params);
  console.log(res.body)
  check(res, { 'status was 200': (r) => r.status == 200 });

  sleep(1);
}

export default function laod() {
  const url = 'http://localhost:8080/dU9xE'
  const res = http.get(url);
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}