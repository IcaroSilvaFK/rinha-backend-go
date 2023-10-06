import {
  parentPort,
  workerData
} from "node:worker_threads";
import axios from 'axios'

const data = workerData.value
let quantityRequest = 0

async function sendRequest() {
  const response = await axios.post('http://localhost:9999/person', data)

  console.log(response.data)
  quantityRequest += 1
}

sendRequest()
console.log(quantityRequest)
// parentPort?.postMessage(someMath);