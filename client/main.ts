import { faker } from '@faker-js/faker';
import axios from 'axios';
import { Worker, isMainThread } from 'node:worker_threads'


async function main() {
  for (let i = 0; i < 1000; i++) {

    const someMath = {
      "name": `randon-${i}`,
      "apelido": `randon-${i}`,
      "nascimento": `randon-${i}`,
      "stack": `randon-${i}`
    };
    const response = await axios.post('http://localhost:9999/person', someMath)
    // if (isMainThread) {
    //   // const someMath = {
    //   //   "name": faker.internet.userName(),
    //   //   "apelido": faker.internet.userName(),
    //   //   "nascimento": faker.internet.protocol(),
    //   //   "stack": faker.internet.protocol()
    //   // };
    //   // new Worker('./worker.ts', {
    //   //   workerData: {
    //   //     value: someMath
    //   //   }
    //   // })

    // }
  }
}
main()