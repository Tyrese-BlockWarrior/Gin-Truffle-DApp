const Cert = artifacts.require("Cert");
const { readFileSync, writeFileSync } = require("fs");

module.exports = function (_deployer) {
  // Use deployer to state migration tasks.
  _deployer
    .deploy(Cert)
    .then((instance) => {
      const envData = readFileSync(".env", "utf8");
      const envLines = envData.split('\n');
      for (let i = 0; i < envLines.length; i++) {
        let line = envLines[i];
        if (line.startsWith('#')) {
          continue;
        }
        if (line.startsWith('CONTRACT_ADDRESS')) {
          line = `CONTRACT_ADDRESS=${instance.address}`
        }
        if (line.startsWith('DEPLOYER_ADDRESS')) {
          line = `DEPLOYER_ADDRESS=${instance.constructor.class_defaults.from}`
        }
        envLines[i] = line;
      }
      const updatedEnvData = envLines.join('\n');
      writeFileSync(".env", updatedEnvData);
      writeFileSync("Cert.abi", JSON.stringify(instance.abi, null, 2));
    })
    .catch((err) => {
      console.log(err);
    });
};
