import http from "k6/http";
import { check, sleep } from "k6";

// Test configuration
export const options = {
  stages: [{ duration: "1s", target: 1000 }],
};

// Simulated user behavior
export default function () {
  let res = http.get("http://localhost:8080");
  // Validate response status
  check(res, { "status was 200": (r) => r.status == 200 });
}
