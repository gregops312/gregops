import {Args, Command, Flags, ux} from '@oclif/core'
import * as dns from 'node:dns'
import * as os from 'node:os'

export default class MyIP extends Command {
  async run(): Promise<void> {
    // const options = {family: 4}

    // dns.lookup(os.hostname(), options, (err, addr) => {
    //   if (err) {
    //     console.error(err);
    //   } else {
    //     console.log(`IPv4 address: ${addr}`);
    //   }
    // })

    const networkInterfaces = os.networkInterfaces();

    this.log(networkInterfaces)
  }
}

// dns.lookup(os.hostname(), options, (err, addr) => {
//   if (err) {
//     console.error(err);
//   } else {
//     console.log(`IPv4 address: ${addr}`);
//   }
// });
